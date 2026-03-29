package repository

import (
	"time"

	"gorm.io/gorm"
)

// Audit struct is used to track the creation, update, and deletion timestamps for database records.
type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Owner struct represents the owner of a seller, including their name, phone number,
// and email address.
type Owner struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	SellerID uint   `gorm:"column:seller_id"`
	Name     string `gorm:"column:name"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
}

// BankAccount struct represents the bank account information associated with a seller,
// including bank code, agency number, and account number.
type BankAccount struct {
	BankCode      string `gorm:"column:bank_code"`
	AgencyNumber  string `gorm:"column:agency_number"`
	AccountNumber string `gorm:"column:account_number"`
}

// Seller struct represents the seller entity in the database, including its fields and relationships.
type Seller struct {
	ID           uint        `gorm:"primaryKey;autoIncrement"`
	Document     string      `gorm:"column:document;uniqueIndex"`
	LegalName    string      `gorm:"column:legal_name"`
	BusinessName string      `gorm:"column:business_name"`
	BankAccount  BankAccount `gorm:"embedded"`
	Owner        []Owner     `gorm:"foreignKey:SellerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Audit        Audit       `gorm:"embedded"`
}
