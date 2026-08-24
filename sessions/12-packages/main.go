package main

import(
	"fmt"
	"github.com/thebilici/go-backend-learning/sessions/12-packages/calculator"
)

func main(){
	result:= calculator.Add(10,5)
	fmt.Println("Result:",result)

}

