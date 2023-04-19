package main

import (
	"bookings/helper"
	"fmt"
	"sync"
)

var wait = sync.WaitGroup{}

var conferenceName = "Go conference"

const conferenceTickets uint = 50

var sliceBookings []string

var users = make([]map[string]string, 0)

var usersDictionary = make([]helper.UserDictionary, 0)

var remainingTickets = conferenceTickets

// have to run main.go and helper.go or go run . (to run files in a folder in that current location

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTicks := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicks, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			if userTicks == remainingTickets {
				remainingTickets = remainingTickets - userTicks
				bookTicket(firstName, lastName, email, userTicks, remainingTickets)
				fmt.Println("Our conference is sold out")
				break
			} else {
				remainingTickets = remainingTickets - userTicks
				bookTicket(firstName, lastName, email, userTicks, remainingTickets)
			}
		} else {
			sendErrorMessage(isValidName, isValidEmail, isValidTicketNumber)
			continue
		}

		var noTicketsAvailable = remainingTickets == 0

		if noTicketsAvailable {
			fmt.Println("Our conference is sold out")
			break
		}
	}
	wait.Wait()
}
