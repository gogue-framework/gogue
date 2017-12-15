package gogue

import (
	"fmt"
	"gogue/ecs"
)

type TerminalInputHandler struct {
	keyHandlers map[int]func(map[string]interface{}, *ecs.Controller)
	handlerArgs map[int]map[string]interface{}
}

func NewTerminalInputHandler() *TerminalInputHandler {
	terminalInputHandler := TerminalInputHandler{}
	terminalInputHandler.keyHandlers = make(map[int]func(map[string]interface{}, *ecs.Controller))
	terminalInputHandler.handlerArgs = make(map[int]map[string]interface{})
	return &terminalInputHandler
}

// RegisterInputHandler registers a new key input handler to the provided key. args can optionally be provided
func (ti TerminalInputHandler) RegisterInputHandler(key int, handlerFunction func(map[string]interface{}, *ecs.Controller), args map[string]interface{}) {
	// First, check to see if this key has already been registered
	if _, ok := ti.keyHandlers[key]; ok {
		fmt.Printf("The key %v has already been registered in this InputHandler. Aborting.", key)
		return
	} else {
		// The key has not already been assigned to a handler
		ti.keyHandlers[key] = handlerFunction
	}

	// Now check to see if there are any arguments necessary (or provided) to this handler
	if len(args) > 0 {
		ti.handlerArgs[key] = args
	}
}

// Process Input takes a key, checks for a registered handler, and then runs that handler, with any provided args
func (ti TerminalInputHandler) ProcessInput(key int, controller *ecs.Controller) {
	// Check to see if the pressed key has a handler. If it does not, do nothing.
	_, ok := ti.keyHandlers[key]

	if ok {
		keyHandlerFunction := ti.keyHandlers[key]
		keyHandlerArgs := ti.handlerArgs[key]

		keyHandlerFunction(keyHandlerArgs, controller)
	}

}

