package booking

import (
	"bookings/helper"
	"fmt"
	"time"
)

var Bookings []UserData
var MaxPeople uint = 80
var RemainingPeople uint = MaxPeople

type UserData struct {
	firstName string
	lastName  string
	email     string
	numPeople uint
}

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

	isValidNumPeople := numPeople > 0 && numPeople > RemainingPeople

	// for now just a basic if statement to validate user data
	if !helper.ValidateUserData(firstName, lastName, userEmailAddress) {
		return 0
	}

	// we can't add more people that what is left, for now we log an error and return 0
	if isValidNumPeople {
		fmt.Printf(
			"We can't reserve %d spaces as there are only %d left\n",
			numPeople,
			RemainingPeople)
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

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     emailAddress,
		numPeople: numPeople,
	}

	// add the user to the bookings array
	Bookings = append(Bookings, userData)

	go sendTicket(userData.numPeople, userData.firstName, userData.lastName)
}

// loops through the bookings and just grabs the names of the bookings
func GetNames() []string {
	names := []string{}

	for _, booking := range Bookings {
		names = append(
			names,
			fmt.Sprintf(
				"%s %s",
				booking.firstName,
				booking.lastName))
	}

	return names
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

// emulate sending ticket via email
func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(5 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending %v\n", ticket)
}
