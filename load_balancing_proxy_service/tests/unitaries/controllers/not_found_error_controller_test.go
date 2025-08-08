package controllers_test

import (
	"fmt"
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
	responseBodyContentLength := fmt.Sprintf(
		"%d",
		len(responseBodyCastedToString))

	requestResponseHeaderContentType :=
		requestResponse.Header.Get(fiber.HeaderContentType)
	requestResponseHeaderContentLength :=
		requestResponse.Header.Get(fiber.HeaderContentLength)

	requestResponseHeaderAccessControlAllowOrigin :=
		requestResponse.Header.Get(fiber.HeaderAccessControlAllowOrigin)
	requestResponseHeaderAccessControlAllowMethods :=
		requestResponse.Header.Get(fiber.HeaderAccessControlAllowMethods)
	requestResponseHeaderAccessControlAllowHeaders :=
		requestResponse.Header.Get(fiber.HeaderAccessControlAllowHeaders)

	assert.Equal(
		t,
		constants.PROTOBUF_CONTENT_TYPE_VALUE,
		requestResponseHeaderContentType)
	assert.Equal(
		t,
		responseBodyContentLength,
		requestResponseHeaderContentLength)
	assert.Equal(
		t,
		constants.ACCESS_CONTROL_ALLOW_ORIGIN_VALUE,
		requestResponseHeaderAccessControlAllowOrigin)
	assert.Equal(
		t,
		constants.ACCESS_CONTROL_ALLOW_METHODS_VALUE,
		requestResponseHeaderAccessControlAllowMethods)
	assert.Equal(
		t,
		constants.ACCESS_CONTROL_ALLOW_HEADERS_VALUE,
		requestResponseHeaderAccessControlAllowHeaders)
	assert.Equal(
		t,
		fiber.StatusNotFound,
		requestResponse.StatusCode)
	assert.Equal(
		t,
		constants.NOT_FOUND_ERROR_MESSAGE_RESPONSE_BODY,
		responseBodyCastedToString)
}
