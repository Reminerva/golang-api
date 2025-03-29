package repository

import (
	"bank-api/models"
	"encoding/json"
	"errors"
	"os"
)

const merchantFile = "merchant.json"

type MerchantRepository struct {
	Data []models.Merchant
}

func (r *MerchantRepository) LoadData() error {
	file, err := os.ReadFile(merchantFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &r.Data)
	if err != nil {
		return err
	}

	return nil
}

func (r *MerchantRepository) SaveData() error {
	byteValue, err := json.MarshalIndent(r.Data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(merchantFile, byteValue, 0644)
}

func (r *MerchantRepository) GetMerchantByID(id string) (*models.Merchant, error) {
	for _, merchant := range r.Data {
		if merchant.ID == id {
			return &merchant, nil
		}
	}
	return nil, errors.New("merchant not found")
}

func (r *MerchantRepository) UpdateBalance(merchantID string, amount float64) error {
	for i, merchant := range r.Data {
		if merchant.ID == merchantID {
			r.Data[i].Balance += amount
			return r.SaveData()
		}
	}
	return errors.New("merchant not found")
}
