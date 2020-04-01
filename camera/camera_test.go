package camera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCamera(t *testing.T) {
	gameCamera, err := NewGameCamera(0, 1, 2, 3)

	assert.NotNil(t, gameCamera)
	assert.Nil(t, err)
	assert.Equal(t, gameCamera.X, 0)
	assert.Equal(t, gameCamera.Y, 1)
	assert.Equal(t, gameCamera.Width, 2)
	assert.Equal(t, gameCamera.Height, 3)

	// Check to ensure values < 0 are not allowed
	gameCamera, err = NewGameCamera(-1, 0, 1, 2)

	assert.NotNil(t, err)
	assert.Nil(t, gameCamera)
}

func TestGameCamera_MoveCamera(t *testing.T) {
	// Set up a new camera, with a viewport of 100x100
	gameCamera, _ := NewGameCamera(0, 0, 100, 100)

	// Move the camera to a location that is already within the cameras viewport. This effectively will not
	// move the viewport
	gameCamera.MoveCamera(25, 25, 100, 100)
	assert.Equal(t, gameCamera.X, 0)
	assert.Equal(t, gameCamera.Y, 0)

	// Move the camera to a location that is more than halfway across the map, and the cameras viewport, on a small
	// map. This should not move the camera, since the entire map is within the cameras viewport
	gameCamera.MoveCamera(51, 51, 100, 100)
	assert.Equal(t, gameCamera.X, 0)
	assert.Equal(t, gameCamera.Y, 0)

	// Move the camera to a location that is outside of its viewport (on a larger map). This should update the
	// the cameras coordinates
	gameCamera.MoveCamera(101, 101, 500, 500)
	assert.Equal(t, gameCamera.X, 51)
	assert.Equal(t, gameCamera.Y, 51)
}

func TestGameCamera_ToCameraCoordinates(t *testing.T) {
	// GameMap coordinates and Camera coordinates are two different things. A GameMap coordinate pair designates a
	// a location on the GameMap, whereas a Camera coordinate pair designates a location within the cameras viewport.
	// ToCameraCoordinates translates a GameMap coordinate pair to a Camera coordinate pair. If the map coordinate pair
	// is outside of the viewport of the camera, (-1, -1) is returned to indicate this
	gameCamera, _ := NewGameCamera(0, 0, 100, 100)

	// With the camera viewport matching the top left corner of the map, the coordinates should be the same for both
	cameraX, cameraY := gameCamera.ToCameraCoordinates(2, 2)
	assert.Equal(t, cameraX, 2)
	assert.Equal(t, cameraY, 2)

	// If the camera is moved beyond the top left corner of the map, the two coordinate systems no longer align
	gameCamera.MoveCamera(50, 70, 500, 500)
	cameraX, cameraY = gameCamera.ToCameraCoordinates(40, 50)
	assert.Equal(t, cameraX, 40)
	assert.Equal(t, cameraY, 30)

	// If a map coordinate pair that is outside of the viewport of the camera is provided, (-1, -1) should be returned
	cameraX, cameraY = gameCamera.ToCameraCoordinates(1, 1)
	assert.Equal(t, cameraX, -1)
	assert.Equal(t, cameraY, -1)

	cameraX, cameraY = gameCamera.ToCameraCoordinates(499, 499)
	assert.Equal(t, cameraX, -1)
	assert.Equal(t, cameraY, -1)
}
