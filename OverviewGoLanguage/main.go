package main

import "log"

type Personnel struct {
	PersonnelID string
	FirstName   string
	LastName    string
}

func main() {
	myPersonnel := make(map[string]Personnel)
	personnel1 := Personnel{
		PersonnelID: "12345",
		FirstName:   "Atakan",
		LastName:    "Tekoğlu",
	}
	myPersonnel["personnel1"] = personnel1
	log.Println(myPersonnel["personnel1"].PersonnelID)
}
