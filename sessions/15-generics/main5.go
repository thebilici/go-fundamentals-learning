package main

import "fmt"

type UserResponse[T any] struct{
	data T
	success bool
}
func main(){

	user:=UserResponse[string]{
		data: "Fatih",
		success: true,
	}

	fmt.Println(user)

	user2:=UserResponse[int]{
		data: 100,
		success: true,
	}
}

