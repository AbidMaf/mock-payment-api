package repositories

import (
	"bank-api/config"
	"bank-api/models"
	"encoding/json"
	"os"
)

func SaveTransactionHistory(history models.TransactionHistory) error {
	histories, err := GetTransactionHistories()
	if err != nil {
		return err
	}

	histories = append(histories, history)

	data, err := json.MarshalIndent(histories, "", "	") // format histories data to json and applies indent
	if err != nil {
		return err
	}

	return os.WriteFile(config.HistoryDB, data, 0644)
}

func GetTransactionHistories() ([]models.TransactionHistory, error) {
	data, err := os.ReadFile(config.HistoryDB) // read from history file
	if err != nil {
		return nil, err
	}

	var histories []models.TransactionHistory
	err = json.Unmarshal(data, &histories) // parse json data to history struct
	if err != nil {
		return nil, err
	}

	return histories, nil
}
