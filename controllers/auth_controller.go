package controllers

import (
	"bank-api/services"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println("Error parsing request body: ", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	token, err := services.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"data":    token,
	})
}

func LogoutController(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid token",
		})
	}

	loggedOutToken, err := services.Logout(token)
	if err != nil {
		fmt.Println("Error logging out: ", err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token" + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logout successful",
		"data": fiber.Map{
			"token": loggedOutToken,
		},
	})
}
