# backend/tasks/analytics.py
from .celery_app import app
from sqlalchemy import create_engine, text
import os
import json
import datetime
import pandas as pd
import logging
from dotenv import load_dotenv
import matplotlib.pyplot as plt
import io
import base64

# Load environment variables
load_dotenv()

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Database connection
db_url = os.getenv('DATABASE_URL', 'postgresql://postgres:postgres@localhost:5432/ghost_kitchen')
engine = create_engine(db_url)

@app.task(name='analytics.analyze_kitchen_performance')
def analyze_kitchen_performance(kitchen_id, start_date=None, end_date=None):
    """
    Analyze kitchen performance metrics
    """
    logger.info(f"Analyzing performance for kitchen {kitchen_id}")
    
    try:
        # Convert string dates to datetime if provided
        if isinstance(start_date, str):
            start_date = datetime.datetime.fromisoformat(start_date.replace('Z', '+00:00'))
        if isinstance(end_date, str):
            end_date = datetime.datetime.fromisoformat(end_date.replace('Z', '+00:00'))
            
        # Use last 30 days if dates not provided
        if not start_date:
            start_date = datetime.datetime.now() - datetime.timedelta(days=30)
        if not end_date:
            end_date = datetime.datetime.now()
        
        # Get kitchen orders data
        query = """
        SELECT o.id, o.created_at, o.required_by, o.status, o.total_price,
               b.name as brand_name, b.cuisine_type
        FROM "order" o
        JOIN brand b ON o.brand_id = b.id
        WHERE b.kitchen_id = :kitchen_id
          AND o.created_at BETWEEN :start_date AND :end_date
        ORDER BY o.created_at
        """
        
        orders_df = pd.read_sql(
            text(query),
            engine,
            params={
                'kitchen_id': kitchen_id,
                'start_date': start_date,
                'end_date': end_date
            }
        )
        
        if orders_df.empty:
            return {
                "status": "success",
                "message": "No orders found for analysis",
                "kitchen_id": kitchen_id,
                "data": {}
            }
        
        # Calculate key metrics
        total_orders = len(orders_df)
        total_revenue = orders_df['total_price'].sum()
        avg_order_value = total_revenue / total_orders if total_orders > 0 else 0
        orders_by_status = orders_df['status'].value_counts().to_dict()
        orders_by_brand = orders_df['brand_name'].value_counts().to_dict()
        orders_by_cuisine = orders_df['cuisine_type'].value_counts().to_dict()
        
        # Calculate daily order counts
        orders_df['date'] = orders_df['created_at'].dt.date
        daily_orders = orders_df.groupby('date').size()
        
        # Get resource utilization data
        resource_query = """
        SELECT kr.name, kr.type, count(ra.id) as usage_count,
               avg(EXTRACT(EPOCH FROM (ra.end_time - ra.start_time)) / 60) as avg_usage_minutes
        FROM resource_allocation ra
        JOIN kitchen_resource kr ON ra.resource_id = kr.id
        JOIN "order" o ON ra.order_id = o.id
        JOIN brand b ON o.brand_id = b.id
        WHERE b.kitchen_id = :kitchen_id
          AND ra.start_time BETWEEN :start_date AND :end_date
        GROUP BY kr.id, kr.name, kr.type
        ORDER BY usage_count DESC
        """
        
        resource_df = pd.read_sql(
            text(resource_query),
            engine,
            params={
                'kitchen_id': kitchen_id,
                'start_date': start_date,
                'end_date': end_date
            }
        )
        
        # Generate charts
        charts = {}
        
        if not orders_df.empty:
            # Daily orders chart
            plt.figure(figsize=(10, 6))
            plt.plot(daily_orders.index, daily_orders.values)
            plt.title('Daily Order Volume')
            plt.xlabel('Date')
            plt.ylabel('Number of Orders')
            plt.grid(True)
            plt.tight_layout()
            
            # Convert plot to base64 encoded string
            buf = io.BytesIO()
            plt.savefig(buf, format='png')
            buf.seek(0)
            daily_orders_chart = base64.b64encode(buf.read()).decode('utf-8')
            plt.close()
            
            charts['daily_orders'] = daily_orders_chart
            
            # Order status distribution
            plt.figure(figsize=(8, 8))
            plt.pie(
                orders_by_status.values(),
                labels=orders_by_status.keys(),
                autopct='%1.1f%%',
                startangle=90
            )
            plt.title('Order Status Distribution')
            plt.axis('equal')
            plt.tight_layout()
            
            buf = io.BytesIO()
            plt.savefig(buf, format='png')
            buf.seek(0)
            status_chart = base64.b64encode(buf.read()).decode('utf-8')
            plt.close()
            
            charts['status_distribution'] = status_chart
        
        # Prepare resource usage data for frontend
        resource_usage = []
        if not resource_df.empty:
            for _, row in resource_df.iterrows():
                resource_usage.append({
                    'name': row['name'],
                    'type': row['type'],
                    'usage_count': int(row['usage_count']),
                    'avg_usage_minutes': float(row['avg_usage_minutes'])
                })
        
        # Calculate kitchen utilization metrics
        utilization_query = """
        WITH resource_times AS (
            SELECT 
                EXTRACT(EPOCH FROM (ra.end_time - ra.start_time)) / 3600 as usage_hours,
                kr.id as resource_id
            FROM resource_allocation ra
            JOIN kitchen_resource kr ON ra.resource_id = kr.id
            WHERE kr.kitchen_id = :kitchen_id
              AND ra.start_time BETWEEN :start_date AND :end_date
        ),
        total_resources AS (
            SELECT COUNT(id) as resource_count
            FROM kitchen_resource
            WHERE kitchen_id = :kitchen_id
        )
        SELECT 
            SUM(usage_hours) as total_usage_hours,
            (SELECT resource_count FROM total_resources) as resource_count
        FROM resource_times
        """
        
        util_df = pd.read_sql(
            text(utilization_query),
            engine,
            params={
                'kitchen_id': kitchen_id,
                'start_date': start_date,
                'end_date': end_date
            }
        )
        
        # Calculate utilization percentage
        total_hours = (end_date - start_date).total_seconds() / 3600
        resource_count = util_df['resource_count'].iloc[0] if not util_df.empty else 0
        total_capacity_hours = total_hours * resource_count
        
        total_usage_hours = util_df['total_usage_hours'].iloc[0] if not util_df.empty else 0
        utilization_percentage = (total_usage_hours / total_capacity_hours * 100) if total_capacity_hours > 0 else 0
        
        # Put everything together for the report
        report = {
            "status": "success",
            "kitchen_id": kitchen_id,
            "period": {
                "start_date": start_date.isoformat(),
                "end_date": end_date.isoformat()
            },
            "summary": {
                "total_orders": total_orders,
                "total_revenue": float(total_revenue),
                "avg_order_value": float(avg_order_value),
                "utilization_percentage": float(utilization_percentage)
            },
            "orders": {
                "by_status": orders_by_status,
                "by_brand": orders_by_brand,
                "by_cuisine": orders_by_cuisine,
                "daily_counts": daily_orders.to_dict()
            },
            "resource_usage": resource_usage,
            "charts": charts
        }
        
        # Save report to database or file system if needed
        # For now, just return it
        logger.info(f"Completed performance analysis for kitchen {kitchen_id}")
        return report
        
    except Exception as e:
        logger.error(f"Error analyzing kitchen performance: {str(e)}")
        return {
            "status": "error",
            "message": str(e),
            "kitchen_id": kitchen_id
        }

