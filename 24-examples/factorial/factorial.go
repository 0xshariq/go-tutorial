package main

import "fmt"

func factorial(num uint64) uint64 {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	var num uint64
	fmt.Println("Enter number : ")
	fmt.Scan(&num)
	result := factorial(num)
	fmt.Printf("Factorial of %d is : %d\n", num, result)
}
