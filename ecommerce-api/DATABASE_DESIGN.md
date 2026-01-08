# Database Schema Design

## Tables

### Users
Stores user account information.
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_shop_owner BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
```

**Fields:**
- `id` - Auto-incrementing primary key
- `email` - Unique identifier for login
- `password` - Hashed password (never store plain text!)
- `is_shop_owner` - Role flag (true = admin)

---

### Products
Stores product catalog.
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    img_url TEXT,
    stock INTEGER NOT NULL DEFAULT 0 CHECK (stock >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_products_created_at ON products(created_at DESC);
```

**Fields:**
- `price` - Using DECIMAL for exact money calculations
- `stock` - Inventory count (added for cart/orders later)
- Constraints ensure price and stock can't be negative

---

## Relationships
```
users (1) ──── (many) orders
products (1) ──── (many) order_items
orders (1) ──── (many) order_items
```

(Cart and Orders tables will be added in later parts)

---

## Indexes Strategy

- `users.email` - Frequent lookups during login
- `products.price` - Common for filtering/sorting
- `products.created_at` - For showing newest products

---

## Migration Strategy

We use `sql-migrate` with up/down migrations:
- **Up** migrations apply changes
- **Down** migrations revert changes

This allows us to safely evolve the database schema.