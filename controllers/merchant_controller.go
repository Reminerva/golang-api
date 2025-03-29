package controllers

import (
	"bank-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	Service *services.MerchantService
}

func NewMerchantController() *MerchantController {
	service := services.NewMerchantService()
	return &MerchantController{Service: service}
}

func (s *MerchantController) CreateMerchant(ctx *gin.Context) {
	var request struct {
		Name     string  `json:"name"`
		Address  string  `json:"address"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := s.Service.CreateMerchant(request.Name, request.Address, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Merchant created successfully"})
}

func (s *MerchantController) GetAllMerchants(ctx *gin.Context) {
	merchants, err := s.Service.GetAllMerchant()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, merchants)
}

// GetMerchantByID menangani request untuk mendapatkan data merchant berdasarkan ID
func (m *MerchantController) GetMerchantByID(ctx *gin.Context) {
	merchantID := ctx.Param("id")
	merchant, err := m.Service.GetMerchantByID(merchantID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, merchant)
}

