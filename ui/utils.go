package ui

import (
	"fmt"
	"strings"
)

// SplitLines takes a text string, and splits into lines that do not exceed the character count represented by width
// This function will try to split on word boundaries
func SplitLines(text string, width int) []string {
	var lines []string

	line := ""

	wordList := strings.Fields(text)

	for _, word := range wordList {
		if len(line) < width {
			// If the line is still less than the total width, attempt to add the new word
			tempLine := ""
			if len(line) == 0 {
				tempLine = word
			} else {
				tempLine = fmt.Sprintf("%s %s", line, word)
			}

			// Check the length of the new line. If it is still less than width, set line to the temp value. If the
			// length exceeds width, append line to the lines list, and start a new line with the current word
			if len(tempLine) < width {
				line = tempLine
			} else {
				lines = append(lines, line)
				line = word
			}
		}
	}

	// Append whatever was left
	lines = append(lines, line)

	return lines
}
