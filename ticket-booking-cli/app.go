package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName string = "Cloud Conference"
const conferenceTicket uint = 100
const ticketPrice float64 = 250.00
var remainingTicket uint = conferenceTicket
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	
	firstName, lastName, email, numberOfTicket := getUserInputs()
	fmt.Println("------------------------------------------------------------------------")

	isValidName, isValidEmail, isValidNumberOfTicket := validateUserInput(firstName, lastName, email, numberOfTicket, remainingTicket)

	if isValidName && isValidEmail && isValidNumberOfTicket {
		bookTickets(firstName, lastName, email, numberOfTicket)
		wg.Add(1)
		go sendTicket(firstName, lastName, email, numberOfTicket)

		firstNames := getFirstNames()
		fmt.Printf("The first names of our bookings data are %v\n", firstNames)
		fmt.Println("Booking Saved Successfully!")
		fmt.Println("------------------------------------------------------------------------")

		
		if remainingTicket == 0 {
			fmt.Println("All tickets have been sold out, we are sorry for the inconvenience.")
			fmt.Println("------------------------------------------------------------------------")
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid name input, it is too short. Please try again!")
			fmt.Println("------------------------------------------------------------------------")
		}
		if !isValidEmail {
			fmt.Println("Invalid email input, please try again!")
			fmt.Println("------------------------------------------------------------------------")
		}
		if !isValidNumberOfTicket {
			fmt.Println("Invalid number of ticket input, please try again!")
			fmt.Println("------------------------------------------------------------------------")
		}
	}
	wg.Wait()
}


func greetUsers() {
	fmt.Println("#########################################################################")
	fmt.Printf("Welcome to %v booking application\n", conferenceName )
	fmt.Printf("We have %v number of ticket for the event, and %v are still unsold\n", conferenceTicket, remainingTicket)
	fmt.Printf("Each ticket cost $%v\n", ticketPrice)
	fmt.Println("Get your ticket here to attend!")
	fmt.Println("#########################################################################")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numberOfTicket uint

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email address?")
	fmt.Scan(&email)

	fmt.Println("How many ticket do you wish to purchase?")
	fmt.Scan(&numberOfTicket)

	return firstName, lastName, email, numberOfTicket
}

func bookTickets(firstName, lastName, email string, numberOfTicket uint) {
	remainingTicket = remainingTicket - numberOfTicket
	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		numberOfTicket: numberOfTicket,
	}
	bookings = append(bookings, userData)
	fmt.Printf("Listing all bookings: %v\n", bookings)

	fmt.Printf("The user %v %v purchased %v ticket(s)\n, a mail will be sent to your email address %v for confirmation.\n", firstName, lastName, numberOfTicket, email)
	fmt.Printf("There are %v tickets remaining for the event.\n", remainingTicket)
}

func sendTicket(firstName, lastName, email string, numberOfTicket uint) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("Hi %v %v, thank you for booking %v ticket(s) for %v. A mail will be sent to your email address %v for confirmation.", firstName, lastName, numberOfTicket, conferenceName, email)
	
	fmt.Println("#########################################################################")
	fmt.Printf("Sending ticket:\n %v \nto the email: %v\n", ticket, email)
	fmt.Println("#########################################################################")
	wg.Done()
}