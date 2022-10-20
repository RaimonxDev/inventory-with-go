package repository

import (
	"context"
	"github.com/RaimonxDev/inventory-with-go/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	// SaveUser Columns of table email, name, password
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

// Repo implementation of Repository interface
type repo struct {
	db *sqlx.DB
}

// New return interface of Repository
func New(db *sqlx.DB) Repository {
	return &repo{db: db}
}
