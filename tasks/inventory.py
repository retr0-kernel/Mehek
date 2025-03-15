# backend/tasks/inventory.py
from .celery_app import app
from sqlalchemy import create_engine, text
import os
import json
import datetime
import pandas as pd
import logging
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Database connection
db_url = os.getenv('DATABASE_URL', 'postgresql://postgres:postgres@localhost:5432/ghost_kitchen')
engine = create_engine(db_url)

@app.task(name='inventory.generate_inventory_report')
def generate_inventory_report(kitchen_id):
    """
    Generate a comprehensive inventory report for a kitchen
    """
    logger.info(f"Generating inventory report for kitchen {kitchen_id}")
    
    try:
        # Get current inventory levels
        query = """
        SELECT 
            i.id as inventory_id,
            i.quantity,
            i.expiration_date,
            ing.id as ingredient_id,
            ing.name as ingredient_name,
            ing.unit,
            ing.cost_per_unit
        FROM inventory_item i
        JOIN ingredient ing ON i.ingredient_id = ing.id
        WHERE i.kitchen_id = :kitchen_id
        ORDER BY ing.name
        """
        
        inventory_df = pd.read_sql(
            text(query),
            engine,
            params={'kitchen_id': kitchen_id}
        )
        
        if inventory_df.empty:
            return {
                "status": "success",
                "message": "No inventory found for this kitchen",
                "kitchen_id": kitchen_id,
                "data": {}
            }
        
        # Calculate inventory metrics
        total_items = len(inventory_df)
        total_value = (inventory_df['quantity'] * inventory_df['cost_per_unit']).sum()
        
        # Check for items nearing expiration (within 7 days)
        today = datetime.datetime.now().date()
        expiring_soon = inventory_df[
            inventory_df['expiration_date'].notna() &
            (inventory_df['expiration_date'].dt.date - today <= datetime.timedelta(days=7))
        ]
        
        expiring_items = []
        for _, row in expiring_soon.iterrows():
            expiring_items.append({
                'id': int(row['inventory_id']),
                'ingredient_name': row['ingredient_name'],
                'quantity': float(row['quantity']),
                'unit': row['unit'],
                'expiration_date': row['expiration_date'].isoformat() if not pd.isna(row['expiration_date']) else None,
                'days_remaining': (row['expiration_date'].date() - today).days if not pd.isna(row['expiration_date']) else None
            })
        
        # Calculate days of supply based on usage patterns
        usage_query = """
        WITH ingredient_usage AS (
            SELECT 
                i.ingredient_id,
                DATE(o.created_at) as usage_date,
                SUM(oi.quantity * mii.quantity) as daily_usage
            FROM "order" o
            JOIN order_item oi ON o.id = oi.order_id
            JOIN menu_item mi ON oi.menu_item_id = mi.id
            JOIN menu_item_ingredients mii ON mi.id = mii.menu_item_id
            JOIN ingredient i ON mii.ingredient_id = i.id
            JOIN brand b ON o.brand_id = b.id
            WHERE b.kitchen_id = :kitchen_id
              AND o.created_at >= NOW() - INTERVAL '30 days'
            GROUP BY i.ingredient_id, DATE(o.created_at)
        )
        SELECT 
            ingredient_id,
            AVG(daily_usage) as avg_daily_usage
        FROM ingredient_usage
        GROUP BY ingredient_id
        """
        
        usage_df = pd.read_sql(
            text(usage_query),
            engine,
            params={'kitchen_id': kitchen_id}
        )
        
        # Merge with inventory data
        inventory_with_usage = pd.merge(
            inventory_df,
            usage_df,
            left_on='ingredient_id',
            right_on='ingredient_id',
            how='left'
        )
        
        inventory_with_usage['avg_daily_usage'].fillna(0, inplace=True)
        inventory_with_usage['days_of_supply'] = inventory_with_usage.apply(
            lambda row: float('inf') if row['avg_daily_usage'] == 0 else row['quantity'] / row['avg_daily_usage'],
            axis=1
        )
        
        # Identify low stock items (less than 7 days of supply)
        low_stock = inventory_with_usage[
            (inventory_with_usage['days_of_supply'] < 7) & 
            (inventory_with_usage['days_of_supply'] != float('inf'))
        ]
        
        low_stock_items = []
        for _, row in low_stock.iterrows():
            low_stock_items.append({
                'id': int(row['inventory_id']),
                'ingredient_name': row['ingredient_name'],
                'quantity': float(row['quantity']),
                'unit': row['unit'],
                'avg_daily_usage': float(row['avg_daily_usage']),
                'days_of_supply': float(row['days_of_supply']),
                'suggested_order': float(row['avg_daily_usage'] * 14 - row['quantity'])  # 2 week supply
            })
        
        # Format inventory items for report
        inventory_items = []
        for _, row in inventory_df.iterrows():
            inventory_items.append({
                'id': int(row['inventory_id']),
                'ingredient_id': int(row['ingredient_id']),
                'ingredient_name': row['ingredient_name'],
                'quantity': float(row['quantity']),
                'unit': row['unit'],
                'cost_per_unit': float(row['cost_per_unit']),
                'total_value': float(row['quantity'] * row['cost_per_unit']),
                'expiration_date': row['expiration_date'].isoformat() if not pd.isna(row['expiration_date']) else None
            })
        
        # Prepare the report
        report = {
            "status": "success",
            "kitchen_id": kitchen_id,
            "generated_at": datetime.datetime.now().isoformat(),
            "summary": {
                "total_items": total_items,
                "total_value": float(total_value),
                "expiring_soon_count": len(expiring_items),
                "low_stock_count": len(low_stock_items)
            },
            "inventory": inventory_items,
            "alerts": {
                "expiring_soon": expiring_items,
                "low_stock": low_stock_items
            }
        }
        
        logger.info(f"Completed inventory report for kitchen {kitchen_id}")
        return report
        
    except Exception as e:
        logger.error(f"Error generating inventory report: {str(e)}")
        return {
            "status": "error",
            "message": str(e),
            "kitchen_id": kitchen_id
        }

