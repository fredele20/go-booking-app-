package main

import "strings"

func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTicketsNumber
}
