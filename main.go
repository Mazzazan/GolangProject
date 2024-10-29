package main

import "fmt"

func main() {

	conferenceName := "Go Conference" // using := is the same as using var but you can't be specific about the types of var
    const conferenceTickets = 50
    var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are available to book\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your conference tickets here")

    var userName string
    var userTickets int
    // ask user for their name
    fmt.Scan(&userName)
 
    userTickets = 2
    fmt.Printf("user %v booked %v tickets\n", userName, userTickets)
}
 