package main

import c "github.com/yagihash/startup/command"

func main() {
	commands := []*c.Command{
		c.NewCommand("ssh-add", []string{"-K"}),
		c.NewCommand("git", []string{"pull"}, c.OptionWorkingDir("/Users/yagihash/dotfiles")),
	}

	for _, cmd := range commands {
		cmd.Run()
	}
}
