package controller

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procFindWindowW = user32.NewProc("FindWindowW")
)

// WinAPIHandler gets the process's handle using Windows API.
type WinAPIHandler struct{}

// Handle returns the handle to the process with the given class name.
func (*WinAPIHandler) Handle(className string) (HWND, error) {
	classPtr, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}

	hWnd := FindWindowW(classPtr, nil)
	if hWnd == 0 {
		return 0, errors.New("Window not found")
	}

	return hWnd, nil
}

// FindWindowW wraps the call to the user32.dll FindWindowW.
// It's taken (rather than imported) from github.com/AllenDang/w32 to prevent gcc compilation.
func FindWindowW(className, windowName *uint16) HWND {
	ret, _, _ := procFindWindowW.Call(
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)))

	return HWND(ret)
}
