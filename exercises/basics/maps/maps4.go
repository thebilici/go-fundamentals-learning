package main

import "fmt"

func main() {
	products := map[string]float64{
		"Mouse":    750,
		"Keyboard": 1200,
		"Monitor":  6500,
		"Headset":  1800,
	}

	ok := increasePrice(products, "Mouse", 10)

	if ok {
		fmt.Println("Mouse price increased successfully")
	} else {
		fmt.Println("Mouse price increase failed")
	}

	expensiveProducts := filterExpensiveProducts(products, 2000)
	fmt.Println("Expensive products:", expensiveProducts)

	totalValue := calculateTotalValue(products)
	fmt.Println("Total value of products:", totalValue)
}

func increasePrice(products map[string]float64, name string, percentage float64) bool {
	price, ok := products[name]

	if !ok {
		return false
	}

	products[name] = price + (price * percentage / 100)

	return true
}

func filterExpensiveProducts(products map[string]float64, limit float64) map[string]float64 {
	expensiveProducts := make(map[string]float64)

	for name, price := range products {
		if price >= limit {
			expensiveProducts[name] = price
		}
	}

	return expensiveProducts
}

func calculateTotalValue(products map[string]float64) float64 {
	totalValue := 0.0

	for _, price := range products {
		totalValue += price
	}

	return totalValue
}