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

		It("returns an error if incorrect number of args supplied", func() {

			output := io_helpers.CaptureOutput(func() {
				callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user"})
			})

			Expect(output[0]).To(Equal("Incorrect usage"))
			Expect(output[1]).To(Equal("cf test-user <username> <password>"))
		})

		It("creates an Organisation", func() {
			output := io_helpers.CaptureOutput(func() {
				callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
			})

			Expect(output[0]).To(Equal("username = me"))
			Expect(output[1]).To(Equal("password = password"))
		})

	})
})
