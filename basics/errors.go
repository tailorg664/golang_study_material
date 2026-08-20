package main
import (
	"fmt"
)

func divide(a,b int) int {
	return a/b
}
func defferedFunc() {
	r := recover()
	if r == nil {
		fmt.Println("No error occurred")
	} else {
		fmt.Println("Error occurred:", r)
	}
}
func main(){
	
	// fmt.Println(divide(1,0))
	
	defer defferedFunc()// deferred function will be executed at the end of the main function
	panic("Something went wrong")// panic will stop the execution of the program and will call the deferred function
	fmt.Println("Division executing...")
	fmt.Println(divide(10,2))
}