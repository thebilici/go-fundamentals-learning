package main

import "fmt"

type User struct{
	ID int
	Name string
	Age int
	IsActive bool
}
func main(){
	/*
	user:=User{
		ID:1,
		Name:"Fatih",
		Age:22,
		IsActive:true,
	}
		*/


	users:=[]User{
		createUser(1,"Fatih",22,true),
		createUser(2,"Ali",17,true),
		createUser(3,"Ayşe",25,false),
		createUser(4,"Veli",30,false),
	}

	activeUsers:=filterActiveUsers(users,true)
	fmt.Println("Active Users:",activeUsers)

	user,found:=findUserByID(users,2)
	if found{
		fmt.Println("User found:",user)
	}else{
		fmt.Println("User not found")
	}
	
}

func createUser(id int,name string,age int,isActive bool)User{
	return User{
		ID:id,
		Name:name,
		Age:age,
		IsActive:isActive,
	}

}

func filterActiveUsers(users []User,isActive bool)[]User{
	var activeUsers []User

	for _,user:=range users{
		if user.IsActive==isActive{
			activeUsers=append(activeUsers,user)
		}
	}
	return activeUsers
}

func findUserByID(users []User,id int)(User,bool){
	
	for _,user:=range users{
		if user.ID==id{
			return user,true
		}
	}
	return User{},false	
}
