package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

//Hello World
func main() {
	fmt.Println("Hello World")

	//Call the function
	a := 10
	b := 5
	result := add(a, b)
	fmt.Println("The sum of ", a, " and ", b, " is ", result)
}
