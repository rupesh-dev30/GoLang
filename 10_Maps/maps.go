package main

import (
	"fmt"
	"maps"
)

func main() {
	// MAPS IN GO
	// A map is a built-in data type that associates keys with values.
	// Maps are similar to dictionaries in Python or hash tables in other languages.

	// Creating a map
	m := make(map[string]string)

	// Adding key-value pairs to the map
	m["name"] = "Alice"
	m["city"] = "Wonderland"

	// Accessing values by key
	name := m["name"]
	city := m["city"]
	fmt.Println("Name:", name)
	fmt.Println("City:", city)

	// If you try to access a key that does not exist, it returns the zero value for the value type
	fmt.Println(m["phone"])

	// DELETE KEYS
	// You can delete a key-value pair from a map using the built-in delete function
	delete(m, "city")
	fmt.Println("After deleting 'city':", m)

	m2 := make(map[string]int)
	m2["age"] = 30
	age := m2["age"]
	fmt.Println("Age:", age)
	fmt.Println(m2["height"]) // returns 0, the zero value for int

	clear(m2)
	fmt.Println(m2)


	m3 := map[string]int{"price": 100, "quantity": 5}
	fmt.Println("Price:", m3["price"])
	fmt.Println("Quantity:", m3["quantity"])


	// Checking if a key exists
	/*
		_, ok := m["name"]
		if ok {
			fmt.Println("Key 'name' exists")
		} else {
			fmt.Println("Key 'name' does not exist")
		}
	*/

	if value, exists := m["name"]; exists {
		fmt.Println("Key 'name' exists with value:", value)
	} else {
		fmt.Println("Key 'name' does not exist")
	}

	m4 := map[string]int{"price": 100, "quantity": 5}
	m5 := map[string]int{"price": 100, "quantity": 5}

	// fmt.Println(m4 == m5)
	fmt.Println("m4 and m5 are equal in content:", maps.Equal(m4, m5))
}