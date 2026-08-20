package main

import (
	"fmt"
	"math/rand"
	"math"
)

func randomNumberGenerator() int {
	return rand.Intn(100)
}

func predictionRate(originalPredictions float64, correctPredictions float64) float64 {
	var rate float64
	delta := correctPredictions - originalPredictions
	rate = math.Abs(delta) / originalPredictions
	return rate
}
func main(){
	fmt.Println("Hello, World!")
	val := randomNumberGenerator()

	fmt.Println(val)
}