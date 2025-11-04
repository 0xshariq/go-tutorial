package main

func main() {
	// Example usage of arrays in Go
	var arr1 [5]int // Declare an array of 5 integers
	arr1[0] = 10    // Assign value to the first element
	arr1[1] = 20    // Assign value to the second element

	// Iterate over an array using a for loop
	for i := range len(arr1) {
		println(arr1[i]) // Print each element of arr1
	}

	// Declare and initialize an array
	arr2 := [5]string{"apple", "banana", "cherry"}


	// Iterate over an array using range
	for index, value := range arr2 {
		println(index, value) // Print index and value of arr2
	}

	// Multidimensional array
	var matrix [2][3]int // Declare a 2x3 matrix
	matrix[0][0] = 1     // Assign value to the first element
	matrix[0][1] = 2     // Assign value to the second element
	matrix[1][2] = 3     // Assign value to the third element

	// Declare and initialize a multidimensional array
	// matrix := [2][3]int{{1,2,3}, {4,5,6}}

	// Print elements of the multidimensional array
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			println(matrix[i][j]) // Print each element of the matrix
		}
	}
}