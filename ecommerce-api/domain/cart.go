package domain

import "time"

type Cart struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CartItem struct {
	ID        int       `json:"id" db:"id"`
	CartID    int       `json:"cart_id" db:"cart_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// FOR API RESPONSE

type CartItemDetail struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"` // price * quantity
}

type CartDetails struct {
	CartID int              `json:"cart_id"`
	UserID int              `json:"user_id"`
	Items  []CartItemDetail `json:"items"`
	Total  float64          `json:"total"`
}
