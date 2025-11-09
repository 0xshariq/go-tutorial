package main

import "fmt"

func main() {
	var num, table int
	fmt.Println("Enter Number : ")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		table = num * i
		fmt.Printf("%d X %d = %d\n", num, i, table)
	}
}
