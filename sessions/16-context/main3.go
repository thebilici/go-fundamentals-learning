package main

import (
	"fmt"
	"context"
	"time"
)
//withDeadline kullanıyoruz çünkü belirli bir süre sonra işlemi iptal etmek istiyoruz. Bu, uzun süren işlemler için faydalıdır ve kaynakları serbest bırakmamıza yardımcı olur.
func main(){
	ctx:=context.Background()
	deadline:=time.Now().Add(2*time.Second)
	ctx,cancel:=context.WithDeadline(ctx,deadline)

	defer cancel()
	doWork(ctx)
}

func doWork(ctx context.Context){

	select{
		for{
		case<-ctx.Done():
			fmt.Println("context canceled",ctx.Err())
			return
		}

	default:
		fmt.Println("working...")
		time.Sleep(500*time.Millisecond)

	}

}