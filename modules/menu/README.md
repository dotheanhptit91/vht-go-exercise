# Menu Module

A complete implementation of the Menu module following VHT-GO Clean Architecture principles with CQRS pattern and RPC integration.

## Overview

The Menu module manages restaurant menus that contain lists of food items. It uses RPC requests to fetch related restaurant and food data from their respective services.

## Architecture

```
modules/menu/
├── domain/                    # Core business entities
│   ├── model.go              # Menu domain model with JSON food_ids
│   └── error.go              # Domain-specific errors
├── dtos/                     # Data Transfer Objects
│   ├── create_menu.dto.go
│   ├── get_menu.dto.go
│   ├── update_menu.dto.go
│   ├── list_menu.dto.go
│   └── delete_menu.dto.go
├── infras/                   # Infrastructure layer
│   ├── repository/           # Data access
│   │   ├── repo.go
│   │   ├── insert.go
│   │   ├── find.go
│   │   ├── update.go
│   │   ├── delete.go
│   │   └── menurpcclient/    # RPC clients for external services
│   │       ├── rpc_client.go
│   │       ├── get_food.rpc.go
│   │       ├── get_foods.rpc.go
│   │       ├── get_restaurant.rpc.go
│   │       └── get_restaurants.rpc.go
│   └── controller/           # HTTP handlers
│       ├── controller.go
│       ├── create_menu_api.go
│       ├── get_menu_api.go
│       ├── list_menu_api.go
│       ├── update_menu_api.go
│       └── delete_menu_api.go
├── service/                  # Business logic (CQRS handlers)
│   ├── create_menu.svc.go
│   ├── get_menu.svc.go
│   ├── list_menu.svc.go
│   ├── update_menu.svc.go
│   └── delete_menu.svc.go
├── module.go                 # Module setup & DI
├── schema.sql               # Database schema
└── README.md                # This file
```

## Database Schema

The menu table uses a **minimal design** with JSON column for food_ids:

```sql
CREATE TABLE IF NOT EXISTS `menus` (
    `id` VARCHAR(36) PRIMARY KEY COMMENT 'UUID v7',
    `restaurant_id` INT NOT NULL COMMENT 'Reference to restaurant',
    `name` VARCHAR(255) NOT NULL COMMENT 'Menu name',
    `description` TEXT NULL COMMENT 'Menu description',
    `food_ids` JSON NOT NULL COMMENT 'Array of food IDs in this menu',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '1=active, 0=inactive',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX `idx_restaurant_id` (`restaurant_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### Design Benefits

- **Minimal Fields**: Only essential columns (7 total)
- **JSON Array**: Flexible food_ids storage without junction table
- **UUID Primary Key**: Time-ordered UUIDs (v7) for better indexing
- **Soft Delete**: Status field for logical deletion
- **Performance**: Indexed foreign keys and status

### Example Data

```json
{
  "id": "01933a7f-1234-7890-abcd-1234567890ab",
  "restaurant_id": 1,
  "name": "Lunch Special Menu",
  "description": "Our best lunch combinations",
  "food_ids": [1, 3, 5, 7],
  "status": 1,
  "created_at": "2025-11-26T10:00:00Z",
  "updated_at": "2025-11-26T10:00:00Z"
}
```

## RPC Integration

The menu module integrates with external services via RPC:

### Food Service RPC
- `FindFoodById(ctx, id)` - Get single food details
- `FindFoodsByIds(ctx, ids)` - Get multiple foods by IDs (batch)

### Restaurant Service RPC
- `FindRestaurantById(ctx, id)` - Get single restaurant details
- `FindRestaurantsByIds(ctx, ids)` - Get multiple restaurants by IDs (batch)

### RPC Response Population

The service layer automatically populates:
- **Get Menu**: Fetches all foods and restaurant data
- **List Menus**: Batch fetches all related foods and restaurants for efficiency

## API Endpoints

### 1. Create Menu
```
POST /v1/menus
```

**Request Body:**
```json
{
  "restaurant_id": 1,
  "name": "Lunch Menu",
  "description": "Special lunch offerings",
  "food_ids": [1, 2, 3, 5]
}
```

**Response (201 Created):**
```json
{
  "data": "01933a7f-1234-7890-abcd-1234567890ab"
}
```

### 2. Get Menu by ID
```
GET /v1/menus/:id
```

**Response (200 OK):**
```json
{
  "data": {
    "id": "01933a7f-1234-7890-abcd-1234567890ab",
    "restaurant_id": 1,
    "name": "Lunch Menu",
    "description": "Special lunch offerings",
    "food_ids": [1, 2, 3],
    "status": 1,
    "created_at": "2025-11-26T10:00:00Z",
    "updated_at": "2025-11-26T10:00:00Z",
    "restaurant": {
      "id": 1,
      "name": "The Best Restaurant"
    },
    "foods": [
      {
        "id": 1,
        "name": "Grilled Chicken",
        "description": "Juicy grilled chicken",
        "price": 15.99
      },
      {
        "id": 2,
        "name": "Caesar Salad",
        "price": 8.99
      }
    ]
  }
}
```

