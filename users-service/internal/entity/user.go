package entity

import (
	"errors"

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

func NewUser(id uuid.UUID, email, username, password, image, bio string) (*User, error) {
	user := &User{
		ID:       id,
		Email:    email,
		Username: username,
		Password: password,
		Image:    image,
		Bio:      bio,
	}

	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) IsValid() error {
	if u.ID == uuid.Nil {
		return errors.New("invalid id")
	}

	if u.Email == "" {
		return errors.New("invalid email")
	}

	if u.Username == "" {
		return errors.New("invalid username")
	}

	if u.Password == "" {
		return errors.New("invalid password")
	}

	return nil
}
