package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gopl.io/ch7/eval"
)

var htmlString = `
<html>
	<head>
		<title>Calculator</title>
	</head>
	<body>
		<form method="get" action="/calc">
			<input name="expr" />
		</form>
	</body>
</html>
`

var templ = template.Must(template.New("html").Parse(htmlString))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/calc", calculate)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if err := templ.Execute(w, ""); err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %v", err), 500)
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %v", err), 500)
		return
	}

	expr := r.Form.Get("expr")

	parsedExpr, err := eval.Parse(expr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %v", err), 500)
	}
	fmt.Fprintf(w, "result = %g\n", parsedExpr.Eval(nil))
}
