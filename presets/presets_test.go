package presets

import (
	"testing"
)

func TestRegistryContainsPresets(t *testing.T) {
	if len(Registry) == 0 {
		t.Error("Registry is empty, expected presets")
	}
}

func TestGetPreset(t *testing.T) {
	preset := GetPreset(0)
	if preset == nil {
		t.Error("GetPreset(0) returned nil")
	}

	outOfRange := GetPreset(9999)
	if outOfRange == nil {
		t.Error("GetPreset(9999) returned nil, expected fallback to first preset")
	}

	negative := GetPreset(-1)
	if negative == nil {
		t.Error("GetPreset(-1) returned nil, expected fallback to first preset")
	}
}

func TestGetPresetByName(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Galaxy", "Galaxy"},
		{"Firework", "Firework"},
		{"Swarm", "Swarm"},
		{"Fountain", "Fountain"},
		{"Chaos", "Chaos"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preset := GetPresetByName(tt.name)
			if preset == nil {
				t.Errorf("GetPresetByName(%q) returned nil", tt.name)
				return
			}
			if preset.Name() != tt.expected {
				t.Errorf("GetPresetByName(%q).Name() = %q, want %q", tt.name, preset.Name(), tt.expected)
			}
		})
	}
}

func TestGetPresetByName_Unknown(t *testing.T) {
	preset := GetPresetByName("nonexistent")
	if preset == nil {
		t.Error("GetPresetByName(nonexistent) returned nil, expected fallback")
	}
}

func TestPresetInterface(t *testing.T) {
	for i, preset := range Registry {
		t.Run(preset.Name(), func(t *testing.T) {
			if preset.Name() == "" {
				t.Errorf("Registry[%d].Name() is empty", i)
			}
			if preset.Description() == "" {
				t.Errorf("Registry[%d].Description() is empty", i)
			}
		})
	}
}

func TestExpectedPresetsExist(t *testing.T) {
	expectedNames := []string{"Galaxy", "Firework", "Swarm", "Fountain", "Chaos"}
	foundNames := make(map[string]bool)

	for _, preset := range Registry {
		foundNames[preset.Name()] = true
	}

	for _, name := range expectedNames {
		if !foundNames[name] {
			t.Errorf("Expected preset %q not found in Registry", name)
		}
	}
}

// TestPresetPalette tests Palette method for all premium presets.
func TestPresetPalette(t *testing.T) {
	// Test Galaxy preset palette
	galaxy := NewGalaxyPreset().(*galaxyPreset)
	palette := galaxy.Palette()
	if palette.Name == "" {
		t.Error("Galaxy.Palette().Name is empty")
	}
	if palette.StartA != 255 {
		t.Errorf("Galaxy.Palette().StartA = %d, want 255", palette.StartA)
	}
	if palette.EndA != 0 {
		t.Errorf("Galaxy.Palette().EndA = %d, want 0", palette.EndA)
	}

	// Test Firework preset palette
	firework := NewFireworkPreset().(*fireworkPreset)
	palette = firework.Palette()
	if palette.Name == "" {
		t.Error("Firework.Palette().Name is empty")
	}

	// Test Swarm preset palette
	swarm := NewSwarmPreset().(*swarmPreset)
	palette = swarm.Palette()
	if palette.Name == "" {
		t.Error("Swarm.Palette().Name is empty")
	}

	// Test Fountain preset palette
	fountain := NewFountainPreset().(*fountainPreset)
	palette = fountain.Palette()
	if palette.Name == "" {
		t.Error("Fountain.Palette().Name is empty")
	}

	// Test Chaos preset palette
	chaos := NewChaosPreset().(*chaosPreset)
	palette = chaos.Palette()
	if palette.Name == "" {
		t.Error("Chaos.Palette().Name is empty")
	}
}
