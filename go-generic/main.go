package main

import "fmt"

func main() {
	result := add[int](5, 2)
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
