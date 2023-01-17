package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//we can get the data type by using %T
	fmt.Printf("conferenceName is type of %T conferenceTickets is type of %T remainingTickets is type of %T", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Welcome to our ", conferenceName, " application :)")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here...")

	var userName string
	var userTickets int

	//ask user for their name
	fmt.Println("Enter your User Name...")
	fmt.Scan(&userName)

	//userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

	bookings := []string{}

	for remainingTickets >= 0 {

		var firstName string
		var lastName string
		var userEmail string
		var numberOfTickets int

		fmt.Println("Please enter your first name : ")
		fmt.Scan(&firstName)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2

		fmt.Println("Please enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter you email ID")
		fmt.Scan(&userEmail)

		isValidEmail := strings.Contains(userEmail, "@")

		fmt.Println("Please enter the number of tickets you want")
		fmt.Scan(&numberOfTickets)

		isValidNumberOfTickets := numberOfTickets > 0 && numberOfTickets <= 50

		if !isValidName || !isValidEmail || !isValidNumberOfTickets {
			fmt.Println("Sorry! you entered wrong input")
			fmt.Println("Press 1 to book tickets again and 2 for exit")
			var choice int
			fmt.Scan(&choice)
			if choice == 1 {
				continue
			} else {
				break
			}
		}

		if numberOfTickets > remainingTickets {
			fmt.Println("sorry we don't have that much ticket left, we are having %v tickets only", remainingTickets)
			fmt.Println("Press 1 to book tickets again and 2 for exit")
			var choice int
			fmt.Scan(&choice)
			if choice == 1 {
				continue
			} else {
				break
			}
		}

		remainingTickets = remainingTickets - numberOfTickets

		fmt.Printf("Congratulations %v %v with email ID %v on grabbing %v tickets\n", firstName, lastName, userEmail, numberOfTickets)
		fmt.Printf("Remaining tickets is %v\n", remainingTickets)

		bookings = append(bookings, firstName+" "+lastName)

		if remainingTickets <= 0 {
			fmt.Println("Sorry! All Tickets are sold")
			break
		}
	}

	fmt.Println(bookings)

	//we want to print first name only while we have full names of users in bookings slice

	for _, booking := range bookings {
		parseFullName := strings.Fields(booking)
		firstName := parseFullName[0]
		println(firstName)

	}

}
