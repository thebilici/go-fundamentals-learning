package main

import "fmt"

type User struct{
	Name   string
	Age    int
	IsActive bool

}

func main(){
	user:=createUser("Fatih",22,true)
	user2:=createUser("Furkan",22,true)
	user3:=createUser("Ali",22,false)
	
	users:=[]User{
		
        user,	
		user2,	
		user3,
		
	}

	for index,user:=range users{
		fmt.Println("User:",index,"Name:",user.Name,"Age:",user.Age,"IsActive:",user.IsActive)
	}

	for index,user:=range users{
		fmt.Println("User:",index,"Name:",user.Name,"Age:",user.Age,"IsActive:",user.IsActive)
		if user.IsActive==true{
			fmt.Println("User is active",user.Name)
		}
	}

	user4:=createUser("Veli",30,true)

	users:=append(users,user4)

	userCount:=len(users)
	fmt.Println("User Count:",userCount)
}

func createUser(name string,age int,isActive bool)User{
	return User{
		Name:name,
		Age:age,
		IsActive:isActive,
	}
}