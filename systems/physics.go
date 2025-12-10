// Package systems provides all ECS system implementations for Particle Symphony.
//
// Systems contain the logic that operates on entities with specific component combinations.
// Each system implements the ecs.System interface with Setup(), Process(), and Teardown() methods.
//
// # System Execution Order
//
// Systems are executed in registration order each frame:
//   1. InputSystem - handles mouse/keyboard input
//   2. EmitterSystem - spawns new particles
//   3. GravitySystem - applies attractor forces
//   4. PhysicsSystem - updates positions and velocities
//   5. LifetimeSystem - ages and removes expired entities
//   6. ColorSystem - interpolates colors and sizes
//   7. RenderSystem - draws entities to screen
//
// # Creating Custom Systems
//
// Implement the ecs.System interface:
//
//	type mySystem struct{}
//	func (s *mySystem) Setup() {}
//	func (s *mySystem) Process(em ecs.EntityManager) int { return ecs.StateEngineContinue }
//	func (s *mySystem) Teardown() {}
package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// PhysicsSystem handles movement, velocity, and screen wrapping for entities.
// It integrates acceleration into velocity, applies damping to simulate friction,
// clamps velocity to a maximum, and updates positions based on velocity.
//
// When entities move off-screen, they wrap around to the opposite edge,
// creating a toroidal topology.
type physicsSystem struct {
	damping     float32
	maxVelocity float32
	width       float32
	height      float32
}

// NewPhysicsSystem creates a new physics system with configurable parameters.
//
// Parameters:
//   - damping: velocity multiplier per frame (0.99 = 1% friction)
//   - maxVelocity: maximum speed in pixels per second
//   - width, height: screen dimensions for edge wrapping
func NewPhysicsSystem(damping, maxVelocity, width, height float32) ecs.System {
	return &physicsSystem{
		damping:     damping,
		maxVelocity: maxVelocity,
		width:       width,
		height:      height,
	}
}

func (s *physicsSystem) Setup() {}

func (s *physicsSystem) Process(em ecs.EntityManager) (state int) {
	dt := rl.GetFrameTime()

	entities := em.FilterByMask(components.MaskPosition | components.MaskVelocity)

	for _, e := range entities {
		pos := e.Get(components.MaskPosition).(*components.Position)
		vel := e.Get(components.MaskVelocity).(*components.Velocity)

		if acc := e.Get(components.MaskAcceleration); acc != nil {
			a := acc.(*components.Acceleration)
			vel.X += a.X * dt
			vel.Y += a.Y * dt
		}

		vel.X *= s.damping
		vel.Y *= s.damping

		mag := vel.Magnitude()
		if mag > s.maxVelocity {
			vel.X = vel.X / mag * s.maxVelocity
			vel.Y = vel.Y / mag * s.maxVelocity
		}

		pos.X += vel.X * dt
		pos.Y += vel.Y * dt

		if pos.X < 0 {
			pos.X = s.width
		}
		if pos.X > s.width {
			pos.X = 0
		}
		if pos.Y < 0 {
			pos.Y = s.height
		}
		if pos.Y > s.height {
			pos.Y = 0
		}
	}

	return ecs.StateEngineContinue
}

func (s *physicsSystem) Teardown() {}
