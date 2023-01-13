package tool

import (
	"fmt"
	"juejinCollections/tool/platform"
	"os/exec"
	"syscall"
)

// https://github.com/tonoy30/openbrowser
func OpenBrowser(url string) {
	var args []string
	goos := platform.GetGOOS()
	isWin := false

	switch goos {
	case platform.Win:
		isWin = true
		args = []string{"cmd", "/c", "start", url}
	case platform.Darwin:
		args = []string{"xdg-open", url}
	case platform.Linux:
		args = []string{"open", url}
	default:
		ShowErrMsg(fmt.Sprintf(`unkonw OS %s cannot open browser`, goos))
		return
	}

	mainArg := args[0]
	otherArgs := args[1:]

	cmd := exec.Command(mainArg, otherArgs...)

	// https://qa.1r1g.com/sf/ask/2975039931/
	// 解决window下会cmd闪烁
	if isWin {
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	}

	cmd.Run()
}
