package main

import (
	"fmt"
	"math"
	"os"
)

const resultFile = "profit_results.txt"

func writeResultToFile(earningBeforeTax, profit, ratio float64) {
	resultText := fmt.Sprintf("Earnings Before Tax: %.1f\nProfit (Earnings After Tax): %.1f\nRatio (Earnings Before Tax / Profit): %.1f\n",
		earningBeforeTax, profit, ratio)
	err := os.WriteFile(resultFile, []byte(resultText), 0644)
	if err != nil {
		fmt.Println("Error writing results to file:", err)
	} else {
		fmt.Println("Results successfully written to", resultFile)
	}
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// Start of the program
	outputCommand("Profit Calculator")
	revenue = getInput("Enter total revenue: ")
	if revenue <= 0 {
		fmt.Println("Error: Revenue must be greater than 0.")
		return
	}
	expenses = getInput("Enter total expenses: ")
	if expenses <= 0 {
		fmt.Println("Error: Expenses must be greater than 0.")
		return
	}
	taxRate = getInput("Enter tax rate (as a percentage): ")
	if taxRate < 0 || taxRate > 100 {
		fmt.Println("Error: Tax rate must be between 0 and 100.")
		return
	}

	// Calculate profit and ratios
	earningBeforeTax, profit, ratio := calculateProfit(revenue, expenses, taxRate)
	writeResultToFile(earningBeforeTax, profit, ratio)

	// Display results
	//outputCommand("\nResults:")
	//outputResult(earningBeforeTax, profit, ratio)

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

func outputResult(earningBeforeTax, profit, ratio float64) {
	fmt.Printf("Earnings Before Tax: %.1f\n", earningBeforeTax)
	fmt.Printf("Profit (Earnings After Tax): %.1f\n", profit)
	fmt.Printf("Ratio (Earnings Before Tax / Profit): %.1f\n", ratio)
}
