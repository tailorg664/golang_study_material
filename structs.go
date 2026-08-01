package main
import "fmt"
type Person struct {
	firstName string
	lastName  string
	age       int
}
func main() {
	person1 := Person{firstName: "John", lastName: "Doe", age: 30}
	person_ptr := &person1
	fmt.Println((*person_ptr).firstName) // Accessing firstName using pointer dereference
	*person_ptr = Person{firstName: "Jane", lastName: "Smith", age: 25}
	fmt.Println("First Name:", person_ptr.firstName)
	fmt.Println("Last Name:", person_ptr.lastName)
	fmt.Println("Age:", person_ptr.age)
}