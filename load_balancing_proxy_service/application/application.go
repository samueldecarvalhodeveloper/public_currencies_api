package application

import (
	application_constants "load_balancing_proxy_service/constants"

	"github.com/gofiber/fiber/v3"
)

type Application struct {
	server *fiber.App
}

func New() Application {
	server := fiber.New()

	ConfigureServer(server)

	application := Application{
		server: server,
	}

	return application
}

func (application Application) RunServer() {
	application.server.Listen(application_constants.SERVER_PORT)
}

func ConfigureServer(serverToBeConfigured *fiber.App) {
	serverToBeConfigured.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	serverToBeConfigured.Use(func(c fiber.Ctx) error {
		c.Status(fiber.StatusNotFound)

		return c.JSON(fiber.Map{
			"message": "Hi John!",
		}) // => 404 "Not Found"
	})
}
