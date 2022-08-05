package main

import (
	"fmt"
	"strings"
)

func validateUserData(userFirstName string, userLastName string, userEmail string) bool {
	isValidName := len(userFirstName) >= minNameLen && len(userLastName) >= minNameLen
	isValidEmail := strings.Contains(userEmail, "@")

	// if it's fine, then just return true
	if isValidName && isValidEmail {
		return true
	}

	if !isValidName {
		fmt.Println("We can't reserve these spaces as your first or last name is not long enough.")
	}

	if !isValidEmail {
		fmt.Println("We can't reserve these spaces as your email address doesn't contain an @.")
	}

	// and return false
	return false
}
