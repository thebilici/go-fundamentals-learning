package main 
import "fmt"

func main(){

  greet("Fatih")
  result:=add(5,10)
  fmt.Println("Result:",result)
  fmt.Println(checkAge(17))
  name ,_:=getUser()
  fmt.Println("Name:",name)
}

func greet(name string) {
  fmt.Println("Hello",name)
}

func add(a int,b int) int{
  return a+b
}

func checkAge(age int)string{
	if age>=18{
		return "Reşit"
	}
	return "Reşit değil"
}

func getUser()(string,int){
	return "Aytuğ",22
}

