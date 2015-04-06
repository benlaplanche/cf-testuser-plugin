package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
	"sort"
	"strconv"
	"strings"
)

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

		// so we can iterate on the map in the desired order
		var keys []int
		for k := range c.Command {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, key := range keys {
			output, status := c.execCommand(cliConnection, c.Command[key])
			c.printMessages(output, status)
			if foundError(status) == true {
				return
			}
		}
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

func (t TestUser) execCommand(cliConnection plugin.CliConnection, command CommandData) (output []string, status []int) {

	output = append(output, command.Message)

	response, err := cliConnection.CliCommandWithoutTerminalOutput(command.ExecutionArguments...)

	if response != nil && strings.Contains(response[2], "already exists") {
		status = append(status, 2)
	} else if err != nil {
		status = append(status, 0)
	} else {
		status = append(status, 1)
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

func (t TestUser) commandCounter() (count string) {
	current := strconv.Itoa(t.CmdRunCount)
	total := strconv.Itoa(CmdTotalCount)

	return "[" + current + "/" + total + "] "
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
