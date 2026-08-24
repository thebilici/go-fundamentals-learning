package main

import "fmt"

type User struct{
	Name   string
	Age int
	IsActive bool
}

func main(){

	user:=User{
		Name:"Fatih",
		Age:22,
		IsActive:true,
	}
	
	user2:=User{
		Name:"Ali",
		Age:16,
		IsActive:true,
	}

	fmt.Println("Name:", user.getName())
	fmt.Println("Is Adult:", user.isAdult())
	fmt.Println("Can Access:", user.canAccess(19))
	fmt.Println(user.greet("Hello"))
	fmt.Println(user2.isAdult())

}

func (u User) getName() string{
	return u.Name
}

func(u User) isAdult() bool{
	return u.Age>=18
}

func (u User) canAccess(minAge int) bool{
	return u.Age >= minAge
}

func(u User) greet(message string)string{
	return message + " " + u.Name
}
