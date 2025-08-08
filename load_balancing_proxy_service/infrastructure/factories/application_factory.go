package factories

import (
	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/application"
	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
	loadbalancer "load_balancing_proxy_service/domains/load_balancer"
)

func GetApplicationInstance() application.Application {
	server := fiber.New()
	loadbalancerInstance :=
		loadbalancer.New(constants.LIST_OF_SERVICES_TO_BE_BALANCED())
	currentCurrencyValuesController :=
		controllers.NewCurrentCurrencyValuesController(loadbalancerInstance)
	notFoundErrorController :=
		controllers.NewNotFoundErrorController()

	return application.New(
		server,
		currentCurrencyValuesController,
		notFoundErrorController)
}
