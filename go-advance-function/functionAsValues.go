package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dbNumbers := transformNumber(&numbers, doubleNumber)
	trNumbers := transformNumber(&numbers, tripleNumber)
	fmt.Println(dbNumbers)
	fmt.Println(trNumbers)
}

func transformNumber(num *[]int, transform transformFunc) []int {
	var dbNumbers []int
	for _, value := range *num {
		dbNumbers = append(dbNumbers, transform(value))
	}
	return dbNumbers
}

func doubleNumber(num int) int {
	return num * 2
}

func tripleNumber(num int) int {
	return num * 3
}
