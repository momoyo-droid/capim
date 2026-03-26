package service

import (
	"context"
	"fmt"

	"github.com/momoyo-droid/capim/api/internal/model"
)

type SellerRepository interface {
	CreateSeller(ctx context.Context, seller model.Seller) error
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
	if err := s.Repository.CreateSeller(ctx, seller); err != nil {
		return fmt.Errorf("create seller error: %w", err)
	}

	return nil
}
