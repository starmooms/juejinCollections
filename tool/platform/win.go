//go:build windows

package platform

import (
	"os/exec"
	"syscall"
)

func RunCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	// https://qa.1r1g.com/sf/ask/2975039931/
	// 解决window下会cmd闪烁
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	return cmd.Run()
}
