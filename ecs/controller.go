package ecs

import (
	"reflect"
	"fmt"
)

type Controller struct {
	systems map[reflect.Type]System
	nextEntityID int
	components map[reflect.Type][]int
	entities map[int]map[reflect.Type]Component
	deadEntities []int
}

// NewController is a convenience/constructor method to properly initialize a new processor
func NewController() *Controller {
	controller := Controller{}
	controller.systems = make(map[reflect.Type]System)
	controller.nextEntityID = 0
	controller.components = make(map[reflect.Type][]int)
	controller.entities = make(map[int]map[reflect.Type]Component)
	controller.deadEntities = []int{}

	return &controller
}

// Create a new entity in the world. An entity is simply a unique integer.
// If any components are provided, they will be associated with the created entity
func (c *Controller) CreateEntity(components []Component) int {
	c.nextEntityID += 1

	if len(components) > 0 {
		for _, v := range components {
			c.AddComponent(c.nextEntityID, v)
		}
	}

	c.entities[c.nextEntityID] = make(map[reflect.Type]Component)

	return c.nextEntityID
}

// DeleteEntity removes an entity, all component instances attached to that entity, and any components associations with
// that entity
func (c *Controller) DeleteEntity(entity int) {
	// First, delete all the component associations for the entity to be removed
	for k, _ := range c.entities[entity] {
		c.RemoveComponent(entity, k)
	}

	// Then, delete the entity itself. The components have already been removed and disassociated with it, so a simple
	// delete will do here
	delete(c.entities, entity)
}

// AddComponent adds a component to an entity. The component is added to the global list of components for the
// processor, and also directly associated with the entity itself. This allows for flexible checking of components,
// as you can check which entites are associated with a component, and vice versa.
func (c *Controller) AddComponent(entity int, component Component) {
	// First, get the type of the component
	componentType := reflect.TypeOf(component)
	fmt.Printf("Adding component of type %v to entity %v\n", componentType, entity)

	// Record that the component type is associated with the entity.
	c.components[componentType] = append(c.components[componentType], entity)

	// Now, check to see if the entity is already tracked in the controller entity list. If it is not, add it, and
	// associate the component with it
	if _, ok := c.entities[entity]; !ok {
		c.entities[entity] = make(map[reflect.Type]Component)
	}

	c.entities[entity][componentType] = component
}

// HasComponent checks a given entity to see if it has a given component associated with it
func (c *Controller) HasComponent(entity int, componentType reflect.Type) bool {
	if _, ok := c.entities[entity][componentType]; ok {
		return true
	} else {
		return false
	}
}

// GetComponent returns the component instance for a component type, if one exists for the provided entity
func (c *Controller) GetComponent(entity int, componentType reflect.Type) Component {
	// Check the given entity has the provided component
	if c.HasComponent(entity, componentType) {
		return c.entities[entity][componentType]
	}

	return nil
}

// GetEntity gets a specific entity, and all of its component instances
func (c *Controller) GetEntity(entity int) map[reflect.Type]Component {
	for i, _ := range c.entities {
		if i == entity {
			return c.entities[entity]
		}
	}

	return nil
}

// GetEntities returns a map of all entities and their component instances
func (c *Controller) GetEntities() map[int]map[reflect.Type]Component {
	return c.entities
}

// DeleteComponent will delete a component instance from an entity, based on component type. It will also remove the
// association between the component and the entity, and remove the component from the processor completely if no
// other entities are using it.
func (c *Controller) RemoveComponent(entity int, componentType reflect.Type) int {
	// Find the index of the entity to operate on in the components slice
	index := -1
	for i, v := range c.components[componentType] {
		if (v == entity) {
			index = i
		}
	}

	// If the component was found on the entity, remove the association between the component and the entity
	if index != -1 {
		c.components[componentType] = append(c.components[componentType][:index], c.components[componentType][index+1:]...)
		// If this was the last entity associated with the component, remove the component entry as well
		if len(c.components[componentType]) == 0 {
			delete(c.components, componentType)
		}
	}

	// Now, remove the component instance from the entity
	delete(c.entities[entity], componentType)

	return entity
}

func (c *Controller) AddSystem(system System, priority int) {
	systemType := reflect.TypeOf(system)

	if _, ok := c.systems[systemType]; !ok {
		// A system of this type has not been added yet, so add it to the systems list
		c.systems[systemType] = system
	} else {
		fmt.Printf("A system of type %v was already added to the controller %v!", systemType, c)
	}
}

// Process kicks off system processing for all systems attached to the controller. Systems will be processed in the
// order they are found, or if they have a priority, in priority order. If there is a mix of systems with priority and
// without, systems with priority will be processed first (in order).
func (c *Controller) Process() {
	for _, system := range c.systems {
		system.Process()
	}
}