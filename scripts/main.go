package main

import (
	"fmt"
	"os"
)
func check_link_validity(){
	if len(os.Args) < 2{
		fmt.Printf("Please provide a link to verify : %s", os.Args[1])
		os.Exit(1)
	}

}
func main(){
	fmt.Printf("Please provide a link to verify : %s", os.Args[1])
	
}
