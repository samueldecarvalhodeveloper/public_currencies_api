package application

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/controllers"
	"load_balancing_proxy_service/infrastructure/adapters"
)

type Application struct {
	server *fiber.App
}

func New(
	server *fiber.App,
	currentCurrencyValuesController controllers.CurrentCurrencyValuesController,
	notFoundErrorController controllers.NotFoundErrorController) Application {
	adapters.ConfigureCurrencyRoutes(
		server,
		currentCurrencyValuesController)
	adapters.ConfigureNotFoundErrorRoute(
		server,
		notFoundErrorController)

	application := Application{
		server: server,
	}

	return application
}

func (application Application) RunServer(port int) {
	serverPort := fmt.Sprintf(":%d", port)

	application.server.Listen(serverPort)
}
