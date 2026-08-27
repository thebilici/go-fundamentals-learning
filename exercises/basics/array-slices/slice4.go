package main

import "fmt"

func main(){
	orders:=[]float64{120, 850, 430, 50, 1200, 300, 760}

	fmt.Println("Orders:",orders)

	smallOrders,largeOrders:=SplitOrders(orders,500)
	fmt.Println("Small Orders:",smallOrders)
	fmt.Println("Large Orders:",largeOrders)

	averageOrder:=CalculateAverageOrder(orders)
	fmt.Println("Average Order:",averageOrder)
}

func SplitOrders(orders []float64,limit float64)([]float64,[]float64){
	var smallOrders []float64
	var largeOrders []float64

	for _,order:=range orders{
		if order<=limit{
			smallOrders=append(smallOrders,order)
		}else{
			largeOrders=append(largeOrders,order)
		}
	}
	return smallOrders,largeOrders
}

func CalculateAverageOrder(orders []float64) float64{

	var totalOrder float64
	for _,order:=range orders{
		totalOrder+=order
	}

	averageOrder:=totalOrder/float64(len(orders))
	return averageOrder
}