package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// simple switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("invalid")
	}

	// multiple condition switch working
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// type switch
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		case float32:
			fmt.Println("Float")
		default:
			fmt.Println("Invalid", t)
		}
	}

	whoami(6.0)
}
