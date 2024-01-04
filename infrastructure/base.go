package infrastructure

import "runtime"

type WindowMap struct {
	title  string
	x      int16
	y      int16
	width  uint16
	height uint16
}

type WindowGetter interface {
	GetWindow() string
	GetWindowsData() []WindowMap
}

func NewWindowGetter() (WindowGetter, error) {
	switch os := runtime.GOOS; os {
	case "windows":
		return new(WindowsWindow), nil
	case "linux":
		return new(LinuxWindow), nil
	case "darwin":
		return new(MacWindow), nil
	default:
		return nil, nil
	}
}
