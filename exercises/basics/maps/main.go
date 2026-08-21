package main

import "fmt"

func main() {
	scores := map[string]int{
		"Ali":   90,
		"Ahmet": 75,
		"Mehmet": 60,
	}

	fmt.Println(scores)

	scores["Ayşe"] = 85
	scores["Mehmet"] = 70

	score, ok := scores["Ahmet"]
	if ok {
		fmt.Println("Ahmet's score:", score)
	} else {
		fmt.Println("Ahmet's score not found")
	}

	score, ok = scores["Ali"]
	if ok {
		fmt.Println("Ali's score:", score)
	} else {
		fmt.Println("Ali's score not found")
	}

	delete(scores, "Mehmet")
	len := len(scores)
	fmt.Println("Number of students:", len)

	fmt.Println("Student count:", len(scores))

	for key, value := range scores {
		fmt.Println("Name:", key, "Score:", value)
	}
}