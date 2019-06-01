package main

import (
	"html/template"
	"log"
	"os"

	"./github"
)

const templ = `
	<h1>{{.TotalCount}}</h1>
	<table>
		<tr style='text-align: left;'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'><{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
	</table>
`

var report = template.Must(template.New("issueslist").Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("main: %v\n", err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatalf("main: %v\n", err)
	}

	// report, err := template.New("report").
	// 	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	// 	Parse(templ)

	// if err != nil {
	// 	log.Fatalf("main: %v\n", err)
	// }
}
