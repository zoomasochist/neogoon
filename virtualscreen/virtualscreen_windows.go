package virtualscreen

import "syscall"

var (
	user32           = syscall.NewLazyDLL("User32.dll")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

func Resolution() (int, int) {
	x, _, _ := getSystemMetrics.Call(uintptr(0))
	y, _, _ := getSystemMetrics.Call(uintptr(1))

	return int(x), int(y)
}
