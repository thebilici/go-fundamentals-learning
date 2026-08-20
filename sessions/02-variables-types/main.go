package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Fatih"
	surname := "Bilici"
	height := 1.75
	var age float64 = 22.3
	var language string = "Go"
	var score float64 = 99.5
	var isLearning bool = true
	var lakap string
	var weight float32
	const appName string = "Go Backend Learning"
	const pi = 3.14
	bmiText := "22"
	iqInteger := 200
	letter := 'A'

	fmt.Println("Benim adım", name, "yaşım", age, "ve öğrenmekte olduğum programlama dili", language)
	fmt.Println("Sınavdan aldığım puan", score, "ve öğreniyor muyum?", isLearning)
	fmt.Println("Boyum", height, "ve soyadım", surname)
	fmt.Println("Lakabım", lakap)
	fmt.Println("Ağırlığım", weight)
	fmt.Println("Uygulama adı", appName)
	fmt.Println("Pi sayısı", pi)

	age = 25.3
	convertedAge := int(age)
	fmt.Println("Yaşım", age)
	fmt.Println("Yaşımın int türüne dönüştürülmüş hali", convertedAge)

	bmi, err := strconv.Atoi(bmiText)
	iq := strconv.Itoa(iqInteger)

	fmt.Println("Bmi değeri", bmi)
	fmt.Println("Hata", err)
	fmt.Println("IQ değeri", iq)

	fmt.Printf("name: %T\n", name)
	fmt.Printf("Age: %T\n", age)
	fmt.Printf("Int Yaş: %T\n", convertedAge)
	fmt.Printf("Height: %T\n", height)
	fmt.Printf("Letter: %T\n", letter)
}
