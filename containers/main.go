package containers

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	switch os.Args[1]{
	case "run":
		run()
	default:
		panic("what??")
	}
}
func run(){
	fmt.Printf("running %v\n", os.Args[2:])
}
func must(err error){
	if err != nil{
		panic(err)
	}
}