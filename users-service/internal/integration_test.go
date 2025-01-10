package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"users-service/internal/entity"
	"users-service/internal/entity/mock_entity"
	"users-service/internal/infrastructure/web"
	"users-service/internal/service"
	"users-service/internal/service/mock_service"
	"users-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock_entity.NewMockUserRepository(mockCtrl)
	mockJwtService := mock_service.NewMockJwtService(mockCtrl)

	router := setupRouter(mockUserRepository, mockJwtService)

	t.Run("Create User Successfully", func(t *testing.T) {
		// Given
		input := usecase.CreateUserInputDTO{
			User: usecase.UserInputDTO{
				Username: "dan",
				Email:    "dan@gmail.com",
				Password: "12345678",
				Image:    "path/to/logo.png",
				Bio:      "cool guy!",
			},
		}

		// When
		mockUserRepository.EXPECT().Save(gomock.Any()).Return(nil)
		mockJwtService.EXPECT().GenerateToken(gomock.Any()).Return("token", nil)
		response := doHttpRequest(router, input)

		// Then
		assert.Equal(t, 201, response.Code)
	})

	t.Run("Should get 400 when required input is not provided", func(t *testing.T) {
		// Given
		input := usecase.CreateUserInputDTO{
			User: usecase.UserInputDTO{
				Email: "dan@gmail.com",
				Image: "path/to/logo.png",
				Bio:   "cool guy!",
			},
		}

		// When
		response := doHttpRequest(router, input)

		// Then
		assert.Equal(t, 400, response.Code)
	})

	t.Run("Should get 422 when password is not greater than 8 digits", func(t *testing.T) {
		// Given
		input := usecase.CreateUserInputDTO{
			User: usecase.UserInputDTO{
				Username: "dan",
				Email:    "dan@gmail.com",
				Password: "123",
				Image:    "path/to/logo.png",
				Bio:      "cool guy!",
			},
		}

		// When
		response := doHttpRequest(router, input)

		// Then
		assert.Equal(t, 500, response.Code)
	})
}

func doHttpRequest(router *gin.Engine, input usecase.CreateUserInputDTO) *httptest.ResponseRecorder {
	js, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(js))
	router.ServeHTTP(w, req)

	return w
}

func setupRouter(userRepository entity.UserRepository, jwtService service.JwtService) *gin.Engine {
	createUserUseCase := usecase.NewCreateUser(userRepository, jwtService)
	handlers := web.NewUserHandlers(createUserUseCase)
	server := web.NewWebServer(":8080", handlers)
	web.SetupRoutes(server)

	return server.Router
}
