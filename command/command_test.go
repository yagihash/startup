package command

import (
	"bufio"
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input struct {
			cmd    string
			option []string
		}
		want string
	}{
		{
			name: "EchoFoo",
			input: struct {
				cmd    string
				option []string
			}{
				cmd:    "echo",
				option: []string{"-n", "foo"},
			},
			want: "foo",
		},
		{
			name: "NoSuchCommand",
			input: struct {
				cmd    string
				option []string
			}{
				cmd:    "nosuchcommand",
				option: []string{},
			},
			want: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var buffer bytes.Buffer
			w := bufio.NewWriter(&buffer)
			command := New(c.input.cmd, c.input.option, OptionStdout(w))
			command.Run()

			got := buffer.String()
			if got != c.want {
				t.Errorf("\ngot:  %v\nwant: %v", got, c.want)
			}
		})
	}
}
