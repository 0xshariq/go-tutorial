package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		fmt.Println("\n===== Simple CLI Menu App =====")
		fmt.Println("1. Factorial")
		fmt.Println("2. Fibonacci Series")
		fmt.Println("3. Check Prime Number")
		fmt.Println("4. Check Armstrong Number")
		fmt.Println("5. Count Vowels in a String")
		fmt.Println("6. Sum of Digits")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Enter your choice (1-7): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var n int
			fmt.Print("\nEnter a number to find factorial: ")
			fmt.Scan(&n)
			fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
			return
		case 2:
			var n int
			fmt.Print("\nEnter number of terms for Fibonacci series: ")
			fmt.Scan(&n)
			fmt.Printf("Fibonacci series of %d terms: ", n)
			fibonacci(n)
			return
		case 3:
			var n int
			fmt.Print("\nEnter a number to check if it is prime: ")
			fmt.Scan(&n)
			if isPrime(n) {
				fmt.Printf("%d is a Prime number\n", n)
			} else {
				fmt.Printf("%d is not a Prime number\n", n)
			}
			return
		case 4:
			var n int
			fmt.Print("\nEnter a number to check if it is Armstrong: ")
			fmt.Scan(&n)
			if isArmstrong(n) {
				fmt.Printf("%d is an Armstrong number\n", n)
			} else {
				fmt.Printf("%d is not an Armstrong number\n", n)
			}
			return
		case 5:
			var str string
			fmt.Print("\nEnter a string to count vowels: ")
			fmt.Scan(&str)
			countVowel(str)
			return
		case 6:
			var n int
			fmt.Print("\nEnter a number to find sum of digits: ")
			fmt.Scan(&n)
			sumOfDigits(n)
			return
		case 7:
			fmt.Println("Exiting program... Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isArmstrong(n int) bool {
	temp := n
	sum := 0
	digits := int(math.Log10(float64(n))) + 1

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func countVowel(str string) {
	count := 0
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	fmt.Printf("The number of vowels in the string \"%s\" is %d\n", str, count)
}

func sumOfDigits(num int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit
		num = num / 10
	}

	fmt.Printf("The sum of digits is: %d\n", sum)
}
