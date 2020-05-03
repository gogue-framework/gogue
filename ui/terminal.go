package ui

import (
	blt "github.com/gogue-framework/bearlibterminalgo"
	"strconv"
)

// Not sure if converting these is going to prove useful or not
// KEY just seems more natural than TK
const (
	KeyClose  = blt.TK_CLOSE
	KeyRight  = blt.TK_RIGHT
	KeyLeft   = blt.TK_LEFT
	KeyUp     = blt.TK_UP
	KeyDown   = blt.TK_DOWN
	KeyA      = blt.TK_A
	KeyB      = blt.TK_B
	KeyC      = blt.TK_C
	KeyD      = blt.TK_D
	KeyE      = blt.TK_E
	KeyF      = blt.TK_F
	KeyG      = blt.TK_G
	KeyH      = blt.TK_H
	KeyI      = blt.TK_I
	KeyJ      = blt.TK_J
	KeyK      = blt.TK_K
	KeyL      = blt.TK_L
	KeyM      = blt.TK_M
	KeyN      = blt.TK_N
	KeyO      = blt.TK_O
	KeyP      = blt.TK_P
	KeyQ      = blt.TK_Q
	KeyR      = blt.TK_R
	KeyS      = blt.TK_S
	KeyT      = blt.TK_T
	KeyU      = blt.TK_U
	KeyV      = blt.TK_V
	KeyW      = blt.TK_W
	KeyX      = blt.TK_X
	KeyY      = blt.TK_Y
	KeyZ      = blt.TK_Z
	KeyComman = blt.TK_COMMA
	KeyEscape = blt.TK_ESCAPE
	KeyEnter  = blt.TK_ENTER
)

var (
	// RuneKeyMapping maps keypresses to the respective Go Rune representations. This can be useful for checking
	// the rune associated with a keypress.
	RuneKeyMapping = map[int]rune{
		blt.TK_A: 'a',
		blt.TK_B: 'b',
		blt.TK_C: 'c',
		blt.TK_D: 'd',
		blt.TK_E: 'e',
		blt.TK_F: 'f',
		blt.TK_G: 'g',
		blt.TK_H: 'h',
		blt.TK_I: 'i',
		blt.TK_J: 'j',
		blt.TK_K: 'k',
		blt.TK_L: 'l',
		blt.TK_M: 'm',
		blt.TK_N: 'n',
		blt.TK_O: 'o',
		blt.TK_P: 'p',
		blt.TK_Q: 'q',
		blt.TK_R: 'r',
		blt.TK_S: 's',
		blt.TK_T: 't',
		blt.TK_U: 'u',
		blt.TK_V: 'v',
		blt.TK_W: 'w',
		blt.TK_X: 'x',
		blt.TK_Y: 'y',
		blt.TK_Z: 'z',
	}
	compositionMode = 0
)

func init() {

}

// InitConsole sets up a BearLibTerminal console window
// The X and Y dimensions, title, and a fullscreen flag can all be provided
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
	blt.Composition(compositionMode)
	blt.Clear()
}

// SetPrimaryFont sets the font size and font to use.
// If this method is not called, the default font and size for BearLibTerminal is used
func SetPrimaryFont(fontSize int, fontPath string) {
	// Next, setup the font config string
	consoleFontSize := "size=" + strconv.Itoa(fontSize)
	font := "font: " + fontPath + ", " + consoleFontSize

	blt.Set(font + ";")
	blt.Clear()
}

// AddFont adds a named font to the console, that can be used when printing text, as an alternative to the
// primary font.
func AddFont(name, path string, fontSize int) {
	consoleFontSize := "size=" + strconv.Itoa(fontSize)
	font := name + " font: " + path + ", " + consoleFontSize

	blt.Set(font + ";")
	blt.Clear()
}

// SetCompositionMode sets the composition mode for drawing glyphs to the terminal. 0 is no composition, meaning the
// entire cell will be replaced by the character drawn. 1 means the character drawn will be composed onto any lower
// level characters. This is used to set the composition mode each time a character is printed.
func SetCompositionMode(mode int) {
	if mode == blt.TK_OFF || mode == blt.TK_ON {
		compositionMode = mode
	}
}

// SetGlobalComposition sets the global terminal state of composition. Off means that each character printed to a cell
// will, graphically, replace all other characters in that cell. Composition on means that each character printed to a
// cell will be composed on top or beneath (depending on layer) any other characters present in that cell.
func SetGlobalComposition(mode int) {
	if mode == blt.TK_OFF || mode == blt.TK_ON {
		blt.Composition(mode)
	}
}

// Refresh calls blt.Refresh on the current console window
func Refresh() {
	blt.Refresh()
}

// CloseConsole calls blt.Close on the current console window
func CloseConsole() {
	blt.Close()
}

// ClearArea clears a (rectangular) area of the terminal, starting at (x, y), and containing the area to
// (x + width, y + height).
func ClearArea(x, y, width, height, layer int) {
	blt.Layer(layer)
	blt.ClearArea(x, y, width, height)
}

// ClearWindow is just a wrapper call around ClearArea that clears the entire terminal window, from (0,0) to
// (WindowWidth, WindowHeight)
func ClearWindow(windowWidth, windowHeight, layer int) {
	ClearArea(0, 0, windowWidth, windowHeight, layer)
}

// PrintGlyph prints out a single character at the x, y coordinates provided, in the color provided,
// and on the layer provided. If no layer is provided, layer 0 is used.
func PrintGlyph(x, y int, g Glyph, backgroundColor string, layer int, useExploredColor ...bool) {
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

// PrintText will print a string of text, starting at the (X, Y) coords provided, using the color/background color
// provided, on the layer provided.
func PrintText(x, y int, text, color, backgroundColor string, layer int, splitWidth int) {
	// Set the layer first. If not provided, this defaults to 0, the base layer in BearLibTerminal
	blt.Layer(layer)

	if backgroundColor != "" {
		// If a background color was provided, set that
		// Background color can only be applied to the lowest layer
		blt.BkColor(blt.ColorFromName(backgroundColor))
	}

	if color != "" {
		// If a color was set, use that, otherwise, default to white
		blt.Color(blt.ColorFromName(color))
	} else {
		blt.Color(blt.ColorFromName("white"))
	}

	// Finally, print the character at the provided coordinates
	if splitWidth > 0 {
		// If a split width is specified, break the string into lines (on word boundaries), and print each line one cell
		// below the previous
		lines := SplitLines(text, splitWidth)

		lineY := y

		for _, line := range lines {
			blt.Print(x, y, line)
			lineY++
		}
	} else {
		blt.Print(x, y, text)
	}
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
		}

		return -1
	}

	// Default behavior is to use blocking reads
	return blt.Read()
}
