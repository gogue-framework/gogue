package ui

import (
	"github.com/jcerise/gogue"
)

const (
	ORDLOWERSTART = 97
)

type MenuList struct {
	Options map[int]string
	Inputs map[rune]int
	Paginated bool
}

func (ml *MenuList) Create(options map[int]string) {
	ml.Options = options

	ml.Inputs = make(map[rune]int)

	ordLower := ORDLOWERSTART

	for identifier := range options {
		if ordLower <= 122 {
			ml.Inputs[rune(ordLower)] = identifier
			ordLower += 1
		}
	}
}

func (ml *MenuList) Print(height, width int) {
	lineStart := 3
	for keyRune, keyIndex := range ml.Inputs {
		gogue.PrintText(3, lineStart, "(" + string(keyRune) + ")" + ml.Options[keyIndex], "", "", 0)
		lineStart += 1
	}
}
