# Menu Module Integration Guide

## Quick Start

### 1. Run the SQL Schema

Execute the SQL schema to create the menus table:

```bash
mysql -u your_user -p your_database < modules/menu/schema.sql
```

Or manually run:

```sql
CREATE TABLE IF NOT EXISTS `menus` (
    `id` VARCHAR(36) PRIMARY KEY COMMENT 'UUID v7',
    `restaurant_id` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NULL,
    `food_ids` JSON NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX `idx_restaurant_id` (`restaurant_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 2. Register Module in main.go

Add the menu module to your application:

```go
package main

import (
    menumodule "vht-go/modules/menu"
    // ... other imports
)

func main() {
    // ... your existing setup ...
    
    // Register modules
    v1 := r.Group("/v1")
    
    // Add menu module
    menumodule.SetupMenuModule(v1, serviceContext)
    
    // ... start server ...
}
```

### 3. Configure Service URIs (if needed)

The module uses these default URIs:
- Food Service: `http://localhost:3600/v1/rpc/foods`
- Restaurant Service: `http://localhost:3600/v1/rpc/restaurants`

To override, use command-line flags:

```bash
./your-app \
  -food-service-uri="http://your-food-service/v1/rpc/foods" \
  -restaurant-service-uri="http://your-restaurant-service/v1/rpc/restaurants"
```

## API Testing

### Create a Menu

```bash
curl -X POST http://localhost:3600/v1/menus \
  -H "Content-Type: application/json" \
  -d '{
    "restaurant_id": 1,
    "name": "Lunch Special Menu",
    "description": "Best lunch combinations",
    "food_ids": [1, 2, 3, 5]
  }'
```

**Expected Response:**
```json
{
  "data": "01933a7f-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
```

### Get Menu by ID

```bash
curl http://localhost:3600/v1/menus/01933a7f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Expected Response:**
```json
{
  "data": {
    "id": "01933a7f-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "restaurant_id": 1,
    "name": "Lunch Special Menu",
    "description": "Best lunch combinations",
    "food_ids": [1, 2, 3, 5],
    "status": 1,
    "restaurant": {
      "id": 1,
      "name": "The Restaurant"
    },
    "foods": [
      {
        "id": 1,
        "name": "Grilled Chicken",
        "price": 15.99
      },
      ...
    ]
  }
}
```

### List Menus

```bash
# All menus
curl http://localhost:3600/v1/menus

# Filter by restaurant
curl http://localhost:3600/v1/menus?restaurant_id=1

# With pagination
curl "http://localhost:3600/v1/menus?restaurant_id=1&limit=10&page=1"

# Active menus only
curl http://localhost:3600/v1/menus?status=1
```

### Update Menu

```bash
curl -X PATCH http://localhost:3600/v1/menus/01933a7f-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Lunch Menu",
    "food_ids": [1, 2, 3, 4, 5, 6]
  }'
```

### Delete Menu (Soft Delete)

```bash
curl -X DELETE http://localhost:3600/v1/menus/01933a7f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

## Module Structure Overview

```
modules/menu/
├── domain/                           # Business entities
│   ├── model.go                      # Menu, FoodIDs, MenuFood, MenuRestaurant
│   └── error.go                      # Error constants
├── dtos/                             # Input/Output DTOs
├── infras/
│   ├── repository/                   # Database operations
│   │   └── menurpcclient/            # RPC clients
│   └── controller/                   # HTTP handlers
├── service/                          # Business logic (CQRS)
└── module.go                         # Dependency injection
```

## Key Design Decisions

### 1. JSON Array for Food IDs

**Why?** Keeps the table minimal and flexible:
- No junction table needed
- Easy to add/remove food items
- Simple to query and update
- Performant for reasonable array sizes (<100 items)

**Custom Type:**
```go
type FoodIDs []int  // Implements sql.Scanner and driver.Valuer
```

### 2. RPC Integration

**Why?** Decoupled services:
- Menu service doesn't directly access food/restaurant databases
- Services can scale independently
- Clear service boundaries

**Batch Fetching:**
- List operations collect all IDs first
- Single RPC call per service (not N+1)
- Efficient even with many menus

### 3. Soft Delete

**Why?** Data preservation:
- Historical data maintained
- Can restore deleted menus
- Audit trail support

## Troubleshooting

### RPC Calls Failing

**Symptom:** Menu returns without restaurant/food data

**Solution:**
1. Check service URIs are correct
2. Verify food/restaurant services are running
3. Check network connectivity
4. Look for RPC endpoint availability

### JSON Parsing Errors

**Symptom:** Error saving/loading food_ids

**Solution:**
1. Ensure MySQL version supports JSON (5.7+)
2. Check food_ids is valid JSON array
3. Verify FoodIDs type is used in model

### Menu Not Found

**Symptom:** 404 when getting menu

**Solution:**
1. Verify UUID format is correct
2. Check menu exists and status=1
3. Ensure database connection is working

## Dependencies

Required packages:
- `github.com/gin-gonic/gin` - HTTP framework
- `github.com/google/uuid` - UUID generation
- `gorm.io/gorm` - ORM
- `resty.dev/v3` - HTTP client for RPC
- `github.com/viettranx/service-context` - Service context

## Performance Tips

1. **Use Indexes**: The schema includes necessary indexes
2. **Batch Operations**: List handler optimizes RPC calls
3. **Limit Array Size**: Keep food_ids under 100 items for best performance
4. **Pagination**: Always use pagination for list operations
5. **Caching**: Consider caching frequently accessed menus

## Next Steps

1. ✅ Run SQL schema
2. ✅ Register module in main.go
3. ✅ Test API endpoints
4. ⬜ Add authentication middleware (if needed)
5. ⬜ Implement caching layer (optional)
6. ⬜ Add monitoring and logging
7. ⬜ Write integration tests

---

For detailed API documentation, see [README.md](README.md)

