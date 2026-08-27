package main

import "fmt"

func main(){

	temperatures:=[]int{18, 22, 35, 41, 29, 16, 38, 25}
	hotDays:=FilterHotDays(temperatures,30)
	fmt.Println("Hot Days:",hotDays)

	fahrenheitTemperatures:=ConvertToFahrenheit(hotDays)
	fmt.Println("hotdays Temperatures:",fahrenheitTemperatures)

}

func FilterHotDays(temperatures []int,limit int)[]int{
	var hotDays []int
	for _,temperature:=range temperatures{
		if temperature>limit{
			hotDays=append(hotDays,temperature)
		}
	}
	return hotDays
}

func ConvertToFahrenheit(temperatures []int)[]float64{
	var fahrenheitTemperatures []float64

	for _,temperature:=range temperatures{
		fahrenheitTemperature:=float64(temperature)*9/5 + 32
		fahrenheitTemperatures=append(fahrenheitTemperatures,fahrenheitTemperature)
	}
	return fahrenheitTemperatures
}