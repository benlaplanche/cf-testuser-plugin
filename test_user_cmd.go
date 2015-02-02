package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
)

type TestUser struct{}

func (c *TestUser) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "TestUser",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "test-user",
				HelpText: "Command to create a test user and development org & space and grant all permissions",
				UsageDetails: plugin.Usage{
					Usage: "test-user username password",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(TestUser))
}

func (c *TestUser) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("running the new test user command")
	fmt.Println(args)
}
