package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	input, err := validateInputRequest(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Service.CreateSeller(context, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create seller"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Seller created successfully"})
}

func (h *SellerHandler) GetAllSellers(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellers, err := h.Service.GetAllSellers(context)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sellers"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sellers retrieved successfully", "sellers": sellers})
}

func (h *SellerHandler) GetSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	id, err := validateSellerID(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	seller, err := h.Service.GetSellerByID(context, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller retrieved successfully", "seller": seller})
}

func (h *SellerHandler) DeleteSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	id, err := validateSellerID(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	err = h.Service.DeleteSellerByID(context, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller deleted successfully"})
}

func (h *SellerHandler) UpdateSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	id, err := validateSellerID(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var request SellerRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var input model.Seller
	copier.Copy(&input, &request)

	err = h.Service.UpdateSellerByID(context, id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller updated successfully"})
}

func (h *SellerHandler) UpdateOwnerByID(ctx *gin.Context) {
	context := ctx.Request.Context()
	defer ctx.Request.Body.Close()

	sellerID := ctx.Param("id")

	id, err := validateSellerID(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var request Owner

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var input model.Owner
	copier.Copy(&input, &request)

	err = h.Service.UpdateOwnerByID(context, id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Owner updated successfully"})

}

func validateSellerID(sellerID string) (uint64, error) {
	if sellerID == "" {
		return 0, fmt.Errorf("seller ID is required")
	}

	id, err := strconv.ParseUint(sellerID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("seller ID must be a valid integer")
	}

	return id, nil
}

func validateInputRequest(request SellerRequest) (model.Seller, error) {
	if request.Document == "" || request.LegalName == "" || request.BusinessName == "" {
		return model.Seller{}, fmt.Errorf("document, legal_name and business_name are required")
	}

	input := model.Seller{
		Document:     request.Document,
		LegalName:    request.LegalName,
		BusinessName: request.BusinessName,
		BankAccount: model.BankAccount{
			BankCode:      request.BankAccount.BankCode,
			AgencyNumber:  request.BankAccount.AgencyNumber,
			AccountNumber: request.BankAccount.AccountNumber,
		},
	}

	copier.Copy(&input.Owner, &request.Owner)

	return input, nil
}
