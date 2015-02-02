package main_test

import (
	. "github.com/benlaplanche/cf-testuser-plugin"
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestUserCmd", func() {
	Describe(".Run", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var callCliCommandPlugin *TestUser

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			callCliCommandPlugin = &TestUser{}
		})

		It("returns the correct output", func() {
			fakeCliConnection.CliCommandReturns([]string{"test-user"}, nil)
			output := io_helpers.CaptureOutput(func() {
				callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user"})
			})

			Expect(output[0]).To(Equal("running the new test user command"))
			Expect(output[1]).To(Equal("[test-user]"))
		})
	})
})
