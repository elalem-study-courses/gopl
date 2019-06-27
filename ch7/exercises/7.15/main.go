package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"./eval"
)

var variables eval.Env

func init() {
	variables = make(eval.Env)
}

func main() {
	run()
}

func run() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		parseLine(scan.Text())
	}
}

func parseLine(line string) {
	tokens := strings.Split(line, "=")
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	switch len(tokens) {
	case 1:
		handleEvaluation(tokens)
	case 2:
		handleAssignment(tokens)
	}
}

func handleEvaluation(tokens []string) {
	result := evaluateExpression(tokens[0])
	fmt.Printf(">> %g\n", result)
}

func handleAssignment(tokens []string) {
	result := evaluateExpression(tokens[1])
	variables[eval.Var(tokens[0])] = result
}

func evaluateExpression(exprString string) float64 {
	expr, err := eval.Parse(exprString)
	if err != nil {
		log.Fatal(err)
	}
	result := expr.Eval(variables)
	return result
}
