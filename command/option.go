package command

import "io"

type Option func(*Command)

func OptionWorkingDir(dir string) Option {
	return func(cmd *Command) {
		cmd.Cmd.Dir = dir
	}
}

func OptionStdout(stdout io.Writer) Option {
	return func(cmd *Command) {
		cmd.Cmd.Stdout = stdout
	}
}

func OptionEnv(env Env) Option {
	return func(cmd *Command) {
		cmd.Cmd.Env = env.parse()
	}
}

type Env map[string]string

func (e Env) parse() []string {
	var parsed []string
	for k, v := range e {
		parsed = append(parsed, k+"="+v)
	}
	return parsed
}
