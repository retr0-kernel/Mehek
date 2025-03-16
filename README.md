# Mehek

A comprehensive system for managing ghost kitchen operations, optimizing resource allocation, and managing multi-tenant food delivery brands.

## System Overview

Mehek is a full-stack application that helps multiple food delivery brands share kitchen resources efficiently. The system provides:

    1. Real-time kitchen resource allocation (cooking stations, equipment, staff)

    2. Predictive ingredient ordering and inventory management
    
    3. Multi-tenant brand management with independent menus
    
    4. Order coordination and batching for kitchen efficiency
    
    5. Performance analytics and kitchen utilization metrics

![Flow](/images/Mehek.png)


## Tech Stack

### Frontend

- **Next.js**: React framework for server-side rendering and static site generation

- **PrimeReact:** Component library for building the user interface

- **WebSockets:** For real-time updates between client and server

### Backend

- **Go:** High-performance language for the API server

- **ent:** Entity framework for Go, providing an ORM-like experience

- **Gin:** Web framework for building the RESTful API

### Task Processing

- **Celery:** Distributed task queue for handling asynchronous operations

- **RabbitMQ:** Message broker for coordinating tasks between services

### Database & Caching

- **PostgreSQL:** Primary relational database for data persistence

- **Redis:** In-memory data store for caching and real-time operations

- **Elasticsearch:** Search and analytics engine for performance metrics

## Data Model

![Data Models](/images/DataModels.png)

### Kitchen Management

- **Kitchen:** Central entity representing a physical kitchen location

- **KitchenResource:** Equipment and facilities in the kitchen

- **Staff:** Personnel working in the kitchen

- **Shift:** Work schedules for staff members

### Brand Management

- **Brand:** Food delivery brands operating in the kitchen

- **Menu:** Collections of food items organized by brand

- **MenuItem:** Individual food items available for order

- **Ingredient:** Raw materials used to prepare menu items

### Order Management

- **Order:** Customer orders for specific brands

- **OrderItem:** Individual items within an order

### Inventory Management

- **InventoryItem:** Stock of ingredients with quantities and expiration dates

### Resource Allocation

- **ResourceAllocation:** Assignments of kitchen resources to specific orders




## Core Features

### Resource Allocation

- Real-time tracking of kitchen resource availability
- Smart allocation of cooking stations and equipment based on order requirements
- Conflict resolution for resource scheduling
- Resource utilization analytics

### Brand Management

- Multi-tenant support for different food delivery brands
- Independent menu and pricing management
- Brand-specific performance metrics
- Customizable order workflows

### Order Coordination

- Order batching for optimized kitchen operations
- Priority-based order scheduling
- Real-time order status tracking
- Integration with delivery services

### Inventory Management

- Automatic inventory deduction based on orders
- Predictive ingredient ordering
- Expiration date tracking
- Inventory utilization analytics

### Analytics Dashboard

- Kitchen utilization metrics
- Brand performance comparisons
- Order fulfillment statistics
- Predictive analytics for demand forecasting

