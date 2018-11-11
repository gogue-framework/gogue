package ui

import (
	"github.com/jcerise/gogue"
	"sort"
)

const (
	ORDLOWERSTART = 97
)

type MenuList struct {
	Options map[int]string
	Inputs map[rune]int
	keys []int
	Paginated bool
}

func (ml *MenuList) Create(options map[int]string) {
	ml.Options = options

	ml.Inputs = make(map[rune]int)

	ordLower := ORDLOWERSTART

	for identifier := range options {
		if ordLower <= 122 {
			ml.Inputs[rune(ordLower)] = identifier
			ml.keys = append(ml.keys, ordLower)
			ordLower += 1
		}
	}
}

func (ml *MenuList) Print(height, width int) {
	lineStart := 3

	// Sort the index slice, this will allow for guaranteed printing order of the two data maps
	sort.Ints(ml.keys)

	for _, keyRune := range ml.keys {
		input := ml.Inputs[rune(keyRune)]
		gogue.PrintText(3, lineStart, "(" + string(keyRune) + ")" + ml.Options[input], "", "", 0)
		lineStart += 1
	}
}
