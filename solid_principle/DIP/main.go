package main

import "fmt"
type greet interface{
	Greet() string
}
type french struct{}
func (f french) Greet() string {
	return "Bonjour"
}
type english struct{}

func (e english) Greet() string {
	return "Hello"
}

func main() {
	e := english{}
	greeter(e)
}
func greeter(g greet){
	fmt.Println(g.Greet())
}