package main
import (
	"fmt"
	"math"
)

// Interfaces: 
type Geometry interface {
	area() float64
	perimeter() float64
}
// Structs:
type Rectangle struct {
	height, width float64
}
type Circle struct {
	radius float64
}
// here we define a method named area that is associated with the Rectangle struct. The method takes a Rectangle receiver and returns the area of the rectangle by multiplying its height and width.
func (r Rectangle) area() float64 {
	return r.height * r.width
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main(){
	rectangle := Rectangle{height: 10, width: 20}
	circle := Circle{radius: 5}
	fmt.Printf("Area: %.2f\n", circle.area())

	fmt.Printf("Perimeter: %.2f\n", circle.perimeter())
	fmt.Printf("Area: %.2f\n", rectangle.area())
	fmt.Printf("Perimeter: %.2f\n", rectangle.perimeter())

	fmt.Println("Interface methods")
	measure(rectangle)
	measure(circle)
}