package main

import (
	"bookings/booking"
	"fmt"
)

const (
	locationName string = "Crookham Street Social"
	emailAddress string = "enquiries@crookhamstreet.co.uk"
)

func main() {
	greetUsers()

	for booking.RemainingPeople > 0 && len(booking.Bookings) < int(booking.MaxPeople) {
		booking.RemainingPeople = booking.RemainingPeople - booking.BookingInformation(emailAddress, locationName)

		if booking.AtCapacity() {
			fmt.Println("At capacity")
			break
		}

		fmt.Printf("We only have room for %d more people\n", booking.RemainingPeople)
	}

	names := booking.GetNames()
	fmt.Printf("These are all of our bookings: %v.\n", names)
}

// greets the users
func greetUsers() {
	fmt.Printf("Welcome to the %s table booking system.\n", locationName)
	fmt.Printf("We can currently only accommodate %d people for each event. This event currently has %d seats left.\n",
		booking.MaxPeople,
		booking.RemainingPeople)
	fmt.Printf("For more information on how to book, please contact us at %s.\n", emailAddress)
}
