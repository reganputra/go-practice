package main

import "fmt"

func main() {
	age := 30

	agePointer := &age // Create a pointer to age
	fmt.Println("Pointer to age:", *agePointer)

	//fmt.Println(age)
	//adult := adultAge(age)
	//fmt.Println(adult)
}

func adultAge(age int) int {
	return age - 18
}
