package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Profit Calculator")
	fmt.Print("Enter total revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter total expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate (as a percentage): ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses

	profit := earningBeforeTax * (1 - taxRate/100)

	ratio := earningBeforeTax / profit

	ratio = math.Round(ratio*100) / 100

	// Display results
	fmt.Println("\nResults:")
	fmt.Printf("Earnings Before Tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Profit (Earnings After Tax): %.2f\n", profit)
	fmt.Printf("Ratio (Earnings Before Tax / Profit): %.2f\n", ratio)
}
