package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return fmt.Sprintf("%s", string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "(%c", u.op)
	fmt.Fprintf(&buf, "%s", u.x.String())
	buf.WriteByte(')')
	return buf.String()
}

func (b binary) String() string {
	var buf bytes.Buffer
	buf.WriteByte('(')
	fmt.Fprintf(&buf, "%s", b.x.String())
	fmt.Fprintf(&buf, " %c ", b.op)
	fmt.Fprintf(&buf, "%s", b.y.String())
	buf.WriteByte(')')
	return buf.String()
}

func (c call) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%s", arg.String())
	}
	buf.WriteByte(')')
	return buf.String()
}
