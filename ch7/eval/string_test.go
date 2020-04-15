package eval_test

import (
	"math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gopl.io/ch7/eval"
)

var _ = Describe("String", func() {
   It("should print variable correctly", func() {
		 v := Var("x")
		 Expect(v.String()).To(Equal("x"))
   })

	 It("should print an expression correctly", func() {
		  expr, err := Parse("sin(pow(x, 3)) + sqrt(A / pi)")
			Expect(err).To(Not(HaveOccurred()))
			Expect(expr.String()).To(Equal("(sin(pow(x, 3)) + sqrt((A / pi)))"))

			expr2, err := Parse(expr.String())
			Expect(err).To(Not(HaveOccurred()))
			Expect(expr2.String()).To(Equal("(sin(pow(x, 3)) + sqrt((A / pi)))"))

			res := expr2.Eval(Env{"x": 8, "A": 3, "pi": math.Pi})
			Expect(res).To(BeNumerically("~", 1.056, 0.01))
	 })
})
