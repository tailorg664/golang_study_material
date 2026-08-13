package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)
func check_link_validity(link string){
	pattern := `^https?://(www\.)?github\.com/[\w-]+/[\w.-]+/?$`
	if !regexp.MustCompile(pattern).MatchString(link) {
		fmt.Printf("Please provide a valid github link to verify : %s", link)
		os.Exit(1)
	}
	cmd := exec.Command("git", "ls-remote", "--exit-code", link)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error verifying link: %s\n", err)
		os.Exit(1)
	}
}

func main(){
	fmt.Printf("Please provide a link to verify : %s", os.Args[1])
	var link string = os.Args[1]

	check_link_validity(link)
}
