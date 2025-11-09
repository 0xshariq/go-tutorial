package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length int
	fmt.Print("Enter the length of the password: ")
	fmt.Scan(&length)

	// All valid characters for the password
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	strLen := len(validChars)

	// Create a local random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(strLen)
		password[i] = validChars[randomIndex]
	}

	fmt.Printf("Generated password: %s\n", string(password))
}
