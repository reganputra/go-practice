package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {

	taxRates := []float64{0.05, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("output_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		job.Process()
	}
}
