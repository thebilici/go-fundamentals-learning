package main

import "fmt"

func main(){

		prices:=[]float64{120.50, 75.25, 300.00, 50.00}
		discountedTotal, total:=calculateTotal(prices)
		fmt.Println("Total:", total)
		fmt.Println("Discounted Total:", discountedTotal)

}

func calculateTotal(prices []float64) (float64, float64){

	total:=0.0
	discountedTotal:=0.0

	for _,price:=range prices{
		
		total+=price
		discountedTotal=total
		if total>=500{
			discountedTotal=total*0.9
		}
		
	}
	return discountedTotal,total
}