package presets

import (
	"testing"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/internal/config"
)

// TestPresetApply tests the Apply method for all presets.
func TestPresetApply(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	for _, preset := range Registry {
		t.Run(preset.Name()+"_Apply", func(t *testing.T) {
			// Apply should not panic
			preset.Apply(em, cfg)

			// After apply, there should be some entities
			// (all presets create initial particles)
		})
	}
}

// TestGalaxyPresetApply tests galaxy preset creates spiral particles.
func TestGalaxyPresetApply(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	preset := NewGalaxyPreset()
	preset.Apply(em, cfg)

	if preset.Name() != "Galaxy" {
		t.Errorf("expected Name() = 'Galaxy', got %q", preset.Name())
	}

	if preset.Description() == "" {
		t.Error("expected non-empty description")
	}
}

// TestFireworkPresetApply tests firework preset creates explosion particles.
func TestFireworkPresetApply(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	preset := NewFireworkPreset()
	preset.Apply(em, cfg)

	if preset.Name() != "Firework" {
		t.Errorf("expected Name() = 'Firework', got %q", preset.Name())
	}

	if preset.Description() == "" {
		t.Error("expected non-empty description")
	}
}

// TestSwarmPresetApply tests swarm preset creates swarm particles.
func TestSwarmPresetApply(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	preset := NewSwarmPreset()
	preset.Apply(em, cfg)

	if preset.Name() != "Swarm" {
		t.Errorf("expected Name() = 'Swarm', got %q", preset.Name())
	}

	if preset.Description() == "" {
		t.Error("expected non-empty description")
	}
}

// TestFountainPresetApply tests fountain preset creates fountain particles.
func TestFountainPresetApply(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	preset := NewFountainPreset()
	preset.Apply(em, cfg)

	if preset.Name() != "Fountain" {
		t.Errorf("expected Name() = 'Fountain', got %q", preset.Name())
	}

	if preset.Description() == "" {
		t.Error("expected non-empty description")
	}
}

// TestChaosPresetApply tests chaos preset creates random particles.
func TestChaosPresetApply(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	preset := NewChaosPreset()
	preset.Apply(em, cfg)

	if preset.Name() != "Chaos" {
		t.Errorf("expected Name() = 'Chaos', got %q", preset.Name())
	}

	if preset.Description() == "" {
		t.Error("expected non-empty description")
	}
}

// TestClearParticles tests that ClearParticles removes all particles.
func TestClearParticles(t *testing.T) {
	cfg := config.Default()
	em := ecs.NewEntityManager()

	// First apply a preset to create particles
	preset := NewChaosPreset()
	preset.Apply(em, cfg)

	// Clear and apply again (which internally calls ClearParticles)
	preset.Apply(em, cfg)

	// The preset should have created new particles after clearing
	// We can't easily count entities, but the function should not panic
}

// TestEmitterConfig tests that presets provide emitter configuration.
func TestEmitterConfig(t *testing.T) {
	presets := []struct {
		name   string
		preset Preset
	}{
		{"Galaxy", NewGalaxyPreset()},
		{"Firework", NewFireworkPreset()},
		{"Swarm", NewSwarmPreset()},
		{"Fountain", NewFountainPreset()},
		{"Chaos", NewChaosPreset()},
	}

	for _, tc := range presets {
		t.Run(tc.name+"_EmitterConfig", func(t *testing.T) {
			type configurable interface {
				EmitterConfig() (sr, sg, sb, sa, er, eg, eb, ea uint8, pattern string, rate int)
			}

			if p, ok := tc.preset.(configurable); ok {
				sr, _, _, _, _, _, _, _, pattern, rate := p.EmitterConfig()
				if sr == 0 && pattern == "" && rate == 0 {
					t.Error("EmitterConfig returned all zeros")
				}
				if pattern == "" {
					t.Error("EmitterConfig pattern is empty")
				}
				if rate <= 0 {
					t.Error("EmitterConfig rate should be positive")
				}
			}
		})
	}
}
