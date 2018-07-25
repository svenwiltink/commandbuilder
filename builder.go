package commandbuilder

import (
	"os/exec"
)

type CommandBuilder struct {
	Prefix    string
	Arguments []string
}

func (c *CommandBuilder) AddArgument(argument string) *CommandBuilder {
	return c.AddArguments(argument)
}

func (c *CommandBuilder) AddArguments(arguments ...string) *CommandBuilder {
	c.Arguments = append(c.Arguments, arguments...)
	return c
}

func (c *CommandBuilder) ToCmd() *exec.Cmd {
	cmd := exec.Command(c.Prefix, c.Arguments...)
	return cmd
}

func New(prefix string) *CommandBuilder {
	instance := new(CommandBuilder)
	instance.Prefix = prefix
	instance.Arguments = make([]string, 0)

	return instance
}
