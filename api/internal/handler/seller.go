package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/service"
)

type Owner struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type BankAccount struct {
	BankCode      string `json:"bank_code"`
	AgencyNumber  string `json:"agency_number"`
	AccountNumber string `json:"account_number"`
}

type SellerRequest struct {
	Document     string      `json:"document"`
	LegalName    string      `json:"legal_name"`
	BusinessName string      `json:"business_name"`
	BankAccount  BankAccount `json:"bank_account"`
	Owner        []Owner     `json:"owner"`
}

type SellerHandler struct {
	Service *service.SellerService
}

func (h *SellerHandler) CreateSeller(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()
	var request SellerRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var input model.Seller
	copier.Copy(&input, &request)

	err := h.Service.CreateSeller(context, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to create seller"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Seller created successfully"})
}

func (h *SellerHandler) GetAllSellers(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellers, err := h.Service.GetAllSellers(context)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to retrieve sellers"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sellers retrieved successfully", "sellers": sellers})
}

func (h *SellerHandler) GetSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	seller, err := h.Service.GetSellerByID(context, sellerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to retrieve seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller retrieved successfully", "seller": seller})
}

func (h *SellerHandler) DeleteSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	err := h.Service.DeleteSellerByID(context, sellerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to delete seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller deleted successfully"})
}

func (h *SellerHandler) UpdateSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	var request SellerRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"details": err.Error(), "error": "Invalid request body"})
		return
	}
	var input model.Seller
	copier.Copy(&input, &request)

	err := h.Service.UpdateSellerByID(context, sellerID, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to update seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller updated successfully"})
}

func (h *SellerHandler) UpdateOwnerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	ownerID := ctx.Param("id")

	var request Owner

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"details": err.Error(), "error": "Invalid request body"})
		return
	}
	var input model.Owner
	copier.Copy(&input, &request)

	err := h.Service.UpdateOwnerByID(context, ownerID, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to update owner"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Owner updated successfully"})

}
