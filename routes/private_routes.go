package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func setPrivateRoutes(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Logged")
	})
}
