package repositories

import (
	"bank-api/config"
	"bank-api/models"
	"encoding/json"
	"errors"
	"os"
)

func GetCustomers() ([]models.Customer, error) {
	data, err := os.ReadFile(config.CustomerDB) // read from customer file
	if err != nil {
		return nil, err
	}

	var customers []models.Customer
	err = json.Unmarshal(data, &customers) // parse json data to customer struct
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func FindCustomerByUsername(username string) (models.Customer, error) {
	customers, err := GetCustomers()
	if err != nil {
		return models.Customer{}, err
	}

	for _, customer := range customers {
		if customer.Username == username {
			return customer, nil
		}
	}

	return models.Customer{}, errors.New("customer not found")
}

func FindCustomerByID(id string) (models.Customer, error) {
	customers, err := GetCustomers()
	if err != nil {
		return models.Customer{}, err
	}

	for _, customer := range customers {
		if customer.ID == id {
			return customer, nil
		}
	}

	return models.Customer{}, errors.New("customer not found")
}

func UpdateCustomer(customer models.Customer) error {
	customers, err := GetCustomers()
	if err != nil {
		return err
	}

	for i, c := range customers {
		if c.ID == customer.ID {
			customers[i] = customer
			break
		}
	}

	data, err := json.MarshalIndent(customers, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.CustomerDB, data, 0644)
}
