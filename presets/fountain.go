package presets

import (
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
)

// FountainPreset creates a water fountain shooting upwards with gravity.
// Particles spawn at the bottom center with upward velocity and fall
// under gravity, creating a realistic fountain arc. Blue-tinted particles
// simulate water droplets.
//
// Keyboard: Press 4 to activate this preset.
type fountainPreset struct{}

// NewFountainPreset creates a new fountain preset instance.
func NewFountainPreset() Preset {
	return &fountainPreset{}
}

func (p *fountainPreset) Name() string { return "Fountain" }

func (p *fountainPreset) Description() string {
	return "Water fountain shooting upwards with gravity"
}

func (p *fountainPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	centerX := float32(cfg.Window.Width) / 2
	bottomY := float32(cfg.Window.Height) - 50

	numParticles := 300
	for i := 0; i < numParticles; i++ {
		x := centerX + (rand.Float32()-0.5)*20
		y := bottomY

		vx := (rand.Float32() - 0.5) * 80
		vy := -200 - rand.Float32()*150

		em.Add(ecs.NewEntity("", []ecs.Component{
			components.NewPosition().With(x, y),
			components.NewVelocity().With(vx, vy),
			components.NewAcceleration().WithY(150),
			components.NewColor().WithGradient(100, 200, 255, 255, 50, 100, 200, 0),
			components.NewLifetime().WithTTL(2.0 + rand.Float32()*2.0),
			components.NewSize().WithRadius(3.0 + rand.Float32()*2.0).WithEndSize(1.0),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *fountainPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	return 100, 200, 255, 255, 50, 100, 200, 0, "center", 100
}
