package infrastructure

type WindowsWindow struct{}

func (w *WindowsWindow) GetWindow() string {
	return "Windows"
}

func (w *WindowsWindow) GetWindowsData() []WindowMap {
	return []WindowMap{}
}
