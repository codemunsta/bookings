package main

import (
	"bookings/helper"
	"fmt"
	"time"
)

func greetUsers() {
	println("welcome to our conference")
	fmt.Printf("the name of the conferenece is %v\n", conferenceName)
	fmt.Println("Welcome to our", conferenceName, "booking application")
}

func getUserInput() (string, string, string, uint) {
	// accepting data
	var firstName string
	var lastName string
	var email string
	var userTicks uint

	fmt.Print("Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter your number of tickets: ")
	fmt.Scan(&userTicks)
	return firstName, lastName, email, userTicks
}

func sendMail(user helper.UserDictionary) {
	var ticketEmailDetails = fmt.Sprintf(
		"Hey %v %v, your %v tickets has been sent to your mail @ %v. \n thank you \n",
		user.FirstName, user.Lastname, user.NumberOfTickets, user.Email,
	)
	time.Sleep(10 * time.Second)
	fmt.Println("###########")
	fmt.Printf("sending email... received %v", ticketEmailDetails)
	fmt.Println("###########")
	wait.Done()
}

func bookTicket(firstName string, lastName string, email string, userTicks uint, remainingTickets uint) {
	user := helper.CreateMap(firstName, lastName, email, userTicks)
	userDic := helper.CreateUserStruct(firstName, lastName, email, userTicks)
	fmt.Printf("Thank you %v %v for booking %v tickets, You would receive a confirmation email at %v. \n", firstName, lastName, userTicks, email)
	fmt.Printf("remaining tickets : %v \n", remainingTickets)
	sliceBookings = append(sliceBookings, firstName+" "+lastName)
	users = append(users, user)
	usersDictionary = append(usersDictionary, userDic)
	firstNames, newFirstnames, firstNamesList := helper.GetFirstNames(sliceBookings, users, usersDictionary)
	fmt.Printf("list of bookings is %v \n", users)
	fmt.Printf("%v\n", firstNames)
	fmt.Printf("%v \n", newFirstnames)
	fmt.Printf("firstnames from struct: %v \n", firstNamesList)
	wait.Add(1)
	go sendMail(userDic)
}

func sendErrorMessage(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
	if !isValidName {
		fmt.Println("Firstname or Lastname enter is too short")
	}
	if !isValidEmail {
		fmt.Println("email enter is not valid")
	}
	if !isValidTicketNumber {
		fmt.Printf("Number of tickets entered cannot be less than 1 or more than %v \n", remainingTickets)
	}
}
