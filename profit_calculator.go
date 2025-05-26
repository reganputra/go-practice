package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// Start of the program
	outputCommand("Profit Calculator")
	revenue = getInput("Enter total revenue: ")
	expenses = getInput("Enter total expenses: ")
	taxRate = getInput("Enter tax rate (as a percentage): ")

	// Calculate profit and ratios
	earningBeforeTax, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	// Display results
	outputCommand("\nResults:")
	outputValues(earningBeforeTax, profit, ratio)

}

func outputCommand(text string) {
	fmt.Println(text)
}

func getInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func calculateProfit(revenue, expenses, taxRate float64) (earningBeforeTax, profit, ratio float64) {
	earningBeforeTax = revenue - expenses
	profit = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / profit
	ratio = math.Round(ratio*100) / 100
	return earningBeforeTax, profit, ratio
}

func outputValues(earningBeforeTax, profit, ratio float64) {
	fmt.Printf("Earnings Before Tax: %.1f\n", earningBeforeTax)
	fmt.Printf("Profit (Earnings After Tax): %.1f\n", profit)
	fmt.Printf("Ratio (Earnings Before Tax / Profit): %.1f\n", ratio)
}
