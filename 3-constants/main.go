package main

import "fmt"

func main(){
	const (
		pi       = 3.14159
		appName  = "GoConstantsApp"
		maxUsers = 1000
	)

	fmt.Println("Application Name:", appName)
	fmt.Println("Value of Pi:", pi)
	fmt.Println("Maximum Users Allowed:", maxUsers)
}