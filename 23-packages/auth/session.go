package auth

import "fmt"

func GetSessionInfo(username string) {
	fmt.Printf("Session info for user: %s\n", username)
}