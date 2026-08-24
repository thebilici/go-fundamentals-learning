package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	
	user:= User{
		Name: "Fatih",
		Age:  23,
	}
	
	age := 22

	fmt.Println("Before:", age)
	fmt.Println("Name:",user.Name)

	changeAge(&age)

	fmt.Println("After:", age)
	user.changeName()
	fmt.Println("Name:",user.Name)
}

func changeAge(age *int) {
	*age = 30
}

func (u *User) changeName(){

	u.Name = "Ahmet"
}