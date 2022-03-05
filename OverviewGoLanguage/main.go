package main

import (
	"fmt"
	"github/atakanteko/golangplayground/helpers"
)

func main() {
	var myVar []helpers.SomeType

	myVar = append(myVar, helpers.SomeType{
		TypeName:   "Gamers",
		TypeNumber: 1,
	}, helpers.SomeType{
		TypeName:   "Developers",
		TypeNumber: 2,
	})
	fmt.Println(myVar)
	for _, info := range myVar {
		fmt.Println("Type Name:", info.TypeName, "Type Number:", info.TypeNumber)
	}
}
