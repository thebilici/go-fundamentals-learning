package main

import "fmt"

func main() {
	prices := []float64{120, 80, 250, 90}
	isPremium := true
	couponPercent := 10.0

	subTotal := CalculateSubTotal(prices)
	fmt.Println("Ara Toplam:", subTotal)

	discountAmount := CalculateDiscount(subTotal, isPremium, couponPercent)
	fmt.Println("İndirim Miktarı:", discountAmount)

	finalPrice, freeDelivery := CalculateFinalPrice(subTotal, discountAmount)
	fmt.Println("Final Price:", finalPrice)
	fmt.Println("Free Delivery:", freeDelivery)
}

func CalculateSubTotal(prices []float64) float64 {
	subTotal := 0.0

	for _, price := range prices {
		subTotal += price
	}

	return subTotal
}

func CalculateDiscount(subTotal float64, isPremium bool, couponPercent float64) float64 {
	discountPercent := couponPercent

	if isPremium {
		discountPercent += 5
	}

	if discountPercent > 25 {
		discountPercent = 25
	}

	discountAmount := subTotal * (discountPercent / 100)

	return discountAmount
}

func CalculateFinalPrice(subTotal float64, discountAmount float64) (float64, bool) {
	subTotalAfterDiscount := subTotal - discountAmount

	if subTotalAfterDiscount >= 400 {
		return subTotalAfterDiscount, true
	}

	return subTotalAfterDiscount + 50, false
}