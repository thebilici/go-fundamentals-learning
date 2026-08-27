package main

import "fmt"

func main(){
		salaries:=[]float64{28000, 35000, 42000, 30000, 50000}
		bonusPercent:=10.0

		averageSalary:=CalculateAverageSalary(salaries)
		fmt.Println("Ortalama Maaş:", averageSalary)

		highestSalary:=HighestSalary(salaries)
		fmt.Println("En Yüksek Maaş:", highestSalary)

		bonusAmount:=CalculateBonus(highestSalary,bonusPercent)
		fmt.Println("Bonus Miktarı:", bonusAmount)
}

func CalculateAverageSalary(salaries []float64)float64{
	totalSalary:=0.0
	for _,salary:=range salaries{
		totalSalary+=salary
	}
	averageSalary:=totalSalary/float64(len(salaries))

	return averageSalary 
}

func HighestSalary(salaries []float64)float64{
	highestSalary:=salaries[0]

	for _,salary:=range salaries{
		if salary>highestSalary{
			highestSalary=salary
		}
	}
	return highestSalary
}

func CalculateBonus(highestSalary float64,bonusPercent float64)float64{
	bonusAmount:=highestSalary*(bonusPercent/100)
	return bonusAmount
}