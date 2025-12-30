package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	//Hello World
	fmt.Println("Hello World")

	//Call the function
	a := 10
	b := 5
	result := add(a, b)
	fmt.Println("The sum of ", a, " and ", b, " is ", result)

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
