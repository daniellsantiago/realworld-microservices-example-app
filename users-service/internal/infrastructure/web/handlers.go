package web

import (
	"net/http"
	"users-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandlers interface {
	CreateUser(c *gin.Context)
}

type userHandlers struct {
	createUserUseCase usecase.CreateUser
}

func NewUserHandlers(createUserUseCase usecase.CreateUser) UserHandlers {
	return &userHandlers{
		createUserUseCase: createUserUseCase,
	}
}

func (h userHandlers) CreateUser(c *gin.Context) {
	var input *usecase.CreateUserInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOutputDTO, err := h.createUserUseCase.Execute(*input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": userOutputDTO})
}
