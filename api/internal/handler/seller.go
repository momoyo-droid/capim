package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/service"
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
}

// CreateSeller handles the HTTP POST request to create a new seller.
// It decodes the request body into a SellerRequest struct, copies the data to a model.Seller struct,
// and calls the CreateSeller method of the SellerService.
// It returns appropriate HTTP responses based on the success or failure of the operation.
func (h *SellerHandler) CreateSeller(ctx *gin.Context) {
	context := ctx.Request.Context()
	var request SellerRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var input model.Seller
	if err := copier.Copy(&input, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy seller data"})
		return
	}

	err := h.Service.CreateSeller(context, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to create seller"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Seller created successfully"})
}

// GetAllSellers handles the HTTP GET request to retrieve all sellers.
// It calls the GetAllSellers method of the SellerService and returns the list of sellers in the response.
// If there is an error during retrieval, it returns an appropriate HTTP error response.
func (h *SellerHandler) GetAllSellers(ctx *gin.Context) {
	context := ctx.Request.Context()

	sellers, err := h.Service.GetAllSellers(context)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to retrieve sellers"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sellers retrieved successfully", "sellers": sellers})
}

// GetSellerByID handles the HTTP GET request to retrieve a seller by its ID.
// It extracts the seller ID from the request parameters, calls the GetSellerByID method of the SellerService,
// and returns the seller information in the response. If there is an error during retrieval, it returns an appropriate HTTP error response.
func (h *SellerHandler) GetSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	sellerID := ctx.Param("id")

	seller, err := h.Service.GetSellerByID(context, sellerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to retrieve seller"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seller retrieved successfully", "seller": seller})
}

// DeleteSellerByID handles the HTTP DELETE request to delete a seller by its ID.
// It extracts the seller ID from the request parameters, calls the DeleteSellerByID method of
// the SellerService, and returns an appropriate HTTP response based on the success or
// failure of the operation.
func (h *SellerHandler) DeleteSellerByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	sellerID := ctx.Param("id")

	err := h.Service.DeleteSellerByID(context, sellerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to delete seller"})
		return
	}

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
		ctx.JSON(http.StatusBadRequest, gin.H{"details": err.Error(), "error": "Invalid request body"})
		return
	}
	var input model.Seller
	if err := copier.Copy(&input, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy seller data"})
		return
	}

	err := h.Service.UpdateSellerByID(context, sellerID, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to update seller"})
		return
	}

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
		ctx.JSON(http.StatusBadRequest, gin.H{"details": err.Error(), "error": "Invalid request body"})
		return
	}
	var input model.Owner
	if err := copier.Copy(&input, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy owner data"})
		return
	}

	err := h.Service.UpdateOwnerByID(context, ownerID, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"details": err.Error(), "error": "Failed to update owner"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Owner updated successfully"})

}
