package gogue

// Glyph represents a single character that can be drawn to the terminal. Char is the physical character representation,
// Color is the display color, and ExploredColor is the color it is shown when not in direct view (usually a darker
// shade of its Color)
type Glyph interface {
	Char() string
	Color() string
	ExploredColor() string
}

type glyph struct {
	char            string
	color           string
	unexploredColor string
}

func (g glyph) Char() string {
	return g.char
}

func (g glyph) Color() string {
	return g.color
}

func (g glyph) ExploredColor() string {
	return g.unexploredColor
}

// NewGlyph returns a new Glyph object. If no ExploredColor is provided, it is set to gray.
func NewGlyph(char string, color, exploredColor string) Glyph {
	if exploredColor == "" {
		exploredColor = "gray"
	}
	return glyph{char, color, exploredColor}
}

// EmptyGlyph is a glyph with an empty string for its Char. This empty glyph is useful for erasing other glyphs, by
// replacing them with an empty glyph (which will show as nothing on the terminal).
var EmptyGlyph = NewGlyph(" ", "white", "")
