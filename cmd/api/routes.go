// backend/cmd/api/routes.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	
	"project/ent"
	"project/internal/analytics"
	"project/internal/auth"
	"project/internal/common"
	"project/internal/inventory"
	"project/internal/order"
	"project/internal/resource"
)

func setupRoutes(r *gin.Engine, client *ent.Client, celeryService *common.CeleryService) {
	// Initialize Redis client for caching
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	
	// Initialize auth service
	authService := auth.NewService()
	authMiddleware := auth.NewMiddleware(authService)
	
	// Initialize resource manager
	resourceManager := order.NewResourceManager(celeryService)
	
	// Public routes
	v1 := r.Group("/api/v1")
	
	// Health check
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})
	
	// Auth routes
	userHandler := auth.NewUserHandler(client, authService)
	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/register", userHandler.RegisterUser)
		authRoutes.POST("/login", userHandler.LoginUser)
	}
	
	// Protected routes
	protected := v1.Group("")
	protected.Use(authMiddleware.AuthRequired())
	
	// Middleware to inject resource manager
	protected.Use(func(c *gin.Context) {
		c.Set("resourceManager", resourceManager)
		c.Next()
	})
	
	// Kitchen routes
	kitchenHandler := resource.NewKitchenHandler(client)
	kitchens := protected.Group("/kitchens")
	{
		kitchens.POST("", authMiddleware.RoleRequired("admin"), kitchenHandler.CreateKitchen)
		kitchens.GET("", kitchenHandler.GetKitchens)
		kitchens.GET("/:id", kitchenHandler.GetKitchenByID)
		kitchens.PUT("/:id", authMiddleware.RoleRequired("admin"), kitchenHandler.UpdateKitchen)
		kitchens.DELETE("/:id", authMiddleware.RoleRequired("admin"), kitchenHandler.DeleteKitchen)
	}
	
	// Resource routes
	resourceHandler := resource.NewResourceHandler(client)
	resources := protected.Group("/resources")
	{
		resources.POST("", authMiddleware.RoleRequired("admin"), resourceHandler.CreateResource)
		resources.GET("/kitchen/:kitchen_id", resourceHandler.GetResourcesByKitchen)
		resources.GET("/:id", resourceHandler.GetResourceByID)
		resources.PUT("/:id", authMiddleware.RoleRequired("admin"), resourceHandler.UpdateResource)
		resources.DELETE("/:id", authMiddleware.RoleRequired("admin"), resourceHandler.DeleteResource)
	}
	
	// Brand routes
	brandHandler := resource.NewBrandHandler(client)
	brands := protected.Group("/brands")
	{
		brands.POST("", authMiddleware.RoleRequired("admin"), brandHandler.CreateBrand)
		brands.GET("", brandHandler.GetBrands)
		brands.GET("/:id", brandHandler.GetBrandByID)
		brands.PUT("/:id", authMiddleware.RoleRequired("admin"), brandHandler.UpdateBrand)
		brands.DELETE("/:id", authMiddleware.RoleRequired("admin"), brandHandler.DeleteBrand)
	}
	
	// Order routes
	orderHandler := order.NewOrderHandler(client)
	orders := protected.Group("/orders")
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("", orderHandler.GetOrders)
		orders.GET("/:id", orderHandler.GetOrderByID)
		orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		orders.PUT("/:id/cancel", orderHandler.CancelOrder)
	}
	
	// Inventory routes
	inventoryHandler := inventory.NewInventoryHandler(client)
	inventories := protected.Group("/inventory")
	{
		// Ingredient management
		inventories.POST("/ingredients", authMiddleware.RoleRequired("admin"), inventoryHandler.CreateIngredient)
		inventories.GET("/ingredients", inventoryHandler.GetIngredients)
		
		// Inventory management
		inventories.POST("", authMiddleware.RoleRequired("admin"), inventoryHandler.AddInventory)
		inventories.GET("/kitchen/:kitchen_id", inventoryHandler.GetKitchenInventory)
		inventories.PUT("/:id", authMiddleware.RoleRequired("admin"), inventoryHandler.UpdateInventory)
		inventories.DELETE("/:id", authMiddleware.RoleRequired("admin"), inventoryHandler.RemoveInventory)
	}
	
	// Analytics routes
	analyticsHandler := analytics.NewAnalyticsHandler(client, celeryService, redisClient)
	analyticsRoutes := protected.Group("/analytics")
	{
		analyticsRoutes.GET("/kitchen/:kitchen_id/performance", analyticsHandler.GenerateKitchenPerformanceReport)
		analyticsRoutes.GET("/kitchen/:kitchen_id/inventory", analyticsHandler.GetInventoryReport)
		analyticsRoutes.GET("/kitchen/:kitchen_id/predict-demand", analyticsHandler.PredictIngredientDemand)
		analyticsRoutes.GET("/kitchen/:kitchen_id/resource-utilization", analyticsHandler.GetResourceUtilization)
		analyticsRoutes.GET("/kitchen/:kitchen_id/revenue", analyticsHandler.GetRevenueAnalysis)
	}
}