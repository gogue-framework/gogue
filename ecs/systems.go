package ecs

// System operates on entities to produce a result. A system is a method of modifying the game world and its contained
// entities in a manner based on the state of the entities (read: what components that have, and what data those
// components contain). A system might move all monster entities that have a MovementComponent, or calcualte poison
// damage for all entities that a PoisonedComponent. They can also be used to allow the player to move: if the player
// has a movement component, the system should accept movement input, and adjust the players position component
// accordingly. If the player is missing a movement component (say, they have been paralyzed), the system should not
// accept input from the user, and should skip the users turn entirely.
// Each system has a Process method that contains the system logic. This Process method will be called when the system
// is processed by the Controller.
type System interface {
	Process()
}
