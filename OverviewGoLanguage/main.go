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
	var numbers []int

	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	log.Println(numbers)

	otherNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(otherNumbers)
}
