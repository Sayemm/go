package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

// This defines what the user service can do
// Handlers depend on this interface (not the concrete implementation)
type Service interface {
	userHandler.Service
}

// This defines what we need from the repository
// Service depends on this interface (not the concrete implementation)
type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
