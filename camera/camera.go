package camera

type GameCamera struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (c *GameCamera) MoveCamera(targetX int, targetY int, mapWidth int, mapHeight int) {
	// Update the camera coordinates to the target coordinates
	x := targetX - c.Width/2
	y := targetY - c.Height/2

	if x < 0 {
		x = 0
	}

	if y < 0 {
		y = 0
	}

	if x > mapWidth-c.Width {
		x = mapWidth - c.Width
	}

	if y > mapHeight-c.Height {
		y = mapHeight - c.Height
	}

	c.X, c.Y = x, y
}

func (c *GameCamera) ToCameraCoordinates(mapX int, mapY int) (cameraX int, cameraY int) {
	// Convert coordinates on the map, to coordinates on the viewport
	x, y := mapX-c.X, mapY-c.Y

	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		return -1, -1
	}

	return x, y
}
