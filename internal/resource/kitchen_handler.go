package resource

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project/ent"
	"project/ent/kitchen"
)

// KitchenHandler handles API requests for kitchens
type KitchenHandler struct {
	client *ent.Client
}

// NewKitchenHandler creates a new kitchen handler
func NewKitchenHandler(client *ent.Client) *KitchenHandler {
	return &KitchenHandler{
		client: client,
	}
}

// CreateKitchen creates a new kitchen
func (h *KitchenHandler) CreateKitchen(c *gin.Context) {
	var input struct {
		Name           string              `json:"name" binding:"required"`
		Location       string              `json:"location" binding:"required"`
		Capacity       int                 `json:"capacity" binding:"required"`
		OperatingHours map[string][]string `json:"operating_hours"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	k, err := h.client.Kitchen.
		Create().
		SetName(input.Name).
		SetLocation(input.Location).
		SetCapacity(input.Capacity).
		SetOperatingHours(input.OperatingHours).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, k)
}

// GetKitchens returns a list of kitchens
func (h *KitchenHandler) GetKitchens(c *gin.Context) {
	kitchens, err := h.client.Kitchen.
		Query().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kitchens)
}

// GetKitchenByID returns a specific kitchen by ID
func (h *KitchenHandler) GetKitchenByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	k, err := h.client.Kitchen.
		Query().
		Where(kitchen.ID(id)).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kitchen not found"})
		return
	}

	c.JSON(http.StatusOK, k)
}

// UpdateKitchen updates a kitchen
func (h *KitchenHandler) UpdateKitchen(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input struct {
		Name           string              `json:"name"`
		Location       string              `json:"location"`
		Capacity       int                 `json:"capacity"`
		OperatingHours map[string][]string `json:"operating_hours"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := h.client.Kitchen.UpdateOneID(id)
	
	if input.Name != "" {
		update.SetName(input.Name)
	}
	if input.Location != "" {
		update.SetLocation(input.Location)
	}
	if input.Capacity != 0 {
		update.SetCapacity(input.Capacity)
	}
	if input.OperatingHours != nil {
		update.SetOperatingHours(input.OperatingHours)
	}

	k, err := update.Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, k)
}

// DeleteKitchen deletes a kitchen
func (h *KitchenHandler) DeleteKitchen(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.client.Kitchen.
		DeleteOneID(id).
		Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kitchen deleted successfully"})
}