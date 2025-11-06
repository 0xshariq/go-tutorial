package main

import (
	"fmt"
	"go-tutorials/23-packages/auth"
	"go-tutorials/23-packages/user"
	"github.com/fatih/color"
)

func main() {
	auth.Authenticate("sharique", "123456789")

	auth.GetSessionInfo("sharique")

	user := user.NewUser("sharique","user@example.com","12345689876")
	fmt.Println(user)
	color.Green("User created successfully! %v", user)

	user.SetSessionID("ABC123XYZ")
	fmt.Println("Updated User with SessionID:", user)
}
