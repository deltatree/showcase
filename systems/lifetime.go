package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// lifetimeSystem ages particles and removes expired ones.
type lifetimeSystem struct{}

// NewLifetimeSystem creates a new lifetime system.
func NewLifetimeSystem() ecs.System {
	return &lifetimeSystem{}
}

func (s *lifetimeSystem) Setup() {}

func (s *lifetimeSystem) Process(em ecs.EntityManager) (state int) {
	dt := rl.GetFrameTime()

	entities := em.FilterByMask(components.MaskLifetime)

	var toRemove []*ecs.Entity

	for _, e := range entities {
		life := e.Get(components.MaskLifetime).(*components.Lifetime)
		life.Age += dt

		if life.Age >= life.TTL {
			life.Expired = true
			toRemove = append(toRemove, e)
		}
	}

	for _, entity := range toRemove {
		em.Remove(entity)
	}

	return ecs.StateEngineContinue
}

func (s *lifetimeSystem) Teardown() {}
