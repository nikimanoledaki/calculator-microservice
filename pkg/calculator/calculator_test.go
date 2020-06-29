package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nikimanoledaki/calculator-microservice/pkg/calculator"
)

var _ = Describe("Calculator", func() {
	Context("Can do basic arithmetic operations and can", func() {
		It("add two numbers", func() {
			Expect(calculator.Sum(1, 1)).To(Equal(2))
		})
		It("substract a number from the first one", func() {
			Expect(calculator.Subtract(2, 1)).To(Equal(1))
		})
		It("multiply two numbers", func() {
			Expect(calculator.Multiply(2, 2)).To(Equal(4))
		})
		It("divide a number by another number", func() {
			Expect(calculator.Divide(2, 2)).To(Equal(1))
		})
	})
})
