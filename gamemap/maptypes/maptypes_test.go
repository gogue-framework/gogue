package maptypes

import (
	"github.com/gogue-framework/gogue/gamemap"
	"github.com/gogue-framework/gogue/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	wallGlyph  ui.Glyph
	floorGlyph ui.Glyph
)

func TestGenerateArena(t *testing.T) {
	wallGlyph = ui.NewGlyph("#", "white", "gray")
	floorGlyph = ui.NewGlyph(".", "white", "gray")

	gameMap := &gamemap.GameMap{Width: 50, Height: 50}
	gameMap.InitializeMap()

	GenerateArena(gameMap, wallGlyph, floorGlyph)

	assert.Equal(t, gameMap.Tiles[0][0].Glyph, wallGlyph)
	assert.Equal(t, gameMap.Tiles[2][2].Glyph, floorGlyph)
}

func TestGenerateCavern(t *testing.T) {
	wallGlyph = ui.NewGlyph("#", "white", "gray")
	floorGlyph = ui.NewGlyph(".", "white", "gray")

	gameMap := &gamemap.GameMap{Width: 50, Height: 50}
	gameMap.InitializeMap()

	GenerateCavern(gameMap, wallGlyph, floorGlyph, 5)

	// Kind of hard to test procedural code, so we'll just check some basic parameters to ensure that a map was
	// actually generated
	assert.Greater(t, len(gameMap.FloorTiles), 0)

	// For a cavern, we seal up all the edges of the map, so when x == 0, or y == 0, or x == width, or y == width,
	// there should never be a floor
	for x := 0; x <= gameMap.Width; x++ {
		for y := 0; y <= gameMap.Height; y++ {
			if x == 0 || x == gameMap.Width-1 || y == 0 || y == gameMap.Height-1 {
				assert.Equal(t, gameMap.Tiles[x][y].Glyph, wallGlyph)
			}
		}
	}
}
