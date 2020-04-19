package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch7/calculator/reader"
	"gopl.io/ch7/eval"
)

//!+parseAndCheck

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	r := reader.NewReader(scanner)

	vars := map[eval.Var]bool{}
	var expr eval.Expr
	var err error
	for {
		fmt.Print("Enter expression: ")
		expr, err = r.ReadExpr(vars)
		if err != nil {
			fmt.Printf("[invalid input] reason: %s; try again\n", err)
			continue
		}
		break
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
