package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	count := 0
	for i := 1; i <= 1000000000; i++ {
		count++ // simple increment
	}

	elapsed := time.Since(start)
	fmt.Println("Looped up to 1 billion!")
	fmt.Println("Time taken:", elapsed)
}
