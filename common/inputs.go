package common

import (
	"fmt"
	"strings"
)

/*
 * Check if the name is valid
 * @param str string to check
 * @return bool true if the name is valid
 */
func IsNameValid(str string) bool {
	return len(str) >= 2
}

func IsEmailValid(email string) bool {
	return len(email) >= 4 && strings.Contains(email, "@") && strings.Contains(email, ".")
}

func IsAvailableTickets(remainingTickets int16, userTickets int16) bool {
	return remainingTickets >= userTickets
}

func SetUserInput(message string, value *string) {
	fmt.Printf("%v", message)
	fmt.Scan(value)
}

func SetUserInputInt(message string, value *int16) {
	fmt.Printf("%v", message)
	fmt.Scan(value)
}
