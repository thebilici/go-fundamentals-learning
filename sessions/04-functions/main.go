package main

import "fmt"

func main() {

	sayHello("Fatih")
	sayHello("Furkan")
	printUser("Davut", 22)

	result := add(5, 10)
	fmt.Println("Result:", result)

	fmt.Println("Result:", add(20, 30))

	fmt.Println("Result2:", multiply(5, 10))

	ageResult := checkAge(17)
	fmt.Println("Age Result:", ageResult)
	getUser("Ali", 25)

	name, _ := getUser("Ali", 25)

	fmt.Println("Name:", name)
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func printUser(name string, age int) {
	fmt.Println("Name:", name, "Age:", age)
}

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func checkAge(age int) string {
	if age >= 18 {
		return "Reşit"
	}
	return "Reşit değil"
}

func getUser(name string, age int) (string, int) {
	return name, age
}
