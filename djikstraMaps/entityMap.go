package DijkstraMaps

import (
	"github.com/jcerise/gogue/gamemap"
)

// An EntityMap is a Dijkstra map that centers around an entity. This could be the player, an item, a monster, a door,
// etc. They are the simplest implementation, as the map radiates values from a single point, setting that point (the
// location of the entity) as the source, meaning it will have the lowest value. These maps can optionally be inverted,
// to make other entities move away from it (fleeing, for example). Each map will keep track of where the source entity
// is, and where it was the previous turn. The map only needs to be recalculated if the source entity moved, otherwise,
// the map can continually be re-used turn after turn without recalculation. Each map will also keep track of its type
// (Player, health potion, mana potion, weapon, pack entity, terrifying, etc), and a master list of all maps can be
// maintained, so each entity that cares about them can utilize the appropriate maps each turn.
// Some examples:
// The most obvious example is stalking the player. The player would be the source entity, and the map would be drawn
// from her location each time movement occurred. Any monster or entity that cared about stalking the player can then
// simply use the map values, rolling downhill towards the player each turn.
//
// A monster that cares about gold will check the map representing any gold entities on the map, and will move towards
// those in the same manner they would move towards the player. If there are multiple entities as the map source, they
// will move towards the closest.
//
// Multiple competing desires. If a monster wants gold, to kill the player, and pick up a health potion, they can
// maintain a weight for each one of those desires (updated each turn, according to whats happening). These weights can
// then be multiplied across all the values of every competing map. A positive number means they want to be far away
// from that entity, 0 is indifference, and negative numbers mean high desire. Multiply values on the map by the desires
// and you end up with a combined set of maps with values that reflect the monsters desires.

type EntityDijkstraMap struct {
	source int // The source entity ID
	sourceX int
	sourceY int
	sourcePrevX int
	sourcePrevY int
	mapType string
	ValuesMap map[gamemap.CoordinatePair]int
}

func NewEntityMap(sourceEntity int, sourceX, sourceY int, mapType string) *EntityDijkstraMap {
	edm := EntityDijkstraMap{}
	edm.ValuesMap = make(map[gamemap.CoordinatePair]int)

	// Set the source position
	edm.sourceX = sourceX
	edm.sourceY = sourceY

	// The previous positions will be the same initially
	edm.sourcePrevX = sourceX
	edm.sourcePrevY = sourceY

	// Set the map type
	edm.mapType = mapType

	return &edm
}

// UpdateSourceCoordinates sets the sourceX and sourceY properties to the latest values available, recording the
// previous coordinates for later use.
func (edm *EntityDijkstraMap) UpdateSourceCoordinates(x, y int) {
	edm.sourcePrevX = edm.sourceX
	edm.sourcePrevY = edm.sourceY

	edm.sourceX = x
	edm.sourceY = y
}

// UpdateMap checks the map to see if the update criteria (location of the source entity has changed) is met. If so,
// the map will be regenerated.
func (edm *EntityDijkstraMap) UpdateMap(surface *gamemap.Map) {
	if (edm.sourceX != edm.sourcePrevX) || (edm.sourceY != edm.sourcePrevY) {
		// The coordinates differ from the last previous set, meaning the entity has moved. Re-generate the map.
		edm.GenerateMap(surface)
	}
}

// GenerateMap will create a Dijkstra map, centered around the source entities current location.
func (edm *EntityDijkstraMap) GenerateMap(surface *gamemap.Map) {
	// Starting from the source, flood fill every tile on the map, incrementing the value for each tile by one,
	// based on how far away from the source it is. Make a visited array first though (everything but the source is
	// unvisited initially. Also mark blocking tiles as visited.
	visited := make(map[gamemap.CoordinatePair]bool)
	for x := 0; x < surface.Width-1; x++ {
		for y := 0; y < surface.Height-1; y++ {
			coordinates := gamemap.CoordinatePair{X: x, Y: y}

			// Mark all blocking tiles as visited, so we don't even bother with them.
			if surface.Tiles[x][y].Blocked {
				visited[coordinates] = true
			} else {
				visited[coordinates] = false
			}
		}
	}

	edm.BreadthFirstSearch(edm.sourceX, edm.sourceY, surface.Width, surface.Height, 0, visited)
}

func (edm *EntityDijkstraMap) BreadthFirstSearch(x, y, n, m, value int, visited map[gamemap.CoordinatePair]bool) {
	// Check if this location has already been visited
	coordinates := gamemap.CoordinatePair{X: x, Y: y}

	// Mark this location as visited, set the value for this location in the EDM, and increase the value by one
	// This will ensure that each subsequently further tile will have an increased value
	edm.ValuesMap[coordinates] = value

	coordQueue := []gamemap.CoordinatePair{coordinates}

	for len(coordQueue) > 0 {
		coords :=  coordQueue[len(coordQueue)-1]
		coordQueue = coordQueue[:len(coordQueue)-1]

		// Check all the immediate neighbors, and set values for them based on the current coordinates value
		// NorthWest
		if ((coords.X - 1) > 0 && (coords.Y - 1) > 0) && ((coords.X - 1) <= n && (coords.Y - 1) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X - 1, Y:coords.Y - 1}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// West
		if ((coords.X - 1) > 0 && (coords.Y) > 0) && ((coords.X - 1) <= n && (coords.Y) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X - 1, Y:coords.Y}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// SouthWest
		if ((coords.X - 1) > 0 && (coords.Y + 1) > 0) && ((coords.X - 1) <= n && (coords.Y + 1) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X - 1, Y:coords.Y + 1}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// South
		if ((coords.X) > 0 && (coords.Y + 1) > 0) && ((coords.X) <= n && (coords.Y + 1) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X, Y:coords.Y + 1}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// SouthEast
		if ((coords.X + 1) > 0 && (coords.Y + 1) > 0) && ((coords.X + 1) <= n && (coords.Y + 1) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X + 1, Y:coords.Y + 1}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// East
		if ((coords.X + 1) > 0 && (coords.Y) > 0) && ((coords.X + 1) <= n && (coords.Y) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X + 1, Y:coords.Y}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// NorthEast
		if ((coords.X + 1) > 0 && (coords.Y - 1) > 0) && ((coords.X + 1) <= n && (coords.Y - 1) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X - 1, Y:coords.Y - 1}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}

		// North
		if ((coords.X) > 0 && (coords.Y - 1) > 0) && ((coords.X) <= n && (coords.Y - 1) <= m) {
			newCoords := gamemap.CoordinatePair{X:coords.X, Y:coords.Y - 1}
			if !visited[newCoords] {
				// This is a valid, un-visited, neighbor. Give it a value of (currentVal + 1), add it to the valueMap, and
				// add it to the coordQueue; We'll check its neighbors soon

				visited[coords] = true
				edm.ValuesMap[newCoords] = edm.ValuesMap[coords] + 1
				coordQueue = append(coordQueue, newCoords)
			}
		}
	}
}

