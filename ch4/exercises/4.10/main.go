package main

import (
	"fmt"
	"log"
	"os"

	"./github"
)

func main() {
	result, err := github.SearchIssuesAgeCategorized(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, text := range []string{"less than a month", "less than a year", "more than a year"} {
		fmt.Println(text)
		singleResult := result[text]
		fmt.Printf("Issues count: %d\n", singleResult.TotalCount)
		fmt.Printf("Issues:\n")
		for _, issue := range singleResult.Items {
			fmt.Printf("#%-5d %9.9s %0.55s %0.55s\n", issue.Number, issue.User.Login, issue.Title, issue.CreatedAt.Local().Format("Jan 02, 2006"))
		}
	}
}
