package main

import "fmt"

/* Array and Slice Practice
1. Create a new array (!) that contains three hobbies you have.
Output (print) that array in the command line.
2. Also output more data about that array:
- The first element (standalone)
- The second and third element combined as a new list
3. Create a slice based on the first element that contains the first and second elements.
Create that slice in two different ways (i.e. create two slices in the end)
4. Re-slice the slice from (3) and change it to contain the second and last element of the original array.
5.Create a "dynamic array" that contains your course goals (at least 2 goals)
6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
7. Bonus: Create a "Product" struct with title, id, price and create a dynamic list of products (at least 2 products).
Then add a third product to the existing list of products.
*/

type Product struct {
	id    int
	title string
	price float64
}

func main() {

	// Step 1
	hobbies := [3]string{"Reading", "Cycling", "Cooking"}
	fmt.Println("Step 1:", hobbies[0])

	// Step 2
	combined := hobbies[1:3]
	fmt.Println("Step 2:", combined)

	// Step 3
	slice1 := hobbies[0:2]
	fmt.Println("Step 3:", slice1)
	// Append
	slice2 := []string{hobbies[0]}
	slice2 = append(slice2, hobbies[0], hobbies[1])

	// Step 4
	reSlice := []string{slice1[1], hobbies[2]}
	fmt.Println("Step 4:", reSlice)

	// Step 5
	goals := []string{"Learn Go", "Build api with Go"}
	fmt.Println("Step 5:", goals)

	// Step 6
	goals[1] = "Master Go"
	// Adding a third goal
	goals = append(goals, "Build a web app with Go")
	fmt.Println("Step 6:", goals)

	// Step 7
	products := []Product{
		{1, "Go Programming Book", 29.99},
		{2, "Go Course", 199.99},
	}
	// Adding a third product
	products = append(products, Product{3, "Go Web Development", 49.99})
	fmt.Println("Step 7:", products)
}
