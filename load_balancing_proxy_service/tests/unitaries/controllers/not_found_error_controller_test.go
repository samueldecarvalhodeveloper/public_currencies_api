package controllers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"

	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/controllers"
)

func TestIfMethodHandleNotFoundErrorRouteHandlesNotFoundErrorRequestResponse(
	t *testing.T) {
	server := fiber.New()
	notFoundErrorController :=
		controllers.NewNotFoundErrorController()

	verificationRequest := httptest.NewRequest(
		constants.REQUEST_METHOD,
		constants.NOT_EXISTING_ROUTE,
		nil)

	server.Use(notFoundErrorController.HandleNotFoundErrorRoute)

	requestResponse, _ := server.Test(
		verificationRequest)

	responseBody := make([]byte, requestResponse.ContentLength)

	requestResponse.Body.Read(responseBody)
	requestResponse.Body.Close()

	responseBodyCastedToString := string(responseBody)

	assert.Equal(
		t,
		fiber.StatusNotFound,
		requestResponse.StatusCode)
	assert.Equal(
		t,
		constants.NOT_FOUND_ERROR_MESSAGE_RESPONSE_BODY,
		responseBodyCastedToString)
}
