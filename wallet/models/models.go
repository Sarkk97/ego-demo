package models

import (
	"fmt"
	"time"
)

const (
	WALLET_CREDIT string = "credit"
	WALLET_DEBIT  string = "debit"
)

//Wallet models a model
type Wallet struct {
	ID        string    `json:"id" gorm:"primary_key;type:varchar(120)"`
	UserID    string    `json:"userId" gorm:"unique_index;type:varchar(120); not null"`
	Balance   int       `json:"balance" gorm:"type:int;not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

//WalletTransaction is a wallet transaction
type WalletTransaction struct {
	ID                string    `json:"id" gorm:"primary_key;type:varchar(120)"`
	WalletID          string    `json:"walletId" gorm:"type:varchar(120); not null"`
	PaystackReference string    `json:"reference" gorm:"index:paystack_reference"`
	PrevBalance       int       `json:"previousBalance" gorm:"type:int; not null"`
	Amount            int       `json:"amount" gorm:"type:int; not null"`
	Balance           int       `json:"balance" gorm:"type:int; not null"`
	Action            string    `json:"action" gorm:"type:varchar(10);not null"`
	Meta              string    `json:"meta" gorm:"type:text;"`
	CreatedAt         time.Time `json:"-"`
}

//TableName sets the default database table name for WalletTransaction records
func (WalletTransaction) TableName() string {
	return "wallet_transactions"
}

//Credit credits a Wallet
func (wallet *Wallet) Credit(amount int) (*Wallet, error) {
	if amount < 0 {
		return nil, fmt.Errorf("Invalid amount to credit wallet: %d", amount)
	}

	wallet.Balance += amount

	return wallet, nil
}

//Debit debits a Wallet
func (wallet *Wallet) Debit(amount int) (*Wallet, error) {
	if amount < 0 {
		return nil, fmt.Errorf("Invalid amount to debit from wallet: %d", amount)
	}

	if amount > wallet.Balance {
		return nil, fmt.Errorf("Insufficient wallet balance: %d", amount)
	}

	wallet.Balance -= amount

	return wallet, nil
}

//HasFundsForTransaction checks if sufficeint funds exist in a wallet
//to carry out a transaction
func (wallet *Wallet) HasFundsForTransaction(amount int) (bool, error) {
	if amount < 0 {
		return false, fmt.Errorf("Invalid amount: %d", amount)
	}

	return wallet.Balance >= amount, nil
}
