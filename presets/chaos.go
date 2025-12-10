package presets

import (
	"math/rand"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
	"github.com/deltatree/showcase/premium"
)

// ChaosPreset creates chaotic random particle movement across the screen.
// Particles spawn at random positions with random velocities and colors,
// creating a vibrant, unpredictable visual effect. High particle count
// and fast movement make this preset visually intense.
//
// Keyboard: Press 5 to activate this preset.
type chaosPreset struct {
	palette premium.ColorPalette
}

// NewChaosPreset creates a new chaos preset instance.
func NewChaosPreset() Preset {
	return &chaosPreset{
		palette: premium.ChaosPalette,
	}
}

func (p *chaosPreset) Name() string { return "Chaos" }

func (p *chaosPreset) Description() string {
	return "Chaotic particles with random colors and velocities"
}

// Palette returns the premium color palette for this preset.
func (p *chaosPreset) Palette() premium.ColorPalette {
	return p.palette
}

func (p *chaosPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
	ClearParticles(em)

	width := float32(cfg.Window.Width)
	height := float32(cfg.Window.Height)
	pal := p.palette

	numParticles := 1000
	for i := 0; i < numParticles; i++ {
		x := rand.Float32() * width
		y := rand.Float32() * height

		vx := (rand.Float32() - 0.5) * 300
		vy := (rand.Float32() - 0.5) * 300

		// Premium palette with electric neon chaos
		choice := rand.Float32()
		var sr, sg, sb, sa, er, eg, eb, ea uint8
		if choice < 0.4 {
			// Primary: Electric Magenta → Cyan
			sr, sg, sb, sa = pal.StartR, pal.StartG, pal.StartB, pal.StartA
			er, eg, eb, ea = pal.EndR, pal.EndG, pal.EndB, pal.EndA
		} else if choice < 0.8 {
			// Alt: Fire Yellow → Red
			sr, sg, sb, sa = pal.AltStartR, pal.AltStartG, pal.AltStartB, pal.AltStartA
			er, eg, eb, ea = pal.AltEndR, pal.AltEndG, pal.AltEndB, pal.AltEndA
		} else {
			// Random neon for extra chaos
			sr = uint8(128 + rand.Intn(128))
			sg = uint8(rand.Intn(256))
			sb = uint8(128 + rand.Intn(128))
			sa = 255
			er, eg, eb, ea = sr, sg, sb, 0
		}

		em.Add(ecs.NewEntity("", []ecs.Component{
			components.NewPosition().With(x, y),
			components.NewVelocity().With(vx, vy),
			components.NewAcceleration(),
			components.NewColor().WithGradient(sr, sg, sb, sa, er, eg, eb, ea),
			components.NewLifetime().WithTTL(3.0 + rand.Float32()*4.0),
			components.NewSize().WithRadius(1.0 + rand.Float32()*4.0).WithEndSize(0.5),
			components.NewParticle(),
		}))
	}
}

// EmitterConfig returns emitter settings for this preset.
func (p *chaosPreset) EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int) {
	pal := p.palette
	return pal.StartR, pal.StartG, pal.StartB, pal.StartA,
		pal.EndR, pal.EndG, pal.EndB, pal.EndA, "random", 150
}
