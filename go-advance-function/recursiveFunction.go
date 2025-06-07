package main

import "fmt"

func main() {

	fact := factorial(10)
	fmt.Println("Factorial of 10 is:", fact)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
