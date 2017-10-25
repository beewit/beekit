package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
)

var commands = map[string]string{
	"windows": "cmd",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	var cmd *exec.Cmd
	if run == "cmd" {
		cmd = exec.Command(run, "/c", "start", uri)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command(run, uri)
	}
	return cmd.Start()
}

func CloseChrome() error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	var cmd *exec.Cmd
	if run == "cmd" {
		cmd = exec.Command("taskkill.exe", "/f", "/t", "/im", "chromedriver.exe")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("sh", "-c", "killall chromedriver")
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Output()
	return nil
}
