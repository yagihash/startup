package main

import "github.com/yagihash/startup/command"

func main() {
	commands := []*command.Command{
		command.NewCommand("ssh-add", []string{"-K"}),
	}

	for _, c := range commands {
		c.Run()
	}
}
