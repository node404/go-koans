package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Println("Welcome to our", conferenceName, "booking application")

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// var bookings = [50]string{}
	// var bookings [50]string
	var bookings []string
	firstNames := []string{}

	for {
		fmt.Printf("We have total of %v tickets and %v are available.\n", conferenceTickets, remainingTickets)

		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name:")
		fmt.Scanln(&firstName)
		if firstName == "" {
			fmt.Println("No first name input, exiting...")
			break
		}
		fmt.Println("Enter your last name:")
		fmt.Scanln(&lastName)
		fmt.Println("Enter your email(optional):")
		fmt.Scan(&email)

		for {
			fmt.Println("Enter how many tickets you want:")
			fmt.Scan(&userTickets)
			if userTickets > 25 {
				fmt.Printf("You can only book 25 tickets in maximum.")
				continue
			} else if userTickets > int(remainingTickets) {
				fmt.Printf("We only have %v remaining, you can't book %v tickets\n", remainingTickets, userTickets)
				continue
			} else {
				break
			}
		}
		userName := firstName + " " + lastName
		bookings = append(bookings, userName)
		remainingTickets = remainingTickets - uint(userTickets)
		fmt.Printf("Thank you %v for booking %v tickets, you will receive email at %v.\n", userName, userTickets, email)

		if remainingTickets == 0 {
			fmt.Println("Ticket are sold out, please comming next year.")
			break
		}

	}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("bookings value are %v\n", bookings)
	fmt.Printf("bookings type is %T\n", bookings)
	fmt.Printf("bookings size is %v\n", len(bookings))
	fmt.Printf("firstNames are %v", firstNames)
}
