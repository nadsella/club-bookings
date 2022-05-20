package main

import "fmt"

const (
	locationName = "Crookham Street Social"
	emailAddress = "enquiries@crookhamstreet.co.uk"
	maxPeople    = 80
)

var bookings []string

func main() {
	var remainingPeople uint = 80

	fmt.Printf("Welcome to the %s table booking system.\n", locationName)
	fmt.Printf("We can currently only accommodate %d people for each event. This event currently has %d seats left.\n",
		maxPeople,
		remainingPeople)
	fmt.Printf("For more information on how to book, please contact us at %s.\n", emailAddress)

	numPeople := bookingInformation()
	remainingPeople = remainingPeople - numPeople

	// let people know we are close to capacity
	// this should go out as an email to notify the bookings team
	if remainingPeople < 5 {
		fmt.Println("We are nearly at capacity!")
	}

	fmt.Printf("We now only have room for %d more people.\n", remainingPeople)
	fmt.Printf("These are all of our bookings: %v.\n", bookings)
}

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
