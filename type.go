package main

import (
	"math/rand"
)

type Accounts struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
}

func NewAccount(firstName, lastName string) *Accounts {
	return &Accounts{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.Intn(10000000)),
	}
}
