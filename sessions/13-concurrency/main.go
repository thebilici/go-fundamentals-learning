package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()

	fmt.Println("All goroutines finished")
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()

	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {
		fmt.Println("Letter:", letter)
		time.Sleep(500 * time.Millisecond)
	}
}