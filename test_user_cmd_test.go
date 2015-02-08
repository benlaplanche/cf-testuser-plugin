package main_test

import (
	"fmt"
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
			fmt.Println(output)
			Expect(output[0]).To(Equal(colorstring.Color("[green][1/10]  Created user me")))

			Expect(output[2]).To(Equal(colorstring.Color("[green][2/10]  Created Organisation development")))

			Expect(output[4]).To(Equal(colorstring.Color("[green][3/10]  Created Space development")))

			Expect(output[6]).To(Equal(colorstring.Color("[green][4/10]  Assigned OrgManager to me in Org development")))

			Expect(output[8]).To(Equal(colorstring.Color("[green][5/10]  Assigned BillingManager to me in Org development")))

			Expect(output[10]).To(Equal(colorstring.Color("[green][6/10]  Assigned OrgAuditor to me in Org development")))

			Expect(output[12]).To(Equal(colorstring.Color("[green][7/10]  Assigned SpaceManager to me in Space development")))

			Expect(output[14]).To(Equal(colorstring.Color("[green][8/10]  Assigned SpaceDeveloper to me in Space development")))

			Expect(output[16]).To(Equal(colorstring.Color("[green][9/10]  Assigned SpaceAuditor to me in Space development")))

			Expect(output[18]).To(Equal(colorstring.Color("[green][10/10]  Logged out and logged in as me")))

		})

	})
})
