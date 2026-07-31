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
func main(){
	// anonymous function declaration
	func(msg string){
		fmt.Println(msg)
	}("Anonymous Hello")
	say_hello()
	// calling log_generator function
	var logs []string= log_generator(random_func)
	// showing logs
	fmt.Println(logs[0:5])
}



// create a function which logs the return values of functions and stores them in a slice.