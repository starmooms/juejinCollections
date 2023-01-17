package tool

import (
	"fmt"
	"juejinCollections/tool/platform"
)

// https://github.com/tonoy30/openbrowser
func OpenBrowser(url string) {
	var args []string
	goos := platform.GetGOOS()

	switch goos {
	case platform.Win:
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

	platform.RunCmd(mainArg, otherArgs...)
}
