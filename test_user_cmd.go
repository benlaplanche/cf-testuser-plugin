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

var assignOrgRoles = []string{
	"OrgManager",
	"BillingManager",
	"OrgAuditor",
}

var assignSpaceRoles = []string{
	"SpaceManager",
	"SpaceDeveloper",
	"SpaceAuditor",
}

const CmdTotalCount = 10
const DefaultOrg = "development"
const DefaultSpace = "development"

func (c *TestUser) commandCounter() (count string) {
	current := strconv.Itoa(c.CmdRunCount)
	total := strconv.Itoa(CmdTotalCount)

	return "[" + current + "/" + total + "] "
}

func searchIntSlice(slice []int, seek int) (result bool) {
	for _, v := range slice {
		if v == seek {
			return true
			break
		}
	}
	return false
}

func foundError(status []int) (response bool) {
	if searchIntSlice(status, 0) == true {
		return true
	}
	return false
}

func (c *TestUser) setProperties(args []string) {

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

		c.setProperties(args)

		c.runCommands(cliConnection)
	}
}

func (c *TestUser) runCommands(cliConnection plugin.CliConnection) (response bool) {

	output, status := c.createUser(cliConnection)
	c.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = c.createOrg(cliConnection)
	c.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = c.createSpace(cliConnection)
	c.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = c.assignOrgRoles(cliConnection)
	c.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = c.assignSpaceRoles(cliConnection)
	c.printMessages(output, status)
	if foundError(status) == true {
		return
	}
	return
}

func (c *TestUser) printMessages(output []string, status []int) {
	for i, v := range output {
		switch status[i] {
		case 0:
			fmt.Println(colorstring.Color("[red]" + c.commandCounter() + v))
		case 1:
			fmt.Println(colorstring.Color("[green]" + c.commandCounter() + v))
		case 2:
			fmt.Println(colorstring.Color("[cyan]" + c.commandCounter() + v))
		}

		c.CmdRunCount++
	}
}

func (c *TestUser) createUser(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created user "+c.UserName)

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", c.UserName, c.Password)

	if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) createOrg(cliConnection plugin.CliConnection) (output []string, status []int) {

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

func (c *TestUser) createSpace(cliConnection plugin.CliConnection) (output []string, status []int) {

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

func (c *TestUser) assignOrgRoles(cliConnection plugin.CliConnection) (output []string, status []int) {

	for _, role := range assignOrgRoles {
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

func (c *TestUser) assignSpaceRoles(cliConnection plugin.CliConnection) (output []string, status []int) {

	for _, role := range assignSpaceRoles {
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
