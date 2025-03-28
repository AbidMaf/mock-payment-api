package repositories

import (
	"bank-api/config"
	"bank-api/models"
	"encoding/json"
	"errors"
	"os"
)

func GetMerchants() ([]models.Merchant, error) {
	data, err := os.ReadFile(config.MerchantDB) // read from merchant file
	if err != nil {
		return nil, err
	}

	var merchants []models.Merchant
	err = json.Unmarshal(data, &merchants) // parse json data to merchant struct
	if err != nil {
		return nil, err
	}

	return merchants, nil
}

func GetMerchantByID(id string) (models.Merchant, error) {
	merchants, err := GetMerchants()

	if err != nil {
		return models.Merchant{}, err
	}

	for _, merchant := range merchants {
		if merchant.ID == id {
			return merchant, nil
		}
	}
	return models.Merchant{}, errors.New("merchant not found")
}
