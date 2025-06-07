package main

import "fmt"

type City struct {
	Name       string
	Population int
}

func main() {
	age := 30

	agePointer := &age                          // Create a pointer to age
	fmt.Println("Pointer to age:", *agePointer) // Dereference the pointer to get the value

	//fmt.Println(age)
	adultAge(agePointer)
	fmt.Println(age)

	city := City{"New York", 8000000}
	city2 := &city
	city2.Name = "London"
	changePopulation(city2, 10000) // Change the population using a pointer to the City struct
	fmt.Println(city)
	city.PrintPopulation()
}

func adultAge(age *int) {
	*age = *age - 18 // Dereference the pointer to modify the value
}

func changePopulation(city *City, population int) {
	// Dereference the pointer to modify the value
	city.Population = population
}

func (populations *City) PrintPopulation() {
	fmt.Println("Population of", populations.Name, "is", populations.Population)
}
