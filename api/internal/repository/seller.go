package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/utils"
	"gorm.io/gorm"
)

// SellerRepository provides methods to interact with the seller data in the database.
// It implements the SellerRepository interface defined in the service package and
// uses GORM for database operations.
type SellerRepository struct {
	Storage *gorm.DB
}

// NewSellerRepository creates a new instance of SellerRepository with the provided GORM database connection.
func NewSellerRepository(storage *gorm.DB) *SellerRepository {
	return &SellerRepository{
		Storage: storage,
	}
}

// CreateSeller creates a new seller record in the database.
// It takes a model.Seller as input and returns an error if the creation fails.
func (r *SellerRepository) CreateSeller(ctx context.Context, seller model.Seller) error {
	model := Seller{
		Document:     seller.Document,
		LegalName:    seller.LegalName,
		BusinessName: seller.BusinessName,
		BankAccount: BankAccount{
			BankCode:      seller.BankAccount.BankCode,
			AgencyNumber:  seller.BankAccount.AgencyNumber,
			AccountNumber: seller.BankAccount.AccountNumber,
		},
	}

	if err := copier.Copy(&model.Owner, &seller.Owner); err != nil {
		return fmt.Errorf("copy seller owners during creation on database: %w", err)
	}

	// Save the seller along with its associated owners in a single transaction
	if err := r.Storage.Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("create seller on database: %w", err)
	}

	return nil
}

// GetAllSellers retrieves all sellers from the database and returns them as a slice of model.Seller.
// It returns an error if there is an issue during retrieval.
func (r *SellerRepository) GetAllSellers(ctx context.Context) ([]model.Seller, error) {
	var sellers []Seller
	// Preload the Owner association to load the related owners for each seller
	result := r.Storage.WithContext(ctx).Preload("Owner").Find(&sellers)
	if result.Error != nil {
		return nil, fmt.Errorf("get all sellers on database: %w", result.Error)
	}

	var modelSellers []model.Seller
	if err := copier.Copy(&modelSellers, &sellers); err != nil {
		return nil, fmt.Errorf("copy sellers during get all from database: %w", err)
	}

	return modelSellers, nil
}

// GetSellerByID retrieves a seller by its ID from the database.
// It returns an error if there is an issue during retrieval or if the seller is not found.
func (r *SellerRepository) GetSellerByID(ctx context.Context, sellerID uint64) (model.Seller, error) {
	var seller Seller
	// Preload the Owner association to load the related owners for each seller
	result := r.Storage.WithContext(ctx).Preload("Owner").First(&seller, sellerID)
	if result.Error != nil {
		return model.Seller{}, fmt.Errorf("get seller by ID on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return model.Seller{}, utils.ErrSellerIDNotFound
	}

	var modelSeller model.Seller
	if err := copier.Copy(&modelSeller, &seller); err != nil {
		return model.Seller{}, fmt.Errorf("copy seller during get by ID from database: %w", err)
	}

	return modelSeller, nil
}

// CheckSellerByDocument checks if a seller with the given document already exists in the database.
// It returns the existing seller if found, or an empty model.Seller if not found.
// It returns an error if there is an issue during the database query.
func (r *SellerRepository) CheckSellerByDocument(ctx context.Context, document string) (model.Seller, error) {
	var seller Seller
	result := r.Storage.WithContext(ctx).Where("document = ?", document).First(&seller)
	if result.Error != nil {
		return model.Seller{}, fmt.Errorf("get seller by document on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return model.Seller{}, nil
	}

	var modelSeller model.Seller
	if err := copier.Copy(&modelSeller, &seller); err != nil {
		return model.Seller{}, fmt.Errorf("copy seller during check by document from database: %w", err)
	}

	return modelSeller, nil
}

// DeleteSellerByID deletes a seller by its ID from the database.
// It returns an error if there is an issue during deletion or if the seller is not found.
func (r *SellerRepository) DeleteSellerByID(ctx context.Context, sellerID uint64) error {
	var seller Seller

	result := r.Storage.WithContext(ctx).Delete(&seller, sellerID)
	if result.Error != nil {
		return fmt.Errorf("delete seller by ID on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return utils.ErrSellerIDNotFound
	}

	return nil
}

// UpdateSellerByID updates a seller's information by its ID in the database.
// It returns an error if there is an issue during the update process or if the seller is not found.
func (r *SellerRepository) UpdateSellerByID(ctx context.Context, sellerID uint64, updatedSeller model.Seller) error {
	var seller Seller
	if err := copier.Copy(&seller, &updatedSeller); err != nil {
		return fmt.Errorf("copy seller during update on database: %w", err)
	}

	result := r.Storage.WithContext(ctx).Where("id = ?", sellerID).Updates(&seller)
	if result.Error != nil {
		return fmt.Errorf("update seller by ID on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return utils.ErrSellerIDNotFound
	}

	return nil
}

// UpdateOwnerByID updates an owner's information by its ID in the database.
// It returns an error if there is an issue during the update process or if the owner is not found.
func (r *SellerRepository) UpdateOwnerByID(ctx context.Context, ownerID uint64, updatedOwner model.Owner) error {
	var owner Owner
	if err := copier.Copy(&owner, &updatedOwner); err != nil {
		return fmt.Errorf("copy owner during update on database: %w", err)
	}

	result := r.Storage.WithContext(ctx).Where("id = ?", ownerID).Updates(&owner)
	if result.Error != nil {
		return fmt.Errorf("update owner by ID on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("owner ID not found")
	}

	return nil
}
