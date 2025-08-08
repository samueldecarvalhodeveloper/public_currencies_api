package integrations_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"

	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
	loadbalancer "load_balancing_proxy_service/domains/load_balancer"
)

func TestIfServerBalancesTheLoadOfServiceServer(
	t *testing.T) {
	server := fiber.New()
	loadbalancerInstance :=
		loadbalancer.New(
			constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK())
	currentCurrencyValuesController :=
		controllers.NewCurrentCurrencyValuesController(
			loadbalancerInstance)

	requestURI := fmt.Sprintf(
		"%s%s",
		constants.BASE_ROUTER,
		constants.CURRENT_CURRENCY_VALUES_ROUTE_PATH)
	verificationRequest := httptest.NewRequest(
		constants.REQUEST_METHOD,
		requestURI,
		nil)

	server.Get(
		requestURI,
		currentCurrencyValuesController.HandleCurrentCurrencyValuesRoute)

	requestResponse, _ :=
		server.Test(verificationRequest)

	redirectedRequestLocation := requestResponse.Header.Get(fiber.HeaderLocation)

	redirectedServerURI :=
		fmt.Sprintf(
			"%s%s",
			constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK()[0],
			constants.CURRENT_CURRENCY_VALUES_ROUTE_PATH)

	assert.Equal(
		t,
		redirectedServerURI,
		redirectedRequestLocation,
	)
	assert.Equal(
		t,
		fiber.StatusPermanentRedirect,
		requestResponse.StatusCode)
}
