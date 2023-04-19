package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// variable and constant declarations
	var conferenceName = "go conference"
	const conferenceTickets = 50
	remainingTickets := 50
	var name string
	var email string
	//var userTickets uint

	// assign values
	//name = "John"
	//userTickets = 10

	// printing strings
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("we have", remainingTickets, "number of thickets remaining")
	//fmt.Println("Get your tickets here to attend", ",your name is", name, "and your tickets is", userTickets)
	//fmt.Println(conferenceName)
	//fmt.Print("Hello")
	//fmt.Println("Hello World")

	// taking user input
	fmt.Scan(&name)
	fmt.Scan(&email)

	// arrays
	var bookings = [5]string{}
	var bookings2 = [5]string{"Felix a", "John b", "Dagba c", "Vincent d"}
	var bookings3 [5]string

	// adding to arrays
	bookings[0] = "Esther"
	bookings3[0] = "Bart"
	bookings2[4] = "Maxwell"

	// slices
	var sliceBookings []string
	sliceBookings2 := []string{}

	// adding to slice
	sliceBookings = append(sliceBookings, "Name")
	//sliceBookings2 = append(sliceBookings2, "name")
	//fmt.Println(sliceBookings)

	// for loops
	//for {
	//	fmt.Printf("%v", bookings)
	//}

	//for _, booking := range bookings2 {
	//	//var firstname = strings.Fields(booking)[0]
	//	//sliceBookings2 = append(sliceBookings2, firstname)
	//	sliceBookings2 = append(sliceBookings2, strings.Fields(booking)[0])
	//}
	fmt.Printf("%v", sliceBookings2)
	count := 10
	statement := true
	for count >= 0 && statement != false {
		fmt.Println("count")
		count = count - 1
	}

	// if else statements

	//if statement == false {
	//	fmt.Println("statement is true")
	//} else if statement == true {
	//	fmt.Println("statement is not true")
	//} else {
	//	fmt.Println("statement is false")
	//}

	// maps
	var myMappie map[string]string
	var myMap = make(map[string]string)
	myMappie["firstname"] = "nami"
	myMap["firstname"] = "mani"
	myMap["lastname"] = "lastie"
	myMap["email"] = "email@example.com"
	var ticks uint = 9
	myMap["tickets"] = strconv.FormatUint(uint64(ticks), 10)

	// struct
	type userDictionary struct {
		firstName   string
		lastname    string
		email       string
		dateOfBirth time.Time
		tickets     uint64
	}
}

func goSwitches() {
	// switch
	city := "London"
	lowerCase := strings.ToLower(city)
	switch lowerCase {
	case "new york":
		print("new york")
	case "london":
		println("london")
	case "singapore", "tokyo":
		fmt.Println("singapore")
	default:
		print("no valid city")
	}
}

func learnGo() {
	// fmt.Printf("conferenceName is of type %T, while conferenceTickets is of type %T and remainingTickets is of type %T. \n", conferenceName, conferenceTickets, remainingTickets)
	// data types
	// strings
	//const conferenceTickets uint = 50
	//remainingTickets := conferenceTickets // short variable declaration, does not work with constants and if you want to explicitly define a type
	//var userName string
	//
	//userName = "Tom"
	//fmt.Println(userName)
	//
	//// integer
	//var userTickets int
	//
	//userTickets = 4
	//fmt.Println(userTickets)
	//fmt.Printf("user %v booked a total of %v tickets.\n", userName, userTickets)
	//
	//// firstName, lastName, email, userTicks := getUserInput()
	//
	//remainingTickets = remainingTickets - userTicks
	//
	//fmt.Printf("Thank you %v %v for booking %v tickets, You would receive a confirmation email at %v. \n", firstName, lastName, userTicks, email)
	//// arrays
	//// arrays in go have fixed sizes
	//
	//var bookings = [50]string{"Josh", "James", "Jason"}
	//// adding elements to array
	//bookings[3] = firstName + " " + lastName
	//
	//fmt.Printf("the whole array is %v\n", bookings)
	//fmt.Printf("the first element is %v\n", bookings[0])
	//fmt.Printf("the fourth element is %v\n", bookings[3])
	//fmt.Printf("the type of the array is %T\n", bookings)
	//fmt.Printf("the length of the array is %v\n", len(bookings))

}
