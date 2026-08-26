package main

import(
	"fmt"
	"context"
)

func main(){
	ctx:=context.Background()
	service(ctx)
}


func service(ctx context.Context){
	fmt.Println("service started")
	repository(ctx)
}

func repository(ctx context.Context){
	fmt.Println("repository called")
	
}