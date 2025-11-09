package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Enter a value to reverse:")
	fmt.Scan(&input)

	// Try to detect type (int, float, or string)
	if intVal, err := strconv.Atoi(input); err == nil {
		// Integer reversal
		reversed := 0
		num := intVal
		for num != 0 {
			digit := num % 10
			reversed = reversed*10 + digit
			num /= 10
		}
		fmt.Printf("Reversed integer: %d\n", reversed)
		return
	}

	if floatVal, err := strconv.ParseFloat(input, 64); err == nil {
		// Float reversal (by converting to string)
		str := fmt.Sprintf("%f", floatVal)
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Printf("Reversed float: %s\n", string(runes))
		return
	}

	// Otherwise, treat as a string
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Printf("Reversed string: %s\n", string(runes))
}
