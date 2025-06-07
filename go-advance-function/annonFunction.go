package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 3
	})

	fmt.Println(transformed)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
