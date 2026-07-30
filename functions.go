package main

import "fmt"


func say_hello(){
	fmt.Println("Hello")
}

// Target : to create function which accepts a function and logs its response
func any_random_func(input int)(int,string){
	var message string
	if input >= 10 { 
		message = "Testing Ok"
	} else {
		message = "Failed, redo the execution."
	}

	return input,message
}
// Log Checker function

func log_generator(input_func func(input int))string{
	_,message := input_func(input)
	fmt.Println("Displayed Logs")
	return message
}

func main(){
	// anonymous function declaration
	func(msg string){
		fmt.Println(msg)
	}("Anonymous Hello")
	say_hello()
	message := log_generator(any_random_func(18))
	fmt.Println(message)


}



