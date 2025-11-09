package main

import (
	"fmt"
	"sync"
)

func main() {
	var N int
	fmt.Println("Enter a number N : ")
	fmt.Scan(&N)

	var wg sync.WaitGroup
	wg.Add(3) // We have three goroutines to wait for

	// Goroutine 1: Count even numbers
	go func() {
		defer wg.Done()
		evenCount := 0
		for i := 1; i <= N; i++ {
			if i%2 == 0 {
				evenCount++
			}
		}
		fmt.Printf("The number of even numbers from 1 to %d is %d\n", N, evenCount)
	}()

	// Goroutine 2: Count odd numbers
	go func() {
		defer wg.Done()
		oddCount := 0
		for i := 1; i <= N; i++ {
			if i%2 != 0 {
				oddCount++
			}
		}
		fmt.Printf("The number of odd numbers from 1 to %d is %d\n", N, oddCount)
	}()

	go func(){
		defer wg.Done()
		primeNumberCount := 0
		for num := 2; num <= N; num++ {
			isPrime := true
			for i := 2; i*i <= num; i++ {
				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primeNumberCount++
				fmt.Printf("%d ", num)
			}
		}
		fmt.Printf("The number of prime numbers from 1 to %d is %d\n", N, primeNumberCount)
	}()

	// Wait for both goroutines to finish
	wg.Wait()
	fmt.Println("All tasks completed.")
}
