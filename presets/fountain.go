package presets

import (
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
	"github.com/deltatree/showcase/premium"
)

// FountainPreset creates a water fountain shooting upwards with gravity.
// Particles spawn at the bottom center with upward velocity and fall
// under gravity, creating a realistic fountain arc. Blue-tinted particles
// simulate water droplets.
//
// Keyboard: Press 4 to activate this preset.
type fountainPreset struct {
	palette premium.ColorPalette
}

// NewFountainPreset creates a new fountain preset instance.
func NewFountainPreset() Preset {
	return &fountainPreset{
		palette: premium.FountainPalette,
	}
}

func (p *fountainPreset) Name() string { return "Fountain" }

func (p *fountainPreset) Description() string {
	return "Water fountain shooting upwards with gravity"
}

// Palette returns the premium color palette for this preset.
func (p *fountainPreset) Palette() premium.ColorPalette {
	return p.palette
}

func (p *fountainPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	centerX := float32(cfg.Window.Width) / 2
	bottomY := float32(cfg.Window.Height) - 50
	pal := p.palette

	numParticles := 300
	for i := 0; i < numParticles; i++ {
		x := centerX + (rand.Float32()-0.5)*20
		y := bottomY

		vx := (rand.Float32() - 0.5) * 80
		vy := -200 - rand.Float32()*150

		// Use premium palette - water blue with occasional white spray
		useAlt := rand.Float32() < 0.2
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
			components.NewAcceleration().WithY(150),
			components.NewColor().WithGradient(sr, sg, sb, sa, er, eg, eb, ea),
			components.NewLifetime().WithTTL(2.0 + rand.Float32()*2.0),
			components.NewSize().WithRadius(3.0 + rand.Float32()*2.0).WithEndSize(1.0),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *fountainPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	pal := p.palette
	return pal.StartR, pal.StartG, pal.StartB, pal.StartA,
		pal.EndR, pal.EndG, pal.EndB, pal.EndA, "center", 100
}
