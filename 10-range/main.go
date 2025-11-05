package main

import "fmt"

func main() {

	// slice example with range
	nums := []int{10, 20, 30, 40, 50} // intializing a slice
	sum := 0 // variable to hold sum
	mul := 1 // variable to hold multiplication

	// iterating over slice using range
	for i, num := range nums {
		fmt.Println(num, i)
		sum += num
		mul *= num
	}
	fmt.Println("Total sum:", sum) // print total sum
	fmt.Println("Total multiplication:", mul) // print total multiplication	




	// map example with range
	m := map[string]int{ // initializing a map
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}

	// iterating over map using range
	for key := range m {
		fmt.Println(key, m[key])
	}

	// string example with range
	str := "Hello"
	for i, ch := range str {
		fmt.Println(i, string(ch)) // ch will print in ascii value / rune, so we have convert into string
	}
}