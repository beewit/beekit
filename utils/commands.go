package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"golang.org/x/sys/windows/registry"
)

var commands = map[string]string{
	"windows": "cmd",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Exec(path string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	var cmd *exec.Cmd
	if run == "cmd" {
		cmd = exec.Command(run, path)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command(run, path)
	}
	return cmd.Start()
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

func CloseSpread() error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	var cmd *exec.Cmd
	if run == "cmd" {
		cmd = exec.Command("taskkill.exe", "/f", "/t", "/im", "spread.exe")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("sh", "-c", "killall spread")
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Output()
	return nil
}

func Close(threadName string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	var cmd *exec.Cmd
	if run == "cmd" {
		cmd = exec.Command("taskkill.exe", "/f", "/t", "/im", fmt.Sprintf("%s.exe", threadName))
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("killall %s", threadName))
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Output()
	return nil
}

func StartQQ() (err error) {
	path, err := QueryRegistry(`Software\\Tencent\\PlatForm_Type_List\\1`, "TypePath")
	if err != nil {
		return
	}
	return CallEXE(path)
}

func QueryRegistry(path, key string) (val string, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		k, err = registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
		if err != nil {
			return
		}
	}
	defer k.Close()
	val, _, err = k.GetStringValue(key)
	if err != nil {
		return
	}
	return
}

func CallEXE(strGameName string) (err error) {
	cmd := exec.Command(strGameName)
	err = cmd.Start()
	if err != nil {
		return
	}
	return
}
