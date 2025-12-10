package premium

import "math/rand"

// ScreenEffects manages screen-wide visual effects.
type ScreenEffects struct {
	shakeIntensity float32
	shakeDuration  float32
	shakeTimer     float32
	pulseScale     float32
	pulseDuration  float32
	pulseTimer     float32
}

// NewScreenEffects creates a new screen effects manager.
func NewScreenEffects() *ScreenEffects {
	return &ScreenEffects{
		pulseScale: 1.0,
	}
}

// ApplyShake triggers a screen shake effect.
func (se *ScreenEffects) ApplyShake(intensity, duration float32) {
	se.shakeIntensity = intensity
	se.shakeDuration = duration
	se.shakeTimer = duration
}

// ApplyPulse triggers a screen pulse (zoom) effect.
func (se *ScreenEffects) ApplyPulse(scale, duration float32) {
	se.pulseScale = scale
	se.pulseDuration = duration
	se.pulseTimer = duration
}

// Update advances effect timers.
func (se *ScreenEffects) Update(dt float32) {
	if se.shakeTimer > 0 {
		se.shakeTimer -= dt
		if se.shakeTimer <= 0 {
			se.shakeIntensity = 0
		}
	}
	if se.pulseTimer > 0 {
		se.pulseTimer -= dt
		if se.pulseTimer <= 0 {
			se.pulseScale = 1.0
		}
	}
}

// GetShakeOffset returns the current shake offset.
func (se *ScreenEffects) GetShakeOffset() (x, y float32) {
	if se.shakeTimer <= 0 {
		return 0, 0
	}
	progress := se.shakeTimer / se.shakeDuration
	intensity := se.shakeIntensity * progress
	x = (rand.Float32()*2 - 1) * intensity
	y = (rand.Float32()*2 - 1) * intensity
	return x, y
}

// GetPulseScale returns the current pulse scale.
func (se *ScreenEffects) GetPulseScale() float32 {
	if se.pulseTimer <= 0 {
		return 1.0
	}
	progress := se.pulseTimer / se.pulseDuration
	return 1.0 + (se.pulseScale-1.0)*progress
}

// IsActive returns true if any effect is active.
func (se *ScreenEffects) IsActive() bool {
	return se.shakeTimer > 0 || se.pulseTimer > 0
}

// Reset clears all effects.
func (se *ScreenEffects) Reset() {
	se.shakeTimer = 0
	se.shakeIntensity = 0
	se.pulseTimer = 0
	se.pulseScale = 1.0
}

// JuiceLevel controls the intensity of visual feedback.
type JuiceLevel int

const (
	JuiceOff JuiceLevel = iota
	JuiceSubtle
	JuiceNormal
	JuiceIntense
)

// JuiceConfig contains juice effect multipliers.
type JuiceConfig struct {
	Level           JuiceLevel
	ShakeMultiplier float32
	PulseMultiplier float32
	FlashEnabled    bool
	ParticleBurst   int
}

var juiceConfigs = map[JuiceLevel]JuiceConfig{
	JuiceOff: {
		Level:           JuiceOff,
		ShakeMultiplier: 0,
		PulseMultiplier: 0,
		FlashEnabled:    false,
		ParticleBurst:   0,
	},
	JuiceSubtle: {
		Level:           JuiceSubtle,
		ShakeMultiplier: 0.3,
		PulseMultiplier: 0.3,
		FlashEnabled:    false,
		ParticleBurst:   5,
	},
	JuiceNormal: {
		Level:           JuiceNormal,
		ShakeMultiplier: 0.7,
		PulseMultiplier: 0.7,
		FlashEnabled:    true,
		ParticleBurst:   15,
	},
	JuiceIntense: {
		Level:           JuiceIntense,
		ShakeMultiplier: 1.2,
		PulseMultiplier: 1.2,
		FlashEnabled:    true,
		ParticleBurst:   30,
	},
}

// GetJuiceConfig returns the juice configuration for a level.
func GetJuiceConfig(level JuiceLevel) JuiceConfig {
	if cfg, ok := juiceConfigs[level]; ok {
		return cfg
	}
	return juiceConfigs[JuiceNormal]
}
