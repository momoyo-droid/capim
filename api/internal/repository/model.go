package repository

import (
	"time"

	"gorm.io/gorm"
)

type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Owner struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	SellerID uint   `gorm:"column:seller_id"`
	Name     string `gorm:"column:name"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
}

type BankAccount struct {
	BankCode      string `gorm:"column:bank_code"`
	AgencyNumber  string `gorm:"column:agency_number"`
	AccountNumber string `gorm:"column:account_number"`
}

type Seller struct {
	ID           uint        `gorm:"primaryKey;autoIncrement"`
	Document     string      `gorm:"column:document;uniqueIndex"`
	LegalName    string      `gorm:"column:legal_name"`
	BusinessName string      `gorm:"column:business_name"`
	BankAccount  BankAccount `gorm:"embedded"`
	Owner        []Owner     `gorm:"foreignKey:SellerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Audit        Audit       `gorm:"embedded"`
}