### 3. List Menus
```
GET /v1/menus?restaurant_id=1&status=1&limit=10&page=1
```

**Query Parameters:**
- `restaurant_id` (optional): Filter by restaurant
- `status` (optional): Filter by status (1=active, 0=inactive)
- `limit` (optional): Items per page
- `page` (optional): Page number

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": "01933a7f-1234-7890-abcd-1234567890ab",
      "restaurant_id": 1,
      "name": "Lunch Menu",
      "food_ids": [1, 2],
      "restaurant": { "id": 1, "name": "Restaurant Name" },
      "foods": [...]
    }
  ],
  "paging": {
    "page": 1,
    "limit": 10,
    "total": 25
  }
}
```

### 4. Update Menu
```
PATCH /v1/menus/:id
```

**Request Body (all fields optional):**
```json
{
  "name": "Updated Lunch Menu",
  "description": "New description",
  "food_ids": [1, 2, 3, 4, 5]
}
```

**Response (200 OK):**
```json
{
  "data": true
}
```

### 5. Delete Menu (Soft Delete)
```
DELETE /v1/menus/:id
```

**Response (200 OK):**
```json
{
  "data": true
}
```

## Module Registration

To register the menu module in your application, add the following to `main.go`:

```go
import (
    menumodule "vht-go/modules/menu"
    // ... other imports
)

func main() {
    // ... database and service context setup
    
    r := gin.Default()
    v1 := r.Group("/v1")
    
    // Register menu module
    menumodule.SetupMenuModule(v1, serviceContext)
    
    // ... other modules and server start
}
```

## Configuration

The module requires the following service URIs in AppConfig:

- `FoodServiceURI`: URI for the food RPC service
- `RestaurantServiceURI`: URI for the restaurant RPC service

**Default Values:**
```
food-service-uri: http://localhost:3600/v1/rpc/foods
restaurant-service-uri: http://localhost:3600/v1/rpc/restaurants
```

## Key Features

### 1. Clean Architecture
- Clear separation of concerns across layers
- Dependency inversion principle
- Interface-based design for testability

### 2. CQRS Pattern
- Separate Command and Query handlers
- Type-safe generic interfaces
- Single responsibility per handler

### 3. RPC Integration
- Automatic population of related entities
- Batch fetching for list operations (performance optimization)
- Error-tolerant (continues if RPC fails)

### 4. Efficient List Operations
- Collects all unique food IDs and restaurant IDs
- Makes single batch RPC calls instead of N+1 queries
- Maps results back to menus efficiently

### 5. JSON Array Support
- Custom `FoodIDs` type with `Scan` and `Value` methods
- Seamless JSON serialization/deserialization
- Database-agnostic JSON handling

### 6. Validation
- DTO-level validation with business rules
- Restaurant ID and Food IDs validation
- Sanitization (trimming, normalization)

## Testing

Example test for creating a menu:

```go
func TestCreateMenu(t *testing.T) {
    // Mock repository
    mockRepo := &MockMenuRepository{}
    handler := menuservice.NewCreateMenuResultCommandHandler(mockRepo)
    
    dto := &menudtos.CreateMenuDTO{
        RestaurantId: 1,
        Name:         "Test Menu",
        FoodIds:      []int{1, 2, 3},
    }
    
    id, err := handler.Handle(context.Background(), &menuservice.CreateMenuResultCommand{DTO: dto})
    
    assert.NoError(t, err)
    assert.NotNil(t, id)
}
```

## Performance Considerations

1. **Batch RPC Calls**: List operation fetches all foods/restaurants in 2 calls instead of N calls
2. **Indexed Queries**: Database indexes on `restaurant_id` and `status`
3. **JSON Column**: Fast array storage without junction table overhead
4. **UUID v7**: Time-ordered UUIDs improve database indexing performance

## Error Handling

Domain errors are defined in `domain/error.go`:

```go
const (
    ErrMenuNotFound        = "menu not found"
    ErrMenuNameRequired    = "menu name is required"
    ErrInvalidRestaurantId = "invalid restaurant id"
    ErrInvalidFoodIds      = "invalid food ids"
    ErrFoodIdsRequired     = "at least one food id is required"
)
```

## Future Enhancements

- Add menu ordering/priority field
- Support menu categories/sections
- Add availability scheduling (time-based menus)
- Implement caching for frequently accessed menus
- Add menu cloning functionality
- Support menu versioning

---

**Created**: 2025-11-26  
**Architecture Version**: 2.0 (CQRS Pattern)  
**Reference Module**: `modules/category`

