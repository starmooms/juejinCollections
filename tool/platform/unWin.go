//go:build !windows

package platform

import (
	"os/exec"
)

func RunCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}
