package presets

import (
	"math"
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
)

// galaxyPreset creates a spiral galaxy simulation.
type galaxyPreset struct{}

// NewGalaxyPreset creates a new galaxy preset.
func NewGalaxyPreset() Preset {
	return &galaxyPreset{}
}

func (p *galaxyPreset) Name() string { return "Galaxy" }

func (p *galaxyPreset) Description() string {
	return "Spiral galaxy with central attractor"
}

func (p *galaxyPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	centerX := float32(cfg.Window.Width) / 2
	centerY := float32(cfg.Window.Height) / 2

	numParticles := 500
	for i := 0; i < numParticles; i++ {
		angle := float32(i) * 0.1
		radius := float32(i) * 0.5
		x := centerX + radius*float32(math.Cos(float64(angle)))
		y := centerY + radius*float32(math.Sin(float64(angle)))

		speed := float32(30.0 + rand.Float64()*20.0)
		vx := -float32(math.Sin(float64(angle))) * speed
		vy := float32(math.Cos(float64(angle))) * speed

		em.Add(ecs.NewEntity("", []ecs.Component{
			components.NewPosition().With(x, y),
			components.NewVelocity().With(vx, vy),
			components.NewAcceleration(),
			components.NewColor().WithGradient(100, 150, 255, 255, 255, 255, 255, 0),
			components.NewLifetime().WithTTL(8.0 + rand.Float32()*4.0),
			components.NewSize().WithRadius(2.0 + rand.Float32()*2.0).WithEndSize(0.5),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *galaxyPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	return 100, 150, 255, 255, 255, 255, 255, 0, "center", 50
}
