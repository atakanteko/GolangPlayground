package main

import "fmt"

type Personnel struct {
	PersonnelID string
	FirstName   string
	LastName    string
}

type App []string

func (personnel *Personnel) printPersonnelInformation() {
	fmt.Println(personnel.PersonnelID, personnel.LastName, personnel.FirstName)
}

func (a App) print() {
	for _, x := range a {
		fmt.Println(x)
	}
}

func main() {
	var personnel Personnel
	personnel.PersonnelID = "12345"
	personnel.FirstName = "Steve"
	personnel.LastName = "Diaz"
	personnel.printPersonnelInformation()
	apps := App{"Facebook", "Instagram", "WhatsApp"}
	apps.print()
}
