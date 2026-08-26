package main

import(
	"fmt"
	"context"
	"time"
)

func main(){
	ctx:=context.Background()
	ctx,cancel:=context.WithTimeout(ctx,3*time.Second)
	defer cancel()
	go worker(ctx)

}

func worker(ctx context.Context){
	for{
		select{
		case<-ctx.Done():
			fmt.Println("context canceled",ctx.Err())
			return
		
	default:
			fmt.Println("Processing...")
			time.Sleep(500*time.Millisecond)
	}
}
	
}