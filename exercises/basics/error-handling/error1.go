package main

import (
	"errors"
	"fmt"
)

type BankAccount struct{
	Owner string
	Balance float64
}

func main(){

	user:=BankAccount{
		Owner:"Fatih Bilici",
		Balance:5000,
	}
	newBalance,err:=user.WithDraw(1500)

	if err!=nil{
		fmt.Println("Error:",err)
		return
	}
	err=nil
	fmt.Println("New Balance:",newBalance)
}

func (a *BankAccount) WithDraw(amount float64) (float64, error){

	if amount<=0{
		return a.Balance,fmt.Errorf("Invalid amount: %.2f. Amount must be greater than zero",amount)
	}
	
	if a.Balance<amount{
		return a.Balance,errors.New("Insufficient balance")
	}
	
	a.Balance-=amount
	return a.Balance,nil
}
