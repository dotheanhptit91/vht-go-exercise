-- Menu Table Schema
-- This table stores menus with a list of food items for each restaurant

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

-- Example of food_ids JSON format: [1, 2, 3, 5, 8]
-- This allows flexible menu composition without additional junction table

