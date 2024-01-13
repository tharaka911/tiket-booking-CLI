package main

import (
	"booking-app/helper"
	"fmt"
	
)

const conferenceTickets int = 50
var remainingTickets uint = 50
var conferenceName = "Go Conference"
//creating the slice
//var bookings = []string{}
//initilizing a list of map
//var bookings = make([]map[string]string,0)
var bookings = make([]UserData,0)

type UserData struct {

	firstName string
	lastName string
	email string
	numberOfTickets uint

}



func main() {

	

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets ,remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book ticket in system
			bookTicket(firstName , lastName , email , userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// exit application if no tickets are left
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}

	// Switch statement example
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}



func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asking for user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {

	remainingTickets = remainingTickets - userTickets

	
	var userData = UserData {
		firstName:firstName,
		lastName:lastName,
		email:email,
		numberOfTickets: userTickets,

	}
	

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
