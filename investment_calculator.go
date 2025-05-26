package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.0
	var investmentAmount float64
	var years float64
	expectedReturn := 7.0

	fmt.Println("Investment Calculator")
	fmt.Print("Enter investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter number of years:")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Println("Future value:", futureValue)
	//fmt.Println("Future real value:", futureRealValue)

	formatRV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formatRFV := fmt.Sprintf("Future real value: %.2f\n", futureRealValue)

	//fmt.Printf("Future value: %.2f\nFuture real value: %.2f\n", futureValue, futureRealValue)
	fmt.Print(formatRV, formatRFV)
}
