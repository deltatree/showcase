package premium_test

import (
	"fmt"

	"github.com/deltatree/showcase/premium"
)

func ExampleGetPalette() {
	palette := premium.GetPalette("Galaxy")
	fmt.Println(palette.Name)
	// Output: Galaxy
}

func ExampleQualityLevel_String() {
	level := premium.QualityHigh
	fmt.Println(level.String())
	// Output: High
}

func ExampleGetQualitySettings() {
	settings := premium.GetQualitySettings(premium.QualityHigh)
	fmt.Println(settings.GlowEnabled)
	// Output: true
}

func ExampleNextQuality() {
	next := premium.NextQuality(premium.QualityLow)
	fmt.Println(next.String())
	// Output: Medium
}

func ExampleNewScreenEffects() {
	effects := premium.NewScreenEffects()
	effects.ApplyShake(5.0, 0.5)
	fmt.Println(effects.IsActive())
	// Output: true
}

func ExampleGetJuiceConfig() {
	config := premium.GetJuiceConfig(premium.JuiceIntense)
	fmt.Println(config.FlashEnabled)
	// Output: true
}

func ExampleNewAudioManager() {
	am := premium.NewAudioManager()
	am.SetMasterVolume(0.5)
	fmt.Println(am.GetMasterVolume())
	// Output: 0.5
}

func ExampleGetSoundConfig() {
	config := premium.GetSoundConfig("Firework")
	fmt.Println(config.PresetName)
	// Output: Firework
}

func ExampleNewUIState() {
	ui := premium.NewUIState()
	fmt.Println(ui.ShowControls)
	// Output: true
}

func ExampleNewUILayout() {
	layout := premium.NewUILayout(1920, 1080)
	x, y := layout.Center()
	fmt.Printf("%d,%d", x, y)
	// Output: 960,540
}
