package main

import "fmt"

func main() {
	n := 3

	out := make(chan int)
	// We want to run a goroutine to multiply n by 2
	go multiplyByTwo(n, out)
	defer close(out)
	result := <-out
	fmt.Println("Result:", result)
}

func multiplyByTwo(num int, c chan<- int) {
	result := num * 2
	c <- result
}
