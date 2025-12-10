package premium

import (
	"testing"
)

func TestGetPalette(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Galaxy", "Galaxy"},
		{"Firework", "Firework"},
		{"Swarm", "Swarm"},
		{"Fountain", "Fountain"},
		{"Chaos", "Chaos"},
		{"Unknown", "Galaxy"},
	}

	for _, tc := range tests {
		p := GetPalette(tc.name)
		if p.Name != tc.expected {
			t.Errorf("GetPalette(%s) = %s, want %s", tc.name, p.Name, tc.expected)
		}
	}
}

func TestColorPaletteValues(t *testing.T) {
	for name, p := range Palettes {
		if p.StartA != 255 {
			t.Errorf("%s: StartA = %d, want 255", name, p.StartA)
		}
		if p.EndA != 0 {
			t.Errorf("%s: EndA = %d, want 0", name, p.EndA)
		}
		if p.GlowIntensity < 0 || p.GlowIntensity > 1 {
			t.Errorf("%s: GlowIntensity = %f, want 0-1", name, p.GlowIntensity)
		}
	}
}

func TestQualityLevelString(t *testing.T) {
	tests := []struct {
		level    QualityLevel
		expected string
	}{
		{QualityLow, "Low"},
		{QualityMedium, "Medium"},
		{QualityHigh, "High"},
		{QualityLevel(99), "Unknown"},
	}

	for _, tc := range tests {
		if got := tc.level.String(); got != tc.expected {
			t.Errorf("%d.String() = %s, want %s", tc.level, got, tc.expected)
		}
	}
}

func TestGetQualitySettings(t *testing.T) {
	tests := []struct {
		level       QualityLevel
		expectedMax int
		glowEnabled bool
		motionBlur  bool
	}{
		{QualityLow, 3000, false, false},
		{QualityMedium, 7000, true, false},
		{QualityHigh, 15000, true, true},
	}

	for _, tc := range tests {
		s := GetQualitySettings(tc.level)
		if s.MaxParticles != tc.expectedMax {
			t.Errorf("Level %s: MaxParticles = %d, want %d", tc.level, s.MaxParticles, tc.expectedMax)
		}
		if s.GlowEnabled != tc.glowEnabled {
			t.Errorf("Level %s: GlowEnabled = %v, want %v", tc.level, s.GlowEnabled, tc.glowEnabled)
		}
		if s.MotionBlur != tc.motionBlur {
			t.Errorf("Level %s: MotionBlur = %v, want %v", tc.level, s.MotionBlur, tc.motionBlur)
		}
	}
}

func TestNextQuality(t *testing.T) {
	tests := []struct {
		current  QualityLevel
		expected QualityLevel
	}{
		{QualityLow, QualityMedium},
		{QualityMedium, QualityHigh},
		{QualityHigh, QualityLow},
	}

	for _, tc := range tests {
		if got := NextQuality(tc.current); got != tc.expected {
			t.Errorf("NextQuality(%s) = %s, want %s", tc.current, got, tc.expected)
		}
	}
}

func TestScreenEffects(t *testing.T) {
	se := NewScreenEffects()

	if se.IsActive() {
		t.Error("NewScreenEffects should not be active")
	}

	se.ApplyShake(5.0, 0.5)
	if !se.IsActive() {
		t.Error("Shake should make effects active")
	}

	se.GetShakeOffset()

	se.ApplyPulse(1.05, 0.3)
	scale := se.GetPulseScale()
	if scale < 1.0 {
		t.Errorf("Pulse scale should be >= 1.0, got %f", scale)
	}

	se.Reset()
	if se.IsActive() {
		t.Error("Reset should deactivate effects")
	}
}

func TestScreenEffectsUpdate(t *testing.T) {
	se := NewScreenEffects()
	se.ApplyShake(10.0, 0.1)

	se.Update(0.2)

	if se.IsActive() {
		t.Error("Effects should decay after duration")
	}

	x, y := se.GetShakeOffset()
	if x != 0 || y != 0 {
		t.Errorf("Shake offset should be (0,0) after decay, got (%f, %f)", x, y)
	}
}

func TestJuiceConfig(t *testing.T) {
	levels := []JuiceLevel{JuiceOff, JuiceSubtle, JuiceNormal, JuiceIntense}
	for _, level := range levels {
		c := GetJuiceConfig(level)
		if c.Level != level {
			t.Errorf("GetJuiceConfig(%d).Level = %d, want %d", level, c.Level, level)
		}
	}

	off := GetJuiceConfig(JuiceOff)
	if off.ShakeMultiplier != 0 {
		t.Errorf("JuiceOff.ShakeMultiplier = %f, want 0", off.ShakeMultiplier)
	}

	intense := GetJuiceConfig(JuiceIntense)
	normal := GetJuiceConfig(JuiceNormal)
	if intense.ShakeMultiplier <= normal.ShakeMultiplier {
		t.Error("JuiceIntense should have higher ShakeMultiplier than JuiceNormal")
	}
}

func TestAudioManager(t *testing.T) {
	am := NewAudioManager()

	if !am.IsEnabled() {
		t.Error("Audio should be enabled by default")
	}
	if am.IsMuted() {
		t.Error("Audio should not be muted by default")
	}

	am.ToggleMute()
	if !am.IsMuted() {
		t.Error("ToggleMute should mute")
	}
	am.ToggleMute()
	if am.IsMuted() {
		t.Error("ToggleMute should unmute")
	}

	am.SetMasterVolume(0.5)
	if am.GetMasterVolume() != 0.5 {
		t.Errorf("GetMasterVolume() = %f, want 0.5", am.GetMasterVolume())
	}

	am.SetMasterVolume(-1.0)
	if am.GetMasterVolume() != 0 {
		t.Errorf("Volume should clamp to 0, got %f", am.GetMasterVolume())
	}
	am.SetMasterVolume(2.0)
	if am.GetMasterVolume() != 1.0 {
		t.Errorf("Volume should clamp to 1, got %f", am.GetMasterVolume())
	}

	am.SetMasterVolume(0.5)
	am.AdjustVolume(0.2)
	if am.GetMasterVolume() != 0.7 {
		t.Errorf("AdjustVolume: expected 0.7, got %f", am.GetMasterVolume())
	}

	am.SetEnabled(false)
	if am.IsEnabled() {
		t.Error("SetEnabled(false) should disable")
	}
}

func TestGetSoundConfig(t *testing.T) {
	tests := []string{"Galaxy", "Firework", "Swarm", "Fountain", "Chaos"}

	for _, name := range tests {
		c := GetSoundConfig(name)
		if c.PresetName != name {
			t.Errorf("GetSoundConfig(%s).PresetName = %s", name, c.PresetName)
		}
		if c.AmbientFile == "" {
			t.Errorf("GetSoundConfig(%s) has empty AmbientFile", name)
		}
	}

	c := GetSoundConfig("Unknown")
	if c.PresetName != "Galaxy" {
		t.Errorf("Unknown preset should return Galaxy config, got %s", c.PresetName)
	}
}

func TestUIState(t *testing.T) {
	ui := NewUIState()

	if !ui.ShowControls {
		t.Error("Controls should be visible by default")
	}
	if ui.GetControlsAlpha() != 1.0 {
		t.Error("Controls alpha should be 1.0 initially")
	}

	if ui.ShowDebug {
		t.Error("Debug should be hidden initially")
	}
	ui.ToggleDebug()
	if !ui.ShowDebug {
		t.Error("ToggleDebug should show debug")
	}

	for i := 0; i < 50; i++ {
		ui.Update(0.1, false)
	}
	if ui.ShowControls {
		t.Error("Controls should auto-hide after inactivity")
	}

	ui.Update(0.1, true)
	if !ui.ShowControls {
		t.Error("Activity should show controls")
	}
}

func TestUILayout(t *testing.T) {
	l := NewUILayout(1280, 720)

	if l.Width != 1280 || l.Height != 720 {
		t.Errorf("Dimensions = (%d, %d), want (1280, 720)", l.Width, l.Height)
	}

	x, y := l.TopLeft()
	if x != l.MarginX || y != l.MarginY {
		t.Errorf("TopLeft = (%d, %d)", x, y)
	}

	x, y = l.BottomCenter()
	if x != 640 {
		t.Errorf("BottomCenter.X = %d, want 640", x)
	}
}
