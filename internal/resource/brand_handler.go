// backend/internal/resource/brand_handler.go
package resource

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project/ent"
	"project/ent/brand"
	"project/ent/kitchen"
)

// BrandHandler handles API requests for brands
type BrandHandler struct {
	client *ent.Client
}

// NewBrandHandler creates a new brand handler
func NewBrandHandler(client *ent.Client) *BrandHandler {
	return &BrandHandler{
		client: client,
	}
}

// CreateBrand creates a new brand
func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var input struct {
		KitchenID   int    `json:"kitchen_id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		CuisineType string `json:"cuisine_type" binding:"required"`
		LogoURL     string `json:"logo_url"`
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

	// Create brand
	b, err := h.client.Brand.
		Create().
		SetKitchenID(input.KitchenID).
		SetName(input.Name).
		SetCuisineType(input.CuisineType).
		SetLogoURL(input.LogoURL).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, b)
}

// GetBrands returns all brands
func (h *BrandHandler) GetBrands(c *gin.Context) {
	// Parse query parameters
	kitchenID := c.Query("kitchen_id")
	
	// Build query
	query := h.client.Brand.Query()
	
	// Filter by kitchen if provided
	if kitchenID != "" {
		id, err := strconv.Atoi(kitchenID)
		if err == nil {
			query = query.Where(brand.HasKitchenWith(kitchen.ID(id)))
		}
	}
	
	// Execute query
	brands, err := query.All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brands)
}

// GetBrandByID returns a specific brand
func (h *BrandHandler) GetBrandByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	b, err := h.client.Brand.
		Query().
		Where(brand.ID(id)).
		WithKitchen().
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// UpdateBrand updates a brand
func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		CuisineType string `json:"cuisine_type"`
		LogoURL     string `json:"logo_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := h.client.Brand.UpdateOneID(id)
	
	if input.Name != "" {
		update.SetName(input.Name)
	}
	if input.CuisineType != "" {
		update.SetCuisineType(input.CuisineType)
	}
	if input.LogoURL != "" {
		update.SetLogoURL(input.LogoURL)
	}

	b, err := update.Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}

// DeleteBrand deletes a brand
func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.client.Brand.
		DeleteOneID(id).
		Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}