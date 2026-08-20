package main

import "fmt"

func main() {
    conferenceName := "Go Conference"
    var totalTickets = 40
    var remainingTickets = 40
    fmt.Printf("Welcome to the %s ticketing app!\n", conferenceName)
    fmt.Printf("We have a total of %d tickets and %d are still available.\n", totalTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")
    // to make a slice, we can use the following syntax
    // var bookings []string
    bookings :=[]string{}
    
    var firstName string
    var lastName string
    var email string
    var userTickets int
    
    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)
    fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)
    fmt.Println("Enter your email address:")
    fmt.Scan(&email)
    fmt.Println("Enter number of tickets:")
    fmt.Scan(&userTickets)
    bookings = append(bookings, firstName + " " + lastName)
    fmt.Printf("The whole array: %v\n", bookings)
    fmt.Printf("The first value: %v\n", bookings[0])
    fmt.Printf("Array type: %T\n", bookings)
    fmt.Printf("Array length: %v\n", len(bookings))
    // fmt.Printf("The slice of bookings: %v\n", bookings[:])
    remainingTickets = remainingTickets - userTickets
    fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s.\n", firstName, lastName, userTickets, email)
    fmt.Printf("%d tickets remaining for %s.\n", remainingTickets, conferenceName)

}