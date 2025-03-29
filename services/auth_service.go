package services

import (
	"bank-api/models"
	"bank-api/repository"
	"errors"
	"math/rand"
	"slices"
	"strconv"
)

type AuthService struct {
	authRepo *repository.AuthRepository
	customerRepo *repository.CustomerRepository
}

// Constructor untuk service
func NewAuthService() *AuthService {
	authRepo := &repository.AuthRepository{}
	customerRepo := &repository.CustomerRepository{}
	return &AuthService{authRepo: authRepo, customerRepo: customerRepo}
}

func (s *AuthService) Login(email, password string) (*models.Auth, error) {
	errAuth := s.authRepo.LoadData()
	errCustomer := s.customerRepo.LoadData()

	if errAuth != nil {
		return nil, errAuth
	}

	if errCustomer != nil {
		return nil, errCustomer
	}

	for _, auth := range s.authRepo.Data {
		if auth.Email == email {
			return nil, errors.New("your account has already logged in")
		}
	}

	newAuth := models.Auth{
		Email:    email,
		Password: password,
	}

	for _, customer := range s.customerRepo.Data {
		if customer.Email == email && customer.Password == password {
			s.authRepo.Data = append(s.authRepo.Data, newAuth)
			s.authRepo.SaveData()
			return &newAuth, nil
		}
	}

	return nil, errors.New("invalid email or password")
}

func (s *AuthService) Register(name, email, password string, balance float64) error {
	err := s.customerRepo.LoadData()
	if err != nil {
		return err
	}

	for _, customer := range s.customerRepo.Data {
		if customer.Email == email {
			return errors.New("email already exists")
		}
	}

	var newID string
	for {
		newID = "c-" + strconv.Itoa(rand.Intn(10000))
		_, err := s.customerRepo.GetCustomerByID(newID)
		if err != nil {
			break
		}
	}

	newCustomer := models.Customer{
		ID:       newID,
		Name:     name,
		Email:    email,
		Password: password,
		Balance:  balance,
	}
	s.customerRepo.Data = append(s.customerRepo.Data, newCustomer)

	err = s.customerRepo.SaveData()
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Logout(email string) error {
	err := s.authRepo.LoadData()
	if err != nil {
		return err
	}

	for i, auth := range s.authRepo.Data {
		if auth.Email == email {
			s.authRepo.Data = slices.Delete(s.authRepo.Data, i, i+1)
			return s.authRepo.SaveData()
		}
	}

	return errors.New("email has not logged in")
}

func (s *AuthService) GetAllAuth() ([]models.Auth, error) {
	err := s.authRepo.LoadData()
	if err != nil {
		return nil, err
	}
	return s.authRepo.Data, nil
}

func (s *AuthService) GetAuthByEmail(email string) (*models.Auth, error) {
	err := s.authRepo.LoadData()
	if err != nil {
		return nil, err
	}
	for _, auth := range s.authRepo.Data {
		if auth.Email == email {
			return &auth, nil
		}
	}
	return nil, errors.New("auth not found")
}