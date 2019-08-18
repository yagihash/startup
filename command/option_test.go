package command

import (
	"bufio"
	"bytes"
	"reflect"
	"sort"
	"testing"
)

func TestOptionWorkingDir(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "GiveValidDirName", input: "/tmp", want: "/tmp"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			command := NewCommand("echo", []string{},
				OptionWorkingDir(c.input),
			)
			got := command.Cmd.Dir
			if got != c.want {
				t.Errorf("\ngot:  %v\nwant: %v", got, c.want)
			}
		})
	}
}

func TestOptionStdout(t *testing.T) {
	t.Parallel()

	var buffer bytes.Buffer
	w := bufio.NewWriter(&buffer)

	want := "TestOptionStdout"
	command := NewCommand("echo", []string{"-n", want}, OptionStdout(w))
	command.Run()

	got := buffer.String()

	if got != want {
		t.Errorf("\ngot:  %v\nwant: %v", got, want)
	}
}

func TestOptionEnv(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input Env
		want  []string
	}{
		{name: "SetSingleEnv", input: Env{"foo": "bar"}, want: []string{"foo=bar"}},
		{name: "SetMultipleEnv", input: Env{"foo1": "bar1", "foo2": "bar2"}, want: []string{"foo1=bar1", "foo2=bar2"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			command := NewCommand("echo", []string{}, OptionEnv(c.input))
			got := command.Cmd.Env

			sort.SliceStable(got, func(i, j int) bool { return got[i] < got[j] })
			sort.SliceStable(c.want, func(i, j int) bool { return c.want[i] < c.want[j] })

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("\ngot:  %v\nwant: %v", got, c.want)
			}
		})
	}
}
