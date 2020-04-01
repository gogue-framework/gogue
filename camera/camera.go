package camera

import "errors"

type GameCamera struct {
	X      int
	Y      int
	Width  int
	Height int
}

// NewGameCamera creates a new GameCamera instance. (x, y) are the starting location for the camera viewport. The
// viewport will be centered over these coordinates on the GameMap. width and height designate the size of the
// cameras viewport, or how much of the GameMap is shown. If the viewport is the same size as, or larger than the
// GameMap, the entire map will be shown. If the viewport is smaller than the GameMap, only the portion of the GameMap
// that fits within the centered viewport will be shown.
func NewGameCamera(x, y, width, height int) (*GameCamera, error) {
	gameCamera := GameCamera{}

	if x < 0 || y < 0 || width < 0 || height < 0 {
		err := errors.New("GameCamera position values must be greater than 0")
		return nil, err
	}

	gameCamera.X = x
	gameCamera.Y = y
	gameCamera.Width = width
	gameCamera.Height = height

	return &gameCamera, nil
}

// MoveCamera centers the GameCamera viewport to the new location specified over the GameMap. If the viewport is
// larger than the GameMap, the viewport is not moved. If the viewport is the same size as the GameMap, the viewport
// is not moved. If the GameMap is larger than the viewport, and coordinates outside of the viewport are requested,
// the viewport is centered over the new coordinate set. The viewport cannot be centered outside the bounds of the
// GameMap.
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

// ToCameraCoordinates translates a GameMap coordinate pair to a GameCamera coordinate pair. GameMap coordinates and
// Camera coordinates are two different things. A GameMap coordinate pair designates a location on the GameMap, whereas
// a Camera coordinate pair designates a location within the cameras viewport. ToCameraCoordinates translates a GameMap
// coordinate pair to a Camera coordinate pair. If the map coordinate pair is outside of the viewport of the camera,
// (-1, -1) is returned to indicate this
func (c *GameCamera) ToCameraCoordinates(mapX int, mapY int) (cameraX int, cameraY int) {
	// Convert coordinates on the map, to coordinates on the viewport
	x, y := mapX-c.X, mapY-c.Y

	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		return -1, -1
	}

	return x, y
}
