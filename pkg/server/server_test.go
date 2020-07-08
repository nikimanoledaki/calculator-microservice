package server_test

import (
	"github.com/hashicorp/go-hclog"
	"github.com/nikimanoledaki/calculator-microservice/pkg/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	Context("Function NewComputation", func() {
		var (
			cs     *server.CalculatorServer
			cs2    *server.CalculatorServer
			logger hclog.Logger
		)

		BeforeEach(func() {
			logger = hclog.Default()
			cs = server.NewComputation(logger)
			cs2 = new(server.CalculatorServer)
		})

		It("returns no error and the operator", func() {
			Expect(cs).Should(BeAssignableToTypeOf(cs2))
		})
	})
})
