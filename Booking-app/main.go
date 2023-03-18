package main

import (
	"booking-app/userInput"
	"fmt"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

func main() {

	greetUsers()

	for remainingTickets > 0 {
		var firstName, lastName, email, tickets = userInput.GetUserData(remainingTickets)
		var user = UserData{firstName: firstName, lastName: lastName, email: email, numberOfTickets: tickets}

		bookTickets(user)
		fmt.Println(bookings)
		printFirstNames()

	}

	fmt.Println("Our conference is booked out. Come back next year.")
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func bookTickets(user UserData) {
	remainingTickets -= user.numberOfTickets
	bookings = append(bookings, user)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a conformation email at %v \n", (user.firstName + " " + user.lastName), user.numberOfTickets, user.email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func printFirstNames() {
	firstNames := []string{}

	for _, user := range bookings {
		firstNames = append(firstNames, user.firstName)
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)
}
