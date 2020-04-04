package gamemap

import (
	"github.com/gogue-framework/gogue"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GenerateArena takes a Map object, and creates a giant room, ringed with walls. This is a very simple type of
// map that contains no features other than the walls.
func GenerateArena(surface *Map, wallGlyph, floorGlyph gogue.Glyph) {
	// Generates a large, empty room, with walls ringing the outside edges
	for x := 0; x <= surface.Width; x++ {
		for y := 0; y <= surface.Height; y++ {
			if x == 0 || x == surface.Width-1 || y == 0 || y == surface.Height-1 {
				surface.Tiles[x][y] = &Tile{Glyph: wallGlyph, Blocked: true, BlocksSight: true, Visited: false, Explored: false, Visible: false, X: x, Y: y}
			} else {
				surface.Tiles[x][y] = &Tile{Glyph: floorGlyph, Blocked: false, BlocksSight: false, Visited: false, Explored: false, Visible: false, X: x, Y: y}

				// Add the tile to the list of floor tiles that have been created. This will be used to add items,
				// monsters, the player, etc
				surface.FloorTiles = append(surface.FloorTiles, surface.Tiles[x][y])
			}
		}
	}
}

func TestTile_IsWall(t *testing.T) {
	glyph := gogue.NewGlyph("#", "white", "white")
	noises := make(map[int]float64)
	tile := Tile{glyph, true, true, true, false, false, false, 1, 1, noises}

	assert.True(t, tile.IsWall())

	tile.Blocked = false
	assert.False(t, tile.IsWall())
}

func TestMap_InitializeMap(t *testing.T) {
	gameMap := Map{Width: 100, Height: 100}

	gameMap.InitializeMap()

	assert.Equal(t, len(gameMap.Tiles), 101)
	assert.Equal(t, len(gameMap.Tiles[0]), 101)
}

func TestMap_IsBlocked(t *testing.T) {
	wallGlyph := gogue.NewGlyph("#", "white", "gray")
	floorGlyph := gogue.NewGlyph(".", "white", "gray")

	gameMap := Map{Width: 100, Height: 100}
	gameMap.InitializeMap()

	// Generate an arena style map
	GenerateArena(&gameMap, wallGlyph, floorGlyph)

	// If map generation went correctly, the Tile at position (0, 0) should be a wall
	topLeftCornerWall := gameMap.Tiles[0][0]
	assert.True(t, topLeftCornerWall.IsWall())
	assert.True(t, gameMap.IsBlocked(0, 0))

	floorTile := gameMap.Tiles[1][1]
	assert.False(t, floorTile.IsWall())
	assert.False(t, gameMap.IsBlocked(1, 1))
}

func TestMap_GetNeighbors(t *testing.T) {
	wallGlyph := gogue.NewGlyph("#", "white", "gray")
	floorGlyph := gogue.NewGlyph(".", "white", "gray")

	gameMap := Map{Width: 100, Height: 100}
	gameMap.InitializeMap()

	// Generate an arena style map
	GenerateArena(&gameMap, wallGlyph, floorGlyph)

	// Get the neighbors of the Tile at (1, 1). This should return eight Tiles, at [(0, 0), (0, 1), (0, 2), (1, 2),
	// (2, 2), (2, 1), (2, 0), (1, 0)]
	neighbors := gameMap.GetNeighbors(1, 1)
	assert.Equal(t, 8, len(neighbors))

	// Get the neighbors of the Tile at (0, 0). This should return three Tiles, at [(0, 1), (1, 1), (1, 0)]
	neighbors = gameMap.GetNeighbors(0, 0)
	assert.Equal(t, 3, len(neighbors))

	// Ensure that edge cases on the opposite end work as well
	neighbors = gameMap.GetNeighbors(99, 99)
	assert.Equal(t, 8, len(neighbors))

	neighbors = gameMap.GetNeighbors(100, 100)
	assert.Equal(t, 3, len(neighbors))

	// And make sure a random value is also correct
	neighbors = gameMap.GetNeighbors(10, 37)
	assert.Equal(t, 8, len(neighbors))
}
