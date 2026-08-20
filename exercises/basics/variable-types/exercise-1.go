package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string = "Fatih"
	var age int = 22
	height := 1.73
	var isStudent bool = true
	const language string = "Go"
	var numberText string = "25"

	ageFloat := float64(age)

	number, err := strconv.Atoi(numberText)
	ageText := strconv.Itoa(age)

	fmt.Println("Benim adım", name, "yaşım", age, "ve öğrenmekte olduğum programlama dili", language)
	fmt.Println("Boyum", height, "ve öğrenci miyim?", isStudent)
	fmt.Println("Yaşımın float64 türüne dönüştürülmüş hali", ageFloat)
	fmt.Println("String türündeki sayı", numberText, "int türüne dönüştürülmüş hali", number)
	fmt.Println("Int türündeki yaşım", age, "string türüne dönüştürülmüş hali", ageText)
	fmt.Println("Hata", err)

	fmt.Printf("name: %T\n", name)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("height: %T\n", height)
	fmt.Printf("isStudent: %T\n", isStudent)
	fmt.Printf("language: %T\n", language)
	fmt.Printf("numberText: %T\n", numberText)
	fmt.Printf("ageFloat: %T\n", ageFloat)
	fmt.Printf("number: %T\n", number)
	fmt.Printf("ageText: %T\n", ageText)
}
