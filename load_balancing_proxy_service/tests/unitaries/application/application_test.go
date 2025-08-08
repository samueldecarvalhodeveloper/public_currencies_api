package application_test

import (
	"io"
	"load_balancing_proxy_service/application"
	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
	loadbalancer "load_balancing_proxy_service/domains/load_balancer"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestIfMethodRunServerRunsServiceServer(
	t *testing.T) {
	server := fiber.New()
	loadbalancerInstance :=
		loadbalancer.New(
			constants.LIST_OF_SERVICES_TO_BE_BALANCED())
	currentCurrencyValuesController :=
		controllers.NewCurrentCurrencyValuesController(
			loadbalancerInstance)
	notFoundErrorController :=
		controllers.NewNotFoundErrorController()

	applicationInstance := application.New(
		server,
		currentCurrencyValuesController,
		notFoundErrorController)

	go func() {
		applicationInstance.RunServer(constants.LOCAL_HOST_PORT)
	}()

	time.Sleep(constants.SERVER_START_UP_DELAY)

	requestResponse, _ := http.Get(
		constants.LOCAL_HOST_URL())
	responseBody, _ := io.ReadAll(
		requestResponse.Body)

	requestResponse.Body.Close()

	responseBodyCastedToString := string(responseBody)

	assert.Equal(
		t,
		constants.NOT_FOUND_ERROR_MESSAGE_RESPONSE_BODY,
		responseBodyCastedToString)
}
