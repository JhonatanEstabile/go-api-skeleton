package controllers

import (
	"api/framework/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type IBaseController interface {
	GetData(*fiber.Ctx, interface{}) []string
	Response(c *fiber.Ctx, statusCode int, data interface{}) error
}

type baseController struct{}

func NewBaseController() *baseController {
	return &baseController{}
}

func (base *baseController) GetData(
	c *fiber.Ctx,
	structData interface{},
) []string {
	var errors []string

	if err := c.BodyParser(structData); err != nil {
		errors = append(errors, "Invalid Payload")

		c.Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"success": false,
				"errors":  errors,
			})

		return errors
	}

	errors = utils.ValUtil.ValidateStruct(structData)
	if len(errors) > 0 {
		c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"success": false,
				"errors":  errors,
			})

		return errors
	}

	return errors
}

func (base *baseController) Response(
	c *fiber.Ctx,
	statusCode int,
	data interface{},
) error {
	var success bool = true

	if statusCode > 299 && statusCode < 505 {
		success = false
	}

	return c.
		Status(statusCode).
		JSON(fiber.Map{
			"success": success,
			"data":    data,
		})
}
