package controllers

import (
	"bank-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	Service *services.TransactionService
}

func NewTransactionController() *TransactionController {
	Service := services.NewTransactionService()
	return &TransactionController{Service: Service}
}

func (t *TransactionController) ProcessTransaction(ctx *gin.Context) {
	var request struct {
		CustomerID string  `json:"customerID"`
		MerchantID string  `json:"merchantID"`
		Amount     float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := t.Service.ProcessTransaction(request.CustomerID, request.MerchantID, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction successful"})
}

func (t *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := t.Service.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

func (t *TransactionController) GetTransactionsByCustomerID(ctx *gin.Context) {
	customerID := ctx.Param("customerID")
	transactions, err := t.Service.GetTransactionsByCustomerID(customerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

func (t *TransactionController) GetTransactionsByMerchantID(ctx *gin.Context) {
	merchantID := ctx.Param("merchantID")
	transactions, err := t.Service.GetTransactionsByMerchantID(merchantID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}
