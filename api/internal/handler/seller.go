package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/service"
	"go.uber.org/zap"
)

// Owner struct represents the owner of a seller, including their name, phone number,
// and email address.
type Owner struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// BankAccount struct represents the bank account information associated with a seller,
// including bank code, agency number, and account number.
type BankAccount struct {
	BankCode      string `json:"bank_code"`
	AgencyNumber  string `json:"agency_number"`
	AccountNumber string `json:"account_number"`
}

// SellerRequest struct represents the request body for creating or updating a seller,
//
//	including its fields and relationships.
type SellerRequest struct {
	Document     string      `json:"document"`
	LegalName    string      `json:"legal_name"`
	BusinessName string      `json:"business_name"`
	BankAccount  BankAccount `json:"bank_account"`
	Owner        []Owner     `json:"owner"`
}

// SellerHandler struct is responsible for handling HTTP requests related to sellers.
// It contains a reference to the SellerService,
// which provides the business logic for managing sellers.
type SellerHandler struct {
	Service *service.SellerService
	Logger  *zap.Logger
}

// CreateSeller handles the HTTP POST request to create a new seller.
// It decodes the request body into a SellerRequest struct, copies the data to a model.Seller struct,
// and calls the CreateSeller method of the SellerService.
// It returns appropriate HTTP responses based on the success or failure of the operation.
func (h *SellerHandler) CreateSeller(ctx *gin.Context) {
	context := ctx.Request.Context()
	var request SellerRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		h.Logger.Error("Failed to decode request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var input model.Seller
	if err := copier.Copy(&input, &request); err != nil {
		h.Logger.Error("Failed to copy seller data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy seller data"})
		return
	}

	h.Logger.Info("Creating seller")

	err := h.Service.CreateSeller(context, input)
	if err != nil {
		h.Logger.Error("Failed to create seller")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create seller"})
		return
	}

	h.Logger.Info("Seller created successfully")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Seller created successfully"})
}

// GetAllSellers handles the HTTP GET request to retrieve all sellers.
// It calls the GetAllSellers method of the SellerService and returns the list of sellers in the response.
// If there is an error during retrieval, it returns an appropriate HTTP error response.
func (h *SellerHandler) GetAllSellers(ctx *gin.Context) {
	context := ctx.Request.Context()
	h.Logger.Info("Retrieving all sellers")
	sellers, err := h.Service.GetAllSellers(context)
	if err != nil {
		h.Logger.Error("Failed to retrieve sellers")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sellers"})
		return
	}
	h.Logger.Info("Sellers retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Sellers retrieved successfully", "sellers": sellers})
}

// GetSellerByID handles the HTTP GET request to retrieve a seller by its ID.
// It extracts the seller ID from the request parameters, calls the GetSellerByID method of the SellerService,
// and returns the seller information in the response. If there is an error during retrieval, it returns an appropriate HTTP error response.
func (h *SellerHandler) GetSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	sellerID := ctx.Param("id")

	h.Logger.Info("Retrieving seller with ID")

	seller, err := h.Service.GetSellerByID(context, sellerID)
	if err != nil {
		h.Logger.Error("Failed to retrieve seller")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve seller"})
		return
	}
	h.Logger.Info("Seller retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Seller retrieved successfully", "seller": seller})
}

// DeleteSellerByID handles the HTTP DELETE request to delete a seller by its ID.
// It extracts the seller ID from the request parameters, calls the DeleteSellerByID method of
// the SellerService, and returns an appropriate HTTP response based on the success or
// failure of the operation.
func (h *SellerHandler) DeleteSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	sellerID := ctx.Param("id")
	h.Logger.Info("Deleting seller with ID")
	err := h.Service.DeleteSellerByID(context, sellerID)
	if err != nil {
		h.Logger.Error("Failed to delete seller")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete seller"})
		return
	}
	h.Logger.Info("Seller deleted successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Seller deleted successfully"})
}

// UpdateSellerByID handles the HTTP PUT request to update a seller's information by its ID.
// It extracts the seller ID from the request parameters, decodes the request body into a SellerRequest struct,
// copies the data to a model.Seller struct, and calls the UpdateSellerByID method of the SellerService.
// It returns appropriate HTTP responses based on the success or failure of the operation.
func (h *SellerHandler) UpdateSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	sellerID := ctx.Param("id")

	var request SellerRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		h.Logger.Error("Failed to decode request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var input model.Seller
	if err := copier.Copy(&input, &request); err != nil {
		h.Logger.Error("Failed to copy seller data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy seller data"})
		return
	}
	h.Logger.Info("Updating seller with ID")
	err := h.Service.UpdateSellerByID(context, sellerID, input)
	if err != nil {
		h.Logger.Error("Failed to update seller")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seller"})
		return
	}
	h.Logger.Info("Seller updated successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Seller updated successfully"})
}

// UpdateOwnerByID handles the HTTP PUT request to update an owner's information by its ID.
// It extracts the owner ID from the request parameters, decodes the request body into an Owner struct,
// copies the data to a model.Owner struct, and calls the UpdateOwnerByID method of the SellerService.
// It returns appropriate HTTP responses based on the success or failure of the operation.
func (h *SellerHandler) UpdateOwnerByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	ownerID := ctx.Param("id")

	var request Owner

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		h.Logger.Error("Failed to decode request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var input model.Owner
	if err := copier.Copy(&input, &request); err != nil {
		h.Logger.Error("Failed to copy owner data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy owner data"})
		return
	}
	h.Logger.Info("Updating owner with ID")
	err := h.Service.UpdateOwnerByID(context, ownerID, input)
	if err != nil {
		h.Logger.Error("Failed to update owner")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner"})
		return
	}
	h.Logger.Info("Owner updated successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Owner updated successfully"})

}
