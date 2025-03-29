package controllers

import (
	"bank-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController() *AuthController {
	Service := services.NewAuthService()
	return &AuthController{Service: Service}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	account, err := c.Service.Login(request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": account})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := c.Service.Register(request.Name, request.Email, request.Password, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := c.Service.Logout(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (c *AuthController) GetAllAuth(ctx *gin.Context) {
	auths, err := c.Service.GetAllAuth()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, auths)
}