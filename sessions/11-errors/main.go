package main
import (
	"fmt"
	errors "errors"
)

func main(){

	result,err:=divide(10,2)

	if err!=nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Result:",result)

	result2,err2:=multiple(10,2)

	if err2!=nil{
		fmt.Println("Error:",err2)
		return
	}
	fmt.Println("Result2:",result2)

	result,err=calculate(10,0)

	if err!=nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Result:",result)
}


func divide(a int,b int )(int,error){
	if b==0{
		return 0,errors.New("Cannot divide by zero")
	}

	return a/b,nil
}

func multiple(a int,b int)(int,error){
	if b==0{
		return 0,fmt.Errorf("Cannot multiply %d by %d: zero",a,b)
	}
	return a*b,nil
}

func calculate(a int,b int)(int,error){
	result,err:=divide(a,b)

	if err!=nil{
		return 0,fmt.Errorf("Calculate Failed: %w",err)
	}
	return result,nil
}