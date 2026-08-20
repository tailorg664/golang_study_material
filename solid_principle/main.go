package main

import "fmt"
// OCP ; use area for all shapes
type shape interface{
	area() float32
}

type Circle struct{
	radius float32
}
type Square struct{
	side float32
}

type Calculator struct{}

//---The below function is flawed due to modification property---
//---------------------
// func (c Calculator) calculateArea(shapes ...interface{}) float32{
// 	var sum float32
// 	for _, shape := range shapes{
// 		switch shape.(type){
// 		case Circle:
// 			r := shape.(Circle).radius
// 			sum += 3.14 * r * r
// 		case Square :
// 			l := shape.(Square).side
// 			sum += l * l
// 		}
// 	}
// 	return sum
// }
//---------------------
func (c Circle) area() float32{
	return 3.14 * c.radius * c.radius
}
func (s Square) area() float32{
	return s.side * s.side
}

func (c Calculator) calculateArea(shapes ...shape) float32{
	var sum float32
	for _, shape := range shapes{
		sum += shape.area()
	}
	return sum
}
func main() {
	fmt.Println("Compiling..")
	c := Circle{radius: 1}
	s := Square{side: 1}
	cal := Calculator{}
	fmt.Println(cal.calculateArea(c,s))

}
