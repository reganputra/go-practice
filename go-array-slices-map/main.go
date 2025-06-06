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

	courses := make(map[string]string, 5)

	courses["Go"] = "Go Programming Language"
	courses["Python"] = "Python Programming Language"
	courses["Java"] = "Java Programming Language"
	courses["JavaScript"] = "JavaScript Programming Language"
	courses["C++"] = "C++ Programming Language"
	fmt.Println(len(courses))

}
