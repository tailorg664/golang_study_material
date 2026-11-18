package main

import "fmt"

type Table struct {
	customerName string
	tableNo      uint16
	time         uint16
}

func list(tables []Table) []Table {
	fmt.Println("Welcome to the Restaurant Reservation System")
	fmt.Println("===========================================")

	var diningLayout []Table
	for _, t := range tables {
		if t.customerName == "" || t.time == 0 {
			fmt.Printf("Table %d is free\n", t.tableNo)
			fmt.Printf("Send the table %d to queue\n", t.tableNo)
		} else {
			fmt.Printf("Table %d is occupied by %s for %d mins\n", t.tableNo, t.customerName, t.time)
			diningLayout = append(diningLayout, t)
		}
	}
	fmt.Println("\n===================================")
	fmt.Printf("These are the Occupied Tables:\n%v\n\n", diningLayout)
	return diningLayout
}
