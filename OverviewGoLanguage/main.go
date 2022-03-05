package main

import (
	"log"
)

type User struct {
	FirstName string
}

func (i *User) printFirstName() string {
	return i.FirstName
}
func main() {
	var myUser User
	myUser.FirstName = "Kemal"
	user := User{
		FirstName: "Atakan",
	}
	log.Println(myUser.printFirstName())
	log.Println(user.printFirstName())
}
