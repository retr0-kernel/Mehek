// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BrandsColumns holds the columns for the "brands" table.
	BrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "cuisine_type", Type: field.TypeString},
		{Name: "logo_url", Type: field.TypeString},
		{Name: "kitchen_brands", Type: field.TypeInt},
	}
	// BrandsTable holds the schema information for the "brands" table.
	BrandsTable = &schema.Table{
		Name:       "brands",
		Columns:    BrandsColumns,
		PrimaryKey: []*schema.Column{BrandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "brands_kitchens_brands",
				Columns:    []*schema.Column{BrandsColumns[4]},
				RefColumns: []*schema.Column{KitchensColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// IngredientsColumns holds the columns for the "ingredients" table.
	IngredientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "unit", Type: field.TypeString},
		{Name: "cost_per_unit", Type: field.TypeFloat64},
	}
	// IngredientsTable holds the schema information for the "ingredients" table.
	IngredientsTable = &schema.Table{
		Name:       "ingredients",
		Columns:    IngredientsColumns,
		PrimaryKey: []*schema.Column{IngredientsColumns[0]},
	}
	// InventoryItemsColumns holds the columns for the "inventory_items" table.
	InventoryItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeFloat64},
		{Name: "expiration_date", Type: field.TypeTime, Nullable: true},
		{Name: "inventory_item_ingredient", Type: field.TypeInt},
		{Name: "kitchen_inventory", Type: field.TypeInt},
	}
	// InventoryItemsTable holds the schema information for the "inventory_items" table.
	InventoryItemsTable = &schema.Table{
		Name:       "inventory_items",
		Columns:    InventoryItemsColumns,
		PrimaryKey: []*schema.Column{InventoryItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "inventory_items_ingredients_ingredient",
				Columns:    []*schema.Column{InventoryItemsColumns[3]},
				RefColumns: []*schema.Column{IngredientsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "inventory_items_kitchens_inventory",
				Columns:    []*schema.Column{InventoryItemsColumns[4]},
				RefColumns: []*schema.Column{KitchensColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// KitchensColumns holds the columns for the "kitchens" table.
	KitchensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "location", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeInt},
		{Name: "operating_hours", Type: field.TypeJSON},
	}
	// KitchensTable holds the schema information for the "kitchens" table.
	KitchensTable = &schema.Table{
		Name:       "kitchens",
		Columns:    KitchensColumns,
		PrimaryKey: []*schema.Column{KitchensColumns[0]},
	}
	// KitchenResourcesColumns holds the columns for the "kitchen_resources" table.
	KitchenResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeInt},
		{Name: "available", Type: field.TypeBool, Default: true},
		{Name: "kitchen_resources", Type: field.TypeInt},
	}
	// KitchenResourcesTable holds the schema information for the "kitchen_resources" table.
	KitchenResourcesTable = &schema.Table{
		Name:       "kitchen_resources",
		Columns:    KitchenResourcesColumns,
		PrimaryKey: []*schema.Column{KitchenResourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "kitchen_resources_kitchens_resources",
				Columns:    []*schema.Column{KitchenResourcesColumns[5]},
				RefColumns: []*schema.Column{KitchensColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "brand_menus", Type: field.TypeInt},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menus_brands_menus",
				Columns:    []*schema.Column{MenusColumns[3]},
				RefColumns: []*schema.Column{BrandsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MenuItemsColumns holds the columns for the "menu_items" table.
	MenuItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "prep_time", Type: field.TypeInt},
		{Name: "equipment_needed", Type: field.TypeString},
		{Name: "menu_items", Type: field.TypeInt},
	}
	// MenuItemsTable holds the schema information for the "menu_items" table.
	MenuItemsTable = &schema.Table{
		Name:       "menu_items",
		Columns:    MenuItemsColumns,
		PrimaryKey: []*schema.Column{MenuItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menu_items_menus_items",
				Columns:    []*schema.Column{MenuItemsColumns[5]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "required_by", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString, Default: "pending"},
		{Name: "total_price", Type: field.TypeFloat64, Default: 0},
		{Name: "brand_orders", Type: field.TypeInt},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_brands_orders",
				Columns:    []*schema.Column{OrdersColumns[5]},
				RefColumns: []*schema.Column{BrandsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrderItemsColumns holds the columns for the "order_items" table.
	OrderItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "special_instructions", Type: field.TypeString, Nullable: true},
		{Name: "order_items", Type: field.TypeInt},
		{Name: "order_item_menu_item", Type: field.TypeInt},
	}
	// OrderItemsTable holds the schema information for the "order_items" table.
	OrderItemsTable = &schema.Table{
		Name:       "order_items",
		Columns:    OrderItemsColumns,
		PrimaryKey: []*schema.Column{OrderItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_items_orders_items",
				Columns:    []*schema.Column{OrderItemsColumns[3]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_items_menu_items_menu_item",
				Columns:    []*schema.Column{OrderItemsColumns[4]},
				RefColumns: []*schema.Column{MenuItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ResourceAllocationsColumns holds the columns for the "resource_allocations" table.
	ResourceAllocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString, Default: "scheduled"},
		{Name: "kitchen_resource_allocations", Type: field.TypeInt},
		{Name: "resource_allocation_order", Type: field.TypeInt},
		{Name: "shift_allocations", Type: field.TypeInt},
	}
	// ResourceAllocationsTable holds the schema information for the "resource_allocations" table.
	ResourceAllocationsTable = &schema.Table{
		Name:       "resource_allocations",
		Columns:    ResourceAllocationsColumns,
		PrimaryKey: []*schema.Column{ResourceAllocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "resource_allocations_kitchen_resources_allocations",
				Columns:    []*schema.Column{ResourceAllocationsColumns[4]},
				RefColumns: []*schema.Column{KitchenResourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "resource_allocations_orders_order",
				Columns:    []*schema.Column{ResourceAllocationsColumns[5]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "resource_allocations_shifts_allocations",
				Columns:    []*schema.Column{ResourceAllocationsColumns[6]},
				RefColumns: []*schema.Column{ShiftsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ShiftsColumns holds the columns for the "shifts" table.
	ShiftsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "staff_shifts", Type: field.TypeInt},
	}
	// ShiftsTable holds the schema information for the "shifts" table.
	ShiftsTable = &schema.Table{
		Name:       "shifts",
		Columns:    ShiftsColumns,
		PrimaryKey: []*schema.Column{ShiftsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shifts_staffs_shifts",
				Columns:    []*schema.Column{ShiftsColumns[3]},
				RefColumns: []*schema.Column{StaffsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StaffsColumns holds the columns for the "staffs" table.
	StaffsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
		{Name: "skills", Type: field.TypeString},
		{Name: "availability", Type: field.TypeJSON},
		{Name: "kitchen_staff", Type: field.TypeInt},
	}
	// StaffsTable holds the schema information for the "staffs" table.
	StaffsTable = &schema.Table{
		Name:       "staffs",
		Columns:    StaffsColumns,
		PrimaryKey: []*schema.Column{StaffsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "staffs_kitchens_staff",
				Columns:    []*schema.Column{StaffsColumns[5]},
				RefColumns: []*schema.Column{KitchensColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "roles", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// MenuItemIngredientsColumns holds the columns for the "menu_item_ingredients" table.
	MenuItemIngredientsColumns = []*schema.Column{
		{Name: "menu_item_id", Type: field.TypeInt},
		{Name: "ingredient_id", Type: field.TypeInt},
	}
	// MenuItemIngredientsTable holds the schema information for the "menu_item_ingredients" table.
	MenuItemIngredientsTable = &schema.Table{
		Name:       "menu_item_ingredients",
		Columns:    MenuItemIngredientsColumns,
		PrimaryKey: []*schema.Column{MenuItemIngredientsColumns[0], MenuItemIngredientsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menu_item_ingredients_menu_item_id",
				Columns:    []*schema.Column{MenuItemIngredientsColumns[0]},
				RefColumns: []*schema.Column{MenuItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "menu_item_ingredients_ingredient_id",
				Columns:    []*schema.Column{MenuItemIngredientsColumns[1]},
				RefColumns: []*schema.Column{IngredientsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BrandsTable,
		IngredientsTable,
		InventoryItemsTable,
		KitchensTable,
		KitchenResourcesTable,
		MenusTable,
		MenuItemsTable,
		OrdersTable,
		OrderItemsTable,
		ResourceAllocationsTable,
		ShiftsTable,
		StaffsTable,
		UsersTable,
		MenuItemIngredientsTable,
	}
)

func init() {
	BrandsTable.ForeignKeys[0].RefTable = KitchensTable
	InventoryItemsTable.ForeignKeys[0].RefTable = IngredientsTable
	InventoryItemsTable.ForeignKeys[1].RefTable = KitchensTable
	KitchenResourcesTable.ForeignKeys[0].RefTable = KitchensTable
	MenusTable.ForeignKeys[0].RefTable = BrandsTable
	MenuItemsTable.ForeignKeys[0].RefTable = MenusTable
	OrdersTable.ForeignKeys[0].RefTable = BrandsTable
	OrderItemsTable.ForeignKeys[0].RefTable = OrdersTable
	OrderItemsTable.ForeignKeys[1].RefTable = MenuItemsTable
	ResourceAllocationsTable.ForeignKeys[0].RefTable = KitchenResourcesTable
	ResourceAllocationsTable.ForeignKeys[1].RefTable = OrdersTable
	ResourceAllocationsTable.ForeignKeys[2].RefTable = ShiftsTable
	ShiftsTable.ForeignKeys[0].RefTable = StaffsTable
	StaffsTable.ForeignKeys[0].RefTable = KitchensTable
	MenuItemIngredientsTable.ForeignKeys[0].RefTable = MenuItemsTable
	MenuItemIngredientsTable.ForeignKeys[1].RefTable = IngredientsTable
}
