package main

import (
	"fmt"
	"userauth/auth"
)

func main() {

	fmt.Println("Multi-File Authentication System")

	var username string
	var password string

	fmt.Scanf("%s %s", &username, &password)

	user := auth.User{
		Username: username,
		Password: password,
	}
	authenticated := auth.Authenticate(user)
	if authenticated {
		fmt.Println("User authenticated successfully")
	} else {
		fmt.Println("Authentication failed")
	}
}
