package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0.05, 0.1, 0.15}
	doneChan := make([]chan bool, len(taxRates))
	errorChan := make([]chan error, len(taxRates))
	for i, taxRate := range taxRates {
		doneChan[i] = make(chan bool)
		errorChan[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go job.Process(doneChan[i])
	}

	for _, errChan := range errorChan {
		errChan <- nil // In a real application, you would handle errors here
	}
	for _, done := range doneChan {
		<-done
		fmt.Println("Done!")
	}
}
