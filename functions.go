package main

import (
	"fmt"
	"math/rand/v2"	
	"time"
)


func say_hello(){
	fmt.Println("Hello")
}

// Target : to create function which accepts a function and logs its response

func random_func() int {
	time.Sleep(40 * time.Millisecond)
	return rand.IntN(100)
}
// create a function which will beautify logs

func beautify_logs(logs int) string{
	log := fmt.Sprintf("Log : %d",logs)  // great way to use custom string formatting
	return log
}

func log_generator(f func() int) []string {
	var logs []string
	for i :=0 ; i<10 ; i++ {
		logs = append(logs,beautify_logs(f()))
	}
	return logs
}

// clauses are used to create anonymous functions and call them immediately. This is a great way to create a function which is only used once and not needed again.
func anonymous_function() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main(){
	// anonymous function declaration
	func(msg string){
		fmt.Println(msg)
	}("Anonymous Hello")
	say_hello()
	print_anonymous := anonymous_function()
	fmt.Printf("Location of anonymous function : %p\n", print_anonymous)
	print_anonymous("Anonymous Hello")
	// calling int_seq function
	seq := int_seq()
	fmt.Printf("Location of int_seq function : %p\n", seq)
	fmt.Println(seq())
	// calling log_generator function
	var logs []string= log_generator(random_func)
	// showing logs
	fmt.Println(logs[0:5])
}



// create a function which logs the return values of functions and stores them in a slice.