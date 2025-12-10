// Package premium provides premium visual and audio enhancements for Particle Symphony.
//
// This package implements Epic E-009: Premium Experience, including:
//   - Curated color palettes for each preset
//   - Quality settings (Low/Medium/High)
//   - Visual effects configuration
//
// # Color Palettes
//
// Each preset has a professionally designed color palette with:
//   - Primary gradient (start → end colors)
//   - Accent colors for variety
//   - Glow colors for bloom effects
//
// # Quality Presets
//
// Three quality levels optimize for different hardware:
//   - Low: Reduced particles, no effects
//   - Medium: Standard particles, simple glow
//   - High: Maximum particles, full effects
package premium

// ColorPalette defines a curated color scheme for a preset.
// Each palette includes primary, accent, and glow colors.
type ColorPalette struct {
	Name string
	// Primary gradient colors (start to end)
	StartR, StartG, StartB, StartA uint8
	EndR, EndG, EndB, EndA         uint8
	// Alternative gradient for variety
	AltStartR, AltStartG, AltStartB, AltStartA uint8
	AltEndR, AltEndG, AltEndB, AltEndA         uint8
	// Glow/bloom color overlay
	GlowR, GlowG, GlowB uint8
	GlowIntensity       float32
}

// GalaxyPalette - Cosmic purple/blue with star-white accents
// Evokes deep space with nebula-like colors
var GalaxyPalette = ColorPalette{
	Name: "Galaxy",
	// Primary: Pink-Magenta → Deep Purple
	StartR: 255, StartG: 107, StartB: 157, StartA: 255,
	EndR: 74, EndG: 0, EndB: 128, EndA: 0,
	// Alt: Cyan → Space Black
	AltStartR: 0, AltStartG: 212, AltStartB: 255, AltStartA: 255,
	AltEndR: 0, AltEndG: 0, AltEndB: 51, AltEndA: 0,
	// Glow: White star core
	GlowR: 255, GlowG: 255, GlowB: 255,
	GlowIntensity: 0.6,
}

// FireworkPalette - Vibrant celebration colors
// Gold, red, green, blue explosions with sparkle
var FireworkPalette = ColorPalette{
	Name: "Firework",
	// Primary: Gold → Orange
	StartR: 255, StartG: 215, StartB: 0, StartA: 255,
	EndR: 255, EndG: 69, EndB: 0, EndA: 0,
	// Alt: Red → Deep Red
	AltStartR: 255, AltStartG: 0, AltStartB: 68, AltStartA: 255,
	AltEndR: 136, AltEndG: 0, AltEndB: 34, AltEndA: 0,
	// Glow: Warm white
	GlowR: 255, GlowG: 240, GlowB: 200,
	GlowIntensity: 0.8,
}

// SwarmPalette - Bioluminescent organic colors
// Teal/green with warm orange accents
var SwarmPalette = ColorPalette{
	Name: "Swarm",
	// Primary: Bioluminescent Teal
	StartR: 0, StartG: 255, StartB: 170, StartA: 255,
	EndR: 0, EndG: 68, EndB: 51, EndA: 0,
	// Alt: Warm Orange
	AltStartR: 255, AltStartG: 136, AltStartB: 0, AltStartA: 255,
	AltEndR: 68, AltEndG: 34, AltEndB: 0, AltEndA: 0,
	// Glow: Soft green
	GlowR: 150, GlowG: 255, GlowB: 200,
	GlowIntensity: 0.5,
}

// FountainPalette - Water and spray colors
// Azure blue with white mist
var FountainPalette = ColorPalette{
	Name: "Fountain",
	// Primary: Azure → Deep Blue
	StartR: 0, StartG: 170, StartB: 255, StartA: 255,
	EndR: 0, EndG: 51, EndB: 102, EndA: 0,
	// Alt: White → Light Blue (spray)
	AltStartR: 255, AltStartG: 255, AltStartB: 255, AltStartA: 255,
	AltEndR: 136, AltEndG: 204, AltEndB: 255, AltEndA: 0,
	// Glow: Soft blue
	GlowR: 100, GlowG: 180, GlowB: 255,
	GlowIntensity: 0.4,
}

// ChaosPalette - Electric neon chaos
// Magenta/Cyan with fire accents
var ChaosPalette = ColorPalette{
	Name: "Chaos",
	// Primary: Electric Magenta → Cyan
	StartR: 255, StartG: 0, StartB: 255, StartA: 255,
	EndR: 0, EndG: 255, EndB: 255, EndA: 0,
	// Alt: Fire Yellow → Red
	AltStartR: 255, AltStartG: 255, AltStartB: 0, AltStartA: 255,
	AltEndR: 255, AltEndG: 0, AltEndB: 0, AltEndA: 0,
	// Glow: Hot white
	GlowR: 255, GlowG: 220, GlowB: 255,
	GlowIntensity: 0.9,
}

// Palettes maps preset names to their color palettes.
var Palettes = map[string]ColorPalette{
	"Galaxy":   GalaxyPalette,
	"Firework": FireworkPalette,
	"Swarm":    SwarmPalette,
	"Fountain": FountainPalette,
	"Chaos":    ChaosPalette,
}

// GetPalette returns the color palette for a preset name.
// Returns GalaxyPalette as default if not found.
func GetPalette(name string) ColorPalette {
	if p, ok := Palettes[name]; ok {
		return p
	}
	return GalaxyPalette
}
