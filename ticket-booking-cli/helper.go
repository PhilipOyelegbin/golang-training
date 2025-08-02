package main

import "strings"

func validateUserInput(firstName, lastName, email string, numberOfTicket, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidNumberOfTicket := numberOfTicket > 0 && numberOfTicket <= remainingTicket

	return isValidName, isValidEmail, isValidNumberOfTicket
}