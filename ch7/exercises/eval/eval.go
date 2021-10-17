package eval

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

var (
	operatorPrecedence map[string]int
	allowedParenthesis = "()"
	allowedOpenParanthesis = "("
	allowedClosedParanthesis = ")"
	allowedOperators = "+/*-"
)

func init() {
	operatorPrecedence = map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
}

func parse(reader io.Reader) []token {
	scan := scanner.Scanner{}
	tokens := make([]token, 0)
	scan.Init(reader)

	scan.Mode = scanner.ScanFloats | scanner.ScanIdents | scanner.ScanInts
	for {
		tokenType := scan.Scan()
		if tokenType == scanner.EOF {
			break
		}

		tokens = append(tokens, token{tokenType: tokenType, token: scan.TokenText()})
	}

	return tokens
}

func Eval(reader io.Reader, vars map[string]float64) float64 {
	tokens := parse(reader)
	postfixTokens := buildPostfixExpression(tokens)
	fmt.Printf("%#v\n", postfixTokens)
	return evaluatePostfix(postfixTokens, vars)
}

func buildPostfixExpression(tokens []token) []token {
	stack := make([]token, 0)
	postfixTokens := make([]token, 0)
	for _, token := range tokens {
		if token.tokenType == scanner.Ident || token.tokenType == scanner.Int || token.tokenType == scanner.Float {
			postfixTokens = append(postfixTokens, token)
		} else {
			if strings.Contains(allowedParenthesis, token.token) {
				if strings.Contains(allowedOpenParanthesis, token.token) {
					stack = append(stack, token)
				} else {
					for !strings.Contains(allowedOpenParanthesis, stack[len(stack) - 1].token) {
						postfixTokens = append(postfixTokens, stack[len(stack) - 1])
						stack = stack[:len(stack) - 1]
					}
				}
			} else {
				for !strings.Contains(allowedOpenParanthesis, stack[len(stack) - 1].token) {
					if t := stack[len(stack) - 1].token; strings.Contains(allowedOperators, t) && operatorPrecedence[t] < operatorPrecedence[token.token] {
						break
					}
					postfixTokens = append(postfixTokens, stack[len(stack) - 1])
					stack = stack[:len(stack) - 1]
				}

				stack = append(stack, token)
			}
		}
	}

	for len(stack) > 0 {
		if !strings.Contains(allowedParenthesis, stack[len(stack) - 1].token) {
			postfixTokens = append(postfixTokens, stack[len(stack) - 1])
		}
		stack = stack[:len(stack) - 1]
	}

	return postfixTokens
}

func evaluatePostfix(tokens []token, vars map[string]float64) float64 {
	stack := make([]token, 0)
	expr := bytes.Buffer{}
	for _, t := range tokens {
		expr.WriteString(t.token)
		if t.tokenType == scanner.Ident || t.tokenType == scanner.Int || t.tokenType == scanner.Float {
			stack = append(stack, t)
		} else {
			operands := stack[len(stack) - 2:]
			stack = stack[:len(stack) - 2]
			res := [2]float64{}
			val := float64(0)
			for i, operand := range operands {
				if operand.tokenType == scanner.Ident {
					res[i] = vars[operand.token]
				} else {
					x, _ := strconv.ParseFloat(operand.token, 64)
					res[i] += x
				}
			}

			switch t.token {
			case "+":
				val = res[0] + res[1]
			case "-":
				val = res[0] - res[1]
			case "/":
				val = res[0] / res[1]
			case "*":
				val = res[0] * res[1]
			}

			stack = append(stack, token{tokenType: scanner.Float, token: strconv.FormatFloat(val, 'g', -1, 64)})
		}
		fmt.Printf("%v\n", stack)
	}

	fmt.Println(expr.String())

	ret, _ := strconv.ParseFloat(stack[0].token, 64)
	return ret
}



func main() {
	scan := scanner.Scanner{}
	scan.Init(os.Stdin)
	scan.Mode = scanner.ScanFloats | scanner.ScanIdents | scanner.ScanInts | scanner.ScanStrings
	for {
		fmt.Printf("%v, %s\n", scan.Scan(), scan.TokenText())
		time.Sleep(1 * time.Second)
	}
}
