package routes

import (
	"api/controller"

	"github.com/gofiber/fiber/v2"
)

func setPublicRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	authController := controller.NewAuthController()
	userController := controller.NewUserController()

	app.Post("/login", authController.Login)
	app.Post("/user", userController.CreateUser)
}
