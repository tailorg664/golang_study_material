package main

import (
	"errors"
	"fmt"
)

// FindTable looks up a table by its number.
// Returns the Table and nil on success, or an empty Table and an error if not found.
func FindTable(tables []Table, tableNo uint16) (Table, error) {
	for _, t := range tables {
		if t.tableNo == tableNo {
			return t, nil
		}
	}
	return Table{}, fmt.Errorf("table %d not found", tableNo)
}

// CalculateBill computes the charge for a table based on minutes spent.
// - Returns duration * ratePerMinute.
// - Returns an error if duration is 0 or ratePerMinute is <= 0.
func CalculateBill(duration uint16, ratePerMinute float64) (float64, error) {
	if duration == 0 {
		return 0, errors.New("duration must be greater than 0")
	}
	if ratePerMinute <= 0 {
		return 0, errors.New("rate per minute must be positive")
	}

	return float64(duration) * ratePerMinute, nil
}

// GetFreeTables filters and returns all tables that are not occupied.
// A table is considered free if customerName is empty or time is 0.
func GetFreeTables(tables []Table) []Table {
	var free []Table
	for _, t := range tables {
		if t.customerName == "" || t.time == 0 {
			free = append(free, t)
		}
	}
	return free
}
