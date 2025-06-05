package main

import "fmt"

func main() {

	var productNames = [4]string{"A Book"}
	fmt.Println(productNames)
	arr := [5]float64{1.23, 2.12, 3.14, 4.20, 55.}
	fmt.Println(arr[0])

	// Display the array elements
	//for i: = 0; i < len(productNames); i++ {
	//	println(productNames[i])
	//}

	featuredArr := arr[1:3]
	fmt.Println(featuredArr)
}
