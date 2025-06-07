package main

import (
	"price-calculator/prices"
)

func main() {

	taxRates := []float64{0.05, 0.1, 0.15}

	for _, rate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(rate)
		job.Process()
	}
}
