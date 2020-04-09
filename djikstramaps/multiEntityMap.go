package dijkstramaps

import (
	"github.com/gogue-framework/gogue/ecs"
	"github.com/gogue-framework/gogue/gamemap"
)

// DMSource represents a single source in a multi-entity Dijkstra map
type DMSource struct {
	entity int // The source entity ID
	X      int
	Y      int
	PrevX  int
	PrevY  int
}

// MultiEntityDijkstraMap is a Dijkstra map that tracks several entities. It is very similar to the single entity map, in that
// each entity radiates a value outwards from it, the entity being the source. The main difference is that there can be
// many sources, each affecting the values of the map.
//
// General idea is that each provided source will act as a single entity map source would, radiating values out from it.
// But, since there will be multiple sources, values from other sources are never overwritten. This means that as the
// values radiate out from a source, they will be calculated until they hit values from another source, and then stop.
// In this way, a map representing the distance for each item will be generated.
//
// It must be noted, this type of map assumes that all sources are of the same entity type, otherwise, it wouldn't
// really make sense.
type MultiEntityDijkstraMap struct {
	sources   map[int]DMSource
	mapType   string
	ValuesMap [][]int
	visited   map[*gamemap.Tile]bool
}

// NewMultiEntityMap creates a new multi-entity Dijkstra map. sourceList is the list of sources for the map to operate
// from, and mapWidth and mapHeight indicate how large the resulting map should be (typically the size of the gamemap).
// mapType is a string identifier for the map.
func NewMultiEntityMap(sourceList map[int]DMSource, mapType string, mapWidth, mapHeight int) *MultiEntityDijkstraMap {
	medm := MultiEntityDijkstraMap{}
	medm.ValuesMap = make([][]int, mapWidth+1)
	medm.visited = make(map[*gamemap.Tile]bool)
	for i := range medm.ValuesMap {
		medm.ValuesMap[i] = make([]int, mapHeight+1)
	}

	medm.sources = sourceList

	// Set the map type
	medm.mapType = mapType

	return &medm
}

// AddSourceEntity adds a new entity to the source list for a multi-entity map
func (medm *MultiEntityDijkstraMap) AddSourceEntity(source DMSource) {
	medm.sources[source.entity] = source
}

// UpdateSourceEntity updates an existing source with a new position, for example, if the source entity moved, the
// Dijkstra map will need to be re-calculated based on the entities new position.
func (medm *MultiEntityDijkstraMap) UpdateSourceEntity(entity, newX, newY int) {
	source := medm.sources[entity]

	source.PrevX = source.X
	source.PrevY = source.X

	source.X = newX
	source.Y = newY
}

// GenerateMap runs through creating a new multi-entity dijkstra map. This will calculate distances for every tile in
// the map, from each source entity, to each other source entity. For example, if two source entities are three tiles
// apart, the max distance for a tile between them would be two, as we only care about the maximum distance from any
// entity.
func (medm *MultiEntityDijkstraMap) GenerateMap(surface *gamemap.GameMap) {
	sourceList := make(map[int][]*gamemap.Tile, len(medm.sources))
	medm.visited = make(map[*gamemap.Tile]bool)

	// Populate the sourceList
	for entity, source := range medm.sources {
		tile := surface.Tiles[source.X][source.Y]
		tileMap := []*gamemap.Tile{tile}

		// Also set the starting value for each source tile to zero
		medm.ValuesMap[source.X][source.Y] = 0

		sourceList[entity] = tileMap
	}

	// Now, iterate over each source, running a single round of BFS. If there are no tiles in the sources tileMap, no
	// further BFS search rounds need to be run. If all sources have no tiles in their tileMaps, then exit the loop,
	// as we're done
	finishedSources := []int{}

	for entity, tileMap := range sourceList {

		// If every source has been added to the finishedSources list, we're done, so exit the loop
		if len(finishedSources) == len(medm.sources) {
			break
		}

		// Check to see if this source has any tiles in its tileList. If it does not, it is done, and should be marked
		// as such. If there are tiles, continue processing
		if len(tileMap) == 0 && !ecs.IntInSlice(entity, finishedSources) {
			finishedSources = append(finishedSources, entity)
		}

		tile := tileMap[0]
		sourceList[entity] = tileMap[1:]
		sourceList[entity] = append(sourceList[entity], medm.SingleRoundBreadthFirstSearch(tile.X, tile.Y, surface)...)
	}
}

// SingleRoundBreadthFirstSearch runs a single round of a breadth first search algorithm. This will calculate distances
// for one level of distance from a source entity, rather than calculating every distance at once. This is useful for
// calculating multiple entity distances, as we can do them one level, and entity, at a time, and the see where the
// distances may overlap.
func (medm *MultiEntityDijkstraMap) SingleRoundBreadthFirstSearch(x, y int, surface *gamemap.GameMap) []*gamemap.Tile {
	// Check if this location has already been visited
	curTile := surface.Tiles[x][y]

	// Mark this location as visited, and increase the value by one
	// This will ensure that each subsequently further tile will have an increased value

	medm.visited[curTile] = true

	tileQueue := []*gamemap.Tile{}

	curValue := medm.ValuesMap[curTile.X][curTile.Y] + 1

	// Check all the immediate neighbors, and set values for them based on the current coordinates value
	// NorthWest
	neighborTile := surface.Tiles[curTile.X-1][curTile.Y-1]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// West
	neighborTile = surface.Tiles[curTile.X-1][curTile.Y]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// SouthWest
	neighborTile = surface.Tiles[curTile.X-1][curTile.Y+1]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// South
	neighborTile = surface.Tiles[curTile.X][curTile.Y+1]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// SouthEast
	neighborTile = surface.Tiles[curTile.X+1][curTile.Y+1]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// East
	neighborTile = surface.Tiles[curTile.X+1][curTile.Y]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// NorthEast
	neighborTile = surface.Tiles[curTile.X+1][curTile.Y-1]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	// North
	neighborTile = surface.Tiles[curTile.X][curTile.Y-1]
	if !medm.visited[neighborTile] && !neighborTile.IsWall() {
		// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
		// add it to the tileQueue; We'll check its neighbors soon

		medm.visited[neighborTile] = true
		medm.ValuesMap[neighborTile.X][neighborTile.Y] = curValue
		tileQueue = append(tileQueue, neighborTile)
	}

	return tileQueue
}
