package main 
import (
	"fmt"
)

func main(){

	ch:= make(chan string)

	go sendMessage(ch)

	message:= <- ch
	
	fmt.Println(message)


}

func sendMessage(ch chan string){
	ch <- "Hello from goroutine"
}