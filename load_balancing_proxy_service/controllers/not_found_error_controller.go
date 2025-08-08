package controllers

import (
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

	context.Set(
		fiber.HeaderContentType,
		constants.PROTOBUF_CONTENT_TYPE_VALUE)

	return context.Status(
		fiber.StatusNotFound).Send(notFoundErrorMessageResponse)
}
