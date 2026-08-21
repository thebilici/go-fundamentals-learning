package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {

	user := User{
		Name:     "Fatih",
		Age:      22,
		IsActive: true,
	}

	fmt.Println("User:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("IsActive:", user.IsActive)

	user.Age = 23
	fmt.Println("User güncellemeden sonra yaş:", user.Age)

	createdUser:= createUser("Furkan", 22)
	fmt.Println("Created User:", createdUser.Name, createdUser.Age, createdUser.IsActive)


	printUser(user)


	users:=[]User{
		{
			Name: "Ali",
			Age:20,
			IsActive:true,
	    },
		{
			Name: "Tahir",
			Age:25,
			IsActive:false,
		},
	}
	fmt.Println("User:",users[1])
	fmt.Println("Users:", users[0].Name, users[0].Age, users[0].IsActive)
	fmt.Println("Users:", users)

	newUser:=User{
		Name:"Veli",
		Age:30,
		IsActive:true,
	}
	users=append(users,newUser)
	fmt.Println("Users:",users)
	
	for _,user:=range users{
		fmt.Println("User:",user.Name,"Age:",user.Age,"IsActive:",user.IsActive)
	}
}

func printUser(user User) {
	fmt.Println("User:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("IsActive:", user.IsActive)
}

func createUser(name string, age int) User {

	return User{
		Name:     name,
		Age:      age,
		IsActive: true,
	}

}