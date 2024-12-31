package main

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

func main() {
	var account BankAccount
	fmt.Println(account)
	account.AccountName = "123456789"

	account.DisplayBalance()

	acc := BankAccount{AccountName: "123456789", Balance: 1000.00}
	fmt.Println(acc.Balance)

	// account2 := &BankAccount{AccountName: "987654321", Balance: 5000.00}

	account3 := BankAccount{AccountName: "987654321", Balance: 5000.00, AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()}}
	fmt.Println(account3.CreatedAt)

	account.DisplayBalance()
	account.Deposit(1000)
	account.DisplayBalance()

	account4 := NewBankAccount("12345", 00.00)
	fmt.Println(account4)

	var acct AccountNumber
	fmt.Printf("%T\n", acct)

	acct2 := BankAccount{AccountName: "9876543210", Balance: 5000.00, AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()}}
	fmt.Println(acct2.AccountName.isValid())

	var balance Balance
	fmt.Printf("%T\n",balance)
}
