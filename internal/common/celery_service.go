// backend/internal/common/celery_service.go
package common

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"crypto/rand"
	"encoding/hex"

	"github.com/streadway/amqp"
)

// CeleryService handles integration with Celery tasks
type CeleryService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewCeleryService creates a new Celery service
func NewCeleryService() (*CeleryService, error) {
	// Get RabbitMQ URL from environment variables
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}

	// Connect to RabbitMQ
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare Celery queue
	_, err = channel.QueueDeclare(
		"celery", // queue name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	return &CeleryService{
		conn:    conn,
		channel: channel,
	}, nil
}

// Close closes the connection to RabbitMQ
func (s *CeleryService) Close() {
	if s.channel != nil {
		s.channel.Close()
	}
	if s.conn != nil {
		s.conn.Close()
	}
}

// generateTaskID generates a unique ID for a Celery task
func generateTaskID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		// If random generation fails, use timestamp as fallback
		return hex.EncodeToString([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	}
	return hex.EncodeToString(bytes)
}

// SendTask sends a task to Celery
func (s *CeleryService) SendTask(ctx context.Context, taskName string, args []interface{}, kwargs map[string]interface{}) error {
	// Create Celery task message
	task := map[string]interface{}{
		"id":      generateTaskID(),
		"task":    taskName,
		"args":    args,
		"kwargs":  kwargs,
		"retries": 0,
	}

	body, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("failed to marshal task: %w", err)
	}

	// Publish message to RabbitMQ
	err = s.channel.Publish(
		"",       // exchange
		"celery", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Sent task %s to Celery", taskName)
	return nil
}

// AllocateResourcesForOrder sends a task to allocate resources for an order
func (s *CeleryService) AllocateResourcesForOrder(ctx context.Context, orderID int) error {
	return s.SendTask(ctx, "resource_allocation.allocate_resources_for_order", []interface{}{orderID}, nil)
}

// OptimizeKitchenResources sends a task to optimize resources for a kitchen
func (s *CeleryService) OptimizeKitchenResources(ctx context.Context, kitchenID int) error {
	return s.SendTask(ctx, "resource_allocation.optimize_kitchen_resources", []interface{}{kitchenID}, nil)
}

// GenerateInventoryReport sends a task to generate an inventory report
func (s *CeleryService) GenerateInventoryReport(ctx context.Context, kitchenID int) error {
	return s.SendTask(ctx, "inventory.generate_inventory_report", []interface{}{kitchenID}, nil)
}

// AnalyzeKitchenPerformance sends a task to analyze kitchen performance
func (s *CeleryService) AnalyzeKitchenPerformance(ctx context.Context, kitchenID int, startDate, endDate time.Time) error {
	return s.SendTask(ctx, "analytics.analyze_kitchen_performance", []interface{}{kitchenID}, map[string]interface{}{
		"start_date": startDate.Format(time.RFC3339),
		"end_date":   endDate.Format(time.RFC3339),
	})
}