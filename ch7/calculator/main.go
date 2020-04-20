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
		if err == nil {
			break
		}
		fmt.Printf("[invalid input] reason: %s; try again\n", err)
	}

	env := make(eval.Env)
	for v := range vars {
		for {
			fmt.Printf("Enter value for %s: ", v)
			val, err := r.ReadVar()
			if err == nil {
				env[v] = val
				break
			}
			fmt.Println("[invalid input] not a number; try again")
		}
	}

	fmt.Printf("Result: %s", strconv.FormatFloat(expr.Eval(env), 'f', -1, 64))
}
