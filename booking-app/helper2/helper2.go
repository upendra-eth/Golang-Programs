package helper2

import "fmt"


func BookTickets(userTickets uint, FirstName string, LastName string, email string ,RemainingTickets *uint ,Bookings *[] string ,partyName string ) {

	*RemainingTickets = *RemainingTickets - userTickets
	*Bookings = append(*Bookings, FirstName+" "+LastName)

	fmt.Printf("Thank you %v %v for booking %v tickets of %v . You will get confirmation Email at %v address\n\n", FirstName, LastName, userTickets, partyName, email)
	fmt.Printf("%v tickets are remaining to book for %v \n\n", RemainingTickets, partyName)

}
