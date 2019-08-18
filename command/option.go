package command

import "io"

// Option is functional option for Command
type Option func(*Command)

// OptionWorkingDir returns functional option with given string describes woking directory to replace Cmd.Dir
func OptionWorkingDir(dir string) Option {
	return func(cmd *Command) {
		cmd.Cmd.Dir = dir
	}
}

// OptionStdout returns functional option with given io.Writer to replace Cmd.Stdout
func OptionStdout(stdout io.Writer) Option {
	return func(cmd *Command) {
		cmd.Cmd.Stdout = stdout
	}
}

// OptionEnv returns functional option with given Env to replace Cmd.Env
func OptionEnv(env Env) Option {
	return func(cmd *Command) {
		cmd.Cmd.Env = env.parse()
	}
}

// Env is alias to map[string]string to express key-value structure of environment variables
type Env map[string]string

func (e Env) parse() []string {
	var parsed []string
	for k, v := range e {
		parsed = append(parsed, k+"="+v)
	}
	return parsed
}
