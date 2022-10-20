package service

import (
	"context"
	"github.com/RaimonxDev/inventory-with-go/internal/model"
	"github.com/RaimonxDev/inventory-with-go/internal/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}

type service struct {
	repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{repository}
}
