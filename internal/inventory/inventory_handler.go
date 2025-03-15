// backend/internal/inventory/inventory_handler.go
package inventory

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"project/ent"
	"project/ent/ingredient"
	"project/ent/inventoryitem"
	"project/ent/kitchen"
)

// InventoryHandler handles API requests for inventory
type InventoryHandler struct {
	client *ent.Client
}

// NewInventoryHandler creates a new inventory handler
func NewInventoryHandler(client *ent.Client) *InventoryHandler {
	return &InventoryHandler{
		client: client,
	}
}

// CreateIngredient creates a new ingredient
func (h *InventoryHandler) CreateIngredient(c *gin.Context) {
	var input struct {
		Name        string  `json:"name" binding:"required"`
		Unit        string  `json:"unit" binding:"required"`
		CostPerUnit float64 `json:"cost_per_unit" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ing, err := h.client.Ingredient.
		Create().
		SetName(input.Name).
		SetUnit(input.Unit).
		SetCostPerUnit(input.CostPerUnit).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ing)
}

// GetIngredients returns all ingredients
func (h *InventoryHandler) GetIngredients(c *gin.Context) {
	ingredients, err := h.client.Ingredient.
		Query().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ingredients)
}

// AddInventory adds inventory of an ingredient to a kitchen
func (h *InventoryHandler) AddInventory(c *gin.Context) {
	var input struct {
		KitchenID      int        `json:"kitchen_id" binding:"required"`
		IngredientID   int        `json:"ingredient_id" binding:"required"`
		Quantity       float64    `json:"quantity" binding:"required,gt=0"`
		ExpirationDate *time.Time `json:"expiration_date"`
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

	// Check if ingredient exists
	_, err = h.client.Ingredient.Get(c.Request.Context(), input.IngredientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient not found"})
		return
	}

	// Check if inventory item already exists, update quantity if it does
	exists, err := h.client.InventoryItem.
		Query().
		Where(
			inventoryitem.HasKitchenWith(kitchen.ID(input.KitchenID)),
			inventoryitem.HasIngredientWith(ingredient.ID(input.IngredientID)),
		).
		Exist(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		// Update existing inventory
		inv, err := h.client.InventoryItem.
			Query().
			Where(
				inventoryitem.HasKitchenWith(kitchen.ID(input.KitchenID)),
				inventoryitem.HasIngredientWith(ingredient.ID(input.IngredientID)),
			).
			Only(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update quantity and expiration date
		update := h.client.InventoryItem.UpdateOneID(inv.ID).
			SetQuantity(inv.Quantity + input.Quantity)

		if input.ExpirationDate != nil {
			update.SetExpirationDate(*input.ExpirationDate)
		}

		inv, err = update.Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "Inventory updated successfully",
			"inventory": inv,
		})
		return
	}

	// Create new inventory item
	invCreate := h.client.InventoryItem.
		Create().
		SetKitchenID(input.KitchenID).
		SetIngredientID(input.IngredientID).
		SetQuantity(input.Quantity)

	if input.ExpirationDate != nil {
		invCreate.SetExpirationDate(*input.ExpirationDate)
	}

	inv, err := invCreate.Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, inv)
}

// GetKitchenInventory returns all inventory items for a kitchen
func (h *InventoryHandler) GetKitchenInventory(c *gin.Context) {
	kitchenID, err := strconv.Atoi(c.Param("kitchen_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kitchen ID"})
		return
	}

	inventory, err := h.client.InventoryItem.
		Query().
		Where(inventoryitem.HasKitchenWith(kitchen.ID(kitchenID))).
		WithIngredient().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// UpdateInventory updates the quantity of an inventory item
func (h *InventoryHandler) UpdateInventory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input struct {
		Quantity       float64    `json:"quantity" binding:"required,gte=0"`
		ExpirationDate *time.Time `json:"expiration_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update inventory
	update := h.client.InventoryItem.
		UpdateOneID(id).
		SetQuantity(input.Quantity)

	if input.ExpirationDate != nil {
		update.SetExpirationDate(*input.ExpirationDate)
	}

	inv, err := update.Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inv)
}

// RemoveInventory removes an inventory item
func (h *InventoryHandler) RemoveInventory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.client.InventoryItem.
		DeleteOneID(id).
		Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item removed successfully"})
}