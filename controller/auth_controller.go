package controller

import (
	"api/configs"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct {
	base IBaseController
}

type LoginData struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

func NewAuthController() *AuthController {
	return &AuthController{
		base: NewBaseController(),
	}
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	login := LoginData{}
	erros := controller.base.GetData(c, &login)
	if len(erros) > 0 {
		return nil
	}

	authTokens := configs.GetTokens()

	data, exists := authTokens[login.Token]

	fmt.Println(login)
	fmt.Println(data)

	if !exists ||
		data.Secret != login.Secret {
		return c.
			Status(http.StatusUnauthorized).
			JSON(fiber.Map{
				"success": false,
				"message": "Invalid credentials",
			})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  data.Name,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
