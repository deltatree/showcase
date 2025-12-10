package presets

import (
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
)

// chaosPreset creates chaotic random particle movement.
type chaosPreset struct{}

// NewChaosPreset creates a new chaos preset.
func NewChaosPreset() Preset {
	return &chaosPreset{}
}

func (p *chaosPreset) Name() string { return "Chaos" }

func (p *chaosPreset) Description() string {
	return "Chaotic particles with random colors and velocities"
}

func (p *chaosPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	width := float32(cfg.Window.Width)
	height := float32(cfg.Window.Height)

	numParticles := 1000
	for i := 0; i < numParticles; i++ {
		x := rand.Float32() * width
		y := rand.Float32() * height

		vx := (rand.Float32() - 0.5) * 300
		vy := (rand.Float32() - 0.5) * 300

		r := uint8(rand.Intn(256))
		g := uint8(rand.Intn(256))
		b := uint8(rand.Intn(256))

		em.Add(ecs.NewEntity("", []ecs.Component{
			components.NewPosition().With(x, y),
			components.NewVelocity().With(vx, vy),
			components.NewAcceleration(),
			components.NewColor().WithGradient(r, g, b, 255, r, g, b, 0),
			components.NewLifetime().WithTTL(3.0 + rand.Float32()*4.0),
			components.NewSize().WithRadius(1.0 + rand.Float32()*4.0).WithEndSize(0.5),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *chaosPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	return 255, 100, 100, 255, 100, 100, 255, 0, "random", 150
}
