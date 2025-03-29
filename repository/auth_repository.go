package repository

import (
	"bank-api/models"
	"encoding/json"
	"errors"
	"os"
	"slices"
)

const authFile = "auth_history.json"

type AuthRepository struct {
	Data []models.Auth
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) LoadData() error {
	file, err := os.ReadFile(authFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &r.Data)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) SaveData() error {
	byteValue, err := json.MarshalIndent(r.Data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(authFile, byteValue, 0644)
}

func (r *AuthRepository) GetCustomerByEmail(email string) (*models.Auth, error) {
	for _, customer := range r.Data {
		if customer.Email == email {
			return &customer, nil
		}
	}
	return nil, errors.New("customer has not logged in")
}

func (r *AuthRepository) DeleteCustomerByEmail(email string) error {
	for i, customer := range r.Data {
		if customer.Email == email {
			r.Data = slices.Delete(r.Data, i, i+1)
			return r.SaveData()
		}
	}
	return errors.New("customer not found")
}