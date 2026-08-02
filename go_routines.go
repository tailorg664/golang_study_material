package main
import (
	"fmt"
	"time"
)

func run1(){
	time.Sleep(2*time.Second)
	fmt.Println("Run-1")
}
func run2(){
	time.Sleep(3*time.Second)
	fmt.Println("Run-2")
}
func run3(){
	time.Sleep(4*time.Second)
	fmt.Println("Run-3")
}
func main(){
	go run1()
	go run2()
	go run3()
	time.Sleep(5*time.Second)// main function will wait for 5 seconds before exiting
	fmt.Println("All functions executed")
}