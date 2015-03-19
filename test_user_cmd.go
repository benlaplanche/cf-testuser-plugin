package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
	"strconv"
	"strings"
)

type TestUser struct{}

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

var CmdRunCount int

func CommandCounter() (count string) {
	current := strconv.Itoa(CmdRunCount)
	total := strconv.FormatInt(CmdTotalCount, 10)

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

		CmdRunCount = 1

		c.RunCommands(cliConnection, args)
	}
}

func (c *TestUser) RunCommands(cliConnection plugin.CliConnection, args []string) (response bool) {

	output, status := c.CreateUser(cliConnection, args)
	PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.CreateOrg(cliConnection, args)
	PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.CreateSpace(cliConnection, args)
	PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.OrgRoles(cliConnection, args)
	PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}

	output, status = c.SpaceRoles(cliConnection, args)
	PrintMessages(output, status)
	if FoundError(status) == true {
		return
	}
	return
}

func PrintMessages(output []string, status []int) {
	for i, v := range output {
		switch status[i] {
		case 0:
			fmt.Println(colorstring.Color("[red]" + CommandCounter() + v))
		case 1:
			fmt.Println(colorstring.Color("[green]" + CommandCounter() + v))
		case 2:
			fmt.Println(colorstring.Color("[cyan]" + CommandCounter() + v))
		}

		CmdRunCount++
	}
}

func (c *TestUser) CreateUser(cliConnection plugin.CliConnection, args []string) (output []string, status []int) {

	output = append(output, "Created user "+args[1])

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", args[1], args[2])

	if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) CreateOrg(cliConnection plugin.CliConnection, args []string) (output []string, status []int) {

	output = append(output, "Created Organisation development")

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", "development")

	if x != nil && strings.Contains(x[0], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) CreateSpace(cliConnection plugin.CliConnection, args []string) (output []string, status []int) {

	output = append(output, "Created Space development")

	x, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", "development", "-o", "development")

	if x != nil && strings.Contains(x[0], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
	}

	return
}

func (c *TestUser) OrgRoles(cliConnection plugin.CliConnection, args []string) (output []string, status []int) {

	for _, v := range OrgRoles {
		output = append(output, "Assigned "+v+" to me in Org development")

		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-org-role", args[1], "development", v)

		if err != nil {
			break
			status = append(status, 0)
		} else {
			status = append(status, 1)
		}
	}
	return
}

func (c *TestUser) SpaceRoles(cliConnection plugin.CliConnection, args []string) (output []string, status []int) {

	for _, v := range SpaceRoles {
		output = append(output, "Assigned "+v+" to me in Space development")

		_, err := cliConnection.CliCommand("set-space-role", args[1], "development", "development", v)

		if err != nil {
			break
			status = append(status, 0)
		} else {
			status = append(status, 1)
		}
	}
	return
}
