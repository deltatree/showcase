// Particle Symphony - Native Desktop Application
//
// This is the main entry point for the native desktop version of Particle Symphony.
// It uses raylib-go for rendering and provides interactive particle effects
// driven by the andygeiss/ecs Entity-Component-System framework.
//
// Controls:
//   - Mouse: Move to guide particles
//   - Left Click: Attract particles
//   - Right Click: Repel particles
//   - Double-Click: Lock attract/repel mode
//   - 1-5: Switch between presets
//   - F3: Toggle debug overlay
//
// Run: go run main.go
package main

import (
	"log"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/internal/config"
	"github.com/deltatree/showcase/presets"
	"github.com/deltatree/showcase/systems"
)

func main() {
	// Load config
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Printf("Config not found, using defaults: %v", err)
		cfg = config.Default()
	}

	// Initialize ECS managers
	em := ecs.NewEntityManager()
	sm := ecs.NewSystemManager()

	// Create systems
	emitterSystem := systems.NewEmitterSystem(
		cfg.Particles.SpawnRate,
		cfg.Particles.MaxCount,
		float32(cfg.Window.Width),
		float32(cfg.Window.Height),
	)

	renderSystem := systems.NewRenderSystem(
		cfg.Window.Width,
		cfg.Window.Height,
		cfg.Window.Title,
	)

	// Connect particle slider to emitter
	renderSystem.SetMaxParticles(cfg.Particles.MaxCount)
	renderSystem.SetOnParticleChange(func(newMax int) {
		emitterSystem.SetMaxParticles(newMax)
	})

	// Preset switcher function
	currentPresetIndex := 0
	presetSwitcher := func(index int) {
		currentPresetIndex = index
		preset := presets.GetPreset(index)
		preset.Apply(em, cfg)
		renderSystem.SetPresetName(preset.Name())

		// Update emitter based on preset
		type presetWithConfig interface {
			EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int)
		}
		if p, ok := preset.(presetWithConfig); ok {
			sr, sg, sb, sa, er, eg, eb, ea, pattern, rate := p.EmitterConfig()
			emitterSystem.SetColors(sr, sg, sb, sa, er, eg, eb, ea)
			emitterSystem.SetSpawnPattern(pattern)
			emitterSystem.SetSpawnRate(rate)
		}
	}

	// Register systems in correct order
	sm.Add(
		systems.NewInputSystem(presetSwitcher),
		emitterSystem,
		systems.NewGravitySystem(),
		systems.NewPhysicsSystem(
			cfg.Physics.Damping,
			cfg.Physics.MaxVelocity,
			float32(cfg.Window.Width),
			float32(cfg.Window.Height),
		),
		systems.NewLifetimeSystem(),
		systems.NewColorSystem(),
		renderSystem,
	)

	// Apply default preset
	_ = currentPresetIndex
	presetSwitcher(0)

	// Create and run engine
	engine := ecs.NewDefaultEngine(em, sm)
	engine.Setup()
	defer engine.Teardown()
	engine.Run()
}
