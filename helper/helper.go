package helper

import (
	"fmt"
	"strconv"
	"strings"
)

type UserDictionary struct {
	FirstName       string
	Lastname        string
	Email           string
	NumberOfTickets uint
}

func ValidateUserInput(firstName string, lastName string, email string, userTicks uint, remainingTickets uint) (bool, bool, bool) {
	// user validation
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@") && strings.Contains(email, ".")
	var isValidTicketNumber = userTicks > 0 && userTicks <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func GetFirstNames(sliceBookings []string, users []map[string]string, userLists []UserDictionary) ([]string, []string, []string) {
	var firstNames []string
	for _, booking := range sliceBookings { // underscores (_) are used to specify variables that you would not use
		var names = strings.Fields(booking)
		var firstNameUser = names[0]
		firstNames = append(firstNames, firstNameUser)
	}
	var newFirstnames []string
	for _, user := range users {
		newFirstnames = append(newFirstnames, user["firstname"])
	}
	var firstNamesList []string
	for _, dic := range userLists {
		firstNamesList = append(firstNamesList, dic.FirstName)
	}
	return firstNames, newFirstnames, firstNamesList
}

func CreateMap(firstName string, lastName string, email string, userTicks uint) map[string]string {
	var userDic = make(map[string]string)
	userDic["firstname"] = firstName
	userDic["lastname"] = lastName
	userDic["email"] = email
	userDic["tickets"] = strconv.FormatUint(uint64(userTicks), 10)
	fmt.Printf("user dic is %v \n", userDic)
	return userDic
}

func CreateUserStruct(firstName string, lastName string, email string, userTicks uint) UserDictionary {
	var userDic = UserDictionary{
		FirstName:       firstName,
		Lastname:        lastName,
		Email:           email,
		NumberOfTickets: userTicks,
	}
	fmt.Printf("user struct dictionary: %v \n", userDic)
	return userDic
}
