package main

import "fmt"

func main() {
	name := "Steve"
	fmt.Println("Name:", name)
	fmt.Println("Memory address:", &name)
	changeNameUsingPointer(&name)
	fmt.Println("Name:", name)
}

func changeNameUsingPointer(s *string) {
	fmt.Println("Memory address:", s)
	newValue := "James"
	*s = newValue
}
