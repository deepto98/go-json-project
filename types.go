package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

//Func to create new account
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		// ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(1000000),
		Balance:   rand.Float64() * 10000,
		CreatedAt: time.Now().UTC(),
	}
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}
