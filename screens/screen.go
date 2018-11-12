package screens

import "fmt"

type Screen interface {
	Enter()
	Exit()
	Render()
	HandleInput()
	UseEcs() bool
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
		fmt.Printf("A screen with name %v was already added to the ScreenManager %v!", screenName, sm)
	}
}

// RemoveScreen will remove a screen from the ScreenManager. This can be useful when a temporary screen needs to be
// created, as it can be quickly added (rather than registering at game creation), and then removed when it is no
// longer needed
func (sm *ScreenManager) RemoveScreen(screenName string, screen Screen) {
	// Check if the given screenName exists in the ScreenManager
	if _, ok := sm.Screens[screenName]; ok {
		delete(sm.Screens, screenName)
	} else {
		// A screen with the given name does not exist
		fmt.Printf("A screen with name %v was not found on ScreenManager %v!", screenName, sm)
	}
}

// SetScreen will set the current screen property of the screen manager to the provided screen
func (sm *ScreenManager) SetScreen(screen Screen) {
	// Call the exit function of the currentScreen, and set the currentScreen as the previousScreen
	// Only do this if there is a currentScreen
	if sm.CurrentScreen != nil {
		sm.CurrentScreen.Exit()
		sm.PreviousScreen = sm.CurrentScreen
	}

	// Set the provided screen as the currentScreen, and call the enter() function of the new currentScreen
	sm.CurrentScreen = screen
	sm.CurrentScreen.Enter()
}

// SetScreenByName takes a string representing the screen desired to navigate to. It will then transition the
// ScreenManager to the specified screen, if one is found.
func (sm *ScreenManager) SetScreenByName(screenName string) {
	// Check if the given screenName exists in the ScreenManager
	if _, ok := sm.Screens[screenName]; ok {
		// Call the exit function of the currentScreen, and set the currentScreen as the previousScreen
		// Only do this if there is a currentScreen
		if sm.CurrentScreen != nil {
			sm.CurrentScreen.Exit()
			sm.PreviousScreen = sm.CurrentScreen
		}

		// Set the provided screen as the currentScreen, and call the enter() function of the new currentScreen
		sm.CurrentScreen = sm.Screens[screenName]
		sm.CurrentScreen.Enter()
	} else {
		// A screen with the given name does not exist
		fmt.Printf("A screen with name %v was not found on ScreenManager %v!", screenName, sm)
	}
}
