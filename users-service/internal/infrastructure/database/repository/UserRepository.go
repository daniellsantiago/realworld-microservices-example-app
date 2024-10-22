package repository

import (
	"users-service/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) entity.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) Save(user entity.User) error {
	dbUser := &UserModel{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
		Bio:      user.Bio,
	}

	if err := u.DB.Create(&dbUser).Error; err != nil {
		return err
	}

	return nil
}

type UserModel struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
	Image    string    `json:"image"`
	Bio      string    `json:"bio"`
}

func (UserModel) TableName() string {
	return "users"
}
