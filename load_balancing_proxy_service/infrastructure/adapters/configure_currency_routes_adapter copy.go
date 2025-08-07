package adapters

import (
	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
)

func ConfigureCurrencyRoutes(
	server *fiber.App,
	currentCurrencyValuesController controllers.CurrentCurrencyValuesController) {
	server.Get(
		constants.CURRENT_CURRENCY_VALUES_ROUTE_PATH,
		currentCurrencyValuesController.HandleCurrentCurrencyValuesRoute)
}
