package main

import "fmt"

func main(){

	 age :=22
	 isStudent := true
	 hasPermission := false

	 fmt.Println("Yaşım:", age==22)
	 fmt.Println("Öğrenci miyim?:", age>=18)

	

	//if else 
     age:=20
	 isStudent:=true
	 if age>=18{
		 fmt.Println("18 yaşından büyüksün ")	
	 } else{
		 fmt.Println("18 yaşından küçüksün")
	 }

	 score:=85
	 if score>=90{
		 fmt.Println("Aldığın not A")
	 }
	 else if score>=80{
		 fmt.Println("Aldığın not B")
	 }
	 else if score>=70{
		 fmt.Println("Aldığın not C")
	 }
	 else{
		fmt.Println("Kaldın")
	 }

	 //switch case
		switch role:= "admin"{
		case "admin":
			fmt.Println("Admin yetkisine sahipsin")
		case "user":
			fmt.Println("User yetkisine sahipsin")
		case "guest,visitor":
		default:
			fmt.Println("Yetkin yok")
		}

		//for 
		for i:=1 ;i<=5 ;i++{
			fmt.Println(i)
		}

		for i:=5 ;i<=1 ;i--{
			fmt.Println(i)
		}

		//while benzeri
		count:=1
		for count<=5{
			fmt.Println(count)
			count++
		}

		//sonsuz
		for {
			fmt.Println("Sonsuz döngü")
		}
		//Çift sayılar
		for i:=0;i<=10;i+=2{
			fmt.Println(i)
		}


		//break
		for i:1;i<=10;i++{
			if i==5{
				break
			}

			fmt.Println(i)
		}

		for i:=1;i<=10;i++{

			if i==3{
				continue
			}
			fmt.Println(i)
		}

		//range

		language:=[]{"Go","Python","JavaScript"}

		for index,value:=range language{
			fmt.Println("Index:",index,"Value:",value)
		}

		for _,value:=range language{
			fmt.Println("Value:",value)
		}

		for index :=range language{
			fmt.Println("Index:",index)
		}
}