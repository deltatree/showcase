package premium

// UIState manages UI visibility and animations.
type UIState struct {
	ShowControls    bool
	ShowDebug       bool
	ShowHelp        bool
	controlsFade    float32
	inactivityTimer float32
	fadeSpeed       float32
	hideDelay       float32
}

// NewUIState creates a new UI state manager.
func NewUIState() *UIState {
	return &UIState{
		ShowControls:    true,
		ShowDebug:       false,
		ShowHelp:        false,
		controlsFade:    1.0,
		inactivityTimer: 0,
		fadeSpeed:       2.0,
		hideDelay:       3.0,
	}
}

// Update advances UI state with delta time.
func (ui *UIState) Update(dt float32, hasActivity bool) {
	if hasActivity {
		ui.inactivityTimer = 0
		ui.ShowControls = true
	} else {
		ui.inactivityTimer += dt
		if ui.inactivityTimer > ui.hideDelay {
			ui.ShowControls = false
		}
	}

	// Animate fade
	targetFade := float32(0)
	if ui.ShowControls {
		targetFade = 1.0
	}
	if ui.controlsFade < targetFade {
		ui.controlsFade += dt * ui.fadeSpeed
		if ui.controlsFade > targetFade {
			ui.controlsFade = targetFade
		}
	} else if ui.controlsFade > targetFade {
		ui.controlsFade -= dt * ui.fadeSpeed
		if ui.controlsFade < targetFade {
			ui.controlsFade = targetFade
		}
	}
}

// GetControlsAlpha returns the current controls opacity.
func (ui *UIState) GetControlsAlpha() float32 {
	return ui.controlsFade
}

// ToggleDebug toggles debug display.
func (ui *UIState) ToggleDebug() {
	ui.ShowDebug = !ui.ShowDebug
}

// ToggleHelp toggles help display.
func (ui *UIState) ToggleHelp() {
	ui.ShowHelp = !ui.ShowHelp
}

// ForceShowControls shows controls immediately.
func (ui *UIState) ForceShowControls() {
	ui.ShowControls = true
	ui.controlsFade = 1.0
	ui.inactivityTimer = 0
}

// UILayout provides UI positioning helpers.
type UILayout struct {
	Width   int
	Height  int
	MarginX int
	MarginY int
}

// NewUILayout creates a UI layout for given dimensions.
func NewUILayout(width, height int) *UILayout {
	return &UILayout{
		Width:   width,
		Height:  height,
		MarginX: 20,
		MarginY: 20,
	}
}

// TopLeft returns top-left position.
func (l *UILayout) TopLeft() (x, y int) {
	return l.MarginX, l.MarginY
}

// TopRight returns top-right position.
func (l *UILayout) TopRight() (x, y int) {
	return l.Width - l.MarginX, l.MarginY
}

// BottomLeft returns bottom-left position.
func (l *UILayout) BottomLeft() (x, y int) {
	return l.MarginX, l.Height - l.MarginY
}

// BottomRight returns bottom-right position.
func (l *UILayout) BottomRight() (x, y int) {
	return l.Width - l.MarginX, l.Height - l.MarginY
}

// BottomCenter returns bottom-center position.
func (l *UILayout) BottomCenter() (x, y int) {
	return l.Width / 2, l.Height - l.MarginY
}

// Center returns center position.
func (l *UILayout) Center() (x, y int) {
	return l.Width / 2, l.Height / 2
}
