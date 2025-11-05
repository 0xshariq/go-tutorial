package main

import "fmt"

// num variable is pass as by value
func changeNumByValue(num int)  {
	num = 5
	fmt.Println("In changeNumByValue func", num) 
}

// num variable is pass as by reference
func changeNumByReference(num *int)  {
	*num = 5
	fmt.Println("In changeNumByReference func", *num)
}

func main()  {
	num := 1

	changeNumByValue(num)
	fmt.Println("After by value In main func", num)

	changeNumByReference(&num)

	fmt.Println("After by value In main func", num)
}
