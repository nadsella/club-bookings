package helper

import (
	"fmt"
	"strings"
)

const (
	minNameLen int  = 2
	MaxPeople  uint = 80
)

var RemainingPeople uint = MaxPeople

func ValidateUserData(userFirstName string, userLastName string, userEmail string) bool {
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

// checks if we're close to capacity, and if we're at capacity will let the user know
// and returns a boolean
func AtCapacity() bool {
	// let people know we are close to capacity
	// this should go out as an email to notify the bookings team
	if RemainingPeople == 0 {
		fmt.Println("We are at capacity!")

		// we return true when were at capacity
		return true
	}

	if RemainingPeople < 5 {
		fmt.Println("We are nearly at capacity!")
	}

	return false
}
