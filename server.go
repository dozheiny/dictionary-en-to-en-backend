package main

import (
	"DictionaryENtoENBackend/config"
	"DictionaryENtoENBackend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	router := fiber.New()
	// use recover function for when server panic, recover server.
	router.Use(recover())
	router.Use(logger.New(logger.Config{TimeFormat: "2006-01-02 15:04:05"}))

	api := router.Group("/api" + config.ApiVersion)
	controllers.RegisterAll(api)
}