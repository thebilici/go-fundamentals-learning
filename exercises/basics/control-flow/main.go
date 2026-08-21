package main

import "fmt"

func main(){
	
	age:=22
	role:="admin"
	isActive:=true

	if age>=18 && isActive{
		fmt.Println("18 yaşından büyüksün ve aktif bir kullanıcı")
	}


	switch role{
	case "admin":
		fmt.Println("Admin yetkisine sahipsin")
	case "user":
		fmt.Println("User yetkisine sahipsin")
		case "Guest":
		fmt.Println("Guest yetkisine sahipsin")
	}
    

		for i:=1 ;i<=10 ;i++{
			
			if i==4{
				continue
			}

			if i==8{
				break
			}
			fmt.Println(i)

		}

		languages:=[]string{"Go","Python","JavaScript"}

		for index,value :=range languages{
			fmt.Println("Index:",index,"Value:",value)
		}

}