package main

import "fmt"

func main() {
	var age int = 25
	var name string = "John"
	var height float32 = 5.9
	var isEmployed bool = true // isEmployed := true (short declaration not used here)
	experience := 30 // Short declaration

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Employed:", isEmployed)
	fmt.Println("Years of Experience:", experience)
}
