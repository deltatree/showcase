package presets

import (
	"math"
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
)

// fireworkPreset creates firework-like explosions.
type fireworkPreset struct{}

// NewFireworkPreset creates a new firework preset.
func NewFireworkPreset() Preset {
	return &fireworkPreset{}
}

func (p *fireworkPreset) Name() string { return "Firework" }

func (p *fireworkPreset) Description() string {
	return "Colorful firework explosions with gravity"
}

func (p *fireworkPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	width := float32(cfg.Window.Width)
	height := float32(cfg.Window.Height)

	numExplosions := 5
	for e := 0; e < numExplosions; e++ {
		explosionX := rand.Float32() * width
		explosionY := height*0.2 + rand.Float32()*height*0.4

		colors := []struct{ r, g, b uint8 }{
			{255, 50, 50},
			{255, 200, 50},
			{50, 255, 50},
			{50, 150, 255},
			{255, 50, 255},
		}
		c := colors[rand.Intn(len(colors))]

		numParticles := 100
		for i := 0; i < numParticles; i++ {
			angle := rand.Float32() * 2 * math.Pi
			speed := 50.0 + rand.Float32()*150.0
			vx := float32(math.Cos(float64(angle))) * speed
			vy := float32(math.Sin(float64(angle)))*speed - 50

			em.Add(ecs.NewEntity("", []ecs.Component{
				components.NewPosition().With(explosionX, explosionY),
				components.NewVelocity().With(vx, vy),
				components.NewAcceleration().WithY(100),
				components.NewColor().WithGradient(c.r, c.g, c.b, 255, c.r, c.g, c.b, 0),
				components.NewLifetime().WithTTL(1.5 + rand.Float32()*1.5),
				components.NewSize().WithRadius(2.0 + rand.Float32()*3.0).WithEndSize(0.5),
				components.NewParticle(),
			}))
		}
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *fireworkPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	return 255, 200, 50, 255, 255, 100, 50, 0, "edges", 20
}
