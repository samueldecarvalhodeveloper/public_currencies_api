package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/infrastructure/concerns"
)

type NotFoundErrorController struct {
}

func NewNotFoundErrorController() NotFoundErrorController {
	return NotFoundErrorController{}
}

func (notFoundErrorController *NotFoundErrorController) HandleNotFoundErrorRoute(
	context fiber.Ctx) error {
	notFoundErrorMessageResponse := concerns.GetNotFoundErrorMessageResponse()
	notFoundErrorMessageResponseContentLength := fmt.Sprintf(
		"%d",
		len(notFoundErrorMessageResponse))

	context.Set(
		fiber.HeaderContentType,
		constants.PROTOBUF_CONTENT_TYPE_VALUE)
	context.Set(
		fiber.HeaderContentLength,
		notFoundErrorMessageResponseContentLength)

	context.Set(
		fiber.HeaderAccessControlAllowOrigin,
		constants.ACCESS_CONTROL_ALLOW_ORIGIN_VALUE)
	context.Set(
		fiber.HeaderAccessControlAllowMethods,
		constants.ACCESS_CONTROL_ALLOW_METHODS_VALUE)
	context.Set(
		fiber.HeaderAccessControlAllowHeaders,
		constants.ACCESS_CONTROL_ALLOW_HEADERS_VALUE)

	return context.Status(
		fiber.StatusNotFound).Send(notFoundErrorMessageResponse)
}
