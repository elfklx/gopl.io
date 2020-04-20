package reader_test

import (
	"bufio"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gopl.io/ch7/calculator/reader"
	"gopl.io/ch7/eval"
)

var _ = Describe("Reader", func() {

	Context("ReadExpr", func() {
		It("should error out if an empty string is provided", func() {
			const src = ""
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			expr, err := reader.ReadExpr(map[eval.Var]bool{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("empty input"))
			Expect(expr).To(BeNil())
		})

		It("should return an error if an invalid input is provided", func() {
			const src = `\t
\x99 Մ Յ
&nbsp;
`
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			expr, err := reader.ReadExpr(map[eval.Var]bool{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unparsable input"))
			Expect(expr).To(BeNil())
		})

		It("should return an error if an invalid expression is provided", func() {
			const src = "myfunc(x + y)"
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			expr, err := reader.ReadExpr(map[eval.Var]bool{})
			Expect(err.Error()).To(ContainSubstring("invalid expr"))
			Expect(expr).To(BeNil())
		})

		It("should return an expression with vars if a valid input is provided", func() {
			const src = "(x + y)"
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			vars := map[eval.Var]bool{}
			expr, err := reader.ReadExpr(vars)
			Expect(err).To(Not(HaveOccurred()))
			Expect(expr).To(Not(BeNil()))
			Expect(expr.String()).To(Equal(src))
			Expect(vars).To(HaveKeyWithValue(eval.Var("x"), true))
			Expect(vars).To(HaveKeyWithValue(eval.Var("y"), true))
		})
	})

	Context("ReadVar", func() {
		It("should return an error if the input is empty", func() {
			const src = ""
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			_, err := reader.ReadVar()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("empty input"))
		})

		It("should return an error if the input is not a number", func() {
			const src = "a"
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			_, err := reader.ReadVar()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("not a number"))
		})

		It("should return a number if the input is valid", func() {
			const src = "1.32\n25"
			s := bufio.NewScanner(strings.NewReader(src))

			reader := NewReader(s)
			val, err := reader.ReadVar()
			Expect(err).To(Not(HaveOccurred()))
			Expect(val).To(BeNumerically("~", 1.32, 0.01))

			val, err = reader.ReadVar()
			Expect(err).To(Not(HaveOccurred()))
			Expect(val).To(BeNumerically("~", 25, 0.01))
		})
	})
})
