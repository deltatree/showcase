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
	width, height    int32
	title            string
	showDebug        bool
	presetName       string
	quality          premium.QualitySettings
	uiState          *premium.UIState
	effects          *premium.ScreenEffects
	palette          premium.ColorPalette
	maxParticles     int32     // For slider
	onParticleChange func(int) // Callback when slider changes
	isFullscreen     bool      // Track fullscreen state
}

// NewRenderSystem creates a new render system for the specified window.
//
// Parameters:
//   - width, height: window dimensions in pixels
//   - title: window title displayed in the title bar
func NewRenderSystem(width, height int32, title string) *renderSystem {
	return &renderSystem{
		width:        width,
		height:       height,
		title:        title,
		showDebug:    true,
		presetName:   "Fountain",
		quality:      premium.GetQualitySettings(premium.QualityMedium), // Default to MEDIUM for better performance
		uiState:      premium.NewUIState(),
		effects:      premium.NewScreenEffects(),
		palette:      premium.GalaxyPalette,
		maxParticles: 10000,
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

	// Toggle fullscreen with F11 or F key
	if rl.IsKeyPressed(rl.KeyF11) || rl.IsKeyPressed(rl.KeyF) {
		s.isFullscreen = !s.isFullscreen
		rl.ToggleFullscreen()
	}

	// ESC to exit fullscreen (not close app)
	if rl.IsKeyPressed(rl.KeyEscape) && s.isFullscreen {
		s.isFullscreen = false
		rl.ToggleFullscreen()
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
			"F3: Debug | Q: Quality | F/F11: Fullscreen | ESC: Exit Fullscreen | 1-5: Presets",
			10, s.height-30, 16, rl.NewColor(150, 150, 150, uiAlpha),
		)
	}

	// Particle count slider (always visible in debug mode)
	if s.showDebug {
		sliderX := s.width - 250
		sliderY := int32(10)
		sliderWidth := int32(200)
		sliderHeight := int32(20)

		// Draw slider background
		rl.DrawRectangle(sliderX, sliderY, sliderWidth, sliderHeight, rl.NewColor(40, 40, 60, 200))
		rl.DrawRectangleLines(sliderX, sliderY, sliderWidth, sliderHeight, rl.NewColor(100, 126, 234, 200))

		// Calculate fill based on maxParticles (500-20000 range)
		fillPercent := float32(s.maxParticles-500) / float32(20000-500)
		fillWidth := int32(float32(sliderWidth-4) * fillPercent)
		rl.DrawRectangle(sliderX+2, sliderY+2, fillWidth, sliderHeight-4, rl.NewColor(102, 126, 234, 200))

		// Draw label
		rl.DrawText(fmt.Sprintf("Max Particles: %d", s.maxParticles), sliderX, sliderY+25, 14, rl.White)

		// Handle slider interaction
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			if mouseX >= sliderX && mouseX <= sliderX+sliderWidth &&
				mouseY >= sliderY && mouseY <= sliderY+sliderHeight {
				// Calculate new value
				percent := float32(mouseX-sliderX) / float32(sliderWidth)
				newMax := int32(500 + percent*19500)
				if newMax < 500 {
					newMax = 500
				}
				if newMax > 20000 {
					newMax = 20000
				}
				if newMax != s.maxParticles {
					s.maxParticles = newMax
					if s.onParticleChange != nil {
						s.onParticleChange(int(newMax))
					}
				}
			}
		}

		// Quality buttons
		qx := s.width - 250
		qy := int32(55)
		rl.DrawText("Quality:", qx, qy, 14, rl.White)

		qualities := []string{"Low", "Med", "High"}
		for i, q := range qualities {
			btnX := qx + 60 + int32(i*55)
			btnW := int32(50)
			btnH := int32(20)

			isActive := int(s.quality.Level) == i
			bgColor := rl.NewColor(40, 40, 60, 200)
			if isActive {
				bgColor = rl.NewColor(102, 126, 234, 255)
			}

			rl.DrawRectangle(btnX, qy, btnW, btnH, bgColor)
			rl.DrawRectangleLines(btnX, qy, btnW, btnH, rl.NewColor(150, 150, 180, 200))
			rl.DrawText(q, btnX+8, qy+3, 14, rl.White)

			// Handle click
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if mouseX >= btnX && mouseX <= btnX+btnW &&
					mouseY >= qy && mouseY <= qy+btnH {
					s.quality = premium.GetQualitySettings(premium.QualityLevel(i))
				}
			}
		}
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

// SetMaxParticles sets the max particles for the slider display.
func (s *renderSystem) SetMaxParticles(max int) {
	s.maxParticles = int32(max)
}

// GetMaxParticles returns the current max particles setting.
func (s *renderSystem) GetMaxParticles() int {
	return int(s.maxParticles)
}

// SetOnParticleChange sets the callback for when particle slider changes.
func (s *renderSystem) SetOnParticleChange(callback func(int)) {
	s.onParticleChange = callback
}
