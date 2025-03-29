package controllers

import (
	"bank-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController() *CustomerController {
	service := services.NewCustomerService()
	return &CustomerController{Service: service}
}

func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := c.Service.GetAllCustomer()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	customerID := ctx.Param("id")
	customer, err := c.Service.GetCustomerByID(customerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) GetCustomerByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	customer, err := c.Service.GetCustomerByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) Deposit(ctx *gin.Context) {
	var request struct {
		CustomerID string  `json:"customerID"`
		Amount     float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := c.Service.Deposit(request.CustomerID, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deposit successful"})
}
