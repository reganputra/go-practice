package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 3
	})

	fmt.Println(transformed)
	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
