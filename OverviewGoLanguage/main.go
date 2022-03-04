package main

import "fmt"

func main() {
	var myFirstMessage string
	var myNumber int
	myFirstMessage = "Hello, Go Lang. This is my first message"
	myNumber = 22
	fmt.Println(myFirstMessage)
	fmt.Printf("My age is %d\n", myNumber)
	fmt.Println(saySomethingWithParameter("Learn Go"))

	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)
}
func saySomething() string {
	return "something"
}

func saySomethingWithParameter(message string) string {
	return message
}

