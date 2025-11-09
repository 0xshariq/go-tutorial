package main

import "fmt"

func main()  {
	var str string
	var count int

	fmt.Println("Enter a string : ")
	fmt.Scan(&str)

	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}

	fmt.Printf("The number of vowels in the string \"%s\" is %d\n", str, count)
}
