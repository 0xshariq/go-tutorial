package main

import "fmt"

func main()  {
	var num int
	fmt.Println("Enter a number : ")
	fmt.Scan(&num)

	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit
		num = num / 10
	}

	fmt.Printf("The sum of digits is: %d\n", sum)
}
