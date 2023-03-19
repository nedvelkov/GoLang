package main

import (
	"booking-app/userInput"
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]userInput.UserData, 0)

func main() {

	greetUsers()

	for remainingTickets > 0 {

		var user = userInput.GetUserData(remainingTickets)

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

func bookTickets(user userInput.UserData) {
	remainingTickets -= user.NumberOfTickets
	bookings = append(bookings, user)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a conformation email at %v \n", (user.FirstName + " " + user.LastName), user.NumberOfTickets, user.Email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func printFirstNames() {
	firstNames := []string{}

	for _, user := range bookings {
		firstNames = append(firstNames, user.FirstName)
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)
}
