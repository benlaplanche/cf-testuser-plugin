package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
	"strconv"
	"strings"
)

type TestUser struct {
	UserName    string
	Password    string
	OrgName     string
	SpaceName   string
	CmdRunCount int
}

var OrgRoles = []string{
	"OrgManager",
	"BillingManager",
	"OrgAuditor",
}

var SpaceRoles = []string{
	"SpaceManager",
	"SpaceDeveloper",
	"SpaceAuditor",
}

const CmdTotalCount = 10
const DefaultOrg = "development"
const DefaultSpace = "development"

func (c *TestUser) CommandCounter() (count string) {
	current := strconv.Itoa(c.CmdRunCount)
	total := strconv.Itoa(CmdTotalCount)

	return "[" + current + "/" + total + "] "
}

func SearchIntSlice(slice []int, seek int) (result bool) {
	for _, v := range slice {
		if v == seek {
			return true
			break
		}
	}
	return false
}

func FoundError(status []int) (response bool) {
	if SearchIntSlice(status, 0) == true {
		return true
	}
	return false
}

func (c *TestUser) SetProperties(args []string) {

	c.UserName = args[1]
	c.Password = args[2]
	c.CmdRunCount = 1

	if len(args) >= 4 {
		c.OrgName = args[3]
	} else {
		c.OrgName = DefaultOrg
	}

	if len(args) >= 5 {
		c.SpaceName = args[4]
	} else {
		c.SpaceName = DefaultSpace
	}
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
				HelpText: "Create a user and assign all possible permissions, organisation and space are created if they do not already exist as well. If no organisation or space name are specified then the default value of 'development' is used",
				UsageDetails: plugin.Usage{
					Usage: "cf test-user <username> <password> <optional org name> <optional space name>",
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

		c.SetProperties(args)

		c.RunCommands(cliConnection)
	}
}

func (c *TestUser) RunCommands(cliConnection plugin.CliConnection) (response bool) {

	output, status := c.CreateUser(cliConnection)
	c.PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.CreateOrg(cliConnection)
	c.PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.CreateSpace(cliConnection)
	c.PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.OrgRoles(cliConnection)
	c.PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.SpaceRoles(cliConnection)
	c.PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}
	return
}

func (c *TestUser) PrintMessages(output []string, status []int) {
	for i, v := range output {
		switch status[i] {
		case 0:
			fmt.Println(colorstring.Color("[red]" + c.CommandCounter() + v))
		case 1:
			fmt.Println(colorstring.Color("[green]" + c.CommandCounter() + v))
		case 2:
			fmt.Println(colorstring.Color("[cyan]" + c.CommandCounter() + v))
		}

		c.CmdRunCount++
	}
}

func (c *TestUser) CreateUser(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created user "+c.UserName)

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", c.UserName, c.Password)

	if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) CreateOrg(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created Organisation "+c.OrgName)

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", c.OrgName)

	if x != nil && strings.Contains(x[2], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) CreateSpace(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created Space "+c.SpaceName)

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", c.OrgName, "-o", c.SpaceName)

	if x != nil && strings.Contains(x[2], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) OrgRoles(cliConnection plugin.CliConnection) (output []string, status []int) {

	for _, role := range OrgRoles {
		output = append(output, "Assigned "+role+" to "+c.UserName+" in Org "+c.OrgName)

		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-org-role", c.UserName, c.OrgName, role)

		if err != nil {
			break
			status = append(status, 0)
		} else {
			status = append(status, 1)
		}
	}
	return
}

func (c *TestUser) SpaceRoles(cliConnection plugin.CliConnection) (output []string, status []int) {

	for _, role := range SpaceRoles {
		output = append(output, "Assigned "+role+" to "+c.UserName+" in Space "+c.SpaceName)

		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-space-role", c.UserName, c.OrgName, c.SpaceName, role)

		if err != nil {
			break
			status = append(status, 0)
		} else {
			status = append(status, 1)
		}
	}
	return
}
