package main

import (
	"bookings/helper"
	"fmt"
	"os"
	"strconv"
)

const (
	locationName string = "Crookham Street Social"
	emailAddress string = "enquiries@crookhamstreet.co.uk"
	maxPeople  uint = 80
)

var bookings []map[string]string
var remainingPeople uint = maxPeople

func main() {
	greetUsers()

	for remainingPeople > 0 && len(bookings) < int(maxPeople) {
		remainingPeople = remainingPeople - bookingInformation()

		if atCapacity() {
			break
		}

		fmt.Printf("We only have room for %d more people\n", remainingPeople)
	}

	names := getNames()
	fmt.Printf("These are all of our bookings: %v.\n", names)
}

// greets the users
func greetUsers() {
	fmt.Printf("Welcome to the %s table booking system.\n", locationName)
	fmt.Printf("We can currently only accommodate %d people for each event. This event currently has %d seats left.\n",
		maxPeople,
		remainingPeople)
	fmt.Printf("For more information on how to book, please contact us at %s.\n", emailAddress)
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
	bookings = append(bookings, userData)
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
func getNames() []string {
	names := []string{}

	for _, booking := range bookings {
		names = append(
			names, 
			fmt.Sprintf(
				"%s %s",
				booking["firstName"],
				booking["lastName"]))
	}

	return names
}
