package main

import (
	"fmt"
	mathpkg "github.com/thebilici/go-backend-learning/exercises/basics/packages-final/mathutil"
)	

func main() {
	sum := mathpkg.Add(5, 6)
	fmt.Println(sum)

	product := mathpkg.Multiply(4, 5)
	fmt.Println(product)

	isPositive := mathpkg.IsPositive(-5)
	fmt.Println(isPositive)
}