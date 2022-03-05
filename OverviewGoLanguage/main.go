package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Atakan",
		LastName:    "Tekoğlu",
		PhoneNumber: "0 5xx xxx xx xx",
	}
	log.Println(user)
}
