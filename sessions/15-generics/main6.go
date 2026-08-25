package main

import "fmt"

type Number interface{
	int | int64 | float32 | float64
}

func main(){

	fmt.Println(Add(10,20))
	fmt.Println(Add(10.5,20.5))
	fmt.Println(Multiply(10,20))

}

func Add(a,b Number)Number{
	return a+b
}

func Multiply(a,b Number) Number{
	return a*b
}