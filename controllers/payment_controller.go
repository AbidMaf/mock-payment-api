package controllers

import (
	"bank-api/repositories"
	"bank-api/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PaymentController(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	var request struct {
		MerchantID string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	transaction, err := services.PaymentService(username, request.MerchantID, request.Amount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := repositories.SaveTransactionHistory(transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save transaction history",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Payment successful",
		"data": fiber.Map{
			"transaction": fiber.Map{
				"id":     transaction.ID,
				"amount": transaction.Amount,
				"customer": fiber.Map{
					"name":    transaction.Customer.Username,
					"balance": transaction.Customer.Balance,
				},
				"merchant": fiber.Map{
					"name": transaction.Merchant.Name,
				},
				"created_at": transaction.CreatedAt.Format(time.RFC3339),
			},
		},
	})
}
