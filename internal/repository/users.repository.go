package repository

import (
	"context"
	"github.com/RaimonxDev/inventory-with-go/internal/entity"
)

const (
	querySaveUser = `INSERT INTO USERS (email, name, password) VALUES (?, ?, ?)`

	queryGetUser = `SELECT 
       id, 
       email ,
       name ,
       password 
FROM USERS WHERE email = ?`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, querySaveUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	// Need reference to entity user
	user := &entity.User{}

	err := r.db.GetContext(ctx, user, queryGetUser, email)

	if err != nil {
		return nil, err
	}

	return user, err
}
