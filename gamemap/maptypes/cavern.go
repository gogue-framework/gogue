package maptypes

import (
	"github.com/gogue-framework/gogue/gamemap"
	"github.com/gogue-framework/gogue/ui"
	"math/rand"
	"sort"
)

type bySize [][]*gamemap.Tile

func (s bySize) Len() int {
	return len(s)
}

func (s bySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bySize) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// GenerateCavern uses a cellular automata algorithm to create a fairly natural looking, 2D, cavern layout. It accepts
// Glyphs representing the walls and floor.
// The algorithm works in 6 steps. Step 1 fills the entire map with random wall and floor tiles, in roughly a 40/60 mix
// Step 2 decides if each tile should remain a wall or become floor, based on its neighbors. Step 3 repeats step 2 a
// number of times to smooth out the generated caverns. Step 4 seals up the edges of the map, so there are no paths off
// the edge of the map. Step 5 uses a flood fill algorithm to find the largest cavern, and finally, step 6, fills all
// smaller caverns. This algorithm is simple and does not connect unconnected caverns, and simply uses the largest
// cavern as the play area.
func GenerateCavern(surface *gamemap.GameMap, wallGlyph, floorGlyph ui.Glyph, smoothingPasses int, explored bool) {
	// Step 1: Fill the map space with a random assortment of walls and floors. This uses a roughly 40/60 ratio in favor
	// of floors, as I've found that to produce the nicest results.

	for x := 0; x < surface.Width; x++ {
		for y := 0; y < surface.Height; y++ {
			state := rand.Intn(100)
			if state < 30 {
				surface.Tiles[x][y] = &gamemap.Tile{Glyph: wallGlyph, Blocked: true, BlocksSight: true, Visited: explored, Explored: explored, Visible: explored, X: x, Y: y, Noises: make(map[int]float64)}
			} else {
				surface.Tiles[x][y] = &gamemap.Tile{Glyph: floorGlyph, Blocked: false, BlocksSight: false, Visited: explored, Explored: explored, Visible: explored, X: x, Y: y, Noises: make(map[int]float64)}
			}
		}
	}

	// Step 2: Decide what should remain as walls. If 5 or more of a tiles immediate (within 1 space) neighbors are
	// walls, then make that tile a wall. If 2 or less of the tiles next closest (2 spaces away) neighbors are walls,
	// then make that tile a wall. Any other scenario, and the tile will become (or stay) a floor tile.
	// Make several passes on this to help smooth out the walls of the cave.
	for i := 0; i < 1; i++ {
		for x := 0; x < surface.Width; x++ {
			for y := 0; y < surface.Height-1; y++ {
				wallOneAway := countWallsNStepsAway(surface, 1, x, y)

				wallTwoAway := countWallsNStepsAway(surface, 1, x, y)

				if wallOneAway >= 5 || wallTwoAway <= 2 {
					surface.Tiles[x][y].Blocked = true
					surface.Tiles[x][y].BlocksSight = true
					surface.Tiles[x][y].Glyph = wallGlyph
				} else {
					surface.Tiles[x][y].Blocked = false
					surface.Tiles[x][y].BlocksSight = false
					surface.Tiles[x][y].Glyph = floorGlyph
				}
			}
		}
	}

	// Step 3: Make a few more passes, smoothing further, and removing any small or single tile, unattached walls.
	for i := 0; i < smoothingPasses; i++ {
		for x := 0; x < surface.Width; x++ {
			for y := 0; y < surface.Height-1; y++ {
				wallOneAway := countWallsNStepsAway(surface, 1, x, y)

				if wallOneAway >= 5 {
					surface.Tiles[x][y].Blocked = true
					surface.Tiles[x][y].BlocksSight = true
					surface.Tiles[x][y].Glyph = wallGlyph
				} else {
					surface.Tiles[x][y].Blocked = false
					surface.Tiles[x][y].BlocksSight = false
					surface.Tiles[x][y].Glyph = floorGlyph
				}
			}
		}
	}

	// Step 4: Seal up the edges of the map, so the player, and the following flood fill passes, cannot go beyond the
	// intended game area
	for x := 0; x < surface.Width; x++ {
		for y := 0; y < surface.Height; y++ {
			if x == 0 || x == surface.Width-1 || y == 0 || y == surface.Height-1 {
				surface.Tiles[x][y].Blocked = true
				surface.Tiles[x][y].BlocksSight = true
				surface.Tiles[x][y].Glyph = wallGlyph
			}
		}
	}

	// Step 5: Flood fill. This will find each individual cavern in the cave system, and add them to a list. It will
	// then find the largest one, and will make that as the main play area. The smaller caverns will be filled in.
	// In the future, it might make sense to tunnel between caverns, and apply a few more smoothing passes, to make
	// larger, more realistic caverns.

	var cavern []*gamemap.Tile
	var totalCavernArea []*gamemap.Tile
	var caverns [][]*gamemap.Tile
	var tile *gamemap.Tile
	var node *gamemap.Tile

	for x := 0; x < surface.Width-1; x++ {
		for y := 0; y < surface.Height-1; y++ {
			tile = surface.Tiles[x][y]

			// If the current tile is a wall, or has already been visited, ignore it and move on
			if !tile.Visited && !tile.IsWall() {
				// This is a non-wall, unvisited tile
				cavern = append(cavern, surface.Tiles[x][y])

				for len(cavern) > 0 {
					// While the current node tile has valid neighbors, keep looking for more valid neighbors off of
					// each one
					node = cavern[len(cavern)-1]
					cavern = cavern[:len(cavern)-1]

					if !node.Visited && !node.IsWall() {
						// Mark the node as visited, and add it to the cavern area for this cavern
						node.Visited = true
						totalCavernArea = append(totalCavernArea, node)

						// Add the tile to the west, if valid
						if node.X-1 > 0 && !surface.Tiles[node.X-1][node.Y].IsWall() {
							cavern = append(cavern, surface.Tiles[node.X-1][node.Y])
						}

						// Add the tile to east, if valid
						if node.X+1 < surface.Width && !surface.Tiles[node.X+1][node.Y].IsWall() {
							cavern = append(cavern, surface.Tiles[node.X+1][node.Y])
						}

						// Add the tile to north, if valid
						if node.Y-1 > 0 && !surface.Tiles[node.X][node.Y-1].IsWall() {
							cavern = append(cavern, surface.Tiles[node.X][node.Y-1])
						}

						// Add the tile to south, if valid
						if node.Y+1 < surface.Height && !surface.Tiles[node.X][node.Y+1].IsWall() {
							cavern = append(cavern, surface.Tiles[node.X][node.Y+1])
						}
					}
				}

				// All non-wall tiles have been found for the current cavern, add it to the list, and start looking for
				// the next one
				caverns = append(caverns, totalCavernArea)
				totalCavernArea = nil
			} else {
				tile.Visited = true
			}
		}
	}

	// Sort the caverns slice by size. This will make the largest cavern last, which will then be removed from the list.
	// Then, fill in any remaining caverns (aside from the main one). This will ensure that there are no areas on the
	// map that the player cannot reach.
	sort.Sort(bySize(caverns))

	// Take the largest cavern (The one being used as the map), and record it as a list of open floor tiles, since thats
	// what it represents. This will be used for content generation.
	surface.FloorTiles = caverns[len(caverns)-1]
	caverns = caverns[:len(caverns)-1]

	for i := 0; i < len(caverns); i++ {
		for j := 0; j < len(caverns[i]); j++ {
			caverns[i][j].Blocked = true
			caverns[i][j].BlocksSight = true
			caverns[i][j].Glyph = wallGlyph
		}
	}
}

func countWallsNStepsAway(surface *gamemap.GameMap, n int, x int, y int) int {
	// Return the number of wall tiles that are within n spaces of the given tile
	wallCount := 0

	for r := -n; r <= n; r++ {
		for c := -n; c <= n; c++ {
			if x+r >= surface.Width || x+r <= 0 || y+c >= surface.Height || y+c <= 0 {
				// Check if the current coordinates would be off the map. Off map coordinates count as a wall.
				wallCount++
			} else if surface.Tiles[x+r][y+c].IsWall() {
				wallCount++
			}
		}
	}

	return wallCount
}
