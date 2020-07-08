package client_test

import (
	"github.com/golang/mock/gomock"
	calc_mock "github.com/nikimanoledaki/calculator-microservice/mock/calculator"
	"github.com/nikimanoledaki/calculator-microservice/pkg/client"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
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

	var _ = Describe("gRPC functions", func() {

		var (
			ctrl           *gomock.Controller
			mockCalcClient *calc_mock.MockCalculatorClient
		)

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockCalcClient = calc_mock.NewMockCalculatorClient(ctrl)
			args = []string{"1", "2"}
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		Context("When passing a SumRequest and CalculatorClient to function PrintSum", func() {

			var (
				response  *protos.SumResponse
				response2 *protos.SumResponse
			)

			JustBeforeEach(func() {
				mockCalcClient.EXPECT().GetSum(gomock.Any(), gomock.Any()).Return(&protos.SumResponse{Result: 3}, nil)
				response, err = client.PrintSum(mockCalcClient, args)
				response2 = new(protos.SumResponse)
			})

			It("it returns a SumResponse", func() {
				Expect(response).Should(BeAssignableToTypeOf(response2))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When passing an AverageRequest and CalculatorClient to function PrintAverage", func() {

			var (
				response  *protos.AverageResponse
				response2 *protos.AverageResponse
			)

			JustBeforeEach(func() {
				mockCalcClient.EXPECT().GetAverage(gomock.Any(), gomock.Any()).Return(&protos.AverageResponse{Result: 2}, nil)
				response, err = client.PrintAverage(mockCalcClient, args)
				response2 = new(protos.AverageResponse)
			})

			It("it returns an AverageResponse", func() {
				Expect(response).Should(BeAssignableToTypeOf(response2))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
