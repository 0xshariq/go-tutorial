package main

import "fmt"

func primeNumber(n int) (string) {
	if n <= 1 {
		return "Not a prime number"
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return "Not a prime number"
		}
	}
	return "Is a prime number"
}

func main()  {
	var num int
	fmt.Println("Enter a number : ")
	fmt.Scan(&num)

	result := primeNumber(num)

	fmt.Printf("The number %d %s\n", num, result)
}