package routes

import (
	"api/framework/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func (r *Route) setPublicRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//authController := controller2.NewAuthController()
	userController := user.NewUserController(r.logger, r.repository)

	//app.Post("/login", authController.Login)

	app.Post("/user", userController.Create)
	//app.Get("/user", userController.List)
	//app.Get("/user/:id", userController.Detail)
	//app.Patch("/user/:id", userController.Update)
	//app.Delete("/user/:id", userController.Delete)
}
