package entity

import (
	"crypto/sha256"
	"fmt"
	"users-service/pkg"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Username string
	Password string
	Image    string
	Bio      string
}

func CreateUser(id uuid.UUID, email, username, password, image, bio string) (*User, error) {
	user := &User{
		ID:       id,
		Email:    email,
		Username: username,
		Password: password,
		Image:    image,
		Bio:      bio,
	}

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	user.encryptPassword()

	return user, nil
}

func (u *User) isValid() error {
	if u.ID == uuid.Nil {
		return fmt.Errorf("invalid id. error_type: %w", pkg.ErrValidation)
	}

	if u.Email == "" {
		return fmt.Errorf("invalid email. error_type: %w", pkg.ErrValidation)
	}

	if u.Username == "" {
		return fmt.Errorf("invalid username. error_type: %w", pkg.ErrValidation)
	}

	if u.Password == "" || len(u.Password) < 8 {
		return fmt.Errorf("invalid password. error_type: %w", pkg.ErrValidation)
	}

	return nil
}

func (u *User) encryptPassword() {
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))
}
