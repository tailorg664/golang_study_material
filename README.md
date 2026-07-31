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

## New Learnings and Additions (upto : 31/07/2026)

- **functions.go:** Higher-order functions and function types: passing functions as arguments (`log_generator(random_func)`), returning functions (`anonymous_function`, `int_seq`), anonymous functions, closures that capture state, `fmt.Sprintf` for custom formatting, and use of `math/rand/v2` and `time.Sleep` for generating random values and simulating delays. See [functions.go](functions.go#L1-L200).
- **data_structures.go:** Maps and lookups: `make(map[string]int)`, `len()` for map size, iterating maps with `range`, launching goroutines to print map entries, and the comma-ok idiom for safe map lookup. The program also reads user input to query a key. See [data_structures.go](data_structures.go#L1-L200).
- **loop.go:** More `for` loop patterns: infinite loops with `break`, classic init/cond/post loops, condition-only loops (like `while`), `range` over slices, and `range` over strings (index + rune). See [loop.go](loop.go#L1-L200).
- **http_go/main.go:** Example of a non-main package layout (`package httpgo`) with a simple `fmt.Printf` example demonstrating package naming and separate folder structure. See [http_go/main.go](http_go/main.go#L1-L50).
- **main.go:** (Already listed above) basic ticketing app: reading user input with `fmt.Scan`, building and inspecting a `[]string` slice of bookings, updating remaining tickets, and formatted output with `fmt.Printf`. See [main.go](main.go#L1-L200).
