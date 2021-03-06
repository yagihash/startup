package main

import (
	"os"

	c "github.com/yagihash/startup/command"
)

func main() {
	commands := []*c.Command{
		c.New("ssh-add", []string{"-K"}),
		c.New("go", []string{"get", "-u", "github.com/yagihash/startup/cmd/startup"}),
		c.New("git", []string{"pull"}, c.OptionWorkingDir(os.Getenv("HOME")+"/dotfiles")),
	}

	for _, cmd := range commands {
		cmd.Run()
	}
}
