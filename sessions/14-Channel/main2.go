package main

import "fmt"
func main(){
	ch:= make(chan string,3)

	ch<-"Hello"
	ch<-"World"
	ch<-"Fatih"

	message1:= <-ch
	message2:=<- ch
	message3:=<- ch

	capacity:= cap(ch)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)

}