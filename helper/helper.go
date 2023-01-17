package helper

import (
	"fmt"
	"strings"
)

// returns a slice of string
func PrintFirstNamesByReturningSlice(bookings []string) []string {

	var temp []string
	for _, booking := range bookings {
		parseFullName := strings.Fields(booking)
		temp = append(temp, parseFullName[0])
	}
	return temp
}

func UserToNumberOfTickets(firstName string, lastName string, numberOfTickets int) {
	var numberOfTicketsForUser map[string]int
	numberOfTicketsForUser = make(map[string]int)
	numberOfTicketsForUser[firstName+" "+lastName] = numberOfTickets
	for key, value := range numberOfTicketsForUser {
		fmt.Printf("%v User has boooked %v tickets\n", key, value)
	}
}
