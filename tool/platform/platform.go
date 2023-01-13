package platform

import "runtime"

const (
	Win    = "windows"
	Darwin = "darwin"
	Linux  = "linux"
)

func GetGOOS() string {
	return runtime.GOOS
}

func IsWin() bool {
	return runtime.GOOS == "windows"
}

func isDarwin() bool {
	return runtime.GOOS == "darwin"
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}
