package services

import (
	"bank-api/models"
	"bank-api/repository"
	"errors"
	"math/rand"
	"strconv"
)

type MerchantService struct {
	repo *repository.MerchantRepository
}

func NewMerchantService() *MerchantService {
	repo := &repository.MerchantRepository{}
	return &MerchantService{repo: repo}
}

func (s *MerchantService) CreateMerchant(name, address string, balance float64) error {
	err := s.repo.LoadData()
	if err != nil {
		return err
	}

	for _, customer := range s.repo.Data {
		if customer.Address == address {
			return errors.New("address already exists")
		}
	}

	var newID string
	for {
		newID = "m-" + strconv.Itoa(rand.Intn(10000))
		_, err := s.repo.GetMerchantByID(newID)
		if err != nil {
			break
		}
	}

	newMerchant := models.Merchant{
		ID:      newID,
		Name:    name,
		Address: address,
		Balance: balance,
	}

	s.repo.Data = append(s.repo.Data, newMerchant)

	err = s.repo.SaveData()
	if err != nil {
		return err
	}

	return nil
}

func (s *MerchantService) GetAllMerchant() ([]models.Merchant, error) {
	err := s.repo.LoadData()
	if err != nil {
		return nil, err
	}
	return s.repo.Data, nil
}

func (s *MerchantService) GetMerchantByID(id string) (*models.Merchant, error) {
	err := s.repo.LoadData()
	if err != nil {
		return nil, err
	}
	return s.repo.GetMerchantByID(id)
}
