package main

import "fmt"

func multipleReturnValues() (string, string) {
	return "Agus", "Yogi"
}

func multipleReturnValuesWithNamedReturn() (firstName string, lastName string) {
	firstName = "Sunday"
	lastName = "Ahad"
	return
}

func main() {
	firstName, lastName := multipleReturnValues()
	fmt.Println("First Name:", firstName, "Last Name:", lastName)

	firstName2, lastName2 := multipleReturnValuesWithNamedReturn()
	fmt.Println("First Name:", firstName2, "Last Name:", lastName2)
}
