package main
import (
	"fmt"
	"sync"
)

func run(ch chan bool){
	<-ch
}
type Count struct{
	value int
	lock sync.Mutex
}
func count(counter *Count,wg *sync.WaitGroup){
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++

	fmt.Println(counter.value)
	wg.Done()

}
func main(){
	// --------- This whole problem can be used using buffered channel as well. ---------
	// ch := make(chan bool)
	// go run(ch)
	// go run(ch)
	// ch<- true
	// ch<- true

	// fmt.Println("Done")
	// ----------------------------------------------------------------------------------
	// ch := make(chan bool, 2)
	// ch <- true
	// ch <- true
	// <- ch
	// fmt.Println("Done")
	// ----------------------------------------------------------------------------------
	counter := Count{}
	wg := sync.WaitGroup{}
	for i := 0;i<100;i++{
		wg.Add(1)
		go count(&counter,&wg)
	}
	wg.Wait()
}







// func run1(){
// 	time.Sleep(2*time.Second)
// 	fmt.Println("Run-1")
// }
// // func run2(){
// // 	time.Sleep(3*time.Second)
// // 	fmt.Println("Run-2")
// // }
// // func run3(){
// // 	time.Sleep(4*time.Second)
// // 	fmt.Println("Run-3")
// // }
// func add(x int,y int, ch chan int, t int){
// 	time.Sleep(time.Duration(t)*time.Second)
// 	fmt.Println(x+y)
// 	ch <- x + y // send the result to channel using <-
// }
// func main(){
// 	// go run1()
// 	// go run2()
// 	// go run3()
	
// 	// time.Sleep(5*time.Second)// main function will wait for 5 seconds before exiting
// 	// fmt.Println("All functions executed")
// 	// Channels : use this method : make(chan dataType)
// 	ch := make(chan int)
// 	ch2 := make(chan int)
// 	go add(10,20,ch,2)
// 	go add(10,-20,ch2,4)
// // x := <- ch // receive the result from channel using <-
// 	for i := 0;i<2;i++{
// 		select{
// 		case x := <-ch:
// 			fmt.Println("Received from channel 1:", x)
// 		case y := <-ch2:
// 			fmt.Println("Received from channel 2:", y)
// 		case <-time.After(3 * time.Second):
// 			fmt.Println("Timeout occurred")
// 	}
// 	}
// }