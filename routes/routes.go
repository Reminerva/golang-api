package routes

import (
	"bank-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authController *controllers.AuthController,
	customerContoller *controllers.CustomerController,
	merchantController *controllers.MerchantController,
	transactionController *controllers.TransactionController,
) *gin.Engine {
	r := gin.Default()
	
	r.GET("/auth", authController.GetAllAuth)
	r.POST("/auth/login", authController.Login)
	r.POST("/auth/register", authController.Register)
	r.POST("/auth/logout", authController.Logout)

	r.GET("/customers", customerContoller.GetAllCustomers)
	r.GET("/customers/id/:id", customerContoller.GetCustomerByID)
	r.GET("/customers/email/:email", customerContoller.GetCustomerByEmail)
	r.POST("/customers/deposit", customerContoller.Deposit)

	r.POST("/merchants", merchantController.CreateMerchant)
	r.GET("/merchants", merchantController.GetAllMerchants)
	r.GET("/merchants/id/:id", merchantController.GetMerchantByID)

	r.POST("/transactions", transactionController.ProcessTransaction)
	r.GET("/transactions", transactionController.GetAllTransactions)
	r.GET("/transactions/customer/:customerID", transactionController.GetTransactionsByCustomerID)
	r.GET("/transactions/merchant/:merchantID", transactionController.GetTransactionsByMerchantID)

	return r
}
