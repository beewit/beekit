// +build !linux

package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

const (
	WINDOWS = "windows"
	LINUX   = "linux"
)

type Command struct {
	name string
	args []string
}

func Open(uri string) error {
	c := map[string]Command{}
	c[WINDOWS] = Command{name: "cmd", args: []string{"/c", "start", uri}}
	c[LINUX] = Command{name: "xdg-open", args: []string{uri}}
	return Cmd(c)
}

func CloseChrome() error {
	return Close("chromedriver")
}

func CloseSpread() error {
	return Close("spread")
}

func Close(threadName string) error {
	c := map[string]Command{}
	c[WINDOWS] = Command{name: "taskkill.exe", args: []string{"/f", "/t", "/im", fmt.Sprintf("%s.exe", threadName)}}
	c[LINUX] = Command{name: "sh", args: []string{"-c", "killall", threadName}}
	return Cmd(c)
}

func Cmd(c map[string]Command) error {
	goos := runtime.GOOS
	cmd := exec.Command(c[goos].name, c[goos].args...)
	if goos == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
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
