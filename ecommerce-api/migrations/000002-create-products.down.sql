-- +migrate Down

DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_products_price;
DROP TABLE IF EXISTS products;