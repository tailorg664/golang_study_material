package main

import "fmt"

func main(){
	// infinite loop
	// count := 100
	// for {
	// 	if count == 0 {
	// 		break
	// 	}
	// 	count--
	// 	fmt.Println("Running..")
	// }
	// simple for loop
	for i :=0; i < 10;i++ {
		fmt.Println(i + 1)
	}
	// for loop with while condition
	count := 10
	for count > 0 {
		fmt.Println(count)
		count--
	}
	// range over slice
	slice := []string{"apple", "banana", "cherry"}
	for _, fruit := range slice {
		fmt.Println(fruit)
	}
	str := "Hello, World!"
	// range over string
	for index, char := range str {
		fmt.Printf("Character at index %d is %c\n", index, char)
	}
}