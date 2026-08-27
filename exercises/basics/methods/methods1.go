package main

import "fmt"

type BankAccount struct{
	Owner string
	Balance float64
}

func main(){
	user:=BankAccount{
		Owner:"Fatih",
		Balance:5000,	
	}
	
	totalBalance:=user.Deposit(1000)
	fmt.Println("Total Balance:",totalBalance)

	newBalance,ok:=user.WithDraw(2000)
	if ok{
		fmt.Println("New Balance:",newBalance)
	}else{
		fmt.Println("Withdrawal failed. Current Balance:",newBalance)
	}

	currentBalance:=user.GetBalance()
	fmt.Println("Current Balance:",currentBalance)


}

func (a *BankAccount)Deposit(amount float64) float64{
	a.Balance+=amount
	return a.Balance
}

func (a *BankAccount) WithDraw(amount float64) (float64, bool){
	if a.Balance<amount{
		fmt.Println("Insufficient balance")
		return a.Balance,false
	}else{
		a.Balance-=amount
		return a.Balance,true
	}
}

func (a BankAccount) GetBalance()float64{
	return a.Balance
}
