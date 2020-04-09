package ui

import (
	"fmt"
	"github.com/gogue-framework/gogue/ecs"
)

// TerminalInputHandler is a container for all keyHandlers and their respective required args. A key handler is
// a function mapped to a keypress event, and the handlerargs are any arguments that need to be passed to that function.
type TerminalInputHandler struct {
	keyHandlers map[int]func(map[string]interface{}, *ecs.Controller)
	handlerArgs map[int]map[string]interface{}
}

// NewTerminalInputHandler creates a new TerminalInputHandler instance. This can be used to register key handlers for
// use within the application
func NewTerminalInputHandler() *TerminalInputHandler {
	terminalInputHandler := TerminalInputHandler{}
	terminalInputHandler.keyHandlers = make(map[int]func(map[string]interface{}, *ecs.Controller))
	terminalInputHandler.handlerArgs = make(map[int]map[string]interface{})
	return &terminalInputHandler
}

// RegisterInputHandler registers a new key input handler to the provided key. args can optionally be provided
func (ti *TerminalInputHandler) RegisterInputHandler(key int, handlerFunction func(map[string]interface{}, *ecs.Controller), args map[string]interface{}) {
	// First, check to see if this key has already been registered
	if _, ok := ti.keyHandlers[key]; ok {
		fmt.Printf("The key %v has already been registered in this InputHandler. Aborting.", key)
		return
	}

	// The key has not already been assigned to a handler
	ti.keyHandlers[key] = handlerFunction

	// Now check to see if there are any arguments necessary (or provided) to this handler
	if len(args) > 0 {
		ti.handlerArgs[key] = args
	}
}

// ProcessInput takes a key, checks for a registered handler, and then runs that handler, with any provided args
func (ti *TerminalInputHandler) ProcessInput(key int, controller *ecs.Controller) {
	// Check to see if the pressed key has a handler. If it does not, do nothing.
	_, ok := ti.keyHandlers[key]

	// If a key handler has been registered, and the state is not currently a menu, go ahead and process the input
	// Otherwise, do nothing.
	if ok {
		keyHandlerFunction := ti.keyHandlers[key]
		keyHandlerArgs := ti.handlerArgs[key]

		keyHandlerFunction(keyHandlerArgs, controller)
	}

}
