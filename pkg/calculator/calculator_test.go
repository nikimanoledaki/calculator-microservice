package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nikimanoledaki/calculator-microservice/pkg/calculator"
)

var _ = Describe("Calculator", func() {
	Context("Can do basic arithmetic operations such as", func() {
		It("adding two numbers", func() {
			Expect(calculator.Sum(1, 1)).To(Equal(2))
		})
		It("substracting a number from the first one", func() {
			Expect(calculator.Subtract(2, 1)).To(Equal(1))
		})
		It("multiplying two numbers", func() {
			Expect(calculator.Multiply(2, 2)).To(Equal(4))
		})
		It("dividing a number by another number", func() {
			Expect(calculator.Divide(2, 2)).To(Equal(1))
		})
	})

	Context("When it receives an operand as a string", func() {
		It("knows if it should sum the numbers", func() {
			Expect(calculator.Compute("sum", 2, 3)).To(Equal(5))
		})
		It("knows if it should subtract from the first number", func() {
			Expect(calculator.Compute("subtract", 4, 3)).To(Equal(1))
		})
	})
})
