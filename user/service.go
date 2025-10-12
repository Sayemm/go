package user

import (
	"ecommerce/domain"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
func (svc *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := svc.userRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

/*
port.go
-------
type Service interface {
    userHandler.Service
}
- It means: “Any type that claims to implement user.Service
             must have all the methods that userHandler.Service has (Create and Find).”

Go is structural, not nominal
-----------------------------
- Go doesn’t use “extends” or “implements” keywords.
- A type (service) automatically implements an interface if it has the required methods.
	=> service has all the methods that Service requires → it automatically implements the interface.

How it's being forced to create Find and Create
-----------------------------------------------
- if we forget to create any funcion (Find/Create) then we won't be able to create the NewService as it's returning Service

Example: var svc Service = &service{} <- this will fail if we don't implement both functions
- Go allows this assignment because service has all methods required by Service interface.
- This is called structural typing: Go only cares about the methods, not explicit declaration.
*/
