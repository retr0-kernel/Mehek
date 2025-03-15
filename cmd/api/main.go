//go:generate go run ./ent_generate.go

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	
	"project/ent"
	"project/internal/common"
	"project/internal/order"
	"github.com/go-redis/redis/v8"
	
	"project/internal/analytics"
	"project/internal/auth"
	"project/internal/inventory"
	"project/internal/resource"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to PostgreSQL
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgresql://postgres:postgres@localhost:5432/ghost_kitchen?sslmode=disable"
	}
	
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run database migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize Celery service
	celeryService, err := common.NewCeleryService()
	if err != nil {
		log.Fatalf("failed to initialize Celery service: %v", err)
	}
	defer celeryService.Close()

	// Initialize resource manager
	resourceManager := order.NewResourceManager(celeryService)

	// Initialize Gin router
	r := gin.Default()

	// Add middleware to inject services
	r.Use(func(c *gin.Context) {
		c.Set("resourceManager", resourceManager)
		c.Next()
	})

	// Add health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	// Add routes for API endpoints
	setupRoutes(r, client, celeryService)

	// Configure graceful shutdown
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		
		log.Printf("Server starting on port %s...", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("Shutting down server...")

	// Give outstanding requests 5 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

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