package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/capim/api/internal/model"
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
	if err := r.Storage.WithContext(ctx).Preload("Owner").Find(&sellers).Error; err != nil {
		return nil, fmt.Errorf("get all sellers on database: %w", err)
	}

	var modelSellers []model.Seller
	copier.Copy(&modelSellers, &sellers)

	return modelSellers, nil
}

func (r *SellerRepository) GetSellerByID(ctx context.Context, sellerID string) (model.Seller, error) {
	var seller Seller

	if err := r.Storage.WithContext(ctx).Preload("Owner").First(&seller, sellerID).Error; err != nil {
		return model.Seller{}, fmt.Errorf("get seller by ID on database: %w", err)
	}

	var modelSeller model.Seller
	copier.Copy(&modelSeller, &seller)

	return modelSeller, nil
}