@app.task(name='inventory.check_inventory_levels')
def check_inventory_levels(kitchen_id):
    """
    Check inventory levels against minimum thresholds and upcoming order requirements
    """
    logger.info(f"Checking inventory levels for kitchen {kitchen_id}")
    
    try:
        # Get current inventory
        inventory_query = """
        SELECT 
            i.id, 
            i.ingredient_id, 
            ing.name as ingredient_name,
            i.quantity,
            ing.unit
        FROM inventory_item i
        JOIN ingredient ing ON i.ingredient_id = ing.id
        WHERE i.kitchen_id = :kitchen_id
        """
        
        inventory_df = pd.read_sql(
            text(inventory_query),
            engine,
            params={'kitchen_id': kitchen_id}
        )
        
        # Get upcoming orders (next 24 hours)
        upcoming_orders_query = """
        WITH upcoming_ingredients AS (
            SELECT 
                i.ingredient_id,
                SUM(oi.quantity * mi_ing.quantity) as required_quantity
            FROM "order" o
            JOIN order_item oi ON o.id = oi.order_id
            JOIN menu_item mi ON oi.menu_item_id = mi.id
            JOIN menu_item_ingredients mi_ing ON mi.id = mi_ing.menu_item_id
            JOIN ingredient i ON mi_ing.ingredient_id = i.id
            JOIN brand b ON o.brand_id = b.id
            WHERE b.kitchen_id = :kitchen_id
              AND o.status = 'pending'
              AND o.required_by <= NOW() + INTERVAL '24 hours'
            GROUP BY i.ingredient_id
        )
        SELECT 
            ui.ingredient_id,
            ing.name as ingredient_name,
            ui.required_quantity,
            ing.unit
        FROM upcoming_ingredients ui
        JOIN ingredient ing ON ui.ingredient_id = ing.id
        """
        
        upcoming_df = pd.read_sql(
            text(upcoming_orders_query),
            engine,
            params={'kitchen_id': kitchen_id}
        )
        
        # Merge to check sufficiency
        if not upcoming_df.empty:
            merged = pd.merge(
                upcoming_df,
                inventory_df,
                on=['ingredient_id', 'ingredient_name', 'unit'],
                how='left'
            )
            
            # Find insufficient ingredients
            merged['sufficient'] = merged['quantity'] >= merged['required_quantity']
            insufficient = merged[~merged['sufficient']]
            
            alerts = []
            for _, row in insufficient.iterrows():
                alerts.append({
                    'ingredient_id': int(row['ingredient_id']),
                    'ingredient_name': row['ingredient_name'],
                    'current_quantity': float(row['quantity']) if not pd.isna(row['quantity']) else 0,
                    'required_quantity': float(row['required_quantity']),
                    'shortage': float(row['required_quantity'] - (row['quantity'] if not pd.isna(row['quantity']) else 0)),
                    'unit': row['unit']
                })
            
            return {
                "status": "success",
                "kitchen_id": kitchen_id,
                "has_alerts": len(alerts) > 0,
                "alerts": alerts
            }
        
        return {
            "status": "success",
            "kitchen_id": kitchen_id,
            "has_alerts": False,
            "alerts": []
        }
        
    except Exception as e:
        logger.error(f"Error checking inventory levels: {str(e)}")
        return {
            "status": "error",
            "message": str(e),
            "kitchen_id": kitchen_id
        }