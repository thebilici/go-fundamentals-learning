package main

import "fmt"

func main(){

	numbers:= []int{1,2,3,4,5}
	names:= []string{"Fatih","Ali","Veli"}

	fmt.Println("First number:",getFirst(numbers))
	fmt.Println("First name:",getFirst(names))
}

func contains[T comparable](arr []T,target T) bool {

	for _,item:=range arr{
		if item==target{
			return true
		}
	}
	return false
}