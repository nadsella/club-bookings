package main

import "fmt"

func main() {
	locationName := "Crookham Street Social"
	const maxPeople = 80
	remainingPeople := 80

	fmt.Println("Welcome to the", locationName, "table booking system.")
	fmt.Printf("We can currently only accommodate %d people for each event. This event currently has %d seats left.\n",
		maxPeople,
		remainingPeople)
	fmt.Println("For more information, please contact us at enquiries@crookhamstreet.co.uk.")
}
