package main

import (
	"fmt"
	"context"
	"time"
)
//cancel işlemi
func main(){

	ctx:=context.Background()//ctx.Background() ile boş bir context oluşturuyoruz

	ctx,cancel:=context.WithCancel(ctx) //iptal sinyali taşıyan context

	go doWork(ctx)

	time.Sleep(2*time.Second)

	cancel()

	time.Sleep(1*time.Second)
	
}

func doWork(ctx context.Context){
	for{
		select{
		case<-ctx.Done()://iptal sinyalini taşıyan channel
			fmt.Println("context canceled")
			return
		

		default:
			fmt.Println("working...")
			time.Sleep(1*time.Second)//main hemen bitmesin diye 1 saniye bekletiyoruz Done çalıştığını görmek için 
	}
}
}