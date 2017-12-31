package gogue

type Glyph interface {
	Char() string
	Color() string
	ExploredColor() string
}

type glyph struct {
	char string
	color string
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

func NewGlyph(char string, color, exploredColor string) Glyph {
	if exploredColor == "" {
		exploredColor = "gray"
	}
	return glyph{char, color, exploredColor}
}

var EmptyGlyph = NewGlyph(" ", "white", "")
