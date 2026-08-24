package main

import "fmt"

type User struct{
	Name string
	Age int
}

func main(){

	user:=User{
		Name:"Fatih",
		Age:22,
	}

	age:=22

	fmt.Println("Before:",age)
	fmt.Println("Address",&age)

	agePointer := &age
	fmt.Println("Address of age pointer:", agePointer)

	agePointerValue := *agePointer
	fmt.Println("Value of age pointer:", agePointerValue)

	*agePointer = 25
	fmt.Println("After:",age)
	fmt.Println("Address of age pointer:", agePointer)
	fmt.Println("Value of age pointer:", *agePointer)

	changeAge(&age)
	fmt.Println("After:",age)

	user.changeName("Ahmet")
	fmt.Println("Name:",user.Name)

}

func changeAge(age *int){
	*age = 30
}

func(u *User) changeName(newName string){

	u.Name=newName
}