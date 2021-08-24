package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"./github"
)

var action = flag.String("a", "", "(*required) Action to perform must be in [create, read, update, close]")
var repository = flag.String("r", "", "(*required) Repository name")
var accessToken = flag.String("t", "", "(*required) Access user for authentication")
var username = flag.String("u", "", "(*required) Username")
var issueNumber = flag.String("i", "", "Issue number")
var issueTitle = flag.String("title", "", "Issue title")
var editor = flag.String("editor", "vim", "Select editor (default vim)")

func main() {
	flag.Parse()
	if *action == "" ||
		*repository == "" ||
		*accessToken == "" ||
		*username == "" {
		log.Fatal("Invalid parameters list")
	}

	var (
		issue *github.Issue
		err   error
	)

	switch *action {
	case "read":
		issue, err = github.ReadIssue(*username, *repository, *issueNumber)
	case "create":
		fpath := os.TempDir() + "/gissuem"
		cmd := exec.Command(*editor, fpath)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		description, _ := ioutil.ReadFile(fpath)
		issue, err = github.CreateIssue(*username, *accessToken, *repository, *issueTitle, string(description))
	case "close":
		issue, err = github.CloseIssue(*username, *accessToken, *repository, *issueNumber)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(issue)
}
