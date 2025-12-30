package main

import (
	"fmt"
	"time"
)

func main() {
	// SIMPLE SWITCH
	i := 6

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Number not between 1 and 5")
	}

	// MULTIPLE CASES
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("I'm an integer: %d\n", t)
		case string:
			fmt.Printf("I'm a string: %s\n", t)
		default:
			fmt.Printf("Unknown type: %T\n", t)
		}
	}
	whoAmI(42)
	whoAmI("Hello, Go!")
	whoAmI(3.14)
}