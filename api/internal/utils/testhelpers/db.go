package testhelpers

import (
	"testing"

	"github.com/momoyo-droid/capim/api/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(
		&repository.Seller{},
		&repository.Owner{},
		&repository.BankAccount{},
	)
	assert.NoError(t, err)

	return db
}
