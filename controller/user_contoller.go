package controller

import (
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

type userReqBody struct {
	Name  string `json:"name" validate:"required,min=3,max=32"`
	Email string `json:"email" validate:"required,email"`
}

func (usrCtrl *UserController) CreateUser(c *fiber.Ctx) error {
	user := userReqBody{}
	errors := usrCtrl.base.GetData(
		c,
		&user,
	)

	if len(errors) > 0 {
		return nil
	}

	_, err := usrCtrl.repository.CreateUser(
		user.Name,
		user.Email,
	)

	if err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(fiber.Map{
			"success": true,
		})
}
