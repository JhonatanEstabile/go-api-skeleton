package main

import (
	"api/dependencies/repositories/mysql"
	"api/framework/routes"
	"api/framework/utils"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	Environment = "ENVIRONMENT"
	App         = "APP"
	Version     = "VERSION"
	LogLevel    = "LOG_LEVEL"

	DbUser     = "DB_USER"
	DbPassword = "DB_PASSWORD"
	DbHost     = "DB_HOST"
	DbName     = "DB_NAME"
)

var db *mysql.Repository
var logger *log.Logger

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db = mysql.NewMysql(
		logger,
		os.Getenv(DbUser),
		os.Getenv(DbPassword),
		os.Getenv(DbHost),
		os.Getenv(DbName),
	)
	db.Connect()
}

func main() {
	defer db.Close()
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	utils.ValUtil = utils.NewValidatorUtil()

	r := routes.NewRoutes(logger, db)
	r.SetRoutes(app)

	app.Listen(":8080")
}
