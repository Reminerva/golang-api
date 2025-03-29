package main

import (
	"bank-api/controllers"
	"bank-api/routes"
)

func main() {
	authController := controllers.NewAuthController()
	customerController := controllers.NewCustomerController()
	merchantController := controllers.NewMerchantController()
	transactionController := controllers.NewTransactionController()

	router := routes.SetupRouter(authController, customerController, merchantController, transactionController)

	router.Run(":8080")
}
