package main

import(
	"fmt"
)

type Number interface{
	int | float64 | float32 | int64
}

type Response[T any] struct{
	Data T
	Success bool
}

func main(){

	fmt.Println(sum(10,20))
	fmt.Println(sum(10.5,20.5))

	user:=Response[string]{
		Data: "Fatih",
		Success: true,
	}

	user2:=Response[int]{
		Data: 100,
		Success: true,
	}

	fmt.Println(user)
	fmt.Println(user2)
}

func sum[T Number](a T,b T)T{
	return a+b
}