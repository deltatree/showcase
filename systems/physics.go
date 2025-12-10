package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// physicsSystem handles movement, velocity, and screen wrapping.
type physicsSystem struct {
	damping     float32
	maxVelocity float32
	width       float32
	height      float32
}

// NewPhysicsSystem creates a new physics system.
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
