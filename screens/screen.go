package screens

import "fmt"

type Screen interface {
	enter()
	exit()
	render()
	handleInput()
	useEcs()
}

type ScreenManager struct {
	Screens map[string]Screen
	CurrentScreen Screen
	PreviousScreen Screen
}

// NewScreenManager is a convenience/constructor method to properly initialize a new ScreenManager
func NewScreenManager() *ScreenManager {
	manager := ScreenManager{}
	manager.Screens = make(map[string]Screen)
	manager.CurrentScreen = nil

	return &manager
}

func (sm *ScreenManager) AddScreen(screenName string, screen Screen) {
	// Check to see if a screen with the given screenName has already been added
	if _, ok := sm.Screens[screenName]; !ok {
		// A screen with the given name does not yet exist on the ScreenManager, go ahead and add it
		sm.Screens[screenName] = screen
	} else {
		fmt.Printf("A screen with name %v was already added to the ScreenManager %v!", systemType, c)
	}
}

func (sm *ScreenManager) SetScreen(screenName string) {
	// Check if the given screenName exists in the ScreenManager
	if _, ok := sm.Screens[screenName]; ok {
		// Call the exit function of the currentScreen, and set the currentScreen as the previousScreen
		sm.CurrentScreen.exit()
		sm.PreviousScreen = sm.CurrentScreen

		// Set the provided screen as the currentScreen, and call the enter() function of the new currentScreen
		sm.CurrentScreen = sm.Screens[screenName]
		sm.CurrentScreen.enter()
	} else {
		// A screen with the given name does not exist
		fmt.Printf("A screen with name %v was not found on ScreenManager %v!", screenName, sm)
	}
}
