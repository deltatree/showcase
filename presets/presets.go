// Package presets provides predefined particle configurations for Particle Symphony.
//
// Each preset defines a unique visual effect by configuring initial particle properties
// and emitter behavior. Presets implement the Preset interface and can be switched
// at runtime using keyboard shortcuts (1-5).
//
// # Available Presets
//
//   - Galaxy: Spiral galaxy with rotating particles
//   - Firework: Colorful explosion bursts with gravity
//   - Swarm: Organic swarm behavior following attractors
//   - Fountain: Water fountain shooting upward
//   - Chaos: Random particles with varied colors and velocities
//
// # Usage
//
// Get and apply a preset:
//
//	preset := presets.GetPreset(0) // Galaxy
//	preset.Apply(entityManager, config)
//
// # Custom Presets
//
// Implement the Preset interface to create custom effects:
//
//	type myPreset struct{}
//	func (p *myPreset) Name() string { return "MyPreset" }
//	func (p *myPreset) Description() string { return "Custom effect" }
//	func (p *myPreset) Apply(em ecs.EntityManager, cfg *config.Config) { ... }
package presets

import (
	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/internal/config"
)

// Preset defines the interface for particle presets.
type Preset interface {
	Name() string
	Description() string
	Apply(em ecs.EntityManager, cfg *config.Config)
}

// Registry holds all available presets.
var Registry = []Preset{
	NewGalaxyPreset(),
	NewFireworkPreset(),
	NewSwarmPreset(),
	NewFountainPreset(),
	NewChaosPreset(),
}

// GetPreset returns a preset by index.
func GetPreset(index int) Preset {
	if index < 0 || index >= len(Registry) {
		return Registry[0]
	}
	return Registry[index]
}

// GetPresetByName returns a preset by name.
func GetPresetByName(name string) Preset {
	for _, p := range Registry {
		if p.Name() == name {
			return p
		}
	}
	return Registry[0]
}

// ClearParticles removes all particle entities from the entity manager.
func ClearParticles(em ecs.EntityManager) {
	particles := em.FilterByMask(components.MaskParticle)
	for _, p := range particles {
		em.Remove(p)
	}
}
