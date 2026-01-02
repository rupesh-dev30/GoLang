package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add1(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "python", "c++";
}

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 2 * a
	}
}

func main() {
	// FUNCTIONS IN GO
	// A function is a block of code that performs a specific task.
	// Functions in Go are defined using the 'func' keyword followed by the function name and parameters.
	fmt.Println("Hello, Functions in Go!")
	sum := add(5, 10)
	fmt.Println("Sum:", sum)

	sum1 := add1(15, 25)
	fmt.Println("Sum1:", sum1)

	lang1, lang2, lang3 := getLanguages()
	fmt.Println("Languages:", lang1, lang2, lang3)

	lang4, lang5, _ := getLanguages()
	fmt.Println("Languages:", lang4, lang5)


	// fn := func(a int) int {
	// 	return 2 * a
	// }
	
	fn := processIt()
	fmt.Println("Result from processIt:", fn(10))
}