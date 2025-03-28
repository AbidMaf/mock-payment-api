package services

import (
	"bank-api/models"
	"bank-api/repositories"
	"errors"
	"time"

	"github.com/google/uuid"
)

func PaymentService(username, merchantID string, amount int) (models.TransactionHistory, error) {
	customer, err := repositories.FindCustomerByUsername(username)

	if err != nil {
		return models.TransactionHistory{}, errors.New("customer not found")
	}

	merchant, err := repositories.GetMerchantByID(merchantID)
	if err != nil {
		return models.TransactionHistory{}, errors.New("merchant not found")
	}

	if customer.Balance < amount {
		return models.TransactionHistory{}, errors.New("balance is not enough")
	}

	customer.Balance -= amount
	err = repositories.UpdateCustomer(customer)
	if err != nil {
		return models.TransactionHistory{}, errors.New("failed to update customer balance. " + err.Error())
	}

	transaction := models.TransactionHistory{
		ID:        uuid.New().String(),
		Merchant:  merchant,
		Customer:  customer,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	err = repositories.SaveTransactionHistory(transaction)
	if err != nil {
		return models.TransactionHistory{}, errors.New("failed to save transaction history. " + err.Error())
	}

	return transaction, nil
}
