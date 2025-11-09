package main

import "fmt"

func fibonacci(fib []int, num int) []int {
	for i := 2; i < num; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func main() {
	var num int
	fmt.Print("Enter number of terms of Fibonacci series: ")
	fmt.Scan(&num)

	// make slice with required length and default capacity
	fib := make([]int, num)
	fib[0] = 0
	fib[1] = 1

	result := fibonacci(fib, num)
	fmt.Println("Fibonacci series:", result)
}
