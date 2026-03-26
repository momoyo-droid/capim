package model

type Owner struct {
	Name  string
	Phone string
	Email string
}

type BankAccount struct {
	BankCode      string
	AgencyNumber  string
	AccountNumber string
}

type Seller struct {
	Document     string
	LegalName    string
	BusinessName string
	BankAccount  BankAccount
	Owner        []Owner
}
