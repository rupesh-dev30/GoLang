package main

import "fmt"

func main() {
	// ARRAYS
	// INT -> by default = 0, string -> by default = "", bool -> by default = false
	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	fmt.Println("Array length:", len(nums))
	fmt.Println("Array contents:", nums)
	fmt.Println("First element:", nums[0])

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	var name [3]string
	name[0] = "Golang"
	fmt.Println(name)

	nums1 := [5]int {1,2,3,4,5}
	fmt.Println("Initialized array:", nums1)

	// 2D ARRAYS
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println("2D Array (Matrix):", matrix)

	matrix1 := [2][2]string { {"a", "b"}, {"c", "d"} }
	fmt.Println("Initialized 2D Array:", matrix1)
}