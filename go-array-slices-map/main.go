package main

import "fmt"

type floatMap map[string]float64 // Type alias

func (f *floatMap) output() {
	fmt.Println(*f)
}

func main() {

	// Slices
	names := make([]string, 3, 5)

	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Jessica"
	names = append(names, "Jack", "Jill")
	//fmt.Println(names)
	//fmt.Println(len(names))
	for index, value := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Maps
	courses := make(floatMap, 5)

	courses["Go"] = 5.33
	courses["Python"] = 15.55
	courses["Java"] = 25.99
	courses["JavaScript"] = 35.75
	courses["C++"] = 45.00
	//courses.output()

	for key, value := range courses {
		fmt.Printf("Course: %s, Price: %.2f\n", key, value)
	}

}
