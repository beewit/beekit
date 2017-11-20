package utils

import (
	"os/exec"
)

type Command struct {
	name string
	args []string
}

func Open(uri string) error {
	c := Command{name: "xdg-open", args: []string{uri}}
	return Cmd(c)
}

func Cmd(c Command) error {
	cmd := exec.Command(c.name, c.args...)
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Output()
	return nil
}
