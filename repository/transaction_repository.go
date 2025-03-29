package repository

import (
	"bank-api/models"
	"encoding/json"
	"os"
	"time"
)

const transactionFile = "transaction.json"

type TransactionRepository struct {
	Data []models.Transaction
}

func (r *TransactionRepository) LoadData() error {
	file, err := os.ReadFile(transactionFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &r.Data)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) SaveData() error {
	byteValue, err := json.MarshalIndent(r.Data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(transactionFile, byteValue, 0644)
}

func (r *TransactionRepository) CreateTransaction(customerID, merchantID string, amount float64) error {
	newTransaction := models.Transaction{
		ID:         time.Now().Format("20010101150405"),
		CustomerID: customerID,
		MerchantID: merchantID,
		Amount:     amount,
		Timestamp:  time.Now(),
	}

	r.Data = append(r.Data, newTransaction)
	return r.SaveData()
}