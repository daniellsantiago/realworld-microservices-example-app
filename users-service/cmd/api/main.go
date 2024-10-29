package main

import (
	"users-service/internal/infrastructure/config"
	"users-service/internal/infrastructure/database"
	"users-service/internal/infrastructure/database/repository"
	"users-service/internal/infrastructure/web"
	"users-service/internal/service"
	"users-service/internal/usecase"
	"users-service/pkg"

	"go.uber.org/zap"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	pkg.InitLogger(configs.Env)

	db, err := database.Connect(configs.DBHost, configs.DBUser, configs.DBPassword, configs.DBName, configs.DBPort)
	if err != nil {
		panic(err)
	}

	jwtService := service.NewJwtService(configs.JwtSecret)

	userRepository := repository.NewUserRepository(db)

	createUserUseCase := usecase.NewCreateUser(userRepository, jwtService)

	handlers := web.NewUserHandlers(*createUserUseCase)

	webserver := web.NewWebServer(configs.WebServerPort, handlers)

	web.SetupRoutes(webserver)

	pkg.Logger.Info("Starting GIN server...", zap.String("port", configs.WebServerPort))

	err = webserver.Start()
	if err != nil {
		panic(err)
	}
}
