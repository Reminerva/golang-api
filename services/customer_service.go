package services

import (
	"bank-api/models"
	"bank-api/repository"
	"errors"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService() *CustomerService {
	repo := &repository.CustomerRepository{}
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAllCustomer() ([]models.Customer, error) {
	err := s.repo.LoadData()
	if err != nil {
		return nil, err
	}
	return s.repo.Data, nil
}

func (s *CustomerService) GetCustomerByID(id string) (*models.Customer, error) {
	err := s.repo.LoadData()
	if err != nil {
		return nil, err
	}
	return s.repo.GetCustomerByID(id)
}

func (s *CustomerService) GetCustomerByEmail(email string) (*models.Customer, error) {
	err := s.repo.LoadData()
	if err != nil {
		return nil, err
	}

	customer, err := s.repo.GetCustomerByEmail(email)
	if err != nil {
		return nil, err
	} 

	return customer, err
}

func (s *CustomerService) Deposit(customerID string, amount float64) error {

	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	err := s.repo.LoadData()
	if err != nil {
		return err
	}

	err = s.repo.UpdateBalance(customerID, amount)
	if err != nil {
		return  err
	}

	return err
}
