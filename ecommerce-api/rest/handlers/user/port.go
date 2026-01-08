package user

import "ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}

// I need these two methods
// I won't implement this, service layer will do this
// this interface will be embedded in servie layer
