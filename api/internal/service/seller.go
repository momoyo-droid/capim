package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/utils"
	"go.uber.org/zap"
)

// SellerRepository defines the interface for interacting with the seller data in the database.
//
//go:generate moq -out service/mocks/seller_repository_mock.go -pkg mocks . SellerRepository
type SellerRepository interface {
	CreateSeller(ctx context.Context, seller model.Seller) error
	GetAllSellers(ctx context.Context) ([]model.Seller, error)
	GetSellerByID(ctx context.Context, sellerID uint64) (model.Seller, error)
	CheckSellerByDocument(ctx context.Context, document string) (model.Seller, error)
	DeleteSellerByID(ctx context.Context, sellerID uint64) error
	UpdateSellerByID(ctx context.Context, sellerID uint64, updatedSeller model.Seller) error
	UpdateOwnerByID(ctx context.Context, ownerID uint64, updatedOwner model.Owner) error
}

// SellerService provides methods to manage sellers, including creating, retrieving, updating,
// and deleting seller records.
// It interacts with the SellerRepository to perform database operations and includes validation logic to ensure data integrity.
type SellerService struct {
	Repository SellerRepository
	Logger     *zap.Logger
}

// NewSellerService creates a new instance of SellerService with the provided SellerRepository.
func NewSellerService(repository SellerRepository, logger *zap.Logger) *SellerService {
	return &SellerService{
		Repository: repository,
		Logger:     logger,
	}
}

// CreateSeller validates the input seller data and checks for existing records before creating a new seller in the database.
// It returns an error if validation fails, if a seller with the same document already exists, or if there is an issue during creation.
func (s *SellerService) CreateSeller(ctx context.Context, seller model.Seller) error {
	s.Logger.Info("Creating new seller")
	if err := validateSeller(seller); err != nil {
		return fmt.Errorf("validate seller error: %w", err)
	}

	// Check if the record already exists in the database
	sellerDocument, err := s.Repository.CheckSellerByDocument(ctx, seller.Document)
	if err != nil {
		return fmt.Errorf("check existing seller error: %w", err)
	}

	if sellerDocument.Document == seller.Document {
		return fmt.Errorf("a seller with the same document already exists")
	}
	s.Logger.Info("Seller validation passed, creating seller in the database")
	if err := s.Repository.CreateSeller(ctx, seller); err != nil {
		return fmt.Errorf("create seller error: %w", err)
	}

	return nil
}

// GetAllSellers retrieves all sellers from the database and returns them as a slice of model.Seller.
// It returns an error if there is an issue during retrieval.
func (s *SellerService) GetAllSellers(ctx context.Context) ([]model.Seller, error) {
	s.Logger.Info("Retrieving all sellers from the database")
	sellers, err := s.Repository.GetAllSellers(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all sellers error: %w", err)
	}

	return sellers, nil
}

// GetSellerByID retrieves a seller by its ID from the database.
// It validates the input ID and returns an error if the ID is invalid, if the seller is not found, or
// if there is an issue during retrieval.
func (s *SellerService) GetSellerByID(ctx context.Context, sellerID string) (model.Seller, error) {
	id, err := validateID(sellerID)
	if err != nil {
		return model.Seller{}, utils.ErrInvalidID
	}
	s.Logger.Info("Retrieving seller by ID from the database")
	seller, err := s.Repository.GetSellerByID(ctx, id)
	if err != nil {
		return model.Seller{}, fmt.Errorf("get seller by ID error: %w", err)
	}

	return seller, nil
}

// DeleteSellerByID deletes a seller by its ID from the database.
// It validates the input ID and returns an error if the ID is invalid, if the seller is not found, or
// if there is an issue during deletion.
func (s *SellerService) DeleteSellerByID(ctx context.Context, sellerID string) error {
	id, err := validateID(sellerID)
	if err != nil {
		return utils.ErrInvalidID
	}
	s.Logger.Info("Deleting seller by ID from the database")
	err = s.Repository.DeleteSellerByID(ctx, id)
	if err != nil {
		return fmt.Errorf("delete seller by ID error: %w", err)
	}

	return nil
}

// UpdateSellerByID updates a seller's information by its ID in the database.
// It validates the input ID and returns an error if the ID is invalid, if the seller is not found, or
// if there is an issue during the update process.
func (s *SellerService) UpdateSellerByID(ctx context.Context, sellerID string, updatedSeller model.Seller) error {
	id, err := validateID(sellerID)
	if err != nil {
		return utils.ErrInvalidID
	}
	s.Logger.Info("Updating seller by ID in the database")
	err = s.Repository.UpdateSellerByID(ctx, id, updatedSeller)
	if err != nil {
		return fmt.Errorf("update seller by ID error: %w", err)
	}

	return nil
}

// UpdateOwnerByID updates an owner's information by its ID in the database.
// It validates the input ID and returns an error if the ID is invalid, if the owner is not found, or
// if there is an issue during the update process.
func (s *SellerService) UpdateOwnerByID(ctx context.Context, ownerID string, updatedOwner model.Owner) error {
	id, err := validateID(ownerID)
	if err != nil {
		return utils.ErrInvalidID
	}
	s.Logger.Info("Updating owner by ID in the database")
	err = s.Repository.UpdateOwnerByID(ctx, id, updatedOwner)
	if err != nil {
		return fmt.Errorf("update owner by ID error: %w", err)
	}

	return nil
}

// validateSeller checks if the provided seller data is valid, ensuring that required fields are present and properly formatted.
// It returns an error if any validation checks fail.
func validateSeller(seller model.Seller) error {
	if seller.Document == "" || seller.LegalName == "" || seller.BusinessName == "" {
		return fmt.Errorf("document, legal name, and business name are required")
	}

	if len(seller.Owner) == 0 {
		return fmt.Errorf("at least one owner is required")
	}

	for index, owner := range seller.Owner {
		if owner.Name == "" || owner.Phone == "" || owner.Email == "" {
			return fmt.Errorf("owner at index %d is missing required fields", index)
		}
	}

	if seller.BankAccount.BankCode == "" || seller.BankAccount.AgencyNumber == "" || seller.BankAccount.AccountNumber == "" {
		return fmt.Errorf("bank account information is required")
	}

	return nil
}

// validateID checks if the provided ID string is a valid unsigned integer and returns it as uint64.
// It returns an error if the ID is empty or cannot be parsed as a valid uint64.
func validateID(idParam string) (uint64, error) {
	if idParam == "" {
		return 0, utils.ErrInvalidID
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, utils.ErrInvalidID
	}

	return id, nil
}
