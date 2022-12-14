package main

import (
	"fmt"
	"strings"
"booking-app/helper"
"booking-app/helper2"
)

var partyName = "New Year Party"

const TotallTickets = 50

var RemainingTickets uint = 50
var Bookings []string

func main() {

	for {

		// call function to great user
		greetUsers()

		// Ask user information to register
		FirstName, LastName, email, userTickets := inputParameters()

		// use INput validation function
		isValidName, isValidEmail, isValidTicketQuantity := helper.ValidateUserInputs(FirstName, LastName, email, userTickets , RemainingTickets)

		if isValidName && isValidEmail && isValidTicketQuantity {

			// book tickets function

			helper2.BookTickets(userTickets, FirstName, LastName, email , &RemainingTickets, &Bookings , partyName)

			// Call get first name function to print it

			fmt.Printf("All bookings we got are have first name %v \n\n", getFirstNames())
			fmt.Printf("Totall persions booked till now %v\n\n", len(Bookings))

			if RemainingTickets == 0 {
				fmt.Println("Our party is booked out.  Please come back next year")
				break
			}
		} else {
			if !isValidTicketQuantity {
				fmt.Printf("sorry we have only %v tickets left for party. And you are trying to book %v \n", RemainingTickets, userTickets)
			}

			if !isValidName {
				fmt.Printf("First and last name should be more then 2 characters \n")
			}

			if !isValidEmail {
				fmt.Printf("You have given invalid email address \n")
			}
		}
	}

}

func greetUsers() {

	fmt.Println("Welcome to", partyName, "booking application")
	fmt.Printf("We have %v tickets are left out of %v tickets.\n", RemainingTickets, TotallTickets)
	fmt.Println("Get your tickets here")

}

func getFirstNames() []string {
	var FirstNames []string
	for _, booking := range Bookings {
		var name = strings.Fields(booking)
		FirstNames = append(FirstNames, name[0])
	}
	return FirstNames
}

func inputParameters() (string, string, string, uint) {
	var FirstName string
	var LastName string
	var email string
	var userTickets uint

	// Ask user information to register

	fmt.Println("Enter your first name:")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&LastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter no. of tickes you want to book")
	fmt.Scan(&userTickets)

	return FirstName, LastName, email, userTickets
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// city := "London"

//  switch city {

// 	case "New York":
// 		//execute code for new york

// 	case "Singapore","Hong Kong":

// 	default:
// 		fmt.Printf("No valid city selected")

// }
