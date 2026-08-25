package main

import "fmt"

func main(){

	numbers:= []int{1,2,3,4,5}
	names:= []string{"Fatih","Ali","Veli"}

	fmt.Println("First number:",getFirst(numbers))
	fmt.Println("First name:",getFirst(names))
}

func getFirst[T any](arr []T) T {
	return arr[0]
}