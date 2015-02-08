package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
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
				HelpText: "Command to create a passed in user and development org & space and grant all permissions",
				UsageDetails: plugin.Usage{
					Usage: "cf test-user <username> <password>",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(TestUser))
}

func (c *TestUser) Run(cliConnection plugin.CliConnection, args []string) {

	if len(args) < 3 {
		fmt.Println("Incorrect usage")
		fmt.Println(c.GetMetadata().Commands[0].UsageDetails.Usage)
	} else {
		c.CreateUser(cliConnection, args)
		c.CreateOrg(cliConnection, args)
		c.CreateSpace(cliConnection, args)
		c.AssignOrgRole(cliConnection, args)
	}
}

func (c *TestUser) CreateUser(cliConnection plugin.CliConnection, args []string) {

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", args[1], args[2])

	if err != nil {
		fmt.Println(colorstring.Color("[red][1/10]  Created user " + args[1]))
		fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][1/10]  Created user " + args[1]))
		fmt.Println(output)
	}

}

func (c *TestUser) CreateOrg(cliConnection plugin.CliConnection, args []string) {

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", "development")

	if err != nil {
		fmt.Println(colorstring.Color("[red][2/10]  Created Organisation development"))
		fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][2/10]  Created Organisation development"))
		fmt.Println(output)
	}
}

func (c *TestUser) CreateSpace(cliConnection plugin.CliConnection, args []string) {

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", "development", "-o", "development")

	if err != nil {
		fmt.Println(colorstring.Color("[red][3/10]  Created Space development"))
		fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][3/10]  Created Space development"))
		fmt.Println(output)
	}
}

func (c *TestUser) AssignOrgRole(cliConnection plugin.CliConnection, args []string) {

	output, err := cliConnection.CliCommandWithoutTerminalOutput("set-org-role", args[0], "OrgManager")

	if err != nil {
		fmt.Println(colorstring.Color("[red][4/10]  Assigned OrgManager to me in Org development"))
		fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][4/10]  Assigned OrgManager to me in Org development"))
		fmt.Println(output)
	}
}
