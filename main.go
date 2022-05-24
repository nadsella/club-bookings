package main

import (
	"fmt"
	"strings"
)

const (
	locationName      = "Crookham Street Social"
	emailAddress      = "enquiries@crookhamstreet.co.uk"
	maxPeople    uint = 80
)

var bookings []string
var remainingPeople uint = 80

func main() {

	fmt.Printf("Welcome to the %s table booking system.\n", locationName)
	fmt.Printf("We can currently only accommodate %d people for each event. This event currently has %d seats left.\n",
		maxPeople,
		remainingPeople)
	fmt.Printf("For more information on how to book, please contact us at %s.\n", emailAddress)

	for {
		numPeople := bookingInformation()
		remainingPeople = remainingPeople - numPeople
		firstNames := getFirstNames()

		fmt.Printf("We now only have room for %d more people.\n", remainingPeople)
		fmt.Printf("These are all of our bookings: %v.\n", firstNames)

		if atCapacity() {
			break
		}
	}
}

// grabs all the booking info from the user
// prints out a message for the booking and returns the number of people for the booking
func bookingInformation() uint {
	var firstName string
	var lastName string
	var userEmailAddress string
	var numPeople uint

	//ask the user for their name
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&userEmailAddress)

	fmt.Println("Please enter the amount of people you need a table for:")
	fmt.Scan(&numPeople)

	// add the user to the bookings array
	bookings = append(bookings, fmt.Sprintf("%s %s", firstName, lastName))

	fmt.Printf(
		"Thank you %s %s for your booking at %s for %d people. You will receive a confirmation email at %s.\n",
		firstName,
		lastName,
		locationName,
		numPeople,
		userEmailAddress)

	return numPeople
}

// checks if we're close to capacity, and if we're at capacity will let the user know
// and returns a boolean
func atCapacity() bool {
	// let people know we are close to capacity
	// this should go out as an email to notify the bookings team
	if remainingPeople == 0 {
		fmt.Println("We are at capacity!")

		// we return true when were at capacity
		return true
	}

	if remainingPeople < 5 {
		fmt.Println("We are nearly at capacity!")
	}

	return false
}

// loops through the bookings and just grabs the first names
func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}
