package concerns

import (
	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
)

func ConfigureRoutes(
	server *fiber.App,
	currentCurrencyValuesController controllers.CurrentCurrencyValuesController,
	notFoundErrorController controllers.NotFoundErrorController) {
	server.Get(
		constants.CURRENT_CURRENCY_VALUES_ROUTE_PATH,
		currentCurrencyValuesController.HandleCurrentCurrencyValuesRoute)

	server.Use(notFoundErrorController.HandleNotFoundErrorRoute)
}
