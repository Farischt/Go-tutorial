package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
	"tutorial/common"
	"tutorial/files"
)

var wg = sync.WaitGroup{}

func sendTicket(tickets int16, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var message = fmt.Sprintf("Sending %v ticket(s) for %v %v", tickets, firstName, lastName)
	fmt.Println("_______________________________________________________")
	fmt.Printf("\n%v to %v \n", message, email)
	fmt.Println("_______________________________________________________")
	wg.Done()
}

func bookTickets() {
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

		if entry, ok := bookings[email]; ok {
			entry.Tickets = entry.Tickets + userTickets
			bookings[email] = entry
		} else {
			bookings[email] = common.Booking{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Tickets:   userTickets,
			}
		}

		fmt.Printf("Thank you %v %v for your purchase. You have bought %v tickets. \n", firstName, lastName, userTickets)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		if remainingTickets == 0 {
			fmt.Println("We are now sold out !")
			break
		}
	}
	wg.Wait()
}

func menu() {
	var choice int
	fmt.Println("Welcome please choose an action")
	fmt.Println("1. Book a ticket")
	fmt.Println("2. Look for go occurence in text file")
	fmt.Println("3. Exit")
	fmt.Printf("Your choice: ")
	fmt.Scan(&choice)

	for reflect.TypeOf(choice).Kind() != reflect.Int || choice < 1 || choice > 3 {
		fmt.Println("Invalid choice, please choose a number between 1 and 3")
		fmt.Printf("Your choice: ")
		fmt.Scan(&choice)
	}

	switch choice {
	case 1:
		fmt.Println("You chose to book a ticket")
		bookTickets()
	case 2:
		files.Main()
	case 3:
		fmt.Println("You chose to exit")
		return
	}
}

func main() {
	menu()
}
