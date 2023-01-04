package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask for user inputs
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Please enter the number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidUserTicketsNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				// firstNames = append(firstNames, strings.Fields(booking)[0]) // This works perfectly
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
	
			if remainingTickets == 0 {
				// end the program
				fmt.Println("Conference is booked out. come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered is a wrong format")
			}
			if !isValidUserTicketsNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}

	city := "London"
	switch city {
	case "New York":
		// execute code for booking New York conference tickets
	case "Singapore", "Berlin":
		// execute code for booking Singapore Berlin conference tickets
	case "London":
		// execute code for booking London conference tickets
	case "Hong Kong":
		// execute code for booking Hong Kong conference tickets
	default:
		fmt.Println("No valid city selected")
	}

}
