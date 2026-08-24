package main	

import "fmt"

type User struct {
	Name string
	Age int 
}

func main() {

	user:=User{
		Name: "Fatih",
		Age: 22,
	}

	user.printName()

	result := user.isAdult()
	fmt.Println("Is Adult:", result)

	user.greet("Hello")

	fmt.Println("Can Access:", user.canAccess(18))
	fmt.Println("Can Access:", user.canAccess(25))

	user.changeName()
	fmt.Println("New Name:", user.Name)//User struct'ı Değişmez pointer'la değişir

}

func(u User) printName(){
	fmt.Println("Name:", u.Name)
}

func (u User) isAdult() bool{
	return u.Age >= 18
}	

func (u User) greet(message string){
	fmt.Println(message, u.Name)
}

func (u User) canAccess(minAge int) bool{
	return u.Age >= minAge
}

func (u User) changeName(){

	u.Name = "Ahmet"
}