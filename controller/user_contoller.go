package controller

import (
	"api/models"
	"api/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewUserController() *UserController {
	return &UserController{
		base:       NewBaseController(),
		repository: repository.NewUserRepository(),
	}
}

type UserController struct {
	base       IBaseController
	repository repository.IUser
}

func (ctrl *UserController) Create(c *fiber.Ctx) error {
	user := models.User{}
	errors := ctrl.base.GetData(
		c,
		&user,
	)

	if len(errors) > 0 {
		return nil
	}

	_, err := ctrl.repository.Create(&user)

	if err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(fiber.Map{
			"success": true,
		})
}

func (ctrl *UserController) List(c *fiber.Ctx) error {
	users := ctrl.repository.List()

	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"success": true,
			"data":    users,
		})
}

func (ctrl *UserController) Detail(c *fiber.Ctx) error {
	id := c.Params("id")
	user := ctrl.repository.Detail(id)

	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"success": true,
			"data":    user,
		})
}

func (ctrl *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := ctrl.repository.Delete(id)

	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"success": false,
				"data":    "Error to delete data",
			})
	}

	return c.
		SendStatus(http.StatusOK)
}

func (ctrl *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	userData := ctrl.repository.Detail(id)
	if userData == nil {
		return ctrl.base.Response(
			c,
			http.StatusNotFound,
			"User not found",
		)
	}

	errors := ctrl.base.GetData(
		c,
		userData,
	)

	if len(errors) > 0 {
		return nil
	}

	err := ctrl.repository.Update(id, userData)

	if err != nil {
		return ctrl.base.Response(
			c,
			http.StatusBadRequest,
			"Error to update data",
		)
	}

	return ctrl.base.Response(
		c,
		http.StatusOK,
		userData,
	)
}
