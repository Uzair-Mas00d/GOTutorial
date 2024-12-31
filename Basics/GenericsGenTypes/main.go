package main

import (
	"fmt"
	"slices"
)

type BankAccount struct {
	accountNumber string
	balance       float64
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type Pair[K,V any] struct {
	KEY K
	VALUE V
}

func (p Pair[K, V]) Describe() string {
	return fmt.Sprintf("Key %v Value %v", p.KEY, p.VALUE)
}

func main() {
	bodyTypes := []string{"Sedan", "SUV", "Convertible", "Coupe"}
	fmt.Println(bodyTypes)

	bodyTypes = slices.DeleteFunc(bodyTypes, func(s string) bool {
		if s == "Convertible" {
			return true
		} else {
			return false
		}
	})

	fmt.Println(bodyTypes)

	accounts := []BankAccount{
		{accountNumber: "123", balance: 1000},
		{accountNumber: "456", balance: 3000},
		{accountNumber: "789", balance: 6000},
	}
	fmt.Println(accounts)

	accounts = slices.DeleteFunc(accounts, func(ba BankAccount) bool {
		return ba.accountNumber == "789"
	})
	fmt.Println(accounts)

	intStringPair := Pair[int, string]{KEY: 1, VALUE: "One"}
	fmt.Println(intStringPair.Describe())

	stringFloatPair := Pair[string, float64]{KEY: "Pi", VALUE: 3.1415}
	fmt.Println(stringFloatPair.Describe())
}
