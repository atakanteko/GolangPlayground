package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Soldier",
		Breed: "German Shephers",
	}
	gorilla := Gorilla{
		Name:          "Big Boi",
		Color:         "Gray",
		NumberOfTeeth: 28,
	}
	PrintInfo(&dog)
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animail says", a.Says(), "and has", a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Woooh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
