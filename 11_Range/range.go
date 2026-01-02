package main

import "fmt"

func main() {
	// RANGE IN GO
	// The 'range' keyword in Go is used to iterate over elements in various data structures like arrays, slices, maps, and strings.

	nums := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(nums); i++ {
		fmt.Println("Index:", i, "Value:", nums[i])
	}

	fmt.Println("-------------------------------")
	fmt.Println("Using range to iterate over nums:")
	fmt.Println("-------------------------------")
	// Using range to iterate over a slice
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	m := map[string]string{"fname":"john", "lname": "doe"}
	fmt.Println("-------------------------------")
	fmt.Println("Using range to iterate over map m:")
	fmt.Println("-------------------------------")
	// Using range to iterate over a map
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}


	for idx, c := range "golang" {
		// UNICODE point rune
		// Starting byte of rune
		// 255 -> 1 byte
		// 256 -> 2 bytes
		
		fmt.Println("Index:", idx, "Character:", c)
	}

	for idx, c := range "golang" {
		fmt.Println("Index:", idx, "Character:", string(c))
	}
}