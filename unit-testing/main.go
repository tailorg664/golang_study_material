package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Here comes the source.")
	currentLayout := []Table{
		{"Gaurav", 1, 30},
		{"Karan", 2, 45},
		{"Simran", 3, 60},
		{"Rahul", 4, 15},
		{"Priya", 5, 20},
	}
	var table_list = list(currentLayout)
	fmt.Println(table_list)
}
