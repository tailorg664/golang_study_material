# Go Conference Ticketing App

This project is a simple Go practice app for booking conference tickets from the terminal.

## What Has Been Done So Far

- Created a Go module for the project.
- Built a basic `main.go` entrypoint.
- Printed a welcome message for the ticketing app.
- Declared conference details such as the conference name, total tickets, and remaining tickets.
- Asked the user for their first name, last name, email address, and number of tickets.
- Stored booking names in a slice.
- Displayed booking information, slice length, type, and the first booking.
- Updated the remaining ticket count after each booking.
- Printed a confirmation message for the user.

## Go Learning So Far

- `package main` and `func main()` are used to create an executable Go program.
- `fmt` is used for printing output and collecting input from the terminal.
- Variables can be declared with `var` or inferred with `:=`.
- Strings, integers, and slices are the main data types used here.
- `fmt.Scan` reads user input from the terminal.
- `append` adds new items to a slice.
- `len` gives the current size of a slice.
- `fmt.Printf` supports formatted output with placeholders like `%s`, `%d`, `%v`, and `%T`.

# Looping in Go
- For-each loop in Go looks like this : 
```
for _, fruit := range slice {
	fmt.Println(fruit)
}
```