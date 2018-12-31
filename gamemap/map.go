package gamemap

import (
	"math/rand"
	"time"
	"github.com/jcerise/gogue"
	"github.com/jcerise/gogue/camera"
)

type BySize [][]*Tile

func (s BySize) Len() int {
	return len(s)
}

func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s BySize) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

type Tile struct {
	Glyph	gogue.Glyph
	Blocked      bool
	BlocksSight bool
	Visited      bool
	Explored     bool
	Visible      bool
	X            int
	Y            int
	Updated 	 bool
}

func (t *Tile) IsWall() bool {
	if t.BlocksSight && t.Blocked {
		return true
	} else {
		return false
	}
}

type Map struct {
	Width  int
	Height int
	Tiles  [][]*Tile
	FloorTiles []*Tile
}

func (m *Map) InitializeMap() {
	// Initialize a two dimensional array that will represent the current game map (of dimensions Width x Height)
	m.Tiles = make([][]*Tile, m.Width)
	for i := range m.Tiles {
		m.Tiles[i] = make([]*Tile, m.Height)
	}

	// Set a seed for procedural generation
	rand.Seed(time.Now().UTC().UnixNano())
}

func (m *Map) Render(gameCamera *camera.GameCamera, newCameraX, newCameraY int) {

	//newCameraX, newCameraY = gameCamera.ToCameraCoordinates(newCameraX, newCameraY)
	gameCamera.MoveCamera(newCameraX, newCameraY, m.Width, m.Height)

	for x := 0; x < gameCamera.Width; x++ {
		for y := 0; y < gameCamera.Height; y++ {

			mapX, mapY := gameCamera.X + x, gameCamera.Y + y

			if mapX < 0 {
				mapX = 0
			}

			if mapY < 0 {
				mapY = 0
			}

			tile := m.Tiles[mapX][mapY]

			// Check if the tile has been updated. This means that its state has changed since the last time it was
			// rendered. If it has been, re-draw it. Otherwise, skip it.
			if tile.Updated {

				// Clear the tile first, and then redraw
				gogue.PrintGlyph(mapX, mapY, gogue.EmptyGlyph, "", i)

				if tile.Visible {
					gogue.PrintGlyph(x, y, tile.Glyph, "", 0)
				} else if tile.Explored {
					gogue.PrintGlyph(x, y, tile.Glyph, "", 0, true)
				}

				tile.Updated = false
			}
		}
	}
}

func (m *Map) IsBlocked(x, y int) bool {
	// Check to see if the provided coordinates contain a blocked tile
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}

func (m *Map) IsVisibleToPlayer(x, y int) bool {
	// Check to see if the given position on the map is visible to the player currently
	if m.Tiles[x][y].Visible {
		return true
	} else {
		return false
	}
}

func (m *Map) IsVisibleAndExplored(x, y int) bool {
	if m.Tiles[x][y].Visible && m.Tiles[x][y].Explored {
		return true
	} else {
		return false
	}
}
