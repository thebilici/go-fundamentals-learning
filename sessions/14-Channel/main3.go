package main

import "fmt"

func main() {
	resultChannel := make(chan int)

	go calculate(10, 20, resultChannel)
	go calculate(5, 7, resultChannel)
	go calculate(100, 50, resultChannel)

	result1 := <-resultChannel
	result2 := <-resultChannel
	result3 := <-resultChannel

	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
	fmt.Println("Result 3:", result3)
}

func calculate(a int, b int, ch chan int) {
	result := a + b
	ch <- result
}