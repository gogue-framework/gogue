package maptypes

import (
	"github.com/gogue-framework/gogue/gamemap"
	"github.com/gogue-framework/gogue/ui"
)

// GenerateArena takes a gamemap.Map object, and creates a giant room, ringed with walls. This is a very simple type of
// map that contains no features other than the walls.
func GenerateArena(surface *gamemap.GameMap, wallGlyph, floorGlyph ui.Glyph) {
	// Generates a large, empty room, with walls ringing the outside edges
	for x := 0; x <= surface.Width; x++ {
		for y := 0; y <= surface.Height; y++ {
			if x == 0 || x == surface.Width-1 || y == 0 || y == surface.Height-1 {
				surface.Tiles[x][y] = &gamemap.Tile{Glyph: wallGlyph, Blocked: true, BlocksSight: true, Visited: false, Explored: false, Visible: false, X: x, Y: y}
			} else {
				surface.Tiles[x][y] = &gamemap.Tile{Glyph: floorGlyph, Blocked: false, BlocksSight: false, Visited: false, Explored: false, Visible: false, X: x, Y: y}

				// Add the tile to the list of floor tiles that have been created. This will be used to add items,
				// monsters, the player, etc
				surface.FloorTiles = append(surface.FloorTiles, surface.Tiles[x][y])
			}
		}
	}
}
