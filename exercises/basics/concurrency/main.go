package main

import (
	"fmt"
	"sync"
)

func main(){

	var wg sync.WaitGroup
	wg.Add(3)
	results:= make(chan int,3)

	go CalculateSquare(2,results,&wg)
	go CalculateSquare(4,results,&wg)
	go CalculateSquare(6,results,&wg)

	go func(){
		wg.Wait()
		close(results)
	}()

	for result:=range results{
		fmt.Println("Result:",result)
	}
}

func CalculateSquare(a int,results chan int,wg *sync.WaitGroup){
	defer wg.Done()
	results<- a*a
}