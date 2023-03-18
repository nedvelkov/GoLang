package userInput

import (
	"fmt"
	"strings"
)

func GetUserData(remainingTickets uint) (string, string, string, uint) {
	var firstName, lastName = getUserName()
	var email = getEmail()
	var tickets = getTickets(remainingTickets)
	return firstName, lastName, email, tickets

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
		fmt.Println("Enter number of tickets you want to buy")
		fmt.Scan(&numberOfTicket)

		if numberOfTicket > 0 && numberOfTicket <= remainingTickets {
			return numberOfTicket
		}

		fmt.Println("Enter ticket number is not valid,try again")
	}
}
