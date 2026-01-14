# Architecture Design

## Overview

This project follows **Clean Architecture** principles with **Domain-Driven Design** patterns.

## Folder Structure
```
ecommerce-backend/
├── cmd/                    # Application entry points
│   └── serve.go            # Server initialization
├── config/                 # Configuration management
│   └── config.go           # Environment variable loader
├── domain/                 # Business entities (User, Product)
│   ├── user.go
│   ├── product.go
│   └── errors.go          # Domain-specific errors
├── infra/                 # Infrastructure layer
│   ├── db/                # Database connection
│   └── logger/            # Logging infrastructure
├── product/               # Product domain logic
│   ├── service.go         # Business logic
│   └── port.go            # Interfaces
├── user/                  # User domain logic
│   ├── service.go
│   └── port.go
├── repo/                  # Data access layer
│   ├── product.go         # Product repository
│   └── user.go            # User repository
├── rest/                  # HTTP layer
│   ├── handlers/          # HTTP handlers
│   ├── middleware/        # HTTP middleware
│   └── server.go          # HTTP server
├── util/                  # Shared utilities
└── migrations/            # Database migrations
```

## Layers

### 1. Domain Layer (`domain/`)
- **Purpose:** Business entities and rules
- **Dependencies:** None (pure Go structs)
- **Example:** `User`, `Product` structs

### 2. Service Layer (`product/`, `user/`)
- **Purpose:** Business logic
- **Dependencies:** Domain layer
- **Example:** Validate product price, create user

### 3. Repository Layer (`repo/`)
- **Purpose:** Data access
- **Dependencies:** Domain layer
- **Example:** Save product to database

### 4. Infrastructure Layer (`infra/`)
- **Purpose:** External services (database, logging)
- **Dependencies:** Config
- **Example:** PostgreSQL connection

### 5. Presentation Layer (`rest/`)
- **Purpose:** HTTP API
- **Dependencies:** Service layer
- **Example:** Handle POST /products

## Dependency Flow
```
HTTP Request
    ↓
Handler (rest/)
    ↓
Service (product/, user/)
    ↓
Repository (repo/)
    ↓
Database
```

## Key Principles

1. **Dependency Inversion:** Higher layers don't depend on lower layers
2. **Interface-based Design:** Use interfaces for testability
3. **Single Responsibility:** Each layer has one job
4. **Separation of Concerns:** HTTP logic ≠ business logic ≠ data access

## Why This Architecture?

**Testable:** Mock interfaces easily  
**Maintainable:** Clear separation of concerns  
**Scalable:** Add features without breaking existing code  
**Flexible:** Swap implementations (e.g., change database)

## Example Flow: Create Product
```
1. POST /products
   ↓
2. ProductHandler.CreateProduct() [rest/handlers/product/]
   - Parse JSON
   - Validate input
   ↓
3. ProductService.Create() [product/]
   - Business validation (price > 0)
   - Call repository
   ↓
4. ProductRepo.Create() [repo/]
   - SQL INSERT
   - Return created product
   ↓
5. Response back through layers
   ↓
6. JSON response to client
```