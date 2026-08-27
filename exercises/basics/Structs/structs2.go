package main

import "fmt"

type Order struct{
	ID int
	Customer string	
	Total float64
	IsPaid bool
}

func main(){
	orders:=[]Order{
		createOrder(1,"Fatih",750,true),
		createOrder(2,"Ali",300,false),
		createOrder(3,"Ayşe",1200,true),
		createOrder(4,"Veli",40,false),
	}

	paidOrders:=filterPaidOrder(orders)
	fmt.Println("Paid Orders:",PaidOrders)

	order,found:=findOrderByID(orders,3)
	if found{
		fmt.Println("Order found:",order)
	}else{
		fmt.Println("Order not found")
	}

	totalRevenue:=calculatePaidReveneu(orders)
	fmt.Println("Total revenue from paid orders:",totalRevenue)

}

func createOrder(id int,customer string,total float64,isPaid bool)Order{
	return Order{
		ID:id,
		Customer:customer,
		Total:total,
		IsPaid:isPaid,
	}
}

func filterPaidOrder(orders []Order)[]Order{
	var paidOrders []Order
	for _,order:=range orders{
		if order.IsPaid==true{
			paidOrders=append(paidOrders,order)
		}
	}
	return paidOrders
}

func findOrderByID(orders []Order,id int)(Order,bool){
	for _,order:=range orders{
		if order.ID==id{
			return order,true
		}
	}
	return Order{},false
}

func calculatePaidReveneu(orders []Order)float64{
	var totalRevenue float64=0
	for _,order:=range orders{
		if order.IsPaid==true{
			totalRevenue+=order.Total
		}
	}
	return totalRevenue
}