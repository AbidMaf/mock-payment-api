package middlewares

import (
	"bank-api/repositories"
	"bank-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid token",
		})
	}

	token = utils.RemoveBearerPrefix(token)

	session, err := repositories.GetSessionByToken(token)
	if err != nil || session.Token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	validationResult, err := utils.ValidateToken(token)
	if err != nil {
		repositories.RemoveSession(token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	c.Locals("username", validationResult.Claims["username"])
	return c.Next()
}
