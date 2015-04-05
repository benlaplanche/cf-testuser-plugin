package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
	"strconv"
	"strings"
)

type CommandData struct {
	Message            string
	ExecutionArguments []string
}

type TestUser struct {
	UserName    string
	Password    string
	OrgName     string
	SpaceName   string
	CmdRunCount int
	Command     map[int]CommandData
}

func (c *TestUser) setCommands() {
	c.Command = make(map[int]CommandData)

	c.Command[1] = CommandData{
		Message: "Created user " + c.UserName,
		ExecutionArguments: []string{
			"create-user",
			c.UserName,
			c.Password,
		},
	}
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

func (t TestUser) commandCounter() (count string) {
	current := strconv.Itoa(t.CmdRunCount)
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
		c.setCommands()

		// c.runCommands(cliConnection)
		// c.execCommand(cliConnection, c.Command[1])
		output, status := c.execCommand(cliConnection, c.Command[1])
		c.printMessages(output, status)
		if foundError(status) == true {
			return
		}
	}
}

func (t TestUser) execCommand(cliConnection plugin.CliConnection, command CommandData) (output []string, status []int) {
	// fmt.Println(command.Message)
	output = append(output, command.Message)

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", t.UserName, t.Password)

	if x != nil && strings.Contains(x[2], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (t TestUser) runCommands(cliConnection plugin.CliConnection) (response bool) {

	output, status := t.createUser(cliConnection)
	t.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = t.createOrg(cliConnection)
	t.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = t.createSpace(cliConnection)
	t.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = t.assignOrgRoles(cliConnection)
	t.printMessages(output, status)
	if foundError(status) == true {
		return
	}

	output, status = t.assignSpaceRoles(cliConnection)
	t.printMessages(output, status)
	if foundError(status) == true {
		return
	}
	return
}

func (t *TestUser) printMessages(output []string, status []int) {
	for i, v := range output {
		switch status[i] {
		case 0:
			fmt.Println(colorstring.Color("[red]" + t.commandCounter() + v))
		case 1:
			fmt.Println(colorstring.Color("[green]" + t.commandCounter() + v))
		case 2:
			fmt.Println(colorstring.Color("[cyan]" + t.commandCounter() + v))
		}

		t.CmdRunCount++
	}
}

func (t TestUser) createUser(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created user "+t.UserName)

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", t.UserName, t.Password)

	if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (t TestUser) createOrg(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created Organisation "+t.OrgName)

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", t.OrgName)

	if x != nil && strings.Contains(x[2], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (t TestUser) createSpace(cliConnection plugin.CliConnection) (output []string, status []int) {

	output = append(output, "Created Space "+t.SpaceName)

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", t.OrgName, "-o", t.SpaceName)

	if x != nil && strings.Contains(x[2], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (t TestUser) assignOrgRoles(cliConnection plugin.CliConnection) (output []string, status []int) {

	for _, role := range assignOrgRoles {
		output = append(output, "Assigned "+role+" to "+t.UserName+" in Org "+t.OrgName)

		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-org-role", t.UserName, t.OrgName, role)

		if err != nil {
			break
			status = append(status, 0)
		} else {
			status = append(status, 1)
		}
	}
	return
}

func (t TestUser) assignSpaceRoles(cliConnection plugin.CliConnection) (output []string, status []int) {

	for _, role := range assignSpaceRoles {
		output = append(output, "Assigned "+role+" to "+t.UserName+" in Space "+t.SpaceName)

		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-space-role", t.UserName, t.OrgName, t.SpaceName, role)

		if err != nil {
			break
			status = append(status, 0)
		} else {
			status = append(status, 1)
		}
	}
	return
}
