package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nikimanoledaki/calculator-microservice"
)

var _ = Describe("Calculator", func() {
	Context("Know when two numbers", func() {
		It("should be added", func() {
			Expect(calculator.Sum(1, 1)).To(Equal(2))
		})
	})
})
