package main

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
	"reflect"
	"sort"
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

func Call(m map[int]interface{}, position int, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[position])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func (c *TestUser) Run(cliConnection plugin.CliConnection, args []string) {

	if len(args) < 3 {
		fmt.Println("Incorrect usage")
		fmt.Println(c.GetMetadata().Commands[0].UsageDetails.Usage)
	} else {
		c.RunCommands2(cliConnection, args)
	}
}

func (c *TestUser) RunCommands(cliConnection plugin.CliConnection, args []string) {

	var keys []int

	commands := map[int]interface{}{
		1: c.CreateUser,
		2: c.CreateOrg,
		3: c.CreateSpace,
		4: c.OrgRoles,
		5: c.SpaceRoles}

	for k := range commands {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		val, _ := Call(commands, k, cliConnection, args)
		if val[0].Bool() == false {
			break
		}
	}

}

func (c *TestUser) RunCommands2(cliConnection plugin.CliConnection, args []string) {

	var keys []int

	commands := map[int]interface{}{
		1: c.CreateUser2,
		2: c.CreateOrg2,
		3: c.CreateSpace2,
		4: c.OrgRoles2,
		5: c.SpaceRoles2}

	for k := range commands {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	// var messages []string
	// var flags []int64

	for _, k := range keys {
		output, _ := Call(commands, k, cliConnection, args)
		fmt.Println(output[0])
		fmt.Println(output[1])
		fmt.Println(len(output))
		// fmt.Println(success)
		// messages = append(messages, ReflectToString(output[0]))
		// flags = append(flags, ReflectToInt(output[1]))
		// messages = append(messages, output...)
		// flags = append(flags, output...)
		// fmt.Println(success.search(1))
		// if success.Search(1) ==  {
		// 	break
		// }

	}

}

func ReflectToString(values []reflect.Value) (output []string) {
	for _, v := range values {
		output = append(output, v.String())
	}
	return
}

func ReflectToInt(values []reflect.Value) (output []int64) {
	for _, v := range values {
		output = append(output, v.Int())
	}
	return
}

func (c *TestUser) CreateUser2(cliConnection plugin.CliConnection, args []string) (output []string, success []int) {

	output = append(output, "Created user "+args[1])

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", args[1], args[2])

	if err != nil {
		success = append(success, 0)
	} else {
		success = append(success, 1)
	}

	return
}

func (c *TestUser) CreateUser(cliConnection plugin.CliConnection, args []string) (success bool) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", args[1], args[2])

	if err != nil {
		fmt.Println(colorstring.Color("[red][1/10]  Created user " + args[1]))
		success = false
	} else {
		fmt.Println(colorstring.Color("[green][1/10]  Created user " + args[1]))
		success = true
	}
	return
}

func (c *TestUser) CreateOrg2(cliConnection plugin.CliConnection, args []string) (output []string, success []int) {

	output = append(output, "Created Organisation development")

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", "development")

	if output != nil && strings.Contains(output[0], "already exists") {
		success = append(success, 2)
	} else if err != nil {
		success = append(success, 0)
	} else {
		success = append(success, 1)
	}

	return
}

func (c *TestUser) CreateOrg(cliConnection plugin.CliConnection, args []string) (success bool) {

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", "development")

	if output != nil && strings.Contains(output[0], "already exists") {
		fmt.Println(colorstring.Color("[cyan][2/10]  Created Organisation development"))

		success = true
	} else if err != nil {
		fmt.Println(colorstring.Color("[red][2/10]  Created Organisation development"))

		success = false
	} else {
		fmt.Println(colorstring.Color("[green][2/10]  Created Organisation development"))
		success = true
	}
	return
}

func (c *TestUser) CreateSpace2(cliConnection plugin.CliConnection, args []string) (output []string, success []int) {

	output = append(output, "Created Space development")

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", "development", "-o", "development")

	if output != nil && strings.Contains(output[0], "already exists") {
		success = append(success, 2)
	} else if err != nil {
		success = append(success, 0)
	} else {
		success = append(success, 1)
	}

	return
}

func (c *TestUser) CreateSpace(cliConnection plugin.CliConnection, args []string) (success bool) {

	output, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", "development", "-o", "development")

	if output != nil && strings.Contains(output[0], "already exists") {
		fmt.Println(colorstring.Color("[cyan][3/10]  Created Space development"))

		success = true
	} else if err != nil {
		fmt.Println(colorstring.Color("[red][3/10]  Created Space development"))
		success = false
	} else {
		fmt.Println(colorstring.Color("[green][3/10]  Created Space development"))
		success = true
	}
	return
}

func (c *TestUser) OrgRoles2(cliConnection plugin.CliConnection, args []string) (output []string, success []int) {

	for _, v := range OrgRoles {
		output = append(output, "Assigned "+v+" to me in Org development")

		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-org-role", args[1], "development", v)

		if err != nil {
			break
			success = append(success, 0)
		} else {
			success = append(success, 1)
		}
	}
	return
}

func (c *TestUser) OrgRoles(cliConnection plugin.CliConnection, args []string) (success bool) {

	for i, v := range OrgRoles {
		_, err := cliConnection.CliCommandWithoutTerminalOutput("set-org-role", args[1], "development", v)
		index := strconv.Itoa(i + 4)

		if err != nil {
			fmt.Println(colorstring.Color("[red][" + index + "/10]  Assigned " + v + " to me in Org development"))
			break
			success = false
		} else {
			fmt.Println(colorstring.Color("[green][" + index + "/10]  Assigned " + v + " to me in Org development"))
			success = true
		}
	}
	return
}

func (c *TestUser) SpaceRoles2(cliConnection plugin.CliConnection, args []string) (output []string, success []int) {

	for _, v := range SpaceRoles {
		output = append(output, "Assigned "+v+" to me in Space development")

		_, err := cliConnection.CliCommand("set-space-role", args[1], "development", "development", v)

		if err != nil {
			break
			success = append(success, 0)
		} else {
			success = append(success, 1)
		}
	}
	return
}

func (c *TestUser) SpaceRoles(cliConnection plugin.CliConnection, args []string) (success bool) {

	for i, v := range SpaceRoles {
		_, err := cliConnection.CliCommand("set-space-role", args[1], "development", "development", v)
		index := strconv.Itoa(i + 7)
		if err != nil {
			fmt.Println(colorstring.Color("[red][" + index + "/10]  Assigned " + v + " to me in Space development"))
			break
			success = false
		} else {

			fmt.Println(colorstring.Color("[green][" + index + "/10]  Assigned " + v + " to me in Space development"))
			success = true
		}
	}
	return
}
