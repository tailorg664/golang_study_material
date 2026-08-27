package main

import "fmt"

func sliceToChannel(nums []int) <-chan int{
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
func squareChannel(in <-chan int) <-chan int {
	out := make(chan int)
	
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
func main() {
	// input 
	nums := []int{2,5,1,7,6}
	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := squareChannel(dataChannel)

	// stage 3

	for n := range finalChannel {
		fmt.Println(n)
	}

}