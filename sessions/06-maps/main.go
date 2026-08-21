package main

import "fmt"

func main(){

	user:=map[string]string{
		"name":"Fatih",
		"lastname":"Kaya",
	}

	fmt.Println("User:",user["name"])
	user["age"]="22"
	user["name"]="Furkan"
	fmt.Println("User güncellemeden sonra",user)


	scores:=map[string]int{
		"Ali":90,
		"Veli":80,
		"Ahmet":70,
	}
	fmt.Println("Scores:",scores["Ali"])
	
	scores["Mehmet"]=85
	fmt.Println("Scores:",scores)

	fmt.Println("Scores:",scores["İso"])
	score,ok:=scores["İso"]
	if ok{
		fmt.Println("Score:",score)
	}
	else{
		fmt.Println("İso isimli bir öğrenci yok")
	}
	
	delete(scores,"Ahmet")
	fmt.Println("Güncellemeden sonra",len(scores))

	for key,value:=range scores{
		fmt.Println("Key:",key,"Value:",value)
	}

}