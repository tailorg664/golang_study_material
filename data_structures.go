package main

import "fmt"

func main(){
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 3
	myMap["cherry"] = 7
	
	fmt.Println("Capacity of myMap:", len(myMap))
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	var findKey string

	fmt.Scan(&findKey)
	quantity, ok := myMap[findKey]
	if ok {
		fmt.Printf("Quantity of %s: %d\n", findKey, quantity)
	} else {
		fmt.Println("Key '", findKey, "' not found in myMap")
	}
}