package infrastructure

import (
	"fmt"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"log"
	"os"
)

const (
	netWmName = "_NET_WM_NAME" // Store window titles (_NET_WM_NAME atom) as property of a window
)

type LinuxWindow struct{}

func (w *LinuxWindow) GetWindow() string {
	return "Linux"
}

func (w *LinuxWindow) GetWindowsData() []WindowMap {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Error connecting to X Sever: %s", err)
	}
	defer X.Close()

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)
	root := screen.Root

	atomName := netWmName
	atomCookie, err := xproto.InternAtom(X, true, uint16(len(atomName)), atomName).Reply()
	if err != nil {
		log.Fatalf("Error getting atom _NET_WM_NAME: %s", err)
	}
	netWmNameAtom := atomCookie.Atom

	children, err := xproto.QueryTree(X, root).Reply()
	if err != nil {
		log.Fatalf("Error gettings list windows: %s", err)
	}

	var windowMap []WindowMap
	for _, win := range children.Children {
		geometry, err := xproto.GetGeometry(X, xproto.Drawable(win)).Reply()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting geometry %d: %s\n", win, err)
			continue
		}

		title, err := xproto.GetProperty(X, false, win, netWmNameAtom, xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting window title %d: %s\n", win, err)
			continue
		}

		windowTitle := string(title.Value)

		windowMap = append(windowMap, WindowMap{
			title:  windowTitle,
			x:      geometry.X,
			y:      geometry.Y,
			width:  geometry.Width,
			height: geometry.Height,
		})
	}
	return windowMap
}
