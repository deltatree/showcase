package systems

import (
	"math"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
)

// gravitySystem applies gravitational forces from attractors to particles.
type gravitySystem struct{}

// NewGravitySystem creates a new gravity system.
func NewGravitySystem() ecs.System {
	return &gravitySystem{}
}

func (s *gravitySystem) Setup() {}

func (s *gravitySystem) Process(em ecs.EntityManager) (state int) {
	attractors := em.FilterByMask(components.MaskPosition | components.MaskMass | components.MaskAttractor)
	particles := em.FilterByMask(components.MaskPosition | components.MaskAcceleration | components.MaskParticle)

	for _, particle := range particles {
		pPos := particle.Get(components.MaskPosition).(*components.Position)
		pAcc := particle.Get(components.MaskAcceleration).(*components.Acceleration)

		pAcc.Reset()

		for _, attractor := range attractors {
			aPos := attractor.Get(components.MaskPosition).(*components.Position)
			aMass := attractor.Get(components.MaskMass).(*components.Mass)

			if aMass.Value == 0 {
				continue
			}

			dx := aPos.X - pPos.X
			dy := aPos.Y - pPos.Y

			dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
			if dist < 10 {
				dist = 10
			}

			force := aMass.Value / (dist * dist) * 500

			pAcc.Add(dx/dist*force, dy/dist*force)
		}
	}

	return ecs.StateEngineContinue
}

func (s *gravitySystem) Teardown() {}
