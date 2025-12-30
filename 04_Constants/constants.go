package main

import "fmt"

const pi = 3.14
var radius float32 = 5.0
// name := "Rupesh" // This line will cause a compile-time error

func main() {
	const name = "Rupesh"
	fmt.Println("Hello, " + name)

	/*
		name = "John" // This will cause a compile-time error
		fmt.Println("Hello, " + name)
	*/

	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println("Server running at", host, "on port", port)
}