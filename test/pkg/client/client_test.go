package client_test

import (
	"github.com/nikimanoledaki/calculator-microservice/pkg/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		err       error
		args      []string
		operation string
	)

	var _ = Describe("Function ParseArguments", func() {
		JustBeforeEach(func() {
			operation, err = client.ParseArguments(args)
		})

		Context("If operator is not `sum` or `average`", func() {
			BeforeEach(func() {
				args = []string{"./client", "multiply", "1", "2"}
			})

			It("returns an error saying the operation is not recognized", func() {
				Expect(err).Should(MatchError("operation not recognized"))
				Expect(operation).To(Equal(""))
			})
		})

		Context("If the arguments are not the binary, an operator, and two digits", func() {
			BeforeEach(func() {
				args = []string{"./client", "sum", "1", "2", "3"}
			})

			It("it returns an error", func() {
				Expect(err).Should(MatchError("expected 'sum' or 'average' with 2 numeric values"))
				Expect(operation).To(Equal(""))
			})
		})

		Context("If there is a client, an operator, and two digits", func() {
			BeforeEach(func() {
				args = []string{"./client", "sum", "1", "2"}
			})

			It("returns no error and the operator", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(operation).To(Equal("sum"))
			})
		})
	})
})
