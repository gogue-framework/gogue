package ecs

import (
	"reflect"
)

// Component is a metadata container used to house information about something an entities state.
// Example Components might look like:
// type PositionComponent struct {
//     X int
//     Y int
// }
// This position component represents where an entity is in the game world. If an entity does not have a position
// component, it can be assumed they are not present in the world.
// Another type of component might look like:
// type CanAttackComponent {}
// CanAttackComponent has no data attached, and acts merely as a flag. If an entity has this component, they can attack
// if an entity is missing this component, they cannot attack.
// Components are a flexible way of attaching metadata to an entity.
type Component interface {
	TypeOf() reflect.Type
}
