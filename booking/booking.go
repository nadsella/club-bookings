package booking

import (
	"bookings/helper"
	"strconv"
	"fmt"
)

var Bookings []map[string]string
var remainingPeople = helper.RemainingPeople

// grabs all the booking info from the user
// prints out a message for the booking and returns the number of people for the booking
func BookingInformation(emailAddress string, locationName string) uint {
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

	isValidNumPeople := numPeople > 0 && numPeople > remainingPeople

	// for now just a basic if statement to validate user data
	if !helper.ValidateUserData(firstName, lastName, userEmailAddress) {
		return 0
	}

	// we can't add more people that what is left, for now we log an error and return 0
	if isValidNumPeople {
		fmt.Printf(
			"We can't reserve %d spaces as there are only %d left\n",
			numPeople,
			remainingPeople)
		return 0
	}

	bookTickets(firstName, lastName, emailAddress, numPeople)

	fmt.Printf(
		"Thank you %s %s for your booking at %s for %d people. You will receive a confirmation email at %s.\n",
		firstName,
		lastName,
		locationName,
		numPeople,
		userEmailAddress)

	return numPeople
}

// appends the name to the bookings
func bookTickets(
	firstName string,
	lastName string,
	emailAddress string,
	numPeople uint) {

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["emailAddress"] = emailAddress
	userData["numPeople"] = strconv.FormatUint(uint64(numPeople), 10)

	// add the user to the bookings array
	Bookings = append(Bookings, userData)
}