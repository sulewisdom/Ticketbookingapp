package main

//first line of the application has to be package main

import (
	"fmt"
	"strings"
)

//importing a formating built in function

//application entry point bellow

func main() {
	conferencName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	//declaration of variables and constants above

	//%v is a place holder
	fmt.Printf("Welcome to %v booking application\n", conferencName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n ", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// (this is the delaration of the slice whic is also an array without a size definition)
		var firstNAme string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstNAme)

		fmt.Println("Enter lastname name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email Address: ")
		fmt.Scan(&email)

		fmt.Println("Enter your Number of tickets: ")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets

		// (this is an array to store user details) bookings[0] = firstNAme + " " + lastName
		bookings = append(bookings, firstNAme+" "+lastName)
		//this is a slice above and its more dynamic and better than an array in storing data that we dont know specifically the numbers of index we are expecting

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstNAme, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencName)
		firstNAmes := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNAmes = append(firstNAmes, names[0])
		}
		fmt.Printf("The firdt names of the bookings are: %v\n", firstNAmes)

		if remainingTickets == 0 {
			// end program
			fmt.Print("Our conference is booked out. Come back next year.")
			break
		}
	}
}
