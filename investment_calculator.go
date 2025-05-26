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
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
