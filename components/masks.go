// Package components provides all ECS component definitions for the Particle Symphony simulation.
//
// Components are pure data containers following the Entity-Component-System pattern.
// Each component implements the ecs.Component interface with a Mask() method that returns
// a unique bitmask for efficient entity filtering.
//
// # Component Types
//
// Position, Velocity, and Acceleration form the physics foundation.
// Color and Size handle visual representation with gradient interpolation.
// Lifetime manages particle aging and automatic cleanup.
// Mass enables gravitational interactions.
// Particle, Emitter, and Attractor are tag components for entity classification.
//
// # Usage
//
// Components are typically created using their New* constructors with method chaining:
//
//	pos := components.NewPosition().With(100, 200)
//	vel := components.NewVelocity().With(10, -5)
//	col := components.NewColor().WithRGBA(255, 128, 0, 255)
//
// # Bitmask Filtering
//
// Use component masks to efficiently query entities:
//
//	entities := em.FilterByMask(components.MaskPosition | components.MaskVelocity)
package components

// Component masks for efficient entity filtering using bitmasks.
const (
	MaskPosition     = uint64(1 << 0)
	MaskVelocity     = uint64(1 << 1)
	MaskAcceleration = uint64(1 << 2)
	MaskColor        = uint64(1 << 3)
	MaskLifetime     = uint64(1 << 4)
	MaskMass         = uint64(1 << 5)
	MaskSize         = uint64(1 << 6)
	MaskEmitter      = uint64(1 << 7)
	MaskAttractor    = uint64(1 << 8)
	MaskParticle     = uint64(1 << 9)
)

// Composite masks for common component combinations.
const (
	MaskMovable      = MaskPosition | MaskVelocity
	MaskPhysics      = MaskMovable | MaskAcceleration
	MaskRenderable   = MaskPosition | MaskColor | MaskSize
	MaskFullParticle = MaskPhysics | MaskColor | MaskLifetime | MaskSize | MaskParticle
)
