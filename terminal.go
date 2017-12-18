package gogue

import (
	blt "bearlibterminal"
	"strconv"
)

const (
	// Not sure if converting these is going to prove useful or not
	// KEY just seems more natural than TK
	KEY_CLOSE = blt.TK_CLOSE
	KEY_RIGHT = blt.TK_RIGHT
	KEY_LEFT = blt.TK_LEFT
	KEY_UP = blt.TK_UP
	KEY_DOWN = blt.TK_DOWN
	KEY_L = blt.TK_L
	KEY_H = blt.TK_H
	KEY_K = blt.TK_K
	KEY_J = blt.TK_J
	KEY_Y = blt.TK_Y
	KEY_U = blt.TK_U
	KEY_B = blt.TK_B
	KEY_N = blt.TK_N
	KEY_X = blt.TK_X
	KEY_COMMA = blt.TK_COMMA
	KEY_I = blt.TK_I
	KEY_D = blt.TK_I
	KEY_ESCAPE = blt.TK_ESCAPE
)

func init() {

}

// InitConsole sets up a BearLibTerminal conosle window
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

// SetFont sets the font size and font to use.
// If this method is not called, the default font and size for BearLibTerminal is used
func SetFont(fontSize int, fontPath string) {
	// Next, setup the font config string
	consoleFontSize := "size=" + strconv.Itoa(fontSize)
	font := "font: " + fontPath + ", " + consoleFontSize

	blt.Set(font + ";")
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

// Print prints out a single character at the x, y coordinates provided, in the color provided,
// and on the layer provided. If no layer is provided, layer 0 is used.
func Print(x, y int, g Glyph, backgroundColor string, layer int, useExploredColor ...bool) {
	// Set the layer first. If not provided, this defaults to 0, the base layer in BearLibTerminal
	blt.Layer(layer)

	if backgroundColor != "" {
		// If a background color was provided, set that
		// Background color can only be applied to the lowest layer
		blt.BkColor(blt.ColorFromName(backgroundColor))
	}

	// Next, set the color to print in
	if len(useExploredColor) > 0 {
		blt.Color(blt.ColorFromName(g.ExploredColor()))
	} else {
		blt.Color(blt.ColorFromName(g.Color()))
	}

	// Finally, print the character at the provided coordinates
	blt.Print(x, y, string(g.Char()))
}

// ReadInput reads the next input event from the Input queue.
// If the queue is empty, it will wait for an event in a blocking manner
// if blocking=false is provided, the blocking behavior will not occur (if not input is found, execution continues,
// rather than blocking execution until input is provided)
func ReadInput(nonBlocking bool) int {
	if nonBlocking {
		// If non blocking reads are desired (say for a realtime game), check if there is an input in the Input queue
		// If there is, return it, otherwise, continue execution
		inputReady := blt.HasInput()

		if inputReady {
			return blt.Read()
		} else {
			return -1
		}
	}

	// Default behavior is to use blocking reads
	return blt.Read()
}
