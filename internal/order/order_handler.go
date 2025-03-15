// backend/internal/order/order_handler.go
package order

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"project/ent"
	"project/ent/order"
)

// OrderHandler handles API requests for orders
type OrderHandler struct {
	client *ent.Client
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(client *ent.Client) *OrderHandler {
	return &OrderHandler{
		client: client,
	}
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input struct {
		BrandID    int       `json:"brand_id" binding:"required"`
		RequiredBy time.Time `json:"required_by" binding:"required"`
		Items      []struct {
			MenuItemID int    `json:"menu_item_id" binding:"required"`
			Quantity   int    `json:"quantity" binding:"required,min=1"`
			Notes      string `json:"special_instructions"`
		} `json:"items" binding:"required,dive"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start a transaction
	tx, err := h.client.Tx(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	// Create order
	o, err := tx.Order.
		Create().
		SetBrandID(input.BrandID).
		SetCreatedAt(time.Now()).
		SetRequiredBy(input.RequiredBy).
		SetStatus("pending").
		SetTotalPrice(0).  // Will be calculated after adding items
		Save(c.Request.Context())

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add items to order
	totalPrice := 0.0
	for _, item := range input.Items {
		// Get menu item to retrieve price
		menuItem, err := tx.MenuItem.Get(c.Request.Context(), item.MenuItemID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Menu item not found: " + strconv.Itoa(item.MenuItemID)})
			return
		}

		// Create order item
		_, err = tx.OrderItem.
			Create().
			SetOrder(o).
			SetMenuItemID(item.MenuItemID).
			SetQuantity(item.Quantity).
			SetSpecialInstructions(item.Notes).
			Save(c.Request.Context())

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order item"})
			return
		}

		// Add to total price
		totalPrice += menuItem.Price * float64(item.Quantity)
	}

	// Update order with calculated total price
	_, err = tx.Order.
		UpdateOne(o).
		SetTotalPrice(totalPrice).
		Save(c.Request.Context())

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order total"})
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Return created order
	createdOrder, err := h.client.Order.
		Query().
		Where(order.ID(o.ID)).
		WithItems().
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order created but failed to retrieve details"})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

// GetOrders retrieves all orders with optional filters
func (h *OrderHandler) GetOrders(c *gin.Context) {
	// Parse query parameters
	brandID := c.Query("brand_id")
	status := c.Query("status")
	
	// Build query
	query := h.client.Order.Query()
	
	// Apply filters if provided
	if brandID != "" {
		id, err := strconv.Atoi(brandID)
		if err == nil {
			query = query.Where(order.BrandID(id))
		}
	}
	
	if status != "" {
		query = query.Where(order.StatusEQ(status))
	}
	
	// Execute query
	orders, err := query.
		WithItems().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrderByID retrieves a specific order
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	o, err := h.client.Order.
		Query().
		Where(order.ID(id)).
		WithItems(func(query *ent.OrderItemQuery) {
			query.WithMenuItem()
		}).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, o)
}

// UpdateOrderStatus updates an order's status
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":   true,
		"preparing": true,
		"ready":     true,
		"completed": true,
		"cancelled": true,
	}

	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	// Update order status
	_, err = h.client.Order.
		UpdateOneID(id).
		SetStatus(input.Status).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}

// CancelOrder cancels an order
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Check if order exists and can be cancelled
	o, err := h.client.Order.
		Query().
		Where(order.ID(id)).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Only pending or preparing orders can be cancelled
	if o.Status != "pending" && o.Status != "preparing" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only pending or preparing orders can be cancelled"})
		return
	}

	// Update order status
	_, err = h.client.Order.
		UpdateOneID(id).
		SetStatus("cancelled").
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}