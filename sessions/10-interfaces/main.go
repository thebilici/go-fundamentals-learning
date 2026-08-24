package main

import "fmt"

type Greeter interface {
	greet() string
	getName() string
}

type User struct {
	Name string
}

type Admin struct {
	Name string
}

func main() {

	user := User{
		Name: "Fatih",
	}

	admin := Admin{
		Name: "Furkan",
	}

	/*
		var greeter Greeter  type üzerinden variable oluşturma
		greeter = user   user greet() string methodunu implement ettiği için interface'e atanabilir
		fmt.Println(greeter.greet())   interface'ler variable olarak da tanımlanabilir user metodu çalışır
	*/

	printGreeting(user)
	printGreeting(admin)
}

func (u User) greet() string {
	return "Hello, " + u.Name
}

func (u User) getName() string {
	return u.Name
}

func (a Admin) greet() string {
	return "Welcome,Admin " + a.Name
}

func (a Admin) getName() string {
	return a.Name
}

func printGreeting(g Greeter) {
	fmt.Println("Name:", g.getName())
	fmt.Println("Greeting:", g.greet())

}
