package gamemap

import (
	"math/rand"
	"time"
	"gogue"
	"gogue/camera"
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

	for x := 0; x < gameCamera.Width; x++ {
		for y := 0; y < gameCamera.Height; y++ {
			// Clear both our primary layers, so we don't get any strange artifacts from one layer or the other getting
			// cleared.
			for i := 0; i <= 2; i++ {

				gogue.PrintGlyph(x, y, gogue.EmptyGlyph, "", i)
			}
		}
	}
	//newCameraX, newCameraY = gameCamera.ToCameraCoordinates(newCameraX, newCameraY)
	gameCamera.MoveCamera(newCameraX, newCameraY, m.Width, m.Height)

	for x := 0; x < gameCamera.Width; x++ {
		for y := 0; y < gameCamera.Height; y++ {

			mapX, mapY := gameCamera.X + x, gameCamera.Y + y

			if m.Tiles[mapX][mapY].Visible {
				gogue.PrintGlyph(x, y, m.Tiles[mapX][mapY].Glyph, "", 0)
			} else if m.Tiles[mapX][mapY].Explored {
				gogue.PrintGlyph(x, y, m.Tiles[mapX][mapY].Glyph, "", 0, true)
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
