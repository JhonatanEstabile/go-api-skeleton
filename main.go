package main

import (
	"api/db"
	"api/routes"
	"api/utils"
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	defer db.Close()

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	utils.ValUtil = utils.NewValidatorUtil()

	routes.SetRoutes(app)

	app.Listen(":8080")
}
