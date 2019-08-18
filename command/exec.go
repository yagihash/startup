package command

import (
	"fmt"
	"os/exec"

	"github.com/martinlindhe/notify"
)

const (
	AppName    = "Start up"
	TitleError = "Error"
	TitleDone  = "Done"
)

type Command struct {
	Cmd *exec.Cmd
}

func NewCommand(cmd string, cmdOption []string, execOption ...Option) *Command {
	c := &Command{
		Cmd: exec.Command(cmd, cmdOption...),
	}

	for _, opt := range execOption {
		opt(c)
	}

	return c
}

func (c *Command) Run() {
	err := c.Cmd.Run()
	if err != nil {
		notify.Notify(
			AppName,
			TitleError,
			fmt.Sprintf("%v", err),
			"",
		)
	}
}
