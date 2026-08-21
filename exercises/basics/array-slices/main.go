package main

import "fmt"

func main() {

	languages := [3]string{"Go", "Python", "Java"}

	fmt.Println("Languages:", languages)
	fmt.Println("Languages[0]:", languages[0])
	fmt.Println("length of languages:", len(languages))

	languages2 := []string{"Go", "Python", "JavaScript"}

	languages2 = append(languages2, "Javascript", "Rust")

	fmt.Println("Length of languages2:", len(languages2))
	fmt.Println("Capacity of languages2:", cap(languages2))

	selected := languages2[1:4]
	fmt.Println("Selected:", selected)

	for index, value := range languages2 {
		fmt.Println("Index:", index, "Value:", value)
	}

	languages2[0] = "Golang"
	fmt.Println(languages2)

}
