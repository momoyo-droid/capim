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

	if err := r.Storage.Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("create seller on database: %w", err)
	}

	return nil
}
