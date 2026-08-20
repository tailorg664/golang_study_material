package main
import "fmt"
type person interface{
	getName() string
}
type human struct{
	name string
}
func (h human) getName() string {
	return h.name
}
type teacher struct{
	human 
	salary float32
}
type student struct{
	human 
	marks float32
}
type printer struct{}
func (printer) info(p person){
	fmt.Println("Name :", p.getName())
}
func main(){
	h:= human{name:"Gaurav"}
	t:= teacher{
		human : human{name:"Karan"},
		salary : 13.1,
	}
	s:=student{
		human : human{name:"Kiran"},
		marks:12,
	}
	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}