package user

import (
	userService "api/application/services/user"
	"api/dependencies/repositories/mysql"
	domain "api/domain/user"
	"api/framework/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type Controller struct {
	baseController controllers.IBaseController
	service        *userService.Service
}

func NewUserController(logger *log.Logger, repository *mysql.Repository) *Controller {
	return &Controller{
		baseController: controllers.NewBaseController(),
		service:        userService.NewService(logger, repository),
	}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	user := domain.User{}
	errors := c.baseController.GetData(
		ctx,
		&user,
	)

	if len(errors) > 0 {
		//logar e setar response
		return nil
	}

	err := c.service.Create(&user)
	if err != nil {
		//logar e setar resposta
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(fiber.Map{
			"success": true,
		})
}

//func (c *Controller) List(c *fiber.Ctx) error {
//	where, args := utils.ParseQueryParams(string(c.Request().URI().QueryString()))
//
//	users := ctrl.repository.List(where, args)
//
//	return c.
//		Status(http.StatusOK).
//		JSON(fiber.Map{
//			"success": true,
//			"data":    users,
//		})
//}
//
//func (c *Controller) Detail(c *fiber.Ctx) error {
//	id := c.Params("id")
//	user := ctrl.repository.Detail(id)
//
//	return c.
//		Status(http.StatusOK).
//		JSON(fiber.Map{
//			"success": true,
//			"data":    user,
//		})
//}
//
//func (c *Controller) Delete(c *fiber.Ctx) error {
//	id := c.Params("id")
//
//	err := ctrl.repository.Delete(id)
//
//	if err != nil {
//		return c.
//			Status(http.StatusInternalServerError).
//			JSON(fiber.Map{
//				"success": false,
//				"data":    "Error to delete data",
//			})
//	}
//
//	return c.
//		SendStatus(http.StatusOK)
//}
//
//func (c *Controller) Update(c *fiber.Ctx) error {
//	id := c.Params("id")
//
//	userData := ctrl.repository.Detail(id)
//	if userData == nil {
//		return ctrl.base.Response(
//			c,
//			http.StatusNotFound,
//			"User not found",
//		)
//	}
//
//	errors := ctrl.base.GetData(
//		c,
//		userData,
//	)
//
//	if len(errors) > 0 {
//		return nil
//	}
//
//	err := ctrl.repository.Update(id, userData)
//
//	if err != nil {
//		return ctrl.base.Response(
//			c,
//			http.StatusBadRequest,
//			"Error to update data",
//		)
//	}
//
//	return ctrl.base.Response(
//		c,
//		http.StatusOK,
//		userData,
//	)
//}
