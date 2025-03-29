package services

import (
	"bank-api/models"
	"bank-api/repository"
	"errors"
	"fmt"
)

type TransactionService struct {
	customerRepo *repository.CustomerRepository
	transactionRepo *repository.TransactionRepository
	merchantRepo *repository.MerchantRepository
	authRepo *repository.AuthRepository
}

func NewTransactionService() *TransactionService {
	customerRepo := &repository.CustomerRepository{}
	transactionRepo := &repository.TransactionRepository{}
	merchantRepo := &repository.MerchantRepository{}
	authRepo := &repository.AuthRepository{}

	return &TransactionService{
		customerRepo: customerRepo,
		transactionRepo: transactionRepo,
		merchantRepo: merchantRepo,
		authRepo: authRepo,
	}
}

func (s *TransactionService) ProcessTransaction(customerID, merchantID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	
	errCustomer := s.customerRepo.LoadData()
	errAuth := s.authRepo.LoadData()
	errMerchant := s.merchantRepo.LoadData()
	if errCustomer != nil {
		return errCustomer
	}
	if errAuth != nil {
		return errAuth
	}
	if errMerchant != nil {
		return errMerchant
	}

	customer, errCustomer := s.customerRepo.GetCustomerByID(customerID)
	if errCustomer != nil {
		return errCustomer
	}

	loggedInUser, errAuth := s.authRepo.GetCustomerByEmail(customer.Email)
	fmt.Println("LoggedInUser:", loggedInUser)
	if errAuth != nil {
		return errAuth
	}
	
	merchant, errMerchant := s.merchantRepo.GetMerchantByID(merchantID)
	if errMerchant != nil {
		return errMerchant
	}
	fmt.Println("Merchant:", merchant)

	if customer.Balance < amount {
		return errors.New("insufficient balance")
	}

	s.customerRepo.UpdateBalance(customer.ID, -amount)

	s.merchantRepo.UpdateBalance(merchant.ID, amount)

	err := s.transactionRepo.CreateTransaction(customerID, merchantID, amount)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	err := s.transactionRepo.LoadData()
	if err != nil {
		return nil, err
	}
	return s.transactionRepo.Data, nil
}

func (s *TransactionService) GetTransactionsByCustomerID(customerID string) ([]models.Transaction, error) {
	err := s.transactionRepo.LoadData()
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for _, transaction := range s.transactionRepo.Data {
		if transaction.CustomerID == customerID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}

func (s *TransactionService) GetTransactionsByMerchantID(merchantID string) ([]models.Transaction, error) {
	err := s.transactionRepo.LoadData()
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for _, transaction := range s.transactionRepo.Data {
		if transaction.MerchantID == merchantID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}

func (s *TransactionService) GetTransactionsByID(customerID, merchantID string) ([]models.Transaction, error) {
	err := s.transactionRepo.LoadData()
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for _, transaction := range s.transactionRepo.Data {
		if transaction.CustomerID == customerID && transaction.MerchantID == merchantID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}
