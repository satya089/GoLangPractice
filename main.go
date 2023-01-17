package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//fmt.Println("Welcome to our ", conferenceName, " application :)")
	greetUsers()

	//we can get the data type by using %T
	//fmt.Printf("conferenceName is type of %T conferenceTickets is type of %T remainingTickets is type of %T\n", conferenceName, conferenceTickets, remainingTickets)

	dataTypeInfo(conferenceName, conferenceTickets, remainingTickets)

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

		fmt.Println("Please enter your last name : ")
		fmt.Scan(&lastName)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2

		fmt.Println("Please enter you email ID")
		fmt.Scan(&userEmail)

		isValidEmail := strings.Contains(userEmail, "@")

		fmt.Println("Please enter the number of tickets you want")
		fmt.Scan(&numberOfTickets)

		isValidNumberOfTickets := numberOfTickets > 0 && numberOfTickets <= 50

		if !isValidName || !isValidEmail || !isValidNumberOfTickets {
			if !isValidName {
				fmt.Printf("Sorry! you entered wrong name : %v\n", firstName+" "+lastName)
			} else if !isValidEmail {
				fmt.Printf("Sorry! you entered wrong emailID : %v\n", userEmail)
			} else {
				fmt.Printf("Sorry! you entered wrong numberOfTickets : %v\n", numberOfTickets)
			}
			fmt.Println("Press 1 to book tickets again and 2 for exit\n")
			var choice int
			fmt.Scan(&choice)
			if choice == 1 {
				continue
			} else {
				break
			}
		}

		helper.UserToNumberOfTickets(firstName, lastName, numberOfTickets)

		if numberOfTickets > remainingTickets {
			fmt.Printf("sorry we don't have that much ticket left, we are having %v tickets only", remainingTickets)
			fmt.Println("Press 1 to book tickets again and 2 for exit")
			var choice int
			fmt.Scan(&choice)
			if choice == 1 {
				continue
			} else {
				break
			}
		}

		remainingTickets = findRemainingTickets(remainingTickets, numberOfTickets)

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

	//for _, booking := range bookings {
	//	parseFullName := strings.Fields(booking)
	//	firstName := parseFullName[0]
	//	println(firstName)
	//
	//}

	printFirstNames(bookings)
	firstNames := helper.PrintFirstNamesByReturningSlice(bookings)
	for _, firstName := range firstNames {
		fmt.Println(firstName)
	}

	//// Switch statement example
	//city := "London"
	//
	//switch city {
	//case "New York":
	//	// booking for New York conference
	//case "Singapore", "Hong Kong":
	//	// booking for Singapore & Hong Kong conference
	//case "London", "Berlin":
	//	// booking for London & Berlin conference
	//case "Mexico City":
	//	// booking for Mexico City conference
	//default:
	//	fmt.Print("No valid city selected")
	//}

}

//func func_name(func_parameters) return_type{
//	function_body
//}

func greetUsers() {
	fmt.Println("Welcome to our ticketing conference :)")
}

// Data type should't be written while passing const as parameter
func dataTypeInfo(conferenceName string, conferenceTickets, remainingTickets int) {
	fmt.Printf("conferenceName is type of %T conferenceTickets is type of %T remainingTickets is type of %T\n", conferenceName, conferenceTickets, remainingTickets)

}

func findRemainingTickets(remainingTickets int, numberOfTickets int) int {
	return remainingTickets - numberOfTickets
}

// returns nothing
func printFirstNames(bookings []string) {
	for _, booking := range bookings {
		parseFullName := strings.Fields(booking)
		firstName := parseFullName[0]
		println(firstName)
	}
}
