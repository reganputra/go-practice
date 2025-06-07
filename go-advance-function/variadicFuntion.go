package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5)) // Output: 15
	fmt.Println(sumAll(10, 20, 30))    // Output: 60
	fmt.Println(sumAll())              // Output: 0

	numbers := []int{10, 20, 30, 40}
	fmt.Println(sumAll(numbers...)) // Output: 100
}
