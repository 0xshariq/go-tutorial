package main

import "fmt"

func printAllPrimeNumbers(N int) int {
	count := 0
	for num := 2; num <= N; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num)
			count++
		}
	}
	fmt.Println()
	return count
}

func main(){
	var N int
	fmt.Println("Enter a number N : ")
	fmt.Scan(&N)

	result := printAllPrimeNumbers(N)
	fmt.Printf("All prime numbers up to %d are: %d\n", N, result)
}
