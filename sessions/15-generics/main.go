package main

import "fmt"

func main() {
	fmt.Println(isEqual(10, 10))
	fmt.Println(isEqual(10, 20))

	fmt.Println(isEqual("Go", "Go"))
	fmt.Println(isEqual("Go", "Java"))
}

func isEqual[T comparable](a T, b T) bool {
	return a == b
}