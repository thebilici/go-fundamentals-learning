package main

import "fmt"
func main(){

	students:=map[string]int{
		"Fatih": 100,
		"Ali": 90,
		"Veli": 80,
	}

	score,exist:=GetValue(students,"Fatih")
	if exist{
		fmt.Println("Score:",score)
	}
	else{
		fmt.Println("Student not found")
	}
}


func GetValue[K comparable,V any](data map[K]V,key K)(V,bool){
	student,exists:=data[key]
	return student,exists
}