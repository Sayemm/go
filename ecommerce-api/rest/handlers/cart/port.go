package cart

import "ecommerce/domain"

type Service interface {
	AddItem(userID, productID, quantity int) (*domain.CartItemDetail, error)
	GetCart(userID int) (*domain.CartDetails, error)
	UpdateItem(itemID, quantity int) error
	RemoveItem(itemID int) error
	ClearCart(userID int) error
}
