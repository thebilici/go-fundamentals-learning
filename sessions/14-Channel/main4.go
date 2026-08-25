package main
import ("fmt")

func main(){

	var wg sync.WaitGroup

	ch:=make(chan int)

	wg.Add(3)

	go calculate(10, 20, ch, &wg)
	go calculate(5, 7, ch, &wg)
	go calculate(100, 50, ch, &wg)

	go func(){
		wg.Wait()
		close(ch)
	}()
	
	for result:=range ch{
		fmt.Println("Result:",result)	
	}

}

func calculate(a int,b int,ch chan int,wg *sync.WaitGroup){
	defer wg.Done()

	result:= a+b
	ch<-result
}