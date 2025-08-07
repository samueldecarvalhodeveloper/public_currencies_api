package application

import (
	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
)

type Application struct {
	server *fiber.App
}

func New(
	server *fiber.App,
	currentCurrencyValuesController controllers.CurrentCurrencyValuesController,
	notFoundErrorController controllers.NotFoundErrorController) Application {
	server.Get(constants.CURRENT_CURRENCY_VALUES_ROUTE_PATH, currentCurrencyValuesController.HandleCurrentCurrencyValuesRoute)

	server.Use(notFoundErrorController.HandleNotFoundErrorRoute)

	application := Application{
		server: server,
	}

	return application
}

func (application Application) RunServer() {
	application.server.Listen(constants.SERVER_PORT)
}
