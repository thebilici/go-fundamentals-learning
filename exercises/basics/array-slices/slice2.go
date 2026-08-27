package main

import "fmt"

func main(){
	stocks:=[]int{12, 0, 5, 0, 23, 3, 8, 0, 15}

	fmt.Println("Stocks:",stocks)

	availableStocks:=filterAvailableStocks(stocks)
	fmt.Println("Available Stocks:",availableStocks)

	lowStocks:=filterLowStocks(stocks,5)
	fmt.Println("Low Stocks:",lowStocks)
}

func filterAvailableStocks(stocks []int)[]int{
	var stocksAvailable []int
	for _,stock:=range stocks{
		if stock>0{
			stocksAvailable=append(stocksAvailable,stock)
		}
	}
	return stocksAvailable
}

func filterLowStocks(stocks []int,limit int)[]int{
	var stocksLow []int
	for _,stock:=range stocks{
		if stock<=limit && stock>0{
			stocksLow=append(stocksLow,stock)
		}
	}
	return stocksLow
}