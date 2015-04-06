package main

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

const CmdTotalCount = 10
const DefaultOrg = "development"
const DefaultSpace = "development"

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

	c.Command[2] = CommandData{
		Message: "Created Organisation " + c.OrgName,
		ExecutionArguments: []string{
			"create-org",
			c.OrgName,
		},
	}

	c.Command[3] = CommandData{
		Message: "Created Space " + c.SpaceName,
		ExecutionArguments: []string{
			"create-space",
			c.SpaceName,
		},
	}

	c.Command[4] = CommandData{
		Message: "Assigned OrgManager to " + c.UserName + " in Org " + c.OrgName,
		ExecutionArguments: []string{
			"set-org-role",
			c.UserName,
			c.OrgName,
			"OrgManager",
		},
	}

	c.Command[5] = CommandData{
		Message: "Assigned BillingManager to " + c.UserName + " in Org " + c.OrgName,
		ExecutionArguments: []string{
			"set-org-role",
			c.UserName,
			c.OrgName,
			"BillingManager",
		},
	}

	c.Command[6] = CommandData{
		Message: "Assigned OrgAuditor to " + c.UserName + " in Org " + c.OrgName,
		ExecutionArguments: []string{
			"set-org-role",
			c.UserName,
			c.OrgName,
			"AuditorManager",
		},
	}

	c.Command[7] = CommandData{
		Message: "Assigned SpaceManager to " + c.UserName + " in Space " + c.SpaceName,
		ExecutionArguments: []string{
			"set-space-role",
			c.UserName,
			c.OrgName,
			c.SpaceName,
			"SpaceManager",
		},
	}

	c.Command[8] = CommandData{
		Message: "Assigned SpaceDeveloper to " + c.UserName + " in Space " + c.SpaceName,
		ExecutionArguments: []string{
			"set-org-role",
			c.UserName,
			c.OrgName,
			c.SpaceName,
			"SpaceDeveloper",
		},
	}

	c.Command[9] = CommandData{
		Message: "Assigned SpaceAuditor to " + c.UserName + " in Space " + c.SpaceName,
		ExecutionArguments: []string{
			"set-org-role",
			c.UserName,
			c.OrgName,
			c.SpaceName,
			"SpaceAuditor",
		},
	}

	c.Command[10] = CommandData{
		Message: "Logged out and logged in as " + c.UserName,
		ExecutionArguments: []string{
			"auth",
			c.UserName,
			c.Password,
		},
	}

}
