package main

import (
	"math/rand"
)

type Accounts struct {
	ID int
	FirstName string
	LastName string
	Number int64
	Balance int64
}

func NewAccount(firstName, lastName string) *Accounts {
	return &Accounts{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.Intn(10000000)),
	}
}
