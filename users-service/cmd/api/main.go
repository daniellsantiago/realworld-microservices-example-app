package main

import (
	"users-service/internal/infrastructure/config"
	"users-service/internal/infrastructure/web"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	handlers := web.NewUserHandlers()

	webserver := web.NewWebServer(configs.WebServerPort, handlers)

	web.SetupRoutes(webserver)

	err = webserver.Start()
	if err != nil {
		panic(err)
	}
}
