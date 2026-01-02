package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices are dynamically-sized, flexible views into the elements of an array.
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).

	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil) // true

	fmt.Println(len(nums))

	var nums1 = make([]int, 2)
	fmt.Println(len(nums1))

	// CAPACITY: Maximum size the slice can grow to (underlying array size)
	var nums2 = make([]int, 2)
	fmt.Println(cap(nums2))

	var nums3 = make([]int, 2, 5)
	fmt.Println(cap(nums3))

	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)
	nums3 = append(nums3, 3)
	nums3 = append(nums3, 4)
	nums3 = append(nums3, 5)

	fmt.Println(nums3)
	fmt.Println(cap(nums3))


	nums4 := []int{}
	fmt.Println(nums4)
	fmt.Println(len(nums4))
	fmt.Println(cap(nums4))


	var nums5 = make([]int, 0, 5)
	nums5 = append(nums5, 1)
	var nums6 = make([]int, len(nums5))


	fmt.Println(nums5,nums6)

	// COPY FUNCTION
	var nums7 = make([]int, 0, 5)
	nums7 = append(nums7, 3)
	var nums8 = make([]int, len(nums5))

	copy(nums8,nums7)

	fmt.Println(nums7,nums8)


	// SLICE OPERATOR
	var nums9 = []int{1,2,3}
	fmt.Println(nums9[0:2])
	fmt.Println(nums9[:2])
	fmt.Println(nums9[0:])


	// SLICE
	var nums10 = []int{1,2,3,4,5}
	var nums11 = []int{1,2,3,4,5}

	fmt.Println(slices.Equal(nums10,nums11))


	// 2D SLICES
	var matrix = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println(matrix)
}