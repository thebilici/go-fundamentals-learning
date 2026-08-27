package main

import "fmt"

func main(){
	students:=map[string]int{
	"Fatih": 85,
	"Ali":   45,
	"Veli":  70,
	"Ayşe":  95,
	}

	fmt.Println("Students:",students)

	note,ok:=findStudent(students,"Ali")
	if ok==true{
		fmt.Println("Student's note:",note)
	}else{
		fmt.Println("Student's note not found")
	}

	passedStudents:=filterPassedStudents(students,50)
	fmt.Println("Passed students:",passedStudents)
}

func findStudent(students map[string]int,name string)(note int,ok bool){

	/*
	note,ok:=students[name]
	return note,ok
	bu da yapılabilir
	*/

	for studentName,studentNote:=range students{
		if studentName==name{
			return students[name],true
		}
	}	
	return 0,false
}

func filterPassedStudents(students map[string]int,passingGrade int)(passedStudents map[string]int){

	var filteredStudents=make(map[string]int)

	for studentName,studentNote:=range students{
		if studentNote>=passingGrade{
			filteredStudents[studentName]=studentNote
		}
	}

	return filteredStudents

}