package systems

import (
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/premium"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// GlowRenderer provides glow/bloom effects for particles.
type GlowRenderer struct {
	enabled bool
	passes  int
	palette premium.ColorPalette
}

// NewGlowRenderer creates a new glow renderer.
func NewGlowRenderer(enabled bool, passes int) *GlowRenderer {
	return &GlowRenderer{
		enabled: enabled,
		passes:  passes,
		palette: premium.GalaxyPalette,
	}
}

// RenderGlow draws glow layers behind a particle.
func (r *GlowRenderer) RenderGlow(
	pos *components.Position,
	col *components.Color,
	size *components.Size,
	shakeX, shakeY float32,
) {
	if !r.enabled || r.palette.GlowIntensity <= 0 {
		return
	}

	drawX := int32(pos.X + shakeX)
	drawY := int32(pos.Y + shakeY)

	// Calculate glow alpha based on particle alpha and palette intensity
	baseAlpha := float32(col.A) * r.palette.GlowIntensity * 0.25

	// Draw glow layers (largest first, then smaller)
	for i := r.passes; i >= 1; i-- {
		layerAlpha := uint8(baseAlpha / float32(i))
		layerSize := size.Radius * (1.5 + float32(i)*0.8)

		// Blend glow color with particle color for more natural look
		glowR := (uint16(r.palette.GlowR) + uint16(col.R)) / 2
		glowG := (uint16(r.palette.GlowG) + uint16(col.G)) / 2
		glowB := (uint16(r.palette.GlowB) + uint16(col.B)) / 2

		glowColor := rl.NewColor(uint8(glowR), uint8(glowG), uint8(glowB), layerAlpha)
		rl.DrawCircle(drawX, drawY, layerSize, glowColor)
	}
}

// SetEnabled enables or disables glow.
func (r *GlowRenderer) SetEnabled(enabled bool) {
	r.enabled = enabled
}

// SetPasses sets the number of glow passes.
func (r *GlowRenderer) SetPasses(passes int) {
	r.passes = passes
}

// SetPalette sets the color palette for glow colors.
func (r *GlowRenderer) SetPalette(palette premium.ColorPalette) {
	r.palette = palette
}

// ApplyQuality applies quality settings to glow rendering.
func (r *GlowRenderer) ApplyQuality(q premium.QualitySettings) {
	r.enabled = q.GlowEnabled
	r.passes = q.GlowPasses
}

// IsEnabled returns whether glow is enabled.
func (r *GlowRenderer) IsEnabled() bool {
	return r.enabled
}
