package main

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

func main() {
	e("ssh-add", "-K")
}

func e(cmd string, option ...string) {
	err := exec.Command(cmd, option...).Run()
	if err != nil {
		notify.Notify(
			AppName,
			TitleError,
			fmt.Sprintf("%v", err),
			"",
		)
	}
}
