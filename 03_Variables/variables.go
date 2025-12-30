package main

import "fmt"

func main(){
	// WITH TYPE DECLARATION
	var name string = "Rupesh"
	fmt.Println("Hello, " + name)

	// WITHOUT TYPE DECLARATION
	var age = 30
	fmt.Println("I am", age, "years old.")

	// SHORT HAND DECLARATION
	country := "India"
	fmt.Println("I live in " + country)

	// MULTIPLE VARIABLE DECLARATION
	var city, state = "Bokaro", "Jharkhand"
	fmt.Println("I live in", city+",", state)

	// MULTIPLE VARIABLE DECLARATION WITH TYPE
	var (
		pincode int    = 123456
		continent string = "Asia"
	)
	fmt.Println("Pincode:", pincode)
	fmt.Println("Continent:", continent)

	// VARIABLE DECLARATION AND ASSIGNMENT IN DIFFERENT LINES
	var school string
	school = "ABC Public School"
	fmt.Println("I studied at", school)


	var price float32 = 50.2
	var price2 float64 = 100.5
	fmt.Println("Price 1:", price)
	fmt.Println("Price 2:", price2)

	college := "XYZ University"
	fmt.Println("I graduated from", college)
}