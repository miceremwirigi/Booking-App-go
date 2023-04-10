package main

import "fmt"

func main (){
	const conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have %v tickets remaining out of %v total tickets\n",remainingTickets, conferenceTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint


	//ask user for these
	fmt.Println("Enter first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter  last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter email:")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %s %s. You have booked %v tickets. A confirmation email sent to %s\n",firstName,lastName, userTickets, email)
	bookings = append(bookings, firstName + " "+ lastName)

	remainingTickets-=userTickets
	fmt.Printf("Remaining tickets: %d\n", remainingTickets)

	fmt.Printf("All bookings: %v\n", bookings)
	



	
	

}
