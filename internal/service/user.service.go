package service

import (
	"context"
	"errors"
	"github.com/RaimonxDev/inventory-with-go/encryption"
	"github.com/RaimonxDev/inventory-with-go/internal/model"
	"log"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *service) RegisterUser(ctx context.Context, email, name, password string) error {

	user, _ := s.Repository.GetUserByEmail(ctx, email)

	// If user already exists
	if user != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	// Password is encrypted and saved in database
	passDecode := encryption.ToBase64(bb)
	log.Printf("Password: %s", passDecode)
	return s.Repository.SaveUser(ctx, email, name, passDecode)
}

func (s *service) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Password se guarda como un string encriptado en base64
	// y se compara con el password que se envia en el login
	bb, err := encryption.FromBase64(user.Password)
	if err != nil {
		return nil, err
	}
	// Password is decrypted
	descrytedPassword, err := encryption.Decrypt(bb)

	// Password is compared with the one that is in the database
	if string(descrytedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	// if password is correct, return user
	return &model.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
