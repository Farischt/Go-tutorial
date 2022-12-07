package common

import (
	"fmt"
)

type Booking struct {
	FirstName string
	LastName  string
	Email     string
	Tickets   int16
}

func GreetUser(name string, conferenceTickets int16, remainingTickets int16) {
	fmt.Printf("Hello, and welcome to %v event. \n", name)
	fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the event")
}

// type City int64

// const (
// 	London City = iota
// 	Singapore
// 	Rabat
// 	Paris
// 	Unknown
// )

// func (city City) String() string {
// 	switch city {
// 	case London:
// 		return "London"
// 	case Singapore:
// 		return "Singapore"
// 	case Rabat:
// 		return "Rabat"
// 	case Paris:
// 		return "Paris"
// 	default:
// 		return "Unknown"
// 	}
// }
