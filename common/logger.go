package common

import (
	"fmt"
	"strings"
)

func PrintBookings(array []string) {
	for _, value := range array {
		if value != "" {
			fmt.Println(strings.Fields(value)[0])
		}
	}
}

func PrintFullBookings(array []string) {
	for _, value := range array {
		if value != "" {
			fmt.Println(value)
		}
	}
}

func PrintBookingsMap(bookingMap map[string]Booking) {
	for _, value := range bookingMap {

		fmt.Println("Booker")
		fmt.Println(value.FirstName)
		fmt.Println(value.LastName)
		fmt.Println(value.Email)
		fmt.Println(value.Tickets)
	}
}
