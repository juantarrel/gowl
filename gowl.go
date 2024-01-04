package gowl

import (
	"github.com/juantarrel/gowl/infrastructure"
)

func CheckOS() (string, error) {
	windowGetter, err := infrastructure.NewWindowGetter()

	if err != nil {
		return "OS not supported", nil
	}

	return windowGetter.GetWindow(), nil
}

func GetWindowsData() []infrastructure.WindowMap {
	wGetter, _ := infrastructure.NewWindowGetter()
	a := wGetter.GetWindowsData()
	return a
}
