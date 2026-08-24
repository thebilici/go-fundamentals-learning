package main

import (
	"fmt"
	"github.com/thebilici/go-backend-learning/exercises/basics/packages/user"
	//başına mesela userpkg ile import edebilirsin isim çakışmasını engeller
)

func main(){
	
	createdUser:=user.CreateUser("Fatih",22)
	fmt.Println("User başarıyla oluşturuldu:",createdUser)

	name:=user.GetName(createdUser)
	fmt.Println("User adı:",name)

	adult:=user.IsAdult(createdUser)
	fmt.Println("User yetişkin mi?:",adult)
}