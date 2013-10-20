package foreground

import (
	"syscall"
	"unsafe"
)

var (
	user                     = syscall.NewLazyDLL("user32.dll")
	procGetForegroundWindow  = user.NewProc("GetForegroundWindow")
	procGetWindowText        = user.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = user.NewProc("GetWindowTextLengthW")
)

func GetForegroundWindow() string {
	window, _, _ := procGetForegroundWindow.Call()
	textLength, _, _ := procGetWindowTextLengthW.Call(uintptr(window))
	textLength += 1
	titleBuffer := make([]uint16, textLength)
	procGetWindowText.Call(uintptr(window), uintptr(unsafe.Pointer(&titleBuffer[0])), textLength)
	return syscall.UTF16ToString(titleBuffer)
}
