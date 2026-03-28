package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/utils"
)

type SellerRepository interface {
	CreateSeller(ctx context.Context, seller model.Seller) error
	GetAllSellers(ctx context.Context) ([]model.Seller, error)
	GetSellerByID(ctx context.Context, sellerID uint64) (model.Seller, error)
	GetSellerByDocument(ctx context.Context, document string) (bool, error)
	DeleteSellerByID(ctx context.Context, sellerID uint64) error
	UpdateSellerByID(ctx context.Context, sellerID uint64, updatedSeller model.Seller) error
	UpdateOwnerByID(ctx context.Context, ownerID uint64, updatedOwner model.Owner) error
}

type SellerService struct {
	Repository SellerRepository
}

func NewSellerService(repository SellerRepository) *SellerService {
	return &SellerService{
		Repository: repository,
	}
}

func (s *SellerService) CreateSeller(ctx context.Context, seller model.Seller) error {
	if err := validateSeller(seller); err != nil {
		return fmt.Errorf("validate seller error: %w", err)
	}

	// Check if the record already exists in the database
	existingSeller, err := s.Repository.GetSellerByDocument(ctx, seller.Document)
	if err != nil {
		return fmt.Errorf("check existing seller error: %w", err)
	}

	if existingSeller {
		return fmt.Errorf("a seller with the same document already exists")
	}

	if err := s.Repository.CreateSeller(ctx, seller); err != nil {
		return fmt.Errorf("create seller error: %w", err)
	}

	return nil
}

func (s *SellerService) GetAllSellers(ctx context.Context) ([]model.Seller, error) {
	sellers, err := s.Repository.GetAllSellers(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all sellers error: %w", err)
	}

	return sellers, nil
}

func (s *SellerService) GetSellerByID(ctx context.Context, sellerID string) (model.Seller, error) {
	id, err := validateID(sellerID)
	if err != nil {
		return model.Seller{}, utils.ErrInvalidID
	}

	seller, err := s.Repository.GetSellerByID(ctx, id)
	if err != nil {
		return model.Seller{}, fmt.Errorf("get seller by ID error: %w", err)
	}

	return seller, nil
}

func (s *SellerService) DeleteSellerByID(ctx context.Context, sellerID string) error {
	id, err := validateID(sellerID)
	if err != nil {
		return utils.ErrInvalidID
	}

	err = s.Repository.DeleteSellerByID(ctx, id)
	if err != nil {
		return fmt.Errorf("delete seller by ID error: %w", err)
	}

	return nil
}

func (s *SellerService) UpdateSellerByID(ctx context.Context, sellerID string, updatedSeller model.Seller) error {
	id, err := validateID(sellerID)
	if err != nil {
		return utils.ErrInvalidID
	}

	err = s.Repository.UpdateSellerByID(ctx, id, updatedSeller)
	if err != nil {
		return fmt.Errorf("update seller by ID error: %w", err)
	}

	return nil
}

func (s *SellerService) UpdateOwnerByID(ctx context.Context, ownerID string, updatedOwner model.Owner) error {
	id, err := validateID(ownerID)
	if err != nil {
		return utils.ErrInvalidID
	}

	err = s.Repository.UpdateOwnerByID(ctx, id, updatedOwner)
	if err != nil {
		return fmt.Errorf("update owner by ID error: %w", err)
	}

	return nil
}

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
