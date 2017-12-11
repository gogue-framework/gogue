package gogue

type Glyph interface {
	Char() rune
	Color() string
}

type glyph struct {
	char rune
	color string
}

func (g glyph) Char() rune {
	return g.char
}

func (g glyph) Color() string {
	return g.color
}

func NewGlyph(char rune, color string) Glyph {
	return glyph{char, color}
}

var EmptyGlyph = NewGlyph(' ', "white")
