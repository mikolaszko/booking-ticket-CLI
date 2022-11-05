package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceName string = "Go Conference"

const conferenceTickets uint8 = 50

var remainingTickets uint8 = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	for len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// wg.Add(1)
			//could add multi thread  with go
			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNamesOfAttendees()
			fmt.Printf("Here are first names of our attendees: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out! Comeback next year")
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address has to contain @ sing")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
		// wg.Wait()
	}

}

func greetUsers() {
	fmt.Printf("Welcome to our booking application for %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets!!!")
}

func getFirstNamesOfAttendees() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	time.Sleep(1 * time.Second)
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want to buy?")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint8, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \n to email address %v \n", ticket, email)
	fmt.Println("###############")
	// wg.Done()
}
