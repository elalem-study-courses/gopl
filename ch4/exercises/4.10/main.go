// package description
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

const (
	LessThanAMonth = iota
	LessThanAYear
	MoreThanAYear
)

var IssuesByAges [3][]*github.Issue
var IssuesAgesText = [...]string{
	LessThanAMonth: "less than a month",
	LessThanAYear:  "less than a year",
	MoreThanAYear:  "more than a year",
}

func differenceLessThanDuration(specifiedDate time.Time, months int) bool {
	past := time.Now().AddDate(0, -1*months, 0)
	if specifiedDate.After(past) {
		return true
	}
	return false
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		switch {
		case differenceLessThanDuration(item.CreatedAt, 1):
			IssuesByAges[LessThanAMonth] = append(IssuesByAges[LessThanAMonth], item)
		case differenceLessThanDuration(item.CreatedAt, 12):
			IssuesByAges[LessThanAYear] = append(IssuesByAges[LessThanAYear], item)
		default:
			IssuesByAges[MoreThanAYear] = append(IssuesByAges[MoreThanAYear], item)
		}
	}

	for i, text := range IssuesAgesText {
		fmt.Printf("Issues %s\n", text)
		fmt.Println("------------------------\n")
		for _, item := range IssuesByAges[i] {
			fmt.Printf("#%-5d %9.9s %0.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

}
