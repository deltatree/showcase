package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
)

// colorSystem interpolates colors and sizes based on lifetime.
type colorSystem struct{}

// NewColorSystem creates a new color system.
func NewColorSystem() ecs.System {
	return &colorSystem{}
}

func (s *colorSystem) Setup() {}

func (s *colorSystem) Process(em ecs.EntityManager) (state int) {
	entities := em.FilterByMask(components.MaskColor | components.MaskLifetime)

	for _, e := range entities {
		col := e.Get(components.MaskColor).(*components.Color)
		life := e.Get(components.MaskLifetime).(*components.Lifetime)

		t := life.Progress()

		col.R = lerp(col.StartR, col.EndR, t)
		col.G = lerp(col.StartG, col.EndG, t)
		col.B = lerp(col.StartB, col.EndB, t)
		col.A = lerp(col.StartA, col.EndA, t)
	}

	sizeEntities := em.FilterByMask(components.MaskSize | components.MaskLifetime)
	for _, e := range sizeEntities {
		size := e.Get(components.MaskSize).(*components.Size)
		life := e.Get(components.MaskLifetime).(*components.Lifetime)

		t := life.Progress()
		size.Radius = lerpF(size.StartSize, size.EndSize, t)
	}

	return ecs.StateEngineContinue
}

func lerp(a, b uint8, t float32) uint8 {
	return uint8(float32(a) + (float32(b)-float32(a))*t)
}

func lerpF(a, b, t float32) float32 {
	return a + (b-a)*t
}

func (s *colorSystem) Teardown() {}
