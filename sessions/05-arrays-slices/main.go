package main

import "fmt"

func main(){

		//arrays
	languages:=[3]string{"Go","Python","JavaScript"}

	fmt.Println("Languages:",languages)

	fmt.Println("Languages[0]:",languages[0])
	fmt.Println("Languages[1]:",languages[1])
	fmt.Println("Languages[2]:",languages[2])

	languages[0]="C++"
	fmt.Println("Languages[0]:",languages[0])

	fmt.Println("Length of languages:",len(languages))

		//slices
	usernames:=[]string{"fatih","aytug","ahmet"}
	fmt.Println("Usernames:",usernames)

	usernames=append(usernames,"mehmet","merve")
	fmt.Println("Usernames:",usernames)

	fmt.Println(usernames[0])
	usernames[0]="Sena"
	fmt.Println("Usernames:",usernames)

	//capacity ve length
	fmt.Println("Length of usernames:",len(usernames))

	fmt.Println("Capacity of usernames:",cap(usernames))

	selected:=usernames[1:4]
	fmt.Println("Selected:",selected)

	selected2:=usernames[:3]
	selected3:=usernames[2:]

	fmt.Println("Selected2:",selected2)
	fmt.Println("Selected3:",selected3)

	for i:=0;i<len(usernames);i++{
		fmt.Println(usernames[i])	
	}

	for index,value:=range usernames{
		fmt.Println("Index:",index,"Value:",value)
	}

	for _,value:=range usernames{
		fmt.Println("Value:",value)
	}



}