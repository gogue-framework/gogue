package gamemap

import (
	"github.com/gogue-framework/gogue"
	"github.com/gogue-framework/gogue/camera"
	"math/rand"
	"time"
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

type CoordinatePair struct {
	X int
	Y int
}

type Tile struct {
	Glyph       gogue.Glyph
	Blocked     bool
	BlocksSight bool
	BlocksSound bool
	Visited     bool
	Explored    bool
	Visible     bool
	X           int
	Y           int
	Noises      map[int]float64
}

// IsWall determines if a tile acts as a wall or not. A wall blocks sight and movement. If both of these criteria are
// true, the tile is said to be a wall.
func (t *Tile) IsWall() bool {
	if t.BlocksSight && t.Blocked {
		return true
	} else {
		return false
	}
}

type Map struct {
	Width      int
	Height     int
	Tiles      [][]*Tile
	FloorTiles []*Tile
}

// InitializeMap sets up a GameMap for use. It sets the Tiles property of the GameMap to a 2D array of Tile objects,
// with a width and height matching those set for the GameMap. It also initializes a random seed to use for map
// generation
func (m *Map) InitializeMap() {
	// Initialize a two dimensional array that will represent the current game map (of dimensions Width x Height)
	m.Tiles = make([][]*Tile, m.Width+1)
	for i := range m.Tiles {
		m.Tiles[i] = make([]*Tile, m.Height+1)
	}

	// Set a seed for procedural generation
	rand.Seed(time.Now().UTC().UnixNano())
}

// Render draws a GameMap to the terminal, within a Camera viewport. It will only draw tiles from the GameMap that
// visible to the player, and within the viewport of the Camera. If a Tile does not meet these criteria, it will not be
// drawn. If a Tile is within the viewport of the Camera, but is outside the players FOV, and has been explored, it will
// be drawn using the Tile.Glyph exploredColor.
func (m *Map) Render(gameCamera *camera.GameCamera, newCameraX, newCameraY int) {

	gameCamera.MoveCamera(newCameraX, newCameraY, m.Width, m.Height)

	for x := 0; x < gameCamera.Width; x++ {
		for y := 0; y < gameCamera.Height; y++ {

			mapX, mapY := gameCamera.X+x, gameCamera.Y+y

			if mapX < 0 {
				mapX = 0
			}

			if mapY < 0 {
				mapY = 0
			}

			tile := m.Tiles[mapX][mapY]
			camX, camY := gameCamera.ToCameraCoordinates(mapX, mapY)

			// Print the tile, if it meets the following criteria:
			// 1. Its visible or explored
			// 2. It hasn't been printed yet. This will prevent over printing due to camera conversion
			if tile.Visible {
				gogue.PrintGlyph(camX, camY, tile.Glyph, "", 0)
			} else if tile.Explored {
				gogue.PrintGlyph(camX, camY, tile.Glyph, "", 0, true)
			}
		}
	}
}

// IsBlocked returns true if the Tile in the GameMap has its blocked property set to true. False otherwise.
func (m *Map) IsBlocked(x, y int) bool {
	// Check to see if the provided coordinates contain a blocked tile
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}

func (m *Map) BlocksNoises(x, y int) bool {
	// Check to see if the provided coordinates contain a tile that blocks noises
	if m.Tiles[x][y].BlocksSound {
		return true
	} else {
		return false
	}
}

// GetNeighbors will return a list of tiles that are directly next to the given coordinates. It can optionally exclude
// blocked tiles
func (m *Map) GetNeighbors(x, y int) []*Tile {
	neighbors := []*Tile{}
	sourceTile := m.Tiles[x][y]

	nX := 0
	nY := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			// Make sure the neighbor we're checking is within the bounds of the map
			if x+i < 0 || x+i > m.Width {
				continue
			} else {
				nX = x + i
			}

			if y+j < 0 || y+j > m.Height {
				continue
			} else {
				nY = y + j
			}

			// Exclude the source Tile
			if m.Tiles[nX][nY] != sourceTile {
				neighbors = append(neighbors, m.Tiles[nX][nY])
			}
		}
	}

	return neighbors
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

func (m *Map) HasNoises(x, y int) bool {
	if len(m.Tiles[x][y].Noises) > 0 {
		return true
	} else {
		return false
	}
}

func (m *Map) GetAdjacentNoisesForEntity(entity, x, y int) map[*Tile]float64 {
	// Get a list of the neighboring tiles for the location
	tiles := m.GetNeighbors(x, y)

	noisyTiles := make(map[*Tile]float64)

	for _, tile := range tiles {
		for noiseEntity, noise := range m.Tiles[x][y].Noises {
			if noiseEntity == entity {
				noisyTiles[tile] = noise
			}
		}
	}

	return noisyTiles
}
