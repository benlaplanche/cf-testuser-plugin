package main_test

import (
	"errors"
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

		Describe("Testing the CLI arguments", func() {
			It("returns an error if incorrect number of args supplied", func() {

				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user"})
				})

				Expect(output[0]).To(Equal("Incorrect usage"))
				Expect(output[1]).To(Equal("cf test-user <username> <password>"))
			})
		})

		Describe("Testing the happy path", func() {

			It("creates a new user", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
				})
				fmt.Println(output)
				Expect(output[0]).To(Equal(colorstring.Color("[green][1/10] Created user me")))

				Expect(output[1]).To(Equal(colorstring.Color("[green][2/10] Created Organisation development")))

				Expect(output[2]).To(Equal(colorstring.Color("[green][3/10] Created Space development")))

				Expect(output[3]).To(Equal(colorstring.Color("[green][4/10] Assigned OrgManager to me in Org development")))

				Expect(output[4]).To(Equal(colorstring.Color("[green][5/10] Assigned BillingManager to me in Org development")))

				Expect(output[5]).To(Equal(colorstring.Color("[green][6/10] Assigned OrgAuditor to me in Org development")))

				Expect(output[6]).To(Equal(colorstring.Color("[green][7/10] Assigned SpaceManager to me in Space development")))

				Expect(output[7]).To(Equal(colorstring.Color("[green][8/10] Assigned SpaceDeveloper to me in Space development")))

				Expect(output[8]).To(Equal(colorstring.Color("[green][9/10] Assigned SpaceAuditor to me in Space development")))

				// Expect(output[9]).To(Equal(colorstring.Color("[green][10/10]  Logged out and logged in as me")))

			})

		})

		Describe("Cannot create a user", func() {

			BeforeEach(func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub =
					func(args ...string) ([]string, error) {
						if args[0] == "create-user" {
							return nil, errors.New("create user failed")
						}
						return nil, nil
					}
			})

			It("returns an error", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
				})

				Expect(output[0]).To(Equal(colorstring.Color("[red][1/10] Created user me")))

				Expect(len(output)).To(Equal(2))
			})
		})

		Describe("Cannot create an org", func() {

			BeforeEach(func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub =
					func(args ...string) ([]string, error) {
						if args[0] == "create-org" {
							return nil, errors.New("create org failed")
						}
						return nil, nil
					}
			})

			It("returns an error", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
				})

				Expect(output[1]).To(Equal(colorstring.Color("[red][2/10] Created Organisation development")))

				Expect(len(output)).To(Equal(3))
			})
		})

		Describe("Cannot create a space", func() {

			BeforeEach(func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub =
					func(args ...string) (output []string, err error) {
						if args[0] == "create-space" {
							return nil, errors.New("create space failed")
						}
						return nil, nil
					}
			})

			It("returns an error", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
				})

				Expect(output[2]).To(Equal(colorstring.Color("[red][3/10] Created Space development")))

				Expect(len(output)).To(Equal(4))
			})
		})

		Describe("Org already exists", func() {

			BeforeEach(func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub =
					func(args ...string) (output []string, err error) {
						if args[0] == "create-org" {
							output = append(output, "Org development already exists")
							return
						}
						return
					}
			})

			It("returns an error", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
				})
				Expect(output[1]).To(Equal(colorstring.Color("[cyan][2/10] Created Organisation development")))

			})
		})

		Describe("Space already exists", func() {

			BeforeEach(func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub =
					func(args ...string) (output []string, err error) {
						if args[0] == "create-space" {
							output = append(output, "Space development already exists")
							return
						}
						return
					}
			})

			FIt("returns an error", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"test-user", "me", "password"})
				})
				Expect(output[2]).To(Equal(colorstring.Color("[cyan][3/10] Created Space development")))

			})
		})

	})
})
