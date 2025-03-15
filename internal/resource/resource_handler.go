// backend/internal/resource/resource_handler.go
package resource

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project/ent"
	"project/ent/kitchenresource"
	"project/ent/kitchen"
)

// ResourceHandler handles API requests for kitchen resources
type ResourceHandler struct {
	client *ent.Client
}

// NewResourceHandler creates a new resource handler
func NewResourceHandler(client *ent.Client) *ResourceHandler {
	return &ResourceHandler{
		client: client,
	}
}

// CreateResource creates a new kitchen resource
func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var input struct {
		KitchenID int    `json:"kitchen_id" binding:"required"`
		Name      string `json:"name" binding:"required"`
		Type      string `json:"type" binding:"required"`
		Capacity  int    `json:"capacity" binding:"required"`
		Available bool   `json:"available"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if kitchen exists
	_, err := h.client.Kitchen.Get(c.Request.Context(), input.KitchenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kitchen not found"})
		return
	}

	r, err := h.client.KitchenResource.
		Create().
		SetKitchenID(input.KitchenID).
		SetName(input.Name).
		SetType(input.Type).
		SetCapacity(input.Capacity).
		SetAvailable(input.Available).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, r)
}

// GetResourcesByKitchen returns all resources for a specific kitchen
func (h *ResourceHandler) GetResourcesByKitchen(c *gin.Context) {
	kitchenID, err := strconv.Atoi(c.Param("kitchen_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kitchen ID"})
		return
	}

	resources, err := h.client.KitchenResource.
    	Query().
    	Where(kitchenresource.HasKitchenWith(kitchen.ID(kitchenID))).
    	All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
}

// GetResourceByID returns a specific resource
func (h *ResourceHandler) GetResourceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	r, err := h.client.KitchenResource.
		Query().
		Where(kitchenresource.ID(id)).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	c.JSON(http.StatusOK, r)
}

// UpdateResource updates a kitchen resource
func (h *ResourceHandler) UpdateResource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Capacity  int    `json:"capacity"`
		Available bool   `json:"available"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := h.client.KitchenResource.UpdateOneID(id)
	
	if input.Name != "" {
		update.SetName(input.Name)
	}
	if input.Type != "" {
		update.SetType(input.Type)
	}
	if input.Capacity != 0 {
		update.SetCapacity(input.Capacity)
	}
	
	// Always update availability status since it's a boolean
	update.SetAvailable(input.Available)

	r, err := update.Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// DeleteResource deletes a kitchen resource
func (h *ResourceHandler) DeleteResource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.client.KitchenResource.
		DeleteOneID(id).
		Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}