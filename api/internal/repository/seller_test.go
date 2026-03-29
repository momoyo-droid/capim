package repository_test

import (
	"context"
	"testing"

	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/repository"
	"github.com/momoyo-droid/capim/api/internal/utils"
	"github.com/momoyo-droid/capim/api/internal/utils/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestSellerRepository_CreateSeller_Success(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	ctx := context.Background()
	repo := repository.NewSellerRepository(db)

	seller := model.Seller{
		Document:     "123456789",
		LegalName:    "Test LTDA",
		BusinessName: "Test LTDA ME",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "1234",
			AccountNumber: "56789-0",
		},
		Owner: []model.Owner{
			{
				Name:  "test",
				Phone: "123",
				Email: "test@example.com",
			},
		},
	}

	err := repo.CreateSeller(ctx, seller)
	assert.NoError(t, err)
}

func TestSellerRepository_GetAllSellers_Success(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	_ = repo.CreateSeller(ctx, model.Seller{
		Document:     "123",
		LegalName:    "Teste",
		BusinessName: "Teste",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "123",
			AccountNumber: "999",
		},
		Owner: []model.Owner{
			{Name: "Ana", Phone: "1", Email: "a@a.com"},
		},
	})

	sellers, err := repo.GetAllSellers(ctx)

	assert.NoError(t, err)
	assert.Len(t, sellers, 1)
}

func TestSellerRepository_GetSellerByID_Success(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	seller := model.Seller{
		Document:     "123",
		LegalName:    "Teste",
		BusinessName: "Teste",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "123",
			AccountNumber: "999",
		},
		Owner: []model.Owner{
			{Name: "Ana", Phone: "1", Email: "a@a.com"},
		},
	}

	_ = repo.CreateSeller(ctx, seller)

	result, err := repo.GetSellerByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, "123", result.Document)
}

func TestSellerRepository_GetSellerByID_Failure(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	_, err := repo.GetSellerByID(ctx, 999)

	assert.Error(t, err)
	assert.EqualError(t, utils.ErrSellerIDNotFound, "seller not found")
}

func TestSellerRepository_DeleteSellerByID_Success(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	_ = repo.CreateSeller(ctx, model.Seller{
		Document:     "123",
		LegalName:    "Teste",
		BusinessName: "Teste",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "123",
			AccountNumber: "999",
		},
		Owner: []model.Owner{
			{Name: "Ana", Phone: "1", Email: "a@a.com"},
		},
	})

	err := repo.DeleteSellerByID(ctx, 1)

	assert.NoError(t, err)
}

func TestSellerRepository_DeleteSellerByID_Failure(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	err := repo.DeleteSellerByID(ctx, 999)

	assert.Error(t, err)
	assert.EqualError(t, utils.ErrSellerIDNotFound, "seller not found")
}

func TestSellerRepository_UpdateSellerByID_Success(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	_ = repo.CreateSeller(ctx, model.Seller{
		Document:     "123",
		LegalName:    "Teste",
		BusinessName: "Teste",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "123",
			AccountNumber: "999",
		},
		Owner: []model.Owner{
			{Name: "Ana", Phone: "1", Email: "a@a.com"},
		},
	})

	err := repo.UpdateSellerByID(ctx, 1, model.Seller{
		LegalName: "Novo Nome",
	})

	assert.NoError(t, err)
}

func TestSellerRepository_UpdateSellerByID_Failured(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	err := repo.UpdateSellerByID(ctx, 999, model.Seller{
		LegalName: "Novo Nome",
	})

	assert.Error(t, err)
	assert.EqualError(t, utils.ErrSellerIDNotFound, "seller not found")
}

func TestSellerRepository_UpdateOwnerByID_Success(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	_ = repo.CreateSeller(ctx, model.Seller{
		Document:     "123",
		LegalName:    "Teste",
		BusinessName: "Teste",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "123",
			AccountNumber: "999",
		},
		Owner: []model.Owner{
			{Name: "Ana", Phone: "1", Email: "a@a.com"},
		},
	})

	err := repo.UpdateOwnerByID(ctx, 1, model.Owner{
		Name:  "Novo Nome",
		Phone: "2",
		Email: "novo@email.com",
	})

	assert.NoError(t, err)
}

func TestSellerRepository_UpdateOwnerByID_Failure(t *testing.T) {
	db := testhelpers.SetupTestDB(t)
	repo := repository.NewSellerRepository(db)

	ctx := context.Background()

	err := repo.UpdateOwnerByID(ctx, 999, model.Owner{
		Name: "Novo",
	})

	assert.Error(t, err)
}
