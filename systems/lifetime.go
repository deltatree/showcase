package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// LifetimeSystem ages particles and removes expired ones from the world.
// It increments each entity's Age by delta time each frame and marks
// entities as Expired when Age >= TTL. Expired entities are immediately
// removed from the EntityManager.
//
// This system enables particle effects with finite durations, preventing
// unbounded entity accumulation and enabling effects like fading trails.
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
