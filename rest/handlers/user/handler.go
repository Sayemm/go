package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	cnf      *config.Config
}

func NewHandler(userRepo repo.UserRepo, cnf *config.Config) *Handler {
	return &Handler{
		userRepo: userRepo,
		cnf:      cnf,
	}
}
