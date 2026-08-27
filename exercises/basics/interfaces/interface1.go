package main	
import "fmt"

type PaymentMethod interface{
	Pay(amount float64) string
}

type CreditCard struct{
	CardHolder string
}

type CashPayment struct{
	Receiver string
}
func main(){

	creditCardPayment:=CreditCard{
		CardHolder:"Fatih Bilici",
	}
	cashPayment:=CashPayment{
		Receiver:"Veli Kabük",
	}

	ProcessPayment(creditCardPayment,100.0)
	ProcessPayment(cashPayment,50.0)

	
}

func (cc CreditCard) Pay(amount float64) string{
	return fmt.Sprintf("%s tarafından kredi kartı ile %.2f TL ödendi",cc.CardHolder,amount)
}

func (cp CashPayment) Pay(amount float64) string{
	return fmt.Sprintf("%s tarafından nakit olarak %.2f TL ödendi",cp.Receiver,amount)
}

func ProcessPayment(p PaymentMethod,amount float64){
	fmt.Println(p.Pay(amount))
}