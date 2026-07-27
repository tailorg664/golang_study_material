package main

import (
	"fmt"
	"time"
)
func main(){
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 3
	myMap["cherry"] = 7
	
	fmt.Println("Capacity of myMap:", len(myMap))
	for key, value := range myMap {
		
		go func() {
			time.Sleep(2 * time.Second) // Sleep for 2 seconds
			fmt.Printf("Key: %s, Value: %d\n", key, value)
		}()
	}
	var findKey string
	fmt.Println("Enter a key to find its quantity:")
	fmt.Scan(&findKey)
	quantity, ok := myMap[findKey]
	if ok {
		fmt.Printf("Quantity of %s: %d\n", findKey, quantity)
	} else {
		fmt.Println("Key '", findKey, "' not found in myMap")
	}
}