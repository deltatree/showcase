package presets

import (
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
)

// SwarmPreset creates organic swarm-like behavior following the mouse.
// Particles are initialized near the center with random velocities,
// creating a cohesive swarm that responds to attractor forces.
// Green-tinted particles suggest organic, bioluminescent creatures.
//
// Keyboard: Press 3 to activate this preset.
type swarmPreset struct{}

// NewSwarmPreset creates a new swarm preset instance.
func NewSwarmPreset() Preset {
	return &swarmPreset{}
}

func (p *swarmPreset) Name() string { return "Swarm" }

func (p *swarmPreset) Description() string {
	return "Organic swarm behavior following the mouse"
}

func (p *swarmPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	centerX := float32(cfg.Window.Width) / 2
	centerY := float32(cfg.Window.Height) / 2

	numParticles := 800
	for i := 0; i < numParticles; i++ {
		x := centerX + (rand.Float32()-0.5)*200
		y := centerY + (rand.Float32()-0.5)*200

		vx := (rand.Float32() - 0.5) * 50
		vy := (rand.Float32() - 0.5) * 50

		em.Add(ecs.NewEntity("", []ecs.Component{
			components.NewPosition().With(x, y),
			components.NewVelocity().With(vx, vy),
			components.NewAcceleration(),
			components.NewColor().WithGradient(50, 255, 150, 255, 100, 200, 100, 0),
			components.NewLifetime().WithTTL(10.0 + rand.Float32()*5.0),
			components.NewSize().WithRadius(3.0 + rand.Float32()*2.0).WithEndSize(1.0),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *swarmPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	return 50, 255, 150, 255, 100, 200, 100, 0, "center", 30
}
