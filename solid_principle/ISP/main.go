package main

import "fmt"
// OCP ; use area for all shapes
type shape interface{
	area() float64

}
type object interface{
	shape
	volume() float64
}
func  areaSum(shapes ...shape) float64{
	var sum float64
	for _, shape := range shapes{
		sum += shape.area()
	}
	return sum
}
func areaVolumeSum(objects ...object) float64{
	var sum float64
	for _, obj := range objects{
		sum += obj.area() + obj.volume()
	}
	return sum
}

type Square struct{
	side float64
}
type Cube struct{
	side float64
}

func (s Square) area() float64{
	return (s.side) * (s.side)
}

func (c Cube) area() float64{
	return 6 * (c.side) * (c.side)
}
func (c Cube) volume() float64{
	return (c.side) * (c.side) * (c.side)
}
func main() {
	fmt.Println("Compiling..")
	c := Cube{side: 1}
	s := Square{side: 1}
	fmt.Println(areaSum(s,c))
	fmt.Println(areaVolumeSum(c))

}
