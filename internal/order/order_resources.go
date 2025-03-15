// backend/internal/order/order_resource.go
package order

import (
	"context"
	"fmt"
	"log"

	"project/internal/common"
)

// ResourceManager handles resource allocation for orders
type ResourceManager struct {
	celeryService *common.CeleryService
}

// NewResourceManager creates a new resource manager
func NewResourceManager(celeryService *common.CeleryService) *ResourceManager {
	return &ResourceManager{
		celeryService: celeryService,
	}
}

// AllocateResourcesForOrder schedules resource allocation for an order
func (m *ResourceManager) AllocateResourcesForOrder(ctx context.Context, orderID int) error {
	log.Printf("Scheduling resource allocation for order %d", orderID)
	
	err := m.celeryService.AllocateResourcesForOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to schedule resource allocation: %w", err)
	}
	
	return nil
}

// OptimizeKitchenResources schedules resource optimization for a kitchen
func (m *ResourceManager) OptimizeKitchenResources(ctx context.Context, kitchenID int) error {
	log.Printf("Scheduling resource optimization for kitchen %d", kitchenID)
	
	err := m.celeryService.OptimizeKitchenResources(ctx, kitchenID)
	if err != nil {
		return fmt.Errorf("failed to schedule resource optimization: %w", err)
	}
	
	return nil
}