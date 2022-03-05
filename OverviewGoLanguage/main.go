package main

import "log"

type Personnel struct {
	PersonnelID string
	FirstName   string
	LastName    string
}

func main() {
	var personnels []Personnel
	personnel1 := Personnel{
		PersonnelID: "12345",
		FirstName:   "Steve",
		LastName:    "Bonac",
	}
	personnel2 := Personnel{
		PersonnelID: "98765",
		FirstName:   "William",
		LastName:    "Heath",
	}
	personnelList := append(personnels, personnel1, personnel2)
	log.Println(personnelList)

}
