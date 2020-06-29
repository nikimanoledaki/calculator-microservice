package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {
	var (
		calculatorBinary  string
		calculatorCommand *exec.Cmd
	)

	BeforeEach(func() {
		var err error

		calculatorBinary, err = gexec.Build("github.com/nikimanoledaki/calculator-microservice", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("When the command line arguments are `sum 1 1`, it prints `2`", func() {
		calculatorCommand = exec.Command(calculatorBinary, "sum 1 1")

		session, err := gexec.Start(calculatorCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("2"))
	})
})
