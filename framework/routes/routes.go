package routes

import (
	"api/dependencies/repositories/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Route struct {
	logger     *log.Logger
	repository *mysql.Repository
}

func NewRoutes(log *log.Logger, repo *mysql.Repository) *Route {
	return &Route{
		logger:     log,
		repository: repo,
	}
}

func (r *Route) SetRoutes(app *fiber.App) {
	r.setPublicRoutes(app)
	r.setPrivateRoutes(app)
}
