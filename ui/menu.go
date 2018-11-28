package ui

import (
	"github.com/jcerise/gogue"
	"log"
	"reflect"
	"sort"
)

const (
	ORDLOWERSTART = 97
	MAXORD = 122
)

type MenuList struct {
	Options map[int]string
	Inputs map[rune]int
	keys []int
	Paginated bool
	highestOrd int
}

// Create builds a new MenuList, given a mapping of options
func (ml *MenuList) Create(options map[int]string) {
	ml.Options = options

	ml.Inputs = make(map[rune]int)

	ordLower := ORDLOWERSTART

	for identifier := range options {
		if ordLower <= MAXORD {
			ml.Inputs[rune(ordLower)] = identifier
			ml.keys = append(ml.keys, ordLower)
			ordLower += 1
			ml.highestOrd = ordLower
		}
	}
}

// Update takes a list of options, compares it to the existing list of options, and updates the menu if the new list
// is different. This allows for updating a menu without creating a new one (which can mess with item ordering).
// Returns true if the options were updated, false otherwise
func (ml *MenuList) Update(options map[int]string) bool {
	// First things first, see if the updated options is the same as the original. If it is, do nothing
	eq := reflect.DeepEqual(options, ml.Options)

	if eq {
		return false
	} else {
		// The two are not equal. We need to rectify the items in the updated list with the original. This is a two step
		// process. First, update the inputs. For each input, if the identifier still exists in the new list, do nothing
		// If it does not exist, we'll clear out the identifier value. Next, we'll iterate over the new list, and
		// for each value that is not mapped to a key, we'll map it to one. This can be either an existing key, or
		// a new key that will be added
		for key, identifier := range ml.Inputs {
			// Check if the keys identifier is still in the updated list
			if _, ok := options[identifier]; !ok {
				// The identifier is no longer present in the updated list, so remove it from the key mapping
				ml.Inputs[key] = -1
			}
		}

		// Now, walk through the updated list, and assign new items to any empty keys. This will fill in any gaps in the
		// menu.
		for identifier := range options {
			// Loop through the inputs, looking for nulled slots (-1 for the identifier), also checking that the current
			// item is not already in the inputs
			placed := false
			if _, ok := ml.Options[identifier]; !ok {
				// This item is not currently in the existing list, and needs a spot in the input map
				for key, keyIdentifier := range ml.Inputs {
					if keyIdentifier == -1 {
						ml.Inputs[key] = identifier
						placed = true
					}
				}

				if !placed {
					// There was no free spot in an existing key, so we'll add a new one, based on the highest ord rune
					// used previously
					if ml.highestOrd <= MAXORD {
						ml.Inputs[rune(ml.highestOrd)] = identifier
						ml.keys = append(ml.keys, ml.highestOrd)
						ml.highestOrd += 1
						placed = true
					}
				}
			}
			// At this point, the item should have been placed. If it has not been, something has gone wrong
			if !placed {
				log.Fatal("Failed to place an item in the menu. Max Length likely exceeded.")
			}
		}

		// Finally, now that all the new items have been placed, and items that need to be removed have been removed,
		// set the menu options to the updated options list
		ml.Options = options

		return true
	}
}

// Print displays the options for the MenuList, sorted by the rune chosen to represent it. yOffset is the number of rows
// to skip before printing, and xOffset, similarly, is the number of columns to skip before printing
func (ml *MenuList) Print(height, width, xOffset, yOffset int) {
	lineStart := yOffset

	// Sort the index slice, this will allow for guaranteed printing order of the two data maps
	sort.Ints(ml.keys)

	for _, keyRune := range ml.keys {
		input := ml.Inputs[rune(keyRune)]
		gogue.PrintText(xOffset, lineStart, "(" + string(keyRune) + ")" + ml.Options[input], "", "", 0)
		lineStart += 1
	}
}
