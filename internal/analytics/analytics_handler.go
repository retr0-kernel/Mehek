// backend/internal/analytics/analytics_handler.go
package analytics

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"project/ent"
	"project/internal/common"
)

// AnalyticsHandler handles analytics API requests
type AnalyticsHandler struct {
	client        *ent.Client
	celeryService *common.CeleryService
	redisClient   *redis.Client
}

// NewAnalyticsHandler creates a new analytics handler
func NewAnalyticsHandler(client *ent.Client, celeryService *common.CeleryService, redisClient *redis.Client) *AnalyticsHandler {
	return &AnalyticsHandler{
		client:        client,
		celeryService: celeryService,
		redisClient:   redisClient,
	}
}

// GenerateKitchenPerformanceReport generates a performance report for a kitchen
func (h *AnalyticsHandler) GenerateKitchenPerformanceReport(c *gin.Context) {
	kitchenID, err := strconv.Atoi(c.Param("kitchen_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kitchen ID"})
		return
	}

	// Parse date range query parameters
	startDateStr := c.DefaultQuery("start_date", "")
	endDateStr := c.DefaultQuery("end_date", "")

	var startDate, endDate time.Time
	if startDateStr != "" {
		startDate, err = time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use RFC3339."})
			return
		}
	} else {
		startDate = time.Now().AddDate(0, 0, -30) // Last 30 days
	}

	if endDateStr != "" {
		endDate, err = time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use RFC3339."})
			return
		}
	} else {
		endDate = time.Now()
	}

	// Generate a unique task ID
	taskID := strconv.Itoa(kitchenID) + "-" + time.Now().Format("20060102150405")

	// Check if we have a cached report
	cacheKey := "analytics:kitchen:" + strconv.Itoa(kitchenID) + ":" + startDate.Format("20060102") + ":" + endDate.Format("20060102")
	cachedReport, err := h.redisClient.Get(context.Background(), cacheKey).Result()

	if err == nil && cachedReport != "" {
		// We have a cached report
		var report map[string]interface{}
		if err := json.Unmarshal([]byte(cachedReport), &report); err == nil {
			c.JSON(http.StatusOK, report)
			return
		}
	}

	// Send task to Celery
	err = h.celeryService.AnalyzeKitchenPerformance(c.Request.Context(), kitchenID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule report generation"})
		return
	}

	// In a real-world scenario, we'd have a callback mechanism or polling
	// For now, we'll just return a response that the task is in progress
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Report generation in progress",
		"task_id": taskID,
		"kitchen_id": kitchenID,
		"parameters": gin.H{
			"start_date": startDate.Format(time.RFC3339),
			"end_date":   endDate.Format(time.RFC3339),
		},
	})
}

// GetInventoryReport generates an inventory report for a kitchen
func (h *AnalyticsHandler) GetInventoryReport(c *gin.Context) {
	kitchenID, err := strconv.Atoi(c.Param("kitchen_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kitchen ID"})
		return
	}

	// Check if we have a cached report from today
	cacheKey := "inventory:kitchen:" + strconv.Itoa(kitchenID) + ":" + time.Now().Format("20060102")
	cachedReport, err := h.redisClient.Get(context.Background(), cacheKey).Result()

	if err == nil && cachedReport != "" {
		// We have a cached report
		var report map[string]interface{}
		if err := json.Unmarshal([]byte(cachedReport), &report); err == nil {
			c.JSON(http.StatusOK, report)
			return
		}
	}

	// Send task to Celery
	err = h.celeryService.GenerateInventoryReport(c.Request.Context(), kitchenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule inventory report generation"})
		return
	}

	// Generate a unique task ID
	taskID := "inventory-" + strconv.Itoa(kitchenID) + "-" + time.Now().Format("20060102150405")

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Inventory report generation in progress",
		"task_id": taskID,
		"kitchen_id": kitchenID,
	})
}

// PredictIngredientDemand predicts ingredient demand for a kitchen
func (h *AnalyticsHandler) PredictIngredientDemand(c *gin.Context) {
	kitchenID, err := strconv.Atoi(c.Param("kitchen_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kitchen ID"})
		return
	}

	// Send task to Celery
	err = h.celeryService.SendTask(
		c.Request.Context(), 
		"analytics.predict_ingredient_demand", 
		[]interface{}{kitchenID}, 
		nil,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule demand prediction"})
		return
	}

	// Generate a unique task ID
	taskID := "predict-" + strconv.Itoa(kitchenID) + "-" + time.Now().Format("20060102150405")

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Ingredient demand prediction in progress",
		"task_id": taskID,
		"kitchen_id": kitchenID,
	})
}