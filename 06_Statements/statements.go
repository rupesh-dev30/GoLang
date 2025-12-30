package main

import "fmt"

func main() {
	age := 20

	if age > 18 {
		fmt.Println("Adult")
	} else if age < 12 {
		fmt.Println("Child")
	} else {
		fmt.Println("Minor")
	}

	switch age {
		case 13, 14, 15, 16, 17:
			fmt.Println("Teenager")
		case 18:
			fmt.Println("Just became an adult")
		default:
			fmt.Println("Age not categorized")
	}

	count := 0
	for count < 5 {

		if count == 3 {
			fmt.Println("Skipping count 3")
			count++
			continue
		}

		fmt.Println("Count is:", count)
		count++
	}

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

	if role == "admin" || hasPermission {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

	if name := "John"; name == "John" {
		fmt.Println("Hello, John!")
	} else {
		fmt.Println("Hello, Guest!")
	}

	// NO TERNERY OPERATOR IN GO
}