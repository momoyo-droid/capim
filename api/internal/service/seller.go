package service

import (
	"context"
	"fmt"

	"github.com/momoyo-droid/capim/api/internal/model"
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
	if seller.Document == "" || seller.LegalName == "" || seller.BusinessName == "" {
		return fmt.Errorf("document, legal name, and business name are required")
	}

	if len(seller.Owner) == 0 {
		return fmt.Errorf("at least one owner is required")
	}

	if seller.BankAccount.BankCode == "" || seller.BankAccount.AgencyNumber == "" || seller.BankAccount.AccountNumber == "" {
		return fmt.Errorf("bank account information is required")
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

func (s *SellerService) GetSellerByID(ctx context.Context, sellerID uint64) (model.Seller, error) {
	seller, err := s.Repository.GetSellerByID(ctx, sellerID)
	if err != nil {
		return model.Seller{}, fmt.Errorf("get seller by ID error: %w", err)
	}

	return seller, nil
}

func (s *SellerService) DeleteSellerByID(ctx context.Context, sellerID uint64) error {
	err := s.Repository.DeleteSellerByID(ctx, sellerID)
	if err != nil {
		return fmt.Errorf("delete seller by ID error: %w", err)
	}

	return nil
}

func (s *SellerService) UpdateSellerByID(ctx context.Context, sellerID uint64, updatedSeller model.Seller) error {
	err := s.Repository.UpdateSellerByID(ctx, sellerID, updatedSeller)
	if err != nil {
		return fmt.Errorf("update seller by ID error: %w", err)
	}

	return nil
}

func (s *SellerService) UpdateOwnerByID(ctx context.Context, ownerID uint64, updatedOwner model.Owner) error {
	err := s.Repository.UpdateOwnerByID(ctx, ownerID, updatedOwner)
	if err != nil {
		return fmt.Errorf("update owner by ID error: %w", err)
	}

	return nil
}
