// package description
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"./github"
)

func printMainDialog() {
	dialog := `
1) Read an issue
2) Create an issue
3) Update an issue
4) Delete an issue
Enter a number (default: EOF):
	`

	fmt.Println(dialog)
}

func handleReading(in *bufio.Scanner) {
	fmt.Println("Enter Owner, Repo, IssueNumber seperrated by spaces")
	in.Scan()
	args := strings.Split(in.Text(), " ")
	issue, _ := github.ReadIssue(args[0], args[1], args[2])
	fmt.Printf("%v\n", issue)
}

func handleCreation(in *bufio.Scanner) {
	fmt.Println("Enter owner")
	in.Scan()
	owner := in.Text()
	fmt.Println("Enter Repository")
	in.Scan()
	repo := in.Text()
	fmt.Println("Enter title")
	in.Scan()
	title := in.Text()

	body := ""
	fmt.Println("Select method to enter body")
	fmt.Println(`
1) stdin
2) editor
	`)
	in.Scan()
	choiceString := in.Text()
	choice, err := strconv.Atoi(choiceString)
	if err != nil {
		log.Fatalf("handleCreation: %v\n", err)
	}
	switch choice {
	case 1:
		fmt.Println("Enter body")
		in.Scan()
		body = in.Text()
	}

	fmt.Println("Enter labels seprated by spaces")
	in.Scan()
	labels := strings.Split(in.Text(), " ")

	github.CreateIssue(owner, repo, title, body, labels)
}

func handleUpdate(in *bufio.Scanner) {
	fmt.Println("Updating")
}

func handleDeletion(in *bufio.Scanner) {
	fmt.Println("Deleting")
}

func runMainDialog(in *bufio.Scanner) {
	fmt.Println()
	choiceText := in.Text()
	choice, err := strconv.Atoi(choiceText)
	if err != nil {
		fmt.Printf("github: %v\n", err)
		os.Exit(1)
	}
	switch choice {
	case 1:
		handleReading(in)
	// Pain to complete because of authroization
	case 2:
		handleCreation(in)
	case 3:
		handleUpdate(in)
	case 4:
		handleDeletion(in)
	}

	fmt.Println("Press any key to exit...")
	in.Scan()
	clear()
	printMainDialog()
}

func clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	printMainDialog()
	for in.Scan() {
		runMainDialog(in)
	}
}
