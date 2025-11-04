package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices -> built on top of arrays, more flexible than arrays
	// Initializing a slice
	nums := make([]int, 0) // make() takes 3 parameters: type, length, capacity (capacity is optional)
	fmt.Println(nums)
	fmt.Println(cap(nums)) // capacity -> maximum no of elements can fit in the slice

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	fmt.Println(nums)
	fmt.Println(cap(nums))

	// Copying slices
	nums1 := make([]int, 0, 5) // length 0, capacity 5
	nums1 = append(nums1, 10, 20, 30, 40, 50)

	nums2 := make([]int, len(nums1))

	fmt.Println(nums1)

	copy(nums2, nums1)
	fmt.Println(nums2)

	// slice operator
	nums3 := nums1[1:4] // from index 1 to index 3 (4-1)
	fmt.Println(nums3)

	// comparing slices
	fmt.Println(slices.Equal(nums1, nums2))

	// multidimensional slices
	matrix := make([][]int, 3) // slice of slices

	for i := range matrix {
		matrix[i] = make([]int, 4) // each inner slice has length 4
	}

	matrix[0][0] = 1
	matrix[1][2] = 5
	matrix[2][3] = 10

	fmt.Println(matrix)
}
