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
			callCliCommandPlugin = &CliCmd{}
		})

		It("returns the correct output", func() {
			io_helpers.CaptureOutput(func() {
				callCliCommandPlugin.run(fakeCliConnection, []string{"test-user"})
			})

			Expect(fakeCliConnection.CliCommandArgsForCall(0)[0]).To(Equal("running the new test user command"))
			Expect(fakeCliConnection.CliCommandArgsForCall(0)[0]).To(Equal("test-user"))
		})
	})
})
