package main

import (
	"price-calculator/cmdmanager"
	"price-calculator/prices"
)

func main() {

	taxRates := []float64{0.05, 0.1, 0.15}

	for _, taxRate := range taxRates {
		cmd := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		job.Process()
	}
}
