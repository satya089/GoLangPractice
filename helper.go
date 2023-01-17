package main

import "strings"

// returns a slice of string
func printFirstNamesByReturningSlice(bookings []string) []string {

	var temp []string
	for _, booking := range bookings {
		parseFullName := strings.Fields(booking)
		temp = append(temp, parseFullName[0])
	}
	return temp
}
