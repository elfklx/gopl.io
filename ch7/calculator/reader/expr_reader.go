package reader

import (
	"bufio"
	"errors"

	"fmt"

	"gopl.io/ch7/eval"
)

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader(s *bufio.Scanner) Reader {
	return Reader{scanner: s}
}

func (e Reader) ReadExpr(vars map[eval.Var]bool) (eval.Expr, error) {
	e.scanner.Scan()
	text := e.scanner.Text()
	if text == "" {
		return nil, errors.New("empty input")
	}

	expr, err := eval.Parse(text)
	if err != nil {
		return nil, fmt.Errorf("unparsable input: %s", err)
	}

	if err := expr.Check(vars); err != nil {
		return nil, fmt.Errorf("invalid expression: %s", err)
	}
	return expr, nil
}

func (e Reader) ReadVar(env eval.Env) error {
	return nil
}
