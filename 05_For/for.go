package main

import "fmt"

// ONLY FOR AVAILABLE IN GO

func main() {
	// IMPLEMENT WHILE LOOP USING FOR
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// INFINITE LOOP
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// CLASSIC FOR LOOP
	for i := 0; i < 5; i++ {
		fmt.Println("For Loop:", i)
	}

	// FOR RANGE LOOP
	for i := range 5 {
		fmt.Println("For Loop:", i)
	}
}