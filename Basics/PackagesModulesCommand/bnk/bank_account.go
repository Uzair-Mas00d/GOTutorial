package bnk

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	CreatedAt    time.Time
	LastModified time.Time
}

type BankAccount struct {
	AccountName AccountNumber
	Balance     float64
	name string
	AuditInfo
}

func (ba BankAccount) DisplayBalance() {
	fmt.Printf("Account Number %s Balance: %.2f\n", ba.AccountName, ba.Balance)
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.Balance < amount {
		return errors.New("insuffient Fund")
	}

	ba.Balance -= amount
	return nil
}

func NewBankAccount(num AccountNumber, bal float64) *BankAccount {
	return &BankAccount{
		AccountName: num,
		Balance:     bal,
		AuditInfo:   AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
}

func (a AccountNumber) isValid() bool {
	return len(string(a)) == 10
}

type AccountNumber string
type Balance = float64
