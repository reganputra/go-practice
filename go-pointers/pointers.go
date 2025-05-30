package main

import "fmt"

func main() {
	age := 30

	agePointer := &age                          // Create a pointer to age
	fmt.Println("Pointer to age:", *agePointer) // Dereference the pointer to get the value

	//fmt.Println(age)
	adultAge(agePointer)
	fmt.Println(age)
}

func adultAge(age *int) {
	*age = *age - 18
}
