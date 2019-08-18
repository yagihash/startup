package command

import (
	"fmt"
	"os/exec"

	"github.com/martinlindhe/notify"
)

const (
	// AppName is shown on notifications as application name
	AppName = "Start up"

	// TitleError is shown on notification as error title
	TitleError = "Error"
)

// Command wraps exec.Cmd and has some command execution method with notification
type Command struct {
	Cmd *exec.Cmd
}

// NewCommand creates Command with given options
func NewCommand(cmd string, cmdOption []string, execOption ...Option) *Command {
	c := &Command{
		Cmd: exec.Command(cmd, cmdOption...),
	}

	for _, opt := range execOption {
		opt(c)
	}

	return c
}

// Run executes command contained in Command and notify when it fails
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
