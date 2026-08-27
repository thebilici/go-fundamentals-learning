package main

import "fmt"

type smsNotifier struct {
	Phone string
}

type emailNotifier struct {
	Email string
}

type Notifier interface {
	Send() string
	GetType() string
}

func main() {
	email := emailNotifier{
		Email: "deneme@example.com",
	}

	phone := smsNotifier{
		Phone: "5551234567",
	}

	sendNotification(email)
	sendNotification(phone)
}

func (s smsNotifier) Send() string {
	return fmt.Sprintf(
		"Merhaba mesajı %s numarasına gönderildi",
		s.Phone,
	)
}

func (s smsNotifier) GetType() string {
	return "SMS"
}

func (e emailNotifier) Send() string {
	return fmt.Sprintf(
		"Merhaba maili %s adresine gönderildi",
		e.Email,
	)
}

func (e emailNotifier) GetType() string {
	return "Email"
}

func sendNotification(n Notifier) {
	fmt.Println("Notification type:", n.GetType())
	fmt.Println(n.Send())
}