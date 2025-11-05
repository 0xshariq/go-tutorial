package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

// Example of function returning multiple values
func getLanguages() (string, string, string) {
	return "golang", "python", "java"
}

// Example of passing a function as an argument
func processIt(fn func(a int) int) {
	fn(5)
}

// Example of returning a function from another function
func process() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func main() {
	var a, b int
	fmt.Println(getLanguages())
	fmt.Println("Enter a and b : ")
	fmt.Scan(&a, &b)                 // taking input from user
	fmt.Println("Sum : ", add(a, b)) // calling add function and printing the result
	fmt.Println("Sub : ", sub(a, b)) // calling sub function and printing the result
	fmt.Println("Mul : ", mul(a, b)) // calling mul function and printing the result
	fmt.Println("Div : ", div(a, b)) // calling div function and printing the result

	fun := func(a int) int {
		return 2
	}

	processIt(fun)

	fun2 := process()
	fmt.Println(fun2(5))
}
