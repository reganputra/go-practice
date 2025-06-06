package main

import "fmt"

func main() {

	//var productNames = [4]string{"A Book"}
	//fmt.Println(productNames)
	//prices := [5]float64{1.23, 2.12, 3.14, 4.20, 5.15}
	//fmt.Println(prices)

	// Display the array elements
	//for i: = 0; i < len(productNames); i++ {
	//	println(productNames[i])
	//}

	//// Slices are a reference to an array
	//featurePrices: = prices[1:]
	//morePrices := featurePrices[:4]
	//fmt.Println(len(morePrices), cap(featurePrices)) // Slices take the original memory of the array

	stock := []float64{31.1, 22.2, 33.3, 44.4, 55.5}
	fmt.Println(stock[0:5])
	stock = append(stock, 66.6)
	fmt.Println(stock)

}
