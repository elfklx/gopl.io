package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopl.io/ch7/eval"
)

//!+parseAndCheck

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var exprTxt string
	for {
		fmt.Print("Enter expression:\n")
		scanner.Scan()
		exprTxt = scanner.Text()
		if exprTxt != "" {
			break
		}
		fmt.Print("empty expression; try again\n")
	}

	expr, err := eval.Parse(exprTxt)
	if err != nil {
		log.Fatalf("invalid expression %s", err)
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		log.Fatalf("invalid errors %s", err)
	}

	env := make(eval.Env)
	for v := range vars {
		for {
			fmt.Printf("Enter value for %s:\n", v)
			scanner.Scan()
			valueStr := scanner.Text()
			value, err := strconv.ParseFloat(valueStr, 32)
			if err == nil {
				env[v] = value
				break
			}
			fmt.Printf("invalid value %s must be float; try again\n", valueStr)
		}
	}

	fmt.Printf("Result:\n%f", expr.Eval(env))
}
