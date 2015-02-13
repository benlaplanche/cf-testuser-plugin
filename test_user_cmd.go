package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
)

type TestUser struct{}

// var commands = []func(){
// 	CreateUser(cliConnection, args),
// 	CreateOrg(cliConnection, args),
// 	CreateSpace(cliConnection, args),
// 	AssignOrgRole(cliConnection, args, "OrgManager", "4"),
// 	AssignOrgRole(cliConnection, args, "BillingManager", "5"),
// 	AssignOrgRole(cliConnection, args, "OrgAuditor", "6"),
// 	AssignSpaceRole(cliConnection, args, "SpaceManager", "7"),
// 	AssignSpaceRole(cliConnection, args, "SpaceDeveloper", "8"),
// 	AssignSpaceRole(cliConnection, args, "SpaceAuditor", "9"),
// 	SwitchUser(cliConnection, args),
// }

var commands = []func(){
	CreateUser(),
	CreateOrg(),
	CreateSpace(),
	AssignOrgRole(),
	AssignOrgRole(),
	AssignOrgRole(),
	AssignSpaceRole(),
	AssignSpaceRole(),
	AssignSpaceRole(),
	SwitchUser(),
}

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
		for _, v := range commands {
			if err := c.v(cliConnection, args); err != nil {
				break
			}
		}
	}
}

func (c *TestUser) CreateUser(cliConnection plugin.CliConnection, args []string) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", args[1], args[2])

	if err != nil {
		fmt.Println(colorstring.Color("[red][1/10]  Created user " + args[1]))
		// fmt.Println(err)
		// os.Exit(1)
	} else {
		fmt.Println(colorstring.Color("[green][1/10]  Created user " + args[1]))
		// fmt.Println(output)
	}

}

func (c *TestUser) CreateOrg(cliConnection plugin.CliConnection, args []string) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", "development")

	if err != nil {
		fmt.Println(colorstring.Color("[red][2/10]  Created Organisation development"))
		// fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][2/10]  Created Organisation development"))
		// fmt.Println(output)
	}
}

func (c *TestUser) CreateSpace(cliConnection plugin.CliConnection, args []string) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", "development", "-o", "development")

	if err != nil {
		fmt.Println(colorstring.Color("[red][3/10]  Created Space development"))
		// fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][3/10]  Created Space development"))
		// fmt.Println(output)
	}
}

func (c *TestUser) AssignOrgRole(cliConnection plugin.CliConnection, args []string, role string, i string) {

	_, err := cliConnection.CliCommand("set-org-role", args[1], "development", role)

	if err != nil {
		fmt.Println(colorstring.Color("[red][" + i + "/10]  Assigned " + role + " to me in Org development"))
		// fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][" + i + "/10]  Assigned " + role + " to me in Org development"))
		// fmt.Println(output)
	}
}

func (c *TestUser) AssignSpaceRole(cliConnection plugin.CliConnection, args []string, role string, i string) {

	_, err := cliConnection.CliCommand("set-space-role", args[1], "development", "development", role)

	if err != nil {
		fmt.Println(colorstring.Color("[red][" + i + "/10]  Assigned " + role + " to me in Space development"))
		// fmt.Println(err)
	} else {
		fmt.Println(colorstring.Color("[green][" + i + "/10]  Assigned " + role + " to me in Space development"))
		// fmt.Println(output)
	}
}

func (c *TestUser) SwitchUser(cliConnection plugin.CliConnection, args []string) {

	fmt.Println(colorstring.Color("[green][10/10]  Logged out and logged in as " + args[1]))
}
