package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/constants"
	loadbalancer "load_balancing_proxy_service/domains/load_balancer"
)

type CurrentCurrencyValuesController struct {
	loadBalancer loadbalancer.LoadBalancer
}

func NewCurrentCurrencyValuesController(loadBalancer loadbalancer.LoadBalancer) CurrentCurrencyValuesController {
	return CurrentCurrencyValuesController{
		loadBalancer: loadBalancer,
	}
}

func (currentCurrencyValuesController *CurrentCurrencyValuesController) HandleCurrentCurrencyValuesRoute(context fiber.Ctx) error {
	currentServerToBeRedirected := currentCurrencyValuesController.loadBalancer.GetCurrentServer()

	currentServerToBeRedirectedURI := fmt.Sprintf("%s%s", currentServerToBeRedirected, constants.CURRENT_CURRENCY_VALUES_ROUTE_PATH)

	return context.Redirect().Status(fiber.StatusPermanentRedirect).To(currentServerToBeRedirectedURI)
}
