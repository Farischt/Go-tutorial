package main

import (
	"fmt"
	"reflect"
	"tutorial/common"
)

func main() {
	const conferenceName string = "GopherCon"
	const conferenceTickets int16 = 100

	var remainingTickets int16 = 100
	var bookings = make(map[string]common.Booking)

	common.GreetUser(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var userTickets int16

		// First name check
		common.SetUserInput("Enter your name: ", &firstName)
		for !common.IsNameValid(firstName) {
			fmt.Printf("Sorry, your name must be at least 2 characters long. \n")
			common.SetUserInput("Enter your name: ", &firstName)
		}

		// Last name check
		common.SetUserInput("Enter your lastname: ", &lastName)
		for !common.IsNameValid(lastName) {
			fmt.Printf("Sorry, your last name must be at least 2 characters long. \n")
			common.SetUserInput("Enter your lastname: ", &lastName)
		}

		// Email check
		common.SetUserInput("Enter your email: ", &email)
		for !common.IsEmailValid(email) {
			fmt.Printf("Sorry, your email must be at least 4 characters long and contain @ and . \n")
			common.SetUserInput("Enter your email: ", &email)
		}

		// Tickets check
		common.SetUserInputInt("Enter the number of tickets you want to buy: ", &userTickets)
		for reflect.TypeOf(userTickets).Kind() != reflect.Int16 || userTickets <= 0 || !common.IsAvailableTickets(remainingTickets, userTickets) {

			if reflect.TypeOf(userTickets).Kind() != reflect.Int16 || userTickets <= 0 {
				fmt.Println("Sorry, you must enter a valid positive number")
			} else {
				fmt.Printf("Sorry, we don't have %v tickets available. We only have %v remaining tickets!\n", userTickets, remainingTickets)
			}
			common.SetUserInputInt("Enter the number of tickets you want to buy: ", &userTickets)
		}

		remainingTickets = remainingTickets - userTickets

		if bookings[email].Email == "" {
			bookings[email] = common.Booking{FirstName: firstName, LastName: lastName, Email: email, Tickets: userTickets}
		} else {
			if entry, ok := bookings[email]; ok {
				entry.Tickets = entry.Tickets + userTickets
				bookings[email] = entry
			}
		}

		fmt.Printf("Thank you %v %v for your purchase. You have bought %v tickets. \n", firstName, lastName, userTickets)

		if remainingTickets == 0 {
			fmt.Println("We are now sold out !")
			break
		}
	}

	common.PrintBookingsMap(bookings)
}
