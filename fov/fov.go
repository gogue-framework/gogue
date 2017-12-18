package fov

import (
	"math"
	"gogue/gamemap"
)

type FieldOfVision struct {
	cosTable    map[int]float64
	sinTable    map[int]float64
	torchRadius int
}

func (f *FieldOfVision) Initialize() {

	f.cosTable = make(map[int]float64)
	f.sinTable = make(map[int]float64)

	for i := 0; i < 360; i++ {
		ax := math.Sin(float64(i) / (float64(180) / math.Pi))
		ay := math.Cos(float64(i) / (float64(180) / math.Pi))

		f.sinTable[i] = ax
		f.cosTable[i] = ay
	}
}

func (f *FieldOfVision) SetTorchRadius(radius int) {
	if radius > 1 {
		f.torchRadius = radius
	}
}

func (f *FieldOfVision) SetAllInvisible(gameMap *gamemap.Map) {
	for x := 0; x < gameMap.Width; x++ {
		for y := 0; y < gameMap.Height; y++ {
			gameMap.Tiles[x][y].Visible = false
		}
	}
}

func (f *FieldOfVision) RayCast(playerX, playerY int, gameMap *gamemap.Map) {
	// Cast out rays each degree in a 360 circle from the player. If a ray passes over a floor (does not block sight)
	// tile, keep going, up to the maximum torch radius (view radius) of the player. If the ray intersects a wall
	// (blocks sight), stop, as the player will not be able to see past that. Every visible tile will get the Visible
	// and Explored properties set to true.

	for i := 0; i < 360; i++ {

		ax := f.sinTable[i]
		ay := f.cosTable[i]

		x := float64(playerX)
		y := float64(playerY)

		// Mark the players current position as explored
		gameMap.Tiles[playerX][playerY].Explored = true
		gameMap.Tiles[playerX][playerY].Visible = true

		for j := 0; j < f.torchRadius; j++ {
			x -= ax
			y -= ay

			roundedX := int(Round(x))
			roundedY := int(Round(y))

			if x < 0 || x > float64(gameMap.Width-1) || y < 0 || y > float64(gameMap.Height-1) {
				// If the ray is cast outside of the map, stop
				break
			}

			gameMap.Tiles[roundedX][roundedY].Explored = true
			gameMap.Tiles[roundedX][roundedY].Visible = true

			//fmt.Printf("Actual X (float64): %d, rounded X (int): %d\n", x, int(Round(x)))

			if gameMap.Tiles[roundedX][roundedY].BlocksSight == true {
				// The ray hit a wall, go no further
				break
			}
		}
	}
}

func Round(f float64) float64 {
	return math.Floor(f + .5)
}