package controllers

import (
	"github.com/gofiber/fiber/v3"

	"load_balancing_proxy_service/constants"
)

type NotFoundErrorController struct {
}

func NewNotFoundErrorController() NotFoundErrorController {
	return NotFoundErrorController{}
}

func (notFoundErrorController *NotFoundErrorController) HandleNotFoundErrorRoute(context fiber.Ctx) error {
	return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
		constants.NOT_FOUND_ERROR_RESPONSE_MESSAGE_KEY: constants.NOT_FOUND_ERROR_RESPONSE_MESSAGE_VALUE,
	})
}
