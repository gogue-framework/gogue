package gogue

import (
	blt "bearlibterminal"
	"strconv"
)

func init() {

}

// initConsole sets up a BearLibTerminal conosle window
// The X and Y dimensions, title, and a fullscreen flag can all be priovided
// The console window is not actually rendered to the screen until Refresh is called
func InitConsole(windowSizeX, windowSizeY int, title string, fullScreen bool) {
	blt.Open()

	// BearLibTerminal uses configuration strings to set itself up, so we need to build these strings here
	// First set up the string for window properties (size and title)
	size := "size=" + strconv.Itoa(windowSizeX) + "x" + strconv.Itoa(windowSizeY)
	consoleTitle := "title='" + title + "'"
	window := "window: " + size + "," + consoleTitle

	if fullScreen {
		consoleFullScreen := "fullscreen=true"
		window += "," + consoleFullScreen
	}

	// Now, put it all together
	blt.Set(window + "; ")
	blt.Clear()
}

// refresh calls blt.Refresh on the current console window
func Refresh() {
	blt.Refresh()
}

// closeConsole calls blt.Close on the current console window
func CloseConsole() {
	blt.Close()
}
