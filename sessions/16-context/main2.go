package main

import (
	"context"
	"fmt"
	"time"
)
//withTimeout kullanıyoruz çünkü belirli bir süre sonra işlemi iptal etmek istiyoruz. Bu, uzun süren işlemler için faydalıdır ve kaynakları serbest bırakmamıza yardımcı olur.

func main2() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	doWork2(ctx)

}

func doWork2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled", ctx.Err())
			return

		default:
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}
