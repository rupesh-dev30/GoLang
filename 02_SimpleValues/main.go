package main

import "fmt"

// SIMPLE VALUES

func main() {
	// INTEGER
	fmt.Println(1)
	fmt.Println(2 + 3)

	// STRING
	fmt.Println("Hello")
	fmt.Println("Hello" + " " + "World")

	// BOOLEAN
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(true && false)

	// FLOAT
	fmt.Println(3.14)
	fmt.Println(1.1 + 2.2)
	fmt.Println(5.0 / 2.0)

	// PRINT MULTIPLE VALUES
	fmt.Println("The answer is", 42)
}