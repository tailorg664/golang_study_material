package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)
func check_link_validity(link string) int{
	pattern := `^https?://(www\.)?github\.com/[\w-]+/[\w.-]+/?$`
	if !regexp.MustCompile(pattern).MatchString(link) {
		fmt.Printf("Please provide a valid github link to verify : %s\n", link)
		fmt.Println("----------------------Error----------------------")
		os.Exit(1)
		return 0
	}
	cmd := exec.Command("git", "ls-remote", "--exit-code", link)
	err := cmd.Run()
	exitError,ok := err.(*exec.ExitError)
	if !ok {
		fmt.Printf("Error verifying link: %s\n", err)
		fmt.Println("----------------------Error----------------------")
		os.Exit(1)
		return 0
	}
	if exitError.ExitCode() == 2 {
		return 1
	}
	if err != nil {
		fmt.Printf("Error verifying link: %s\n", err)
		fmt.Println("----------------------Error----------------------")
		os.Exit(1)
		return 0
	}
	return 1
}

func init_git_repo(link string){
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error initializing git repository: %s\n", err)
		os.Exit(1)
	}
	cmd = exec.Command("git", "add", ".")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error adding files to git repository: %s\n", err)
		os.Exit(1)
	}
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error committing files to git repository: %s\n", err)
		os.Exit(1)
	}
	cmd = exec.Command("git", "branch", "-M", "main")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error creating main branch: %s\n", err)
		os.Exit(1)
	}
	cmd = exec.Command("git", "remote", "add", "origin", link)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error adding remote origin: %s\n", err)
		os.Exit(1)
	}
	cmd = exec.Command("git", "push", "-u", "origin", "main")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error pushing to remote repository: %s\n", err)
		os.Exit(1)
	}

}
func main(){
	fmt.Println("----------------------Started----------------------")
	var link string = os.Args[1]
	
	if check_link_validity(link) == 1 {
		init_git_repo(link)
	}
	fmt.Println("----------------------Completed----------------------")
}
