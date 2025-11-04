package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 1; i <= 5; i++ {
		// break : exits the loop when i equals 3
		if i == 3 {
			break
		}
		// continue : skips the iteration when i equals 2
		if i == 2 {
			continue
		}
		fmt.Println(" For Iteration:", i)
	}

	// While-like loop using for
	// while keyword does not exist in Go
	i := 1
	for i <= 5 {
		fmt.Println(" While Iteration:", i)
		i++
	}
	// infinite loop
	// Uncomment the following lines to see an infinite loop in action
	/*
		for {
			fmt.Println(" Infinite Loop")
		}
	*/

	// range
	for i := range 3 {
		fmt.Println(" Range Iteration:", i)
	}
}
