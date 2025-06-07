package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}
	dbNumbers := transformNumber(&numbers, doubleNumber)
	trNumbers := transformNumber(&numbers, tripleNumber)
	fmt.Println(dbNumbers)
	fmt.Println(trNumbers)

	transform1 := getTransformNumber(&numbers)
	trNumbers2 := getTransformNumber(&moreNumbers)

	transform3 := transformNumber(&numbers, transform1)
	transform4 := transformNumber(&moreNumbers, trNumbers2)
	fmt.Println(transform3)
	fmt.Println(transform4)

}

func transformNumber(num *[]int, transform transformFunc) []int {
	var dbNumbers []int
	for _, value := range *num {
		dbNumbers = append(dbNumbers, transform(value))
	}
	return dbNumbers
}

func getTransformNumber(num *[]int) transformFunc {
	if (*num)[0] == 1 {
		return doubleNumber
	} else {
		return tripleNumber
	}
}

func doubleNumber(num int) int {
	return num * 2
}

func tripleNumber(num int) int {
	return num * 3
}
