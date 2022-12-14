package main

import ("strings"
"fmt"
)



func validateUserInputs(FirstName string, LastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(FirstName) > 2 && len(LastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketQuantity := userTickets <= RemainingTickets && userTickets > 0
	return isValidName, isValidEmail, isValidTicketQuantity

}

func bookTickets(userTickets uint, FirstName string, LastName string, email string  ) {

	RemainingTickets = RemainingTickets - userTickets
	var userRecord = UserData {

		FirstName: FirstName,
		LastName: LastName,
		email: email,
		numberOfTickets: userTickets,

	}

	Bookings = append(Bookings,userRecord)
	fmt.Printf("Information of all the uses who have booked tickets are %v \n", Bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets of %v . You will get confirmation Email at %v address\n\n", FirstName, LastName, userTickets, partyName, email)
	fmt.Printf("%v tickets are remaining to book for %v \n\n", RemainingTickets, partyName)

}
