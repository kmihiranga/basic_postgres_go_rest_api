package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
		Balance:   int64(rand.Intn(1000)),
	}
}
