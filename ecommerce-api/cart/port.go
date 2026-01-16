package cart

import (
	"ecommerce/domain"
	cartHandler "ecommerce/rest/handlers/cart"
)

type Service interface {
	cartHandler.Service
}

type CartRepo interface {
	GetOrCreate(userID int) (*domain.Cart, error)
	AddItem(cartID, productID, quantity int) (*domain.CartItem, error)
	GetItems(cartID int) ([]domain.CartItemDetail, error)
	UpdateItemQuantity(itemID, quantity int) error
	RemoveItem(itemID int) error
	Clear(cartID int) error
	GetItem(itemID int) (*domain.CartItem, error)
}

type ProductRepo interface {
	Get(id int) (*domain.Product, error)
}
