package main

import (
	"users-service/internal/infrastructure/config"
	"users-service/internal/infrastructure/web"
	"users-service/pkg"

	"go.uber.org/zap"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	pkg.InitLogger(configs.Env)

	handlers := web.NewUserHandlers()

	webserver := web.NewWebServer(configs.WebServerPort, handlers)

	web.SetupRoutes(webserver)

	pkg.Logger.Info("Starting GIN server...", zap.String("port", configs.WebServerPort))

	err = webserver.Start()
	if err != nil {
		panic(err)
	}
}
