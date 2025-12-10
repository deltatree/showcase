package presets

import (
	"math"
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
	"github.com/deltatree/showcase/premium"
)

// FireworkPreset creates colorful firework explosions with gravity.
// Multiple explosion bursts spawn with randomly colored particles that
// explode outward and fall under gravity, simulating real fireworks.
//
// Keyboard: Press 2 to activate this preset.
type fireworkPreset struct {
	palette premium.ColorPalette
}

// NewFireworkPreset creates a new firework preset instance.
func NewFireworkPreset() Preset {
	return &fireworkPreset{
		palette: premium.FireworkPalette,
	}
}

func (p *fireworkPreset) Name() string { return "Firework" }

func (p *fireworkPreset) Description() string {
	return "Colorful firework explosions with gravity"
}

// Palette returns the premium color palette for this preset.
func (p *fireworkPreset) Palette() premium.ColorPalette {
	return p.palette
}

func (p *fireworkPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	width := float32(cfg.Window.Width)
	height := float32(cfg.Window.Height)
	pal := p.palette

	numExplosions := 5
	for e := 0; e < numExplosions; e++ {
		explosionX := rand.Float32() * width
		explosionY := height*0.2 + rand.Float32()*height*0.4

		// Premium palette colors with variation
		colors := []struct{ r, g, b uint8 }{
			{pal.StartR, pal.StartG, pal.StartB},          // Gold
			{pal.AltStartR, pal.AltStartG, pal.AltStartB}, // Red
			{50, 255, 100},  // Green sparkle
			{100, 180, 255}, // Blue sparkle
			{255, 255, 255}, // White sparkle
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
	pal := p.palette
	return pal.StartR, pal.StartG, pal.StartB, pal.StartA,
		pal.EndR, pal.EndG, pal.EndB, pal.EndA, "edges", 20
}
