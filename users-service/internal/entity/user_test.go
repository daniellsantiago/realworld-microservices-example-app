package entity

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Run("returns a valid User with encrypted password when valid params are provided", func(t *testing.T) {
		// Given
		id := uuid.New()
		email := "email@gmail.com"
		username := "username"
		password := "1233456677"
		image := "path/to/image.png"
		bio := "Cool guy!"
		encryptedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

		// When
		user, err := CreateUser(id, email, username, password, image, bio)

		// Then
		assert.Nil(t, err)
		assert.Equal(t, id, user.ID)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, encryptedPassword, user.Password)
		assert.Equal(t, image, user.Image)
		assert.Equal(t, bio, user.Bio)
	})

	t.Run("returns error when a required param is not provided", func(t *testing.T) {
		// Given
		id := uuid.New()
		username := "username"
		password := "123456778"
		image := "path/to/image.png"
		bio := "Cool guy!"

		// When
		_, err := CreateUser(id, "", username, password, image, bio)

		// Then
		assert.Error(t, err)
	})

	t.Run("returns error when password is less than 8 digits", func(t *testing.T) {
		// Given
		id := uuid.New()
		email := "email@gmail.com"
		username := "username"
		password := "123"
		image := "path/to/image.png"
		bio := "Cool guy!"

		// When
		_, err := CreateUser(id, email, username, password, image, bio)

		// Then
		assert.Error(t, err)
	})
}
