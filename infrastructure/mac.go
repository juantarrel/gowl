package infrastructure

type MacWindow struct{}

func (w *MacWindow) GetWindow() string {
	return "MacOS"
}

func (w *MacWindow) GetWindowsData() []WindowMap {
	return []WindowMap{}
}
