package main

import "fmt"

func main() {
	names := make([]string, 3, 5)

	//names = append(names, "John", "Jane", "Jessica")
	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Jessica"
	names = append(names, "Jack", "Jill")
	fmt.Println(names)
	fmt.Println(len(names))

}