@app.task(name='analytics.predict_ingredient_demand')
def predict_ingredient_demand(kitchen_id):
    """
    Predict ingredient demand for the next 7 days based on historical data
    """
    logger.info(f"Predicting ingredient demand for kitchen {kitchen_id}")
    
    try:
        # Get historical order data for last 30 days
        query = """
        SELECT 
            i.ingredient_id, 
            ing.name as ingredient_name,
            ing.unit,
            SUM(oi.quantity) as ordered_quantity,
            DATE(o.created_at) as order_date
        FROM "order" o
        JOIN order_item oi ON o.id = oi.order_id
        JOIN menu_item mi ON oi.menu_item_id = mi.id
        JOIN menu_item_ingredients i ON mi.id = i.menu_item_id
        JOIN ingredient ing ON i.ingredient_id = ing.id
        JOIN brand b ON o.brand_id = b.id
        WHERE b.kitchen_id = :kitchen_id
          AND o.created_at >= NOW() - INTERVAL '30 days'
        GROUP BY i.ingredient_id, ing.name, ing.unit, DATE(o.created_at)
        ORDER BY order_date, ingredient_name
        """
        
        df = pd.read_sql(
            text(query),
            engine,
            params={'kitchen_id': kitchen_id}
        )
        
        if df.empty:
            return {
                "status": "success",
                "message": "No historical data found for prediction",
                "kitchen_id": kitchen_id,
                "predictions": []
            }
        
        # Create time series for each ingredient
        predictions = []
        for ingredient_id in df['ingredient_id'].unique():
            ingredient_df = df[df['ingredient_id'] == ingredient_id]
            name = ingredient_df['ingredient_name'].iloc[0]
            unit = ingredient_df['unit'].iloc[0]
            
            # Create a complete date range
            date_range = pd.date_range(
                end=datetime.datetime.now().date(),
                periods=30,
                freq='D'
            )
            
            # Create complete time series with zeros for missing dates
            ts = pd.DataFrame({'date': date_range})
            ts['ingredient_id'] = ingredient_id
            ts = ts.merge(
                ingredient_df[['order_date', 'ordered_quantity']],
                left_on='date',
                right_on='order_date',
                how='left'
            )
            ts['ordered_quantity'].fillna(0, inplace=True)
            
            # Simple moving average forecast for next 7 days
            avg_daily = ts['ordered_quantity'].mean()
            
            # Add 20% buffer for safety stock
            predicted_amount = avg_daily * 7 * 1.2
            
            predictions.append({
                'ingredient_id': int(ingredient_id),
                'name': name,
                'unit': unit,
                'avg_daily_usage': float(avg_daily),
                'predicted_amount_7_days': float(predicted_amount)
            })
        
        return {
            "status": "success",
            "kitchen_id": kitchen_id,
            "predictions": predictions
        }
        
    except Exception as e:
        logger.error(f"Error predicting ingredient demand: {str(e)}")
        return {
            "status": "error",
            "message": str(e),
            "kitchen_id": kitchen_id
        }