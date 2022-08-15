package main

import (
	"bookings/helper"
	"bookings/booking"
	"fmt"
)

const (
	locationName string = "Crookham Street Social"
	emailAddress string = "enquiries@crookhamstreet.co.uk"
)

var bookings = booking.Bookings
var remainingPeople = helper.RemainingPeople

func main() {
	greetUsers()

	for remainingPeople > 0 && len(bookings) < int(helper.MaxPeople) {
		remainingPeople = remainingPeople - booking.BookingInformation(emailAddress, locationName)

		if helper.AtCapacity() {
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
		helper.MaxPeople,
		remainingPeople)
	fmt.Printf("For more information on how to book, please contact us at %s.\n", emailAddress)
}

// loops through the bookings and just grabs the names of the bookings
func getNames() []string {
	names := []string{}

	for _, booking := range booking.Bookings {
		names = append(
			names, 
			fmt.Sprintf(
				"%s %s",
				booking["firstName"],
				booking["lastName"]))
	}

	return names
}
