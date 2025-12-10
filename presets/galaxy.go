package presets

import (
	"math"
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
	"github.com/deltatree/showcase/premium"
)

// GalaxyPreset creates a spiral galaxy simulation with orbital particle motion.
// Particles are arranged in a spiral pattern with tangential velocities,
// creating a rotating galaxy effect. Works beautifully with a central attractor.
//
// Keyboard: Press 1 to activate this preset.
type galaxyPreset struct {
	palette premium.ColorPalette
}

// NewGalaxyPreset creates a new galaxy preset instance.
func NewGalaxyPreset() Preset {
	return &galaxyPreset{
		palette: premium.GalaxyPalette,
	}
}

func (p *galaxyPreset) Name() string { return "Galaxy" }

func (p *galaxyPreset) Description() string {
	return "Spiral galaxy with central attractor"
}

// Palette returns the premium color palette for this preset.
func (p *galaxyPreset) Palette() premium.ColorPalette {
	return p.palette
}

func (p *galaxyPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	centerX := float32(cfg.Window.Width) / 2
	centerY := float32(cfg.Window.Height) / 2
	pal := p.palette

	numParticles := 500
	for i := 0; i < numParticles; i++ {
		angle := float32(i) * 0.1
		radius := float32(i) * 0.5
		x := centerX + radius*float32(math.Cos(float64(angle)))
		y := centerY + radius*float32(math.Sin(float64(angle)))

		speed := float32(30.0 + rand.Float64()*20.0)
		vx := -float32(math.Sin(float64(angle))) * speed
		vy := float32(math.Cos(float64(angle))) * speed

		// Use premium palette colors with slight variation
		useAlt := rand.Float32() < 0.3
		var sr, sg, sb, sa, er, eg, eb, ea uint8
		if useAlt {
			sr, sg, sb, sa = pal.AltStartR, pal.AltStartG, pal.AltStartB, pal.AltStartA
			er, eg, eb, ea = pal.AltEndR, pal.AltEndG, pal.AltEndB, pal.AltEndA
		} else {
			sr, sg, sb, sa = pal.StartR, pal.StartG, pal.StartB, pal.StartA
			er, eg, eb, ea = pal.EndR, pal.EndG, pal.EndB, pal.EndA
		}

		em.Add(ecs.NewEntity("", []ecs.Component{
			components.NewPosition().With(x, y),
			components.NewVelocity().With(vx, vy),
			components.NewAcceleration(),
			components.NewColor().WithGradient(sr, sg, sb, sa, er, eg, eb, ea),
			components.NewLifetime().WithTTL(8.0 + rand.Float32()*4.0),
			components.NewSize().WithRadius(2.0 + rand.Float32()*2.0).WithEndSize(0.5),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *galaxyPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	pal := p.palette
	return pal.StartR, pal.StartG, pal.StartB, pal.StartA,
		pal.EndR, pal.EndG, pal.EndB, pal.EndA, "center", 50
}
