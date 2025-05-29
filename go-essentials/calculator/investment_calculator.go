package main

import (
	"fmt"
	"math"
)

const inflationRate = 3.0

func main() {
	var investmentAmount float64
	var years float64
	expectedReturn := 7.0

	outputText("Investment Calculator")
	fmt.Print("Enter investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter number of years:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, years, expectedReturn)

	//futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Println("Future value:", futureValue)
	//fmt.Println("Future real value:", futureRealValue)

	formatRV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formatRFV := fmt.Sprintf("Future real value: %.2f\n", futureRealValue)

	//fmt.Printf("Future value: %.2f\nFuture real value: %.2f\n", futureValue, futureRealValue)
	fmt.Print(formatRV, formatRFV)
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValue(investmentAmount, years, expectedReturn float64) (fv float64, rvf float64) {
	fv = investmentAmount * math.Pow(1+expectedReturn/100, years)
	rvf = fv / math.Pow(1+inflationRate/100, years)
	return fv, rvf
}
