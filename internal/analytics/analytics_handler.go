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

// GetResourceUtilization retrieves resource utilization metrics for a kitchen
func (h *AnalyticsHandler) GetResourceUtilization(c *gin.Context) {
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

	// Check cache
	cacheKey := "resource:utilization:" + strconv.Itoa(kitchenID) + ":" + startDate.Format("20060102") + ":" + endDate.Format("20060102")
	cachedData, err := h.redisClient.Get(context.Background(), cacheKey).Result()

	if err == nil && cachedData != "" {
		// We have cached data
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &data); err == nil {
			c.JSON(http.StatusOK, data)
			return
		}
	}

	// Send task to Celery
	err = h.celeryService.SendTask(
		c.Request.Context(),
		"analytics.get_resource_utilization",
		[]interface{}{kitchenID},
		map[string]interface{}{
			"start_date": startDate.Format(time.RFC3339),
			"end_date":   endDate.Format(time.RFC3339),
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve resource utilization"})
		return
	}

	// For now, just return sample data
	sampleData := gin.H{
		"kitchen_id": kitchenID,
		"period": gin.H{
			"start_date": startDate.Format(time.RFC3339),
			"end_date":   endDate.Format(time.RFC3339),
		},
		"utilization": gin.H{
			"overall_percentage": 68.5,
			"by_resource_type": map[string]float64{
				"stove":      75.3,
				"oven":       62.1,
				"grill":      81.7,
				"prep_table": 55.2,
				"fryer":      68.9,
			},
			"by_time_of_day": map[string]float64{
				"morning":   45.2,
				"lunch":     92.6,
				"afternoon": 38.4,
				"dinner":    87.3,
				"late":      32.1,
			},
		},
		"bottlenecks": []gin.H{
			{
				"resource_type": "grill",
				"time_of_day":   "lunch",
				"utilization":   98.7,
				"duration_minutes": 120,
			},
			{
				"resource_type": "oven",
				"time_of_day":   "dinner",
				"utilization":   95.2,
				"duration_minutes": 90,
			},
		},
		"recommendations": []string{
			"Consider adding additional grill capacity during lunch hours",
			"Schedule food prep outside of peak hours to optimize prep table usage",
			"Reorganize oven usage during dinner rush to improve efficiency",
		},
	}

	// Cache this sample data for future requests
	jsonData, _ := json.Marshal(sampleData)
	h.redisClient.Set(context.Background(), cacheKey, jsonData, 1*time.Hour)

	c.JSON(http.StatusOK, sampleData)
}

// GetRevenueAnalysis retrieves revenue analysis for a kitchen
func (h *AnalyticsHandler) GetRevenueAnalysis(c *gin.Context) {
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

	// Check cache
	cacheKey := "revenue:analysis:" + strconv.Itoa(kitchenID) + ":" + startDate.Format("20060102") + ":" + endDate.Format("20060102")
	cachedData, err := h.redisClient.Get(context.Background(), cacheKey).Result()

	if err == nil && cachedData != "" {
		// We have cached data
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &data); err == nil {
			c.JSON(http.StatusOK, data)
			return
		}
	}

	// Send task to Celery
	err = h.celeryService.SendTask(
		c.Request.Context(),
		"analytics.get_revenue_analysis",
		[]interface{}{kitchenID},
		map[string]interface{}{
			"start_date": startDate.Format(time.RFC3339),
			"end_date":   endDate.Format(time.RFC3339),
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve revenue analysis"})
		return
	}

	// For now, just return sample data
	sampleData := gin.H{
		"kitchen_id": kitchenID,
		"period": gin.H{
			"start_date": startDate.Format(time.RFC3339),
			"end_date":   endDate.Format(time.RFC3339),
		},
		"summary": gin.H{
			"total_revenue": 42850.75,
			"average_daily_revenue": 1428.36,
			"growth_from_previous_period": 8.5,
		},
		"by_brand": []gin.H{
			{
				"brand_id": 1,
				"brand_name": "Spicy Dragon",
				"cuisine_type": "Asian Fusion",
				"revenue": 15320.50,
				"percentage": 35.8,
				"growth": 12.3,
			},
			{
				"brand_id": 2,
				"brand_name": "Burger Baron",
				"cuisine_type": "American",
				"revenue": 12450.25,
				"percentage": 29.1,
				"growth": 5.7,
			},
			{
				"brand_id": 3,
				"brand_name": "Pizza Paradise",
				"cuisine_type": "Italian",
				"revenue": 9875.30,
				"percentage": 23.0,
				"growth": 7.2,
			},
			{
				"brand_id": 4,
				"brand_name": "Taco Time",
				"cuisine_type": "Mexican",
				"revenue": 5204.70,
				"percentage": 12.1,
				"growth": 9.8,
			},
		},
		"by_day_of_week": map[string]float64{
			"Monday":    4285.10,
			"Tuesday":   4712.55,
			"Wednesday": 5142.10,
			"Thursday":  5785.25,
			"Friday":    8570.15,
			"Saturday":  8998.45,
			"Sunday":    5357.15,
		},
		"by_time_of_day": map[string]float64{
			"morning":   3428.05,
			"lunch":     12855.25,
			"afternoon": 4285.10,
			"dinner":    17140.30,
			"late":      5142.05,
		},
		"top_menu_items": []gin.H{
			{
				"item_id": 42,
				"name": "Spicy Dragon Roll",
				"revenue": 4285.10,
				"quantity_sold": 428,
			},
			{
				"item_id": 17,
				"name": "Double Bacon Burger",
				"revenue": 3642.35,
				"quantity_sold": 386,
			},
			{
				"item_id": 29,
				"name": "Supreme Pizza",
				"revenue": 3214.05,
				"quantity_sold": 321,
			},
		},
	}

	// Cache this sample data for future requests
	jsonData, _ := json.Marshal(sampleData)
	h.redisClient.Set(context.Background(), cacheKey, jsonData, 1*time.Hour)

	c.JSON(http.StatusOK, sampleData)
}