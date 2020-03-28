package ecs

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type PositionComponent struct {
	X int
	Y int
}

func (pc PositionComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(pc)
}

type AppearanceComponent struct {
	Appearance string
}

func (ac AppearanceComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(ac)
}

type RandomComponent struct {
	prop string
}

func (rc RandomComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(rc)
}

func TestNewController(t *testing.T) {
	controller := NewController()

	assert.Equal(t, controller.nextEntityID, 0, "controller did not init properly")
}

func TestCreateNewEntity(t *testing.T) {
	controller := NewController()

	controller.CreateEntity([]Component{})

	assert.Equal(t, controller.nextEntityID, 1, "expected nextEntityId to increment")

	assert.Equal(t, len(controller.entities), 1, "expected entities array to contain 1 entity")

	// The global list of components should be empty, as we didn't init this entity with
	// any components
	components := controller.components
	assert.Equal(t, len(components), 0, "components list should be empty")

	// Checking for a component on the new entity should return false, as it was not init'd
	// with one
	assert.False(t, controller.HasComponent(0, reflect.TypeOf(PositionComponent{})), "entity should not have component PositionComponent")
}

func TestGetEntity(t *testing.T) {
	controller := NewController()

	entity := controller.CreateEntity([]Component{})

	assert.NotNil(t, controller.GetEntity(entity), "entity not retrieved")
	assert.Nil(t, controller.GetEntity(1000), "incorrect entity retrieved")
}

func TestCreateNewEntityWithComponent(t *testing.T) {
	controller := NewController()

	entity := controller.CreateEntity([]Component{PositionComponent{}})

	assert.Equal(t, entity, 0, "expected entity to be 0")
	assert.Equal(t, controller.nextEntityID, 1, "expected nextEntityId to increment")

	components := controller.components
	assert.Equal(t, len(components), 1, "components should contain 1 component")
	assert.True(t, controller.HasComponent(0, reflect.TypeOf(PositionComponent{})))

	// Also check that we can add new components after an entity has been created
	secondEntity := controller.CreateEntity([]Component{})
	controller.AddComponent(secondEntity, PositionComponent{})

	components = controller.components
	assert.Equal(t, len(components), 1, "components should contain 1 component")
	assert.True(t, controller.HasComponent(1, reflect.TypeOf(PositionComponent{})))
}

func TestDeleteEntity(t *testing.T) {
	controller := NewController()

	entity := controller.CreateEntity([]Component{})
	assert.Equal(t, entity, 0, "entity was not properly created")
	assert.Equal(t, len(controller.entities), 1, "entity not added to controller")

	controller.DeleteEntity(entity)
	assert.Equal(t, len(controller.entities), 0, "entity was not deleted")
	assert.False(t, controller.HasComponent(entity, reflect.TypeOf(PositionComponent{})), "entity has component")
}

func TestComponentMapping(t *testing.T) {
	controller := NewController()

	controller.MapComponentClass("position", PositionComponent{})
	assert.True(t, controller.HasMappedComponent("position"), "component position not present")
	assert.False(t, controller.HasMappedComponent("movement"), "component movement present")
}

func TestAddComponent(t *testing.T) {
	controller := NewController()

	entity := controller.CreateEntity([]Component{})

	controller.AddComponent(entity, PositionComponent{})

	assert.True(t, controller.HasComponent(entity, reflect.TypeOf(PositionComponent{})), "position component missing")
	assert.NotNil(t, controller.GetComponent(entity, reflect.TypeOf(PositionComponent{})), "could not get position component")
}

func TestGetEntitiesWithComponent(t *testing.T) {
	controller := NewController()

	entity := controller.CreateEntity([]Component{PositionComponent{}, AppearanceComponent{}})
	entity2 := controller.CreateEntity([]Component{AppearanceComponent{}})
	entity3 := controller.CreateEntity([]Component{PositionComponent{}})

	entityList := controller.GetEntitiesWithComponent(PositionComponent{}.TypeOf())
	assert.Equal(t, entityList, []int{entity, entity3})

	entityList = controller.GetEntitiesWithComponent(AppearanceComponent{}.TypeOf())
	assert.Equal(t, entityList, []int{entity, entity2})

	entityList = controller.GetEntitiesWithComponent(RandomComponent{}.TypeOf())
	assert.Empty(t, entityList)

}

func TestUpdateComponent(t *testing.T) {
	controller := NewController()

	playerPosition := PositionComponent{
		X: 0,
		Y: 0,
	}

	entity := controller.CreateEntity([]Component{playerPosition})
	component := controller.GetComponent(entity, PositionComponent{}.TypeOf()).(PositionComponent)
	assert.Equal(t, component.X, 0)
	assert.Equal(t, component.Y, 0)

	// Change the value of X and Y in the PositionComponent
	playerPosition = PositionComponent{
		X: 10,
		Y: 5,
	}

	entity = controller.UpdateComponent(entity, PositionComponent{}.TypeOf(), playerPosition)
	newComponent := controller.GetComponent(entity, PositionComponent{}.TypeOf()).(PositionComponent)
	assert.Equal(t, newComponent.X, 10)
	assert.Equal(t, newComponent.Y, 5)
	assert.True(t, controller.HasComponent(entity, PositionComponent{}.TypeOf()))

	monsterPosition := PositionComponent{
		X: 1,
		Y: 0,
	}

	// Check calling UpdateComponent with a component the entity does not have.
	entity2 := controller.CreateEntity([]Component{})
	entity = controller.UpdateComponent(entity, PositionComponent{}.TypeOf(), monsterPosition)
	assert.False(t, controller.HasComponent(entity2, PositionComponent{}.TypeOf()))

}

func TestRemoveComponent(t *testing.T) {
	controller := NewController()

	entity := controller.CreateEntity([]Component{PositionComponent{}})
	assert.True(t, controller.HasComponent(entity, PositionComponent{}.TypeOf()))

	controller.RemoveComponent(entity, PositionComponent{}.TypeOf())
	assert.False(t, controller.HasComponent(entity, PositionComponent{}.TypeOf()))

	// Re-add the component to ensure it does not get removed in the next section
	controller.AddComponent(entity, PositionComponent{})
	assert.True(t, controller.HasComponent(entity, PositionComponent{}.TypeOf()))

	// Attempt to remove a component that doesn't exist. No errors should be raised
	entity2 := controller.CreateEntity([]Component{})
	assert.False(t, controller.HasComponent(entity2, PositionComponent{}.TypeOf()))
	controller.RemoveComponent(entity2, PositionComponent{}.TypeOf())
	assert.False(t, controller.HasComponent(entity2, PositionComponent{}.TypeOf()))
	assert.True(t, controller.HasComponent(entity, PositionComponent{}.TypeOf()))
}

// System tests

type TestSystem struct {
	SystemRun bool
}

func (ts *TestSystem) Process() {
	ts.SystemRun = true
}

type AnotherSystem struct {
	SystemRun bool
}

func (as *AnotherSystem) Process() {
	as.SystemRun = true
}

type OneMoreSystem struct {
	SystemRun bool
}

func (oms *OneMoreSystem) Process() {
	oms.SystemRun = true
}

func TestAddSystem(t *testing.T) {
	controller := NewController()

	// Add two systems, with different priorities
	controller.AddSystem(&TestSystem{SystemRun: false}, 1)
	controller.AddSystem(&AnotherSystem{SystemRun: false}, 2)
	controller.AddSystem(&OneMoreSystem{SystemRun: false}, 1)

	assert.True(t, controller.HasSystem(reflect.TypeOf(&TestSystem{})))
	assert.True(t, controller.HasSystem(reflect.TypeOf(&AnotherSystem{})))

	// Make sure the ordering of systems by priority is correct
	assert.Equal(t, len(controller.sortedSystems[1]), 2)
	assert.Equal(t, len(controller.sortedSystems[2]), 1)
}

func TestProcessSystems(t *testing.T) {
	controller := NewController()

	system1 := &TestSystem{SystemRun: false}
	system2 := &AnotherSystem{SystemRun: false}
	system3 := &OneMoreSystem{SystemRun: false}

	controller.AddSystem(system1, 1)
	controller.AddSystem(system2, 2)
	controller.AddSystem(system3, 3)

	controller.Process([]reflect.Type{})

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.True(t, system3.SystemRun)

	// Create a new controller, and add all systems in the priority
	controller2 := NewController()

	system1 = &TestSystem{SystemRun: false}
	system2 = &AnotherSystem{SystemRun: false}
	system3 = &OneMoreSystem{SystemRun: false}

	controller2.AddSystem(system1, 1)
	controller2.AddSystem(system2, 1)
	controller2.AddSystem(system3, 1)

	controller2.Process([]reflect.Type{})

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.True(t, system3.SystemRun)

	// Create a new controller, and add all systems, but exclude one from processing
	controller3 := NewController()

	system1 = &TestSystem{SystemRun: false}
	system2 = &AnotherSystem{SystemRun: false}
	system3 = &OneMoreSystem{SystemRun: false}

	controller3.AddSystem(system1, 1)
	controller3.AddSystem(system2, 1)
	controller3.AddSystem(system3, 1)

	// Exclude OneMoreSystem systems from processing
	controller3.Process([]reflect.Type{reflect.TypeOf(&OneMoreSystem{})})

	assert.True(t, system1.SystemRun)
	assert.True(t, system2.SystemRun)
	assert.False(t, system3.SystemRun)

}

func TestProcessSingleSystem(t *testing.T) {
	controller := NewController()

	system1 := &TestSystem{SystemRun: false}
	system2 := &AnotherSystem{SystemRun: false}
	system3 := &OneMoreSystem{SystemRun: false}

	controller.AddSystem(system1, 1)
	controller.AddSystem(system2, 1)

	controller.ProcessSystem(reflect.TypeOf(&TestSystem{}))

	assert.True(t, system1.SystemRun)
	assert.False(t, system2.SystemRun)

	controller.ProcessSystem(reflect.TypeOf(&OneMoreSystem{}))
	assert.False(t, system3.SystemRun)
}
