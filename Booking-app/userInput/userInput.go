package userInput

import (
	"fmt"
	"strings"
)

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

func GetUserData(remainingTickets uint) UserData {
	var firstName, lastName = getUserName()
	var email = getEmail()
	var tickets = getTickets(remainingTickets)
	return UserData{firstName, lastName, email, tickets}

}

func getUserName() (string, string) {
	var firstName string
	var lastName string
	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		if len(firstName) >= 2 && len(lastName) >= 2 {
			break
		}

		fmt.Println("Enter name is not valid,try again")
	}
	return firstName, lastName
}

func getEmail() string {
	var email string
	for {
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		if strings.Contains(email, "@") {
			return email
		}

		fmt.Println("Enter email is not valid,try again")
	}
}

func getTickets(remainingTickets uint) uint {
	var numberOfTicket uint
	for {
		fmt.Println("Enter numer of tickets you want to buy")
		fmt.Scan(&numberOfTicket)

		if numberOfTicket > 0 && numberOfTicket <= remainingTickets {
			return numberOfTicket
		}

	}
}
