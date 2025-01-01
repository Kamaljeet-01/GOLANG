package main

import "fmt"

func intrestCalculator() float64 {
	var principal float64
	var rate float64
	var time float64
	fmt.Scan(&principal)
	fmt.Scan(&rate)
	fmt.Scan(&time)
	var intrest float64 = (principal * rate) / time
	fmt.Println("Total Amount after Intrest :")
	return intrest + principal
}

