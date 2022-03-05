package main

import (
	"fmt"
	"github/atakanteko/golangplayground/helpers"
)

const numPool = 100

func CalculateFunction(c chan int) {
	randomNumber := helpers.RandomValue(numPool)
	c <- randomNumber
}

func main() {
	channelExample := make(chan int)
	defer close(channelExample)

	go CalculateFunction(channelExample)
	num := <-channelExample
	fmt.Println(num)
}
