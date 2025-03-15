from .celery_app import app
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base
import os
import json
import datetime
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
Session = sessionmaker(bind=engine)
Base = declarative_base()

@app.task(name='resource_allocation.allocate_resources_for_order')
def allocate_resources_for_order(order_id):
    """
    Allocate kitchen resources for a specific order
    """
    logger.info(f"Starting resource allocation for order {order_id}")
    
    try:
        session = Session()
        
        # Get order details from database
        order_query = """
        SELECT o.id, o.brand_id, o.required_by, b.kitchen_id
        FROM "order" o
        JOIN brand b ON o.brand_id = b.id
        WHERE o.id = %s
        """
        
        order_result = engine.execute(order_query, (order_id,)).fetchone()
        if not order_result:
            logger.error(f"Order {order_id} not found")
            return {"status": "error", "message": "Order not found"}
        
        order = {
            "id": order_result[0],
            "brand_id": order_result[1],
            "required_by": order_result[2],
            "kitchen_id": order_result[3]
        }
        
        # Get order items
        items_query = """
        SELECT oi.id, oi.menu_item_id, oi.quantity, mi.prep_time, mi.equipment_needed
        FROM order_item oi
        JOIN menu_item mi ON oi.menu_item_id = mi.id
        WHERE oi.order_id = %s
        """
        
        items_result = engine.execute(items_query, (order_id,))
        order_items = []
        
        for row in items_result:
            order_items.append({
                "id": row[0],
                "menu_item_id": row[1],
                "quantity": row[2],
                "prep_time": row[3],
                "equipment_needed": row[4]
            })
        
        # Find available resources in the kitchen
        resources_query = """
        SELECT id, name, type, capacity, available
        FROM kitchen_resource
        WHERE kitchen_id = %s AND available = TRUE
        """
        
        resources_result = engine.execute(resources_query, (order["kitchen_id"],))
        available_resources = []
        
        for row in resources_result:
            available_resources.append({
                "id": row[0],
                "name": row[1],
                "type": row[2],
                "capacity": row[3],
                "available": row[4]
            })
        
        # Find staff on shift
        staff_query = """
        SELECT s.id, sh.id as shift_id
        FROM staff s
        JOIN shift sh ON s.id = sh.staff_id
        WHERE s.kitchen_id = %s
        AND sh.start_time <= %s AND sh.end_time >= %s
        """
        
        now = datetime.datetime.now()
        staff_result = engine.execute(staff_query, (order["kitchen_id"], now, now))
        available_staff = []
        
        for row in staff_result:
            available_staff.append({
                "staff_id": row[0],
                "shift_id": row[1]
            })
            
        if not available_staff:
            logger.error(f"No staff available for order {order_id}")
            return {"status": "error", "message": "No staff available"}
        
        # Simple allocation strategy - assign first available resource of each required type
        allocations = []
        required_equipment = set()
        
        for item in order_items:
            equipment_list = item["equipment_needed"].split(",")
            for eq in equipment_list:
                required_equipment.add(eq.strip())
        
        for equipment in required_equipment:
            for resource in available_resources:
                if resource["type"] == equipment and resource["available"]:
                    # Allocate this resource
                    start_time = datetime.datetime.now()
                    end_time = start_time + datetime.timedelta(minutes=max(item["prep_time"] for item in order_items))
                    
                    allocations.append({
                        "resource_id": resource["id"],
                        "order_id": order_id,
                        "shift_id": available_staff[0]["shift_id"],  # Assign to first available staff member
                        "start_time": start_time,
                        "end_time": end_time,
                        "status": "scheduled"
                    })
                    break
        
        # Save allocations to database
        for allocation in allocations:
            insert_query = """
            INSERT INTO resource_allocation 
            (resource_id, order_id, shift_id, start_time, end_time, status)
            VALUES (%s, %s, %s, %s, %s, %s)
            """
            
            engine.execute(
                insert_query, 
                (
                    allocation["resource_id"],
                    allocation["order_id"], 
                    allocation["shift_id"],
                    allocation["start_time"],
                    allocation["end_time"],
                    allocation["status"]
                )
            )
        
        logger.info(f"Completed resource allocation for order {order_id}, created {len(allocations)} allocations")
        return {
            "status": "success", 
            "order_id": order_id, 
            "allocations": len(allocations)
        }
        
    except Exception as e:
        logger.error(f"Error allocating resources for order {order_id}: {str(e)}")
        return {"status": "error", "message": str(e)}
    
    finally:
        session.close()

@app.task(name='resource_allocation.optimize_kitchen_resources')
def optimize_kitchen_resources(kitchen_id):
    """
    Optimize resource allocation across a kitchen based on upcoming orders
    """
    logger.info(f"Starting kitchen resource optimization for kitchen {kitchen_id}")
    
    try:
        # Detailed implementation would include:
        # 1. Get all upcoming orders
        # 2. Calculate resource usage patterns
        # 3. Identify bottlenecks
        # 4. Suggest resource reallocation
        
        # Simplified placeholder
        return {
            "status": "success",
            "kitchen_id": kitchen_id,
            "message": "Optimization complete"
        }
        
    except Exception as e:
        logger.error(f"Error optimizing resources for kitchen {kitchen_id}: {str(e)}")
        return {"status": "error", "message": str(e)}