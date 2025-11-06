package auth

import "fmt"

func Authenticate(user, password string) {
	fmt.Println("Login user with", user, password)
}