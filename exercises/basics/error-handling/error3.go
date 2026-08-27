package main

import (
	"fmt"
	errors "errors"
)

type Product struct{
	Name string
	Stock int
	Price float64
}

func main(){

	products := []Product{
	{Name: "Keyboard", Stock: 10, Price: 1500},
	{Name: "Mouse", Stock: 5, Price: 750},
	{Name: "Monitor", Stock: 2, Price: 6000},
}

totalPrice, err := ProccessOrder(products,"Mouse",3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Total Price: %.2f\n", totalPrice)
	}

}

func SellProduct(products []Product, name string, quantity int)(float64,error){
	if quantity <=0{
		return 0,errors.New("Invalid quantity. Quantity must be greater than zero")
	}

	for _, product := range products {
		if product.Name == name {
			if product.Stock < quantity {
				return 0,errors.New("Insufficient stock")
			}
			productPrice := product.Price * float64(quantity)
			return productPrice, nil
		}
	}
	return 0, errors.New("Product not found")
}

func ProccessOrder(products []Product,name string,quantity int)(float64,error){
	
	totalPrice,err:=SellProduct(products,name,quantity)

	if err!=nil{
		return 0,fmt.Errorf("order processing failed: %w", err)
	}
	return totalPrice,nil
}
