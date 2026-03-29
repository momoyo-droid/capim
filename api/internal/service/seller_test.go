package service_test

import (
	"context"
	"testing"

	"github.com/momoyo-droid/capim/api/internal/model"
	"github.com/momoyo-droid/capim/api/internal/service"
	"github.com/momoyo-droid/capim/api/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSellerService_CreateSeller_Success(t *testing.T) {
	ctx := context.Background()

	seller := &model.Seller{
		Document:     "12345678000199",
		LegalName:    "Empresa LTDA",
		BusinessName: "Empresa Teste",
		BankAccount: model.BankAccount{
			BankCode:      "001",
			AgencyNumber:  "1234",
			AccountNumber: "56789-0",
		},
		Owner: []model.Owner{
			{
				Name:  "Test da Silva",
				Phone: "11999999999",
				Email: "test@empresa.com",
			},
		},
	}

	repository := &mocks.SellerRepositoryMock{
		CheckSellerByDocumentFunc: func(ctx context.Context, document string) (model.Seller, error) {
			return model.Seller{}, nil
		},
		CreateSellerFunc: func(ctx context.Context, seller model.Seller) error {
			return nil
		},
	}
	logger, _ := zap.NewDevelopment()
	service := service.NewSellerService(repository, logger)

	err := service.CreateSeller(ctx, *seller)

	assert.NoError(t, err)

}

func TestSellerService_CreateSeller_Failure(t *testing.T) {
	type fields struct {
		Repository *mocks.SellerRepositoryMock
	}

	type args struct {
		seller model.Seller
		ctx    context.Context
	}

	tests := []struct {
		name          string
		expectedError string
		args          args
		fields        fields
	}{
		{
			name: "Should return error when Seller required fields are missing",
			args: args{
				seller: model.Seller{
					Document:     "",
					LegalName:    "Abacate LTDA",
					BusinessName: "Abacate Teste",
				},
				ctx: context.Background(),
			},
			fields: fields{
				Repository: &mocks.SellerRepositoryMock{},
			},
			expectedError: "validate seller error: document, legal name, and business name are required",
		},
		{
			name: "Should return error when Seller has no owners",
			args: args{
				seller: model.Seller{
					Document:     "12345678000199",
					LegalName:    "Abacate LTDA",
					BusinessName: "Abacate Teste",
					Owner:        []model.Owner{},
				},
				ctx: context.Background(),
			},
			fields: fields{
				Repository: &mocks.SellerRepositoryMock{},
			},
			expectedError: "validate seller error: at least one owner is required",
		},
		{
			name: "Should return error when an owner is missing required fields",
			args: args{
				seller: model.Seller{
					Document:     "12345678000199",
					LegalName:    "Abacate LTDA",
					BusinessName: "Abacate Teste",
					Owner: []model.Owner{
						{
							Name:  "",
							Phone: "11999999999",
							Email: "",
						},
					},
				},
				ctx: context.Background(),
			},
			fields: fields{
				Repository: &mocks.SellerRepositoryMock{},
			},
			expectedError: "validate seller error: owner at index 0 is missing required fields",
		},
		{
			name: "Should return error when bank account is missing required fields",
			args: args{
				seller: model.Seller{
					Document:     "12345678000199",
					LegalName:    "Abacate LTDA",
					BusinessName: "Abacate Teste",
					Owner: []model.Owner{
						{
							Name:  "Test da Silva",
							Phone: "11999999999",
							Email: "test@empresa.com",
						},
					},
					BankAccount: model.BankAccount{
						BankCode:      "",
						AgencyNumber:  "",
						AccountNumber: "",
					},
				},
				ctx: context.Background(),
			},
			fields: fields{
				Repository: &mocks.SellerRepositoryMock{},
			},
			expectedError: "validate seller error: bank account information is required",
		},
		{
			name: "Should return error when a seller with the same document already exists",
			args: args{
				seller: model.Seller{
					Document:     "12345678000199",
					LegalName:    "Abacate LTDA",
					BusinessName: "Abacate Teste",
					Owner: []model.Owner{
						{
							Name:  "Test da Silva",
							Phone: "11999999999",
							Email: "test@empresa.com",
						},
					},
					BankAccount: model.BankAccount{
						BankCode:      "001",
						AgencyNumber:  "1234",
						AccountNumber: "56789-0",
					},
				},
				ctx: context.Background(),
			},
			fields: fields{
				Repository: &mocks.SellerRepositoryMock{
					CheckSellerByDocumentFunc: func(ctx context.Context, document string) (model.Seller, error) {
						return model.Seller{Document: "12345678000199"}, nil
					},
				},
			},
			expectedError: "a seller with the same document already exists",
		},
	}
	logger, _ := zap.NewDevelopment()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := service.NewSellerService(tt.fields.Repository, logger)

			err := service.CreateSeller(tt.args.ctx, tt.args.seller)

			assert.EqualError(t, err, tt.expectedError)
		})
	}
}

func TestSellerService_GetAllSellers_Success(t *testing.T) {
	ctx := context.Background()

	expected := []model.Seller{
		{Document: "123"},
	}

	repo := &mocks.SellerRepositoryMock{
		GetAllSellersFunc: func(ctx context.Context) ([]model.Seller, error) {
			return expected, nil
		},
	}
	logger, _ := zap.NewDevelopment()

	service := service.NewSellerService(repo, logger)

	result, err := service.GetAllSellers(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSellerService_GetSellerByID_Success(t *testing.T) {
	ctx := context.Background()

	expected := model.Seller{
		Document: "123",
	}

	repo := &mocks.SellerRepositoryMock{
		GetSellerByIDFunc: func(ctx context.Context, id uint64) (model.Seller, error) {
			return expected, nil
		},
	}
	logger, _ := zap.NewDevelopment()

	service := service.NewSellerService(repo, logger)

	result, err := service.GetSellerByID(ctx, "1")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSellerService_DeleteSellerByID_Success(t *testing.T) {
	ctx := context.Background()

	repo := &mocks.SellerRepositoryMock{
		DeleteSellerByIDFunc: func(ctx context.Context, id uint64) error {
			return nil
		},
	}
	logger, _ := zap.NewDevelopment()

	service := service.NewSellerService(repo, logger)

	err := service.DeleteSellerByID(ctx, "1")

	assert.NoError(t, err)
}

func TestSellerService_UpdateSellerByID_Success(t *testing.T) {
	ctx := context.Background()

	updated := model.Seller{
		Document: "123",
	}

	repo := &mocks.SellerRepositoryMock{
		UpdateSellerByIDFunc: func(ctx context.Context, id uint64, s model.Seller) error {
			return nil
		},
	}
	logger, _ := zap.NewDevelopment()

	service := service.NewSellerService(repo, logger)

	err := service.UpdateSellerByID(ctx, "1", updated)

	assert.NoError(t, err)
}

func TestSellerService_UpdateOwnerByID_Success(t *testing.T) {
	ctx := context.Background()

	updated := model.Owner{
		Name:  "Test da Silva",
		Phone: "11999999999",
		Email: "novo@email.com",
	}

	repo := &mocks.SellerRepositoryMock{
		UpdateOwnerByIDFunc: func(ctx context.Context, id uint64, o model.Owner) error {
			return nil
		},
	}
	logger, _ := zap.NewDevelopment()

	service := service.NewSellerService(repo, logger)

	err := service.UpdateOwnerByID(ctx, "1", updated)

	assert.NoError(t, err)
}
