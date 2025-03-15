from celery import Celery
import os
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

# Configure Celery
app = Celery(
    'ghost_kitchen',
    broker=os.getenv('RABBITMQ_URL', 'amqp://guest:guest@localhost:5672/'),
    include=[
        'tasks.resource_allocation',
        'tasks.order_processing',
        'tasks.inventory',
        'tasks.analytics'
    ]
)

# Configure Celery settings
app.conf.update(
    result_backend=os.getenv('REDIS_URL', 'redis://localhost:6379/0'),
    task_serializer='json',
    accept_content=['json'],
    result_serializer='json',
    timezone='UTC',
    enable_utc=True,
    task_track_started=True,
    task_time_limit=600,  # 10 minutes
    worker_concurrency=4,
    worker_prefetch_multiplier=1,
)

if __name__ == '__main__':
    app.start()