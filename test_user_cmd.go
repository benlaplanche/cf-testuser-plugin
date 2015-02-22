package main

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
	"reflect"
	"sort"
	"strconv"
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
	// fmt.Println(result)
	return
}

func (c *TestUser) Run(cliConnection plugin.CliConnection, args []string) {

	if len(args) < 3 {
		fmt.Println("Incorrect usage")
		fmt.Println(c.GetMetadata().Commands[0].UsageDetails.Usage)
	} else {
		// insert clever loop logic here
		var keys []int

		funcs := map[int]interface{}{
			1: c.CreateUser,
			2: c.CreateOrg,
			3: c.CreateSpace,
			4: c.Orgs,
			5: c.Spaces}

		for k := range funcs {
			keys = append(keys, k)
		}

		sort.Ints(keys)

		for _, k := range keys {
			val, _ := Call(funcs, k, cliConnection, args)
			// fmt.Println(val[0].Bool())
			// fmt.Println(err)
			if val[0].Bool() == false {
				fmt.Printf("break caused on key: %s", k)
				break
			}
		}
	}
}

func (c *TestUser) Orgs(cliConnection plugin.CliConnection, args []string) (success bool) {

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

func (c *TestUser) Spaces(cliConnection plugin.CliConnection, args []string) (success bool) {

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

func (c *TestUser) CreateUser(cliConnection plugin.CliConnection, args []string) (success bool) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-user", args[1], args[2])

	if err != nil {
		fmt.Println(colorstring.Color("[red][1/10]  Created user " + args[1]))
		// fmt.Println(err)
		// os.Exit(1)
		return false
	} else {
		fmt.Println(colorstring.Color("[green][1/10]  Created user " + args[1]))
		// fmt.Println(output)
		return true
	}

}

func (c *TestUser) CreateOrg(cliConnection plugin.CliConnection, args []string) (success bool) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-org", "development")

	if err != nil {
		fmt.Println(colorstring.Color("[red][2/10]  Created Organisation development"))
		// fmt.Println(err)
		return false
	} else {
		fmt.Println(colorstring.Color("[green][2/10]  Created Organisation development"))
		// fmt.Println(output)
		return true
	}
}

func (c *TestUser) CreateSpace(cliConnection plugin.CliConnection, args []string) (success bool) {

	_, err := cliConnection.CliCommandWithoutTerminalOutput("create-space", "development", "-o", "development")

	if err != nil {
		fmt.Println(colorstring.Color("[red][3/10]  Created Space development"))
		// fmt.Println(err)
		return false
	} else {
		fmt.Println(colorstring.Color("[green][3/10]  Created Space development"))
		// fmt.Println(output)
		return true
	}
}

func (c *TestUser) SwitchUser(cliConnection plugin.CliConnection, args []string) {

	fmt.Println(colorstring.Color("[green][10/10]  Logged out and logged in as " + args[1]))
}
