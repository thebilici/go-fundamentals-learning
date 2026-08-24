package main	

import "fmt"

type EmailNotifier struct{
	Address string
}

type SMSNotifier struct{
	Phone string
}

type Notifier interface{
	send() string
}

func main(){

	eMail:=EmailNotifier{
		Address:"example@example.com",
	}

	phone:=SMSNotifier{
		Phone:"1234567890",
	}

	sendNotification(eMail)
	sendNotification(phone)

	var notifier Notifier
	
	notifier = eMail
	fmt.Println(notifier.send())
	notifier=phone
	fmt.Println(notifier.send())

}

func (e EmailNotifier) send() string{
 return e.Address + " adresine mail gönderildi"
}

func (s SMSNotifier) send() string{
	return s.Phone + " numarasına sms gönderildi"
}

func sendNotification(n Notifier){
fmt.Println(n.send())
}


