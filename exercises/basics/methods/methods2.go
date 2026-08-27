package main

import "fmt"

type Product struct{
	Name string
	Stock int
	Price float64
}
func main(){
	product1:=Product{
		Name:"Keyboard",
		Stock:10,
		Price: 1500,
	}

	newStock:=product1.AddStock(5)
	fmt.Println("New Stock:",newStock)

	remainingStock,ok:=product1.Sell(12)
	if ok{
		fmt.Println("Remaining Stock:",remainingStock)
	}else{
		fmt.Println("Sale Failed")
	}
	fmt.Println("Inventory Value:",product1.GetInventoryValue())
}

func (p *Product) AddStock(amount int) int{
	p.Stock+=amount

	return p.Stock
}

func (p *Product) Sell(quantity int) (int,bool){

	if p.Stock<quantity{
		fmt.Println("Insufficient stock")
		return p.Stock,false
	}else{
		p.Stock-=quantity
		return p.Stock,true
	}
}

func (p Product) GetInventoryValue()float64{
	if  p.Stock>0{
		return p.Price*float64(p.Stock)
	}
	return 0
}