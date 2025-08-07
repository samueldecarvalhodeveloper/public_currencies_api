package adapters

import (
	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/controllers"
)

func ConfigureNotFoundErrorRoute(
	server *fiber.App,
	notFoundErrorController controllers.NotFoundErrorController) {
	server.Use(notFoundErrorController.HandleNotFoundErrorRoute)
}
