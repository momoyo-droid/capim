package model

// Owner struct represents the owner of a seller, including their name, phone number,
// and email address.
type Owner struct {
	Name  string
	Phone string
	Email string
}

// BankAccount struct represents the bank account information associated with a seller,
// including bank code, agency number, and account number.
type BankAccount struct {
	BankCode      string
	AgencyNumber  string
	AccountNumber string
}

// Seller struct represents the seller entity, including its fields and relationships.
type Seller struct {
	Document     string
	LegalName    string
	BusinessName string
	BankAccount  BankAccount
	Owner        []Owner
}
