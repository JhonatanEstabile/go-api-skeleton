package controller

import (
	"api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type IBaseController interface {
	GetData(*fiber.Ctx, interface{}) []string
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
