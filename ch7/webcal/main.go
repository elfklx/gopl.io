package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gopl.io/ch7/calculator/reader"
	"gopl.io/ch7/eval"
)

func cal(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	scanner := bufio.NewScanner(strings.NewReader(r.Form.Get("expr")))

	reader := reader.NewReader(scanner)
	vars := map[eval.Var]bool{}
	expr, err := reader.ReadExpr(vars)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(vars) > 0 {
		http.Error(w, "Undefined variables: "+expr.String(), http.StatusBadRequest)
		return
	}

	env := make(eval.Env)
	fmt.Fprintf(w, "Result: %s", strconv.FormatFloat(expr.Eval(env), 'f', -1, 64))
}

func main() {
	http.HandleFunc("/cal", cal)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
