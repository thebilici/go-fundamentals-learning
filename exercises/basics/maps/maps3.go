package main

import "fmt"

func main(){
	products:=map[string]int{
	"Mouse":    10,
	"Keyboard": 5,
	"Monitor":  2,
	}

	fmt.Println("Products:",products)

	products=updateStocks(products,"Mouse",5)
	fmt.Println("Products after update:",products)

	products=removeProduct(products,"Keyboard")
	fmt.Println("Products after removal:",products)

	_,quantity,ok:=getStock(products,"Monitor")
	if ok==true{
		fmt.Println("Monitor stock quantity:",quantity)
	}else{
		fmt.Println("Monitor stock not found")
	}

}

func updateStocks(products map[string]int,name string,quantity int)map[string]int{
	updatedQuantity:=products[name]+quantity
	products[name]=updatedQuantity

	return products
}

func removeProduct(products map[string]int,name string)map[string]int{
	delete(products,name)

	return products
}

func getStock(products map[string]int,name string)(name string,quantity int,ok bool){

	quantity,ok=products[name]

	return name,quantity,ok

}
