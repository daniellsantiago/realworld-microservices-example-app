package usecase

import (
	"fmt"
	"users-service/internal/entity"
	"users-service/internal/service"
	"users-service/pkg"

	"github.com/google/uuid"
)

type CreateUser struct {
	UserRepository entity.UserRepository
	JwtService     service.JwtService
}

func NewCreateUser(userRepository entity.UserRepository, jwtService service.JwtService) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
		JwtService:     jwtService,
	}
}

func (c *CreateUser) Execute(createDTO CreateUserInputDTO) (CreateUserOutputDTO, error) {
	user, err := entity.NewUser(uuid.New(), createDTO.User.Email, createDTO.User.Username, createDTO.User.Password, "", "")
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	err = c.UserRepository.Save(*user)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	token, err := c.JwtService.GenerateToken(user.ID)
	if err != nil {
		return CreateUserOutputDTO{}, fmt.Errorf("failed to generate token. error_type: %w", pkg.ErrInternal)
	}

	return CreateUserOutputDTO{
		User: UserOutputDTO{
			Username: user.Username,
			Email:    user.Email,
			Image:    user.Image,
			Bio:      user.Bio,
			Token:    token,
		},
	}, nil
}

type CreateUserInputDTO struct {
	User UserInputDTO `json:"user" binding:"required"`
}

type UserInputDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type CreateUserOutputDTO struct {
	User UserOutputDTO `json:"user"`
}

type UserOutputDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Bio      string `json:"bio"`
	Token    string `json:"token"`
}
