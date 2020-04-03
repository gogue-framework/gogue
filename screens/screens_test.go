package screens

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Create a few screens for testing purposes
type MenuScreen struct {
	exitCalled   bool
	enterCalled  bool
	renderCalled bool
}

func (ms *MenuScreen) Enter()       { ms.enterCalled = true }
func (ms *MenuScreen) Exit()        { ms.exitCalled = true }
func (ms *MenuScreen) Render()      { ms.renderCalled = true }
func (ms *MenuScreen) HandleInput() {}
func (ms *MenuScreen) UseEcs() bool { return false }

type GameScreen struct {
	exitCalled   bool
	enterCalled  bool
	renderCalled bool
}

func (gs *GameScreen) Enter()       { gs.enterCalled = true }
func (gs *GameScreen) Exit()        { gs.exitCalled = true }
func (gs *GameScreen) Render()      { gs.renderCalled = true }
func (gs *GameScreen) HandleInput() {}
func (gs *GameScreen) UseEcs() bool { return false }

func TestNewScreenManager(t *testing.T) {
	manager := NewScreenManager()

	assert.Nil(t, manager.CurrentScreen)
	assert.Equal(t, 0, len(manager.Screens))
}

func TestScreenManager_AddScreen(t *testing.T) {
	manager := NewScreenManager()

	menu := &MenuScreen{}
	game := &GameScreen{}

	err := manager.AddScreen("menu", menu)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(manager.Screens))

	expectedScreensMap := map[string]Screen{
		"menu": menu,
	}
	assert.Equal(t, expectedScreensMap, manager.Screens)

	err = manager.AddScreen("menu", menu)
	assert.NotNil(t, err)

	err = manager.AddScreen("game", game)
	expectedScreensMap["game"] = game

	assert.Nil(t, err)
	assert.Equal(t, 2, len(manager.Screens))
	assert.Equal(t, expectedScreensMap, manager.Screens)
}

func TestScreenManager_RemoveScreen(t *testing.T) {
	manager := NewScreenManager()

	menu := &MenuScreen{}
	game := &GameScreen{}

	_ = manager.AddScreen("menu", menu)
	_ = manager.AddScreen("game", game)

	expectedScreensMap := map[string]Screen{
		"menu": menu,
		"game": game,
	}

	assert.Equal(t, expectedScreensMap, manager.Screens)

	// Now, remove a screen
	err := manager.RemoveScreen("game")
	delete(expectedScreensMap, "game")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(manager.Screens))
	assert.Equal(t, expectedScreensMap, manager.Screens)

	// Attempt to remove a screen that is not on the ScreenManager
	err = manager.RemoveScreen("does_not_exist")
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(manager.Screens))
}

func TestScreenManager_setScreen(t *testing.T) {
	manager := NewScreenManager()

	menu := &MenuScreen{}
	game := &GameScreen{}

	_ = manager.AddScreen("menu", menu)
	_ = manager.AddScreen("game", game)

	manager.setScreen(menu)
	// Since there is no current screen to exit, exit will not be called
	assert.False(t, menu.exitCalled)
	assert.True(t, menu.enterCalled)
	assert.Equal(t, menu, manager.CurrentScreen)
	assert.Nil(t, manager.PreviousScreen)

	manager.setScreen(game)
	assert.False(t, game.exitCalled)
	assert.True(t, menu.exitCalled)
	assert.True(t, game.enterCalled)
	assert.Equal(t, game, manager.CurrentScreen)
	assert.Equal(t, menu, manager.PreviousScreen)
}

func TestScreenManager_SetScreenByName(t *testing.T) {
	manager := NewScreenManager()

	menu := &MenuScreen{}
	game := &GameScreen{}

	_ = manager.AddScreen("menu", menu)
	_ = manager.AddScreen("game", game)

	err := manager.SetScreenByName("menu")
	assert.Nil(t, err)
	// Since there is no current screen to exit, exit will not be called
	assert.False(t, menu.exitCalled)
	assert.True(t, menu.enterCalled)
	assert.Equal(t, menu, manager.CurrentScreen)
	assert.Nil(t, manager.PreviousScreen)

	err = manager.SetScreenByName("game")
	assert.Nil(t, err)
	assert.False(t, game.exitCalled)
	assert.True(t, menu.exitCalled)
	assert.True(t, game.enterCalled)
	assert.Equal(t, game, manager.CurrentScreen)
	assert.Equal(t, menu, manager.PreviousScreen)

	err = manager.SetScreenByName("nonExistant")
	assert.NotNil(t, err)
}
