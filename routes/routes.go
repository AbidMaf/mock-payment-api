package routes

import (
	"bank-api/controllers"
	"bank-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/login", controllers.LoginController)
	api.Post("/payment", middlewares.AuthMiddleware, controllers.PaymentController)
	api.Post("/logout", middlewares.AuthMiddleware, controllers.LogoutController)
}
