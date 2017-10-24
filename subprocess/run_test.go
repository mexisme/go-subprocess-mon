package subprocess_test

import (
	. "bitbucket.org/mexisme/get-secrets/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("subprocess.run", func() {
	var command *Details

	BeforeEach(func() {
		command = New()
	})

	It("running 'true' succeeds", func() {
		Expect(command.WithCommand([]string{"true"}).Run()).NotTo(HaveOccurred())
	})

	It("running 'false' fails", func() {
		Expect(command.WithCommand([]string{"false"}).Run()).To(HaveOccurred())
	})
})
