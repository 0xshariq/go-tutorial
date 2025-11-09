package main

import "fmt"

func main()  {
	var N float32
	fmt.Println("Enter a temperature in Celsius : ")
	fmt.Scan(&N)

	fahrenheit := (N * 9 / 5) + 32

	fmt.Printf("The temperature in Fahrenheit is %f\n", fahrenheit)
}