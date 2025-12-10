package presets_test

import (
	"fmt"

	"github.com/deltatree/showcase/presets"
)

// ExampleGetPreset demonstrates getting a preset by index.
func ExampleGetPreset() {
	preset := presets.GetPreset(0)
	fmt.Printf("Preset 0: %s\n", preset.Name())
	// Output: Preset 0: Galaxy
}

// ExampleGetPresetByName demonstrates getting a preset by name.
func ExampleGetPresetByName() {
	preset := presets.GetPresetByName("Firework")
	fmt.Printf("Found: %s - %s\n", preset.Name(), preset.Description())
	// Output: Found: Firework - Colorful firework explosions with gravity
}

// Example demonstrates listing all available presets.
func Example() {
	names := []string{"Galaxy", "Firework", "Swarm", "Fountain", "Chaos"}
	for i, name := range names {
		preset := presets.GetPresetByName(name)
		fmt.Printf("%d: %s\n", i+1, preset.Name())
	}
	// Output:
	// 1: Galaxy
	// 2: Firework
	// 3: Swarm
	// 4: Fountain
	// 5: Chaos
}

// ExampleGetPalette demonstrates getting a preset's color palette.
func ExampleGetPalette() {
	preset := presets.GetPresetByName("Galaxy")
	palette := presets.GetPalette(preset)
	fmt.Printf("Palette: %s, Glow: %.1f\n", palette.Name, palette.GlowIntensity)
	// Output: Palette: Galaxy, Glow: 0.6
}
