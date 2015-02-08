package main_test

import (
	. "github.com/benlaplanche/cf-testuser-plugin"
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	"github.com/mitchellh/colorstring"
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

		It("creates a new user", func() {
			output := io_helpers.CaptureOutput(func() {
				callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
			})

			Expect(output[0]).To(Equal(colorstring.Color("[green][1/10]  Created user me")))

			// Expect(output[1]).To(Equal("Creating user me as admin..."))
			// Expect(output[1]).To(Equal("OK"))
			// Expect(output[2]).To(Equal(""))
			// Expect(output[3]).To(Equal("TIP: Assign roles with 'cf set-org-role' and 'cf set-space-role'"))
		})

	})
})
