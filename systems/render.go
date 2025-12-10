package systems

import (
	"fmt"

	"github.com/andygeiss/ecs"
	"github.com/deltatree/showcase/components"
	"github.com/deltatree/showcase/premium"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// RenderSystem handles window management and rendering of all visible entities.
// It initializes the raylib window, handles window close events, and draws
// particles as filled circles with their current color and size.
//
// Debug overlay (toggle with F3) displays:
//   - FPS counter
//   - Active entity count
//   - Current preset name
//   - Mouse coordinates
//   - Control hints
type renderSystem struct {
	width, height int32
	title         string
	showDebug     bool
	presetName    string
	quality       premium.QualitySettings
	uiState       *premium.UIState
	effects       *premium.ScreenEffects
	palette       premium.ColorPalette
}

// NewRenderSystem creates a new render system for the specified window.
//
// Parameters:
//   - width, height: window dimensions in pixels
//   - title: window title displayed in the title bar
func NewRenderSystem(width, height int32, title string) *renderSystem {
	return &renderSystem{
		width:      width,
		height:     height,
		title:      title,
		showDebug:  true,
		presetName: "Galaxy",
		quality:    premium.GetQualitySettings(premium.QualityMedium),
		uiState:    premium.NewUIState(),
		effects:    premium.NewScreenEffects(),
		palette:    premium.GalaxyPalette,
	}
}

func (s *renderSystem) Setup() {
	rl.InitWindow(s.width, s.height, s.title)
	rl.SetTargetFPS(60)
}

func (s *renderSystem) Process(em ecs.EntityManager) (state int) {
	if rl.WindowShouldClose() {
		return ecs.StateEngineStop
	}

	if rl.IsKeyPressed(rl.KeyF3) {
		s.showDebug = !s.showDebug
	}

	// Toggle quality with Q key
	if rl.IsKeyPressed(rl.KeyQ) {
		s.quality = premium.GetQualitySettings(premium.NextQuality(s.quality.Level))
	}

	// Update effects
	dt := rl.GetFrameTime()
	s.effects.Update(dt)
	s.uiState.Update(dt, rl.GetMouseDelta().X != 0 || rl.GetMouseDelta().Y != 0)

	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(10, 10, 20, 255))

	particles := em.FilterByMask(components.MaskRenderable)

	// Apply screen shake offset
	shakeX, shakeY := s.effects.GetShakeOffset()

	for _, e := range particles {
		pos := e.Get(components.MaskPosition).(*components.Position)
		col := e.Get(components.MaskColor).(*components.Color)
		size := e.Get(components.MaskSize).(*components.Size)

		drawX := int32(pos.X + shakeX)
		drawY := int32(pos.Y + shakeY)
		drawColor := rl.NewColor(col.R, col.G, col.B, col.A)

		// Glow effect (if enabled)
		if s.quality.GlowEnabled && s.palette.GlowIntensity > 0 {
			glowAlpha := uint8(float32(col.A) * s.palette.GlowIntensity * 0.3)
			glowColor := rl.NewColor(s.palette.GlowR, s.palette.GlowG, s.palette.GlowB, glowAlpha)
			// Draw glow layers
			for i := 0; i < s.quality.GlowPasses; i++ {
				glowSize := size.Radius * (2.0 + float32(i)*1.5)
				rl.DrawCircle(drawX, drawY, glowSize, glowColor)
			}
		}

		// Main particle
		rl.DrawCircle(drawX, drawY, size.Radius, drawColor)
	}

	// UI with fade alpha
	uiAlpha := uint8(s.uiState.GetControlsAlpha() * 255)

	if s.showDebug {
		rl.DrawFPS(10, 10)
		rl.DrawText(
			fmt.Sprintf("Entities: %d", len(particles)),
			10, 35, 20, rl.White,
		)
		rl.DrawText(
			fmt.Sprintf("Preset: %s", s.presetName),
			10, 60, 20, rl.White,
		)
		rl.DrawText(
			fmt.Sprintf("Quality: %s", s.quality.Level.String()),
			10, 85, 20, rl.NewColor(100, 255, 100, 255),
		)
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		rl.DrawText(
			fmt.Sprintf("Mouse: (%d, %d)", mouseX, mouseY),
			10, 110, 16, rl.Gray,
		)
	}

	// Controls hint with fade
	if uiAlpha > 10 {
		rl.DrawText(
			"F3: Debug | Q: Quality | LMB: Attract | RMB: Repel | 1-5: Presets",
			10, s.height-30, 16, rl.NewColor(150, 150, 150, uiAlpha),
		)
	}

	rl.EndDrawing()

	return ecs.StateEngineContinue
}

func (s *renderSystem) Teardown() {
	rl.CloseWindow()
}

// SetPresetName sets the current preset name for display.
func (s *renderSystem) SetPresetName(name string) {
	s.presetName = name
}

// SetPalette sets the color palette for glow effects.
func (s *renderSystem) SetPalette(palette premium.ColorPalette) {
	s.palette = palette
}

// SetQuality sets the quality level.
func (s *renderSystem) SetQuality(level premium.QualityLevel) {
	s.quality = premium.GetQualitySettings(level)
}

// GetQuality returns the current quality settings.
func (s *renderSystem) GetQuality() premium.QualitySettings {
	return s.quality
}

// ApplyShake triggers a screen shake effect.
func (s *renderSystem) ApplyShake(intensity, duration float32) {
	s.effects.ApplyShake(intensity, duration)
}

// ApplyPulse triggers a screen pulse effect.
func (s *renderSystem) ApplyPulse(scale, duration float32) {
	s.effects.ApplyPulse(scale, duration)
}
