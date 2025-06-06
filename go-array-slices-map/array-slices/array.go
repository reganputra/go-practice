package main

import "fmt"

func main() {

	var productNames = [4]string{"A Book"}
	fmt.Println(productNames)
	prices := [5]float64{1.23, 2.12, 3.14, 4.20, 5.15}
	fmt.Println(prices)

	// Display the array elements
	//for i: = 0; i < len(productNames); i++ {
	//	println(productNames[i])
	//}

	// Slices are a reference to an array
	featurePrices := prices[1:]
	morePrices := featurePrices[:4]
	fmt.Println(len(morePrices), cap(featurePrices)) // Slices take the original memory of the array

	stock := []float64{31.1, 22.2, 33.3, 44.4, 55.5}
	fmt.Println(stock[0:5])
	stock = append(stock, 66.6)
	fmt.Println(stock)

	// Unpacking list values from a slice
	number := []int{1, 2, 3, 4, 5}
	fmt.Println(number[0:1])

	number = append(number, 6, 7, 8)
	fmt.Println(number)

	addtionalNumbers := []int{9, 10, 11, 12, 13}
	number = append(number, addtionalNumbers...) // Unpacking the slice / Merging slices
	fmt.Println("Unpacked:", number)

}
