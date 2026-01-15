-- +migrate Down

-- Remove stock column from products
ALTER TABLE products DROP COLUMN IF EXISTS stock;

-- Drop indexes
DROP INDEX IF EXISTS idx_cart_items_product_id;
DROP INDEX IF EXISTS idx_cart_items_cart_id;
DROP INDEX IF EXISTS idx_carts_user_id;

-- Drop tables
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;