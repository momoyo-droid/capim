package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/utils"
	"gorm.io/gorm"
)

type SellerRepository struct {
	Storage *gorm.DB
}

func NewSellerRepository(storage *gorm.DB) *SellerRepository {
	return &SellerRepository{
		Storage: storage,
	}
}

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

	copier.Copy(&model.Owner, &seller.Owner)
	// Save the seller along with its associated owners in a single transaction
	if err := r.Storage.Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("create seller on database: %w", err)
	}

	return nil
}

func (r *SellerRepository) GetAllSellers(ctx context.Context) ([]model.Seller, error) {
	var sellers []Seller
	// Preload the Owner association to load the related owners for each seller
	result := r.Storage.WithContext(ctx).Preload("Owner").Find(&sellers)
	if result.Error != nil {
		return nil, fmt.Errorf("get all sellers on database: %w", result.Error)
	}

	var modelSellers []model.Seller
	copier.Copy(&modelSellers, &sellers)

	return modelSellers, nil
}

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
	copier.Copy(&modelSeller, &seller)

	return modelSeller, nil
}

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
	copier.Copy(&modelSeller, &seller)

	return modelSeller, nil
}

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

func (r *SellerRepository) UpdateSellerByID(ctx context.Context, sellerID uint64, updatedSeller model.Seller) error {
	var seller Seller
	copier.Copy(&seller, &updatedSeller)

	result := r.Storage.WithContext(ctx).Where("id = ?", sellerID).Updates(&seller)
	if result.Error != nil {
		return fmt.Errorf("update seller by ID on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return utils.ErrSellerIDNotFound
	}

	return nil
}

func (r *SellerRepository) UpdateOwnerByID(ctx context.Context, ownerID uint64, updatedOwner model.Owner) error {
	var owner Owner
	copier.Copy(&owner, &updatedOwner)

	result := r.Storage.WithContext(ctx).Where("id = ?", ownerID).Updates(&owner)
	if result.Error != nil {
		return fmt.Errorf("update owner by ID on database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("owner ID not found")
	}

	return nil
}
