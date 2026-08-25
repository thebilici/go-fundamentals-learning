package main

import "fmt"

func main(){


	fmt.Println(printPair("Name","Fatih"))
	fmt.Println(printPair(1,100))
}

func printPair[K any,V any](key K,value V) bool {
	fmt.Println("Key:",key,"Value:",value)
	return true
}