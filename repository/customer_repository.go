package repository

import (
	"bank-api/models"
	"encoding/json"
	"errors"
	"os"
)

const customerFile = "customer.json"

type CustomerRepository struct {
	Data []models.Customer
}

func (r *CustomerRepository) LoadData() error {
	file, err := os.ReadFile(customerFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &r.Data)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) SaveData() error {
	byteValue, err := json.MarshalIndent(r.Data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(customerFile, byteValue, 0644)
}

func (r *CustomerRepository) GetCustomerByEmail(email string) (*models.Customer, error) {
	for _, customer := range r.Data {
		if customer.Email == email {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}

func (r *CustomerRepository) GetCustomerByID(customerID string) (*models.Customer, error) {
	for _, customer := range r.Data {
		if customer.ID == customerID {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}

func (r *CustomerRepository) UpdateBalance(customerID string, amount float64) error {
	for i, customer := range r.Data {
		if customer.ID == customerID {
			r.Data[i].Balance += amount
			return r.SaveData()
		}
	}
	return errors.New("customer not found")
}
