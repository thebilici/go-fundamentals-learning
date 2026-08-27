package main

import "fmt"

func main(){

	notes:=[]int{45, 70, 90, 30, 85, 60}
	fmt.Println("Notes:", notes)
	avg,highestNote,succes:=analyzeNotes(notes)
	fmt.Println("Average:", avg)
	fmt.Println("Highest Note:", highestNote)
	fmt.Println("Succes Count:", succes)
}

func analyzeNotes(notes []int)(float64,int,int){
	avg := 0.0
	highestNote := notes[0]
	succes := 0

	for _,note:=range notes{
		
		avg+=(float64(note)/float64(len(notes)))

		if(note>=60){
			succes++
		}
		if(note>highestNote){
			highestNote=note
		}

	}
	return avg,highestNote,succes
}