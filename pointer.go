package main
import "fmt"

func main(){
	var x int = 20
	var p *int
	p = &x
	value := *p
	fmt.Println("Value of x : ",x)
	fmt.Println("Address of x : ",p)
	fmt.Println("Value of x using pointer : ",value)
	*p = 30
	fmt.Println("Value of x after changing using pointer : ",x)
}