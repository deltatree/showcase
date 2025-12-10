package premium

// SoundConfig defines audio settings for a preset.
type SoundConfig struct {
	PresetName    string
	AmbientFile   string
	AmbientVolume float32
	AttractSound  string
	RepelSound    string
	TransitionSFX string
}

var soundConfigs = map[string]SoundConfig{
	"Galaxy": {
		PresetName:    "Galaxy",
		AmbientFile:   "sounds/ambient_space.ogg",
		AmbientVolume: 0.4,
		AttractSound:  "sounds/whoosh_deep.ogg",
		RepelSound:    "sounds/whoosh_high.ogg",
		TransitionSFX: "sounds/transition_cosmic.ogg",
	},
	"Firework": {
		PresetName:    "Firework",
		AmbientFile:   "sounds/ambient_crowd.ogg",
		AmbientVolume: 0.3,
		AttractSound:  "sounds/firework_launch.ogg",
		RepelSound:    "sounds/firework_burst.ogg",
		TransitionSFX: "sounds/transition_bang.ogg",
	},
	"Swarm": {
		PresetName:    "Swarm",
		AmbientFile:   "sounds/ambient_nature.ogg",
		AmbientVolume: 0.35,
		AttractSound:  "sounds/swarm_gather.ogg",
		RepelSound:    "sounds/swarm_scatter.ogg",
		TransitionSFX: "sounds/transition_organic.ogg",
	},
	"Fountain": {
		PresetName:    "Fountain",
		AmbientFile:   "sounds/ambient_water.ogg",
		AmbientVolume: 0.5,
		AttractSound:  "sounds/water_splash.ogg",
		RepelSound:    "sounds/water_spray.ogg",
		TransitionSFX: "sounds/transition_flow.ogg",
	},
	"Chaos": {
		PresetName:    "Chaos",
		AmbientFile:   "sounds/ambient_electric.ogg",
		AmbientVolume: 0.45,
		AttractSound:  "sounds/electric_charge.ogg",
		RepelSound:    "sounds/electric_discharge.ogg",
		TransitionSFX: "sounds/transition_glitch.ogg",
	},
}

// GetSoundConfig returns the sound configuration for a preset.
func GetSoundConfig(name string) SoundConfig {
	if cfg, ok := soundConfigs[name]; ok {
		return cfg
	}
	return soundConfigs["Galaxy"]
}

// AudioManager handles audio playback state.
type AudioManager struct {
	enabled       bool
	muted         bool
	masterVolume  float32
	currentPreset string
}

// NewAudioManager creates a new audio manager.
func NewAudioManager() *AudioManager {
	return &AudioManager{
		enabled:      true,
		muted:        false,
		masterVolume: 0.8,
	}
}

// SetEnabled enables or disables audio.
func (am *AudioManager) SetEnabled(enabled bool) {
	am.enabled = enabled
}

// IsEnabled returns true if audio is enabled.
func (am *AudioManager) IsEnabled() bool {
	return am.enabled
}

// SetMuted sets the muted state.
func (am *AudioManager) SetMuted(muted bool) {
	am.muted = muted
}

// IsMuted returns true if audio is muted.
func (am *AudioManager) IsMuted() bool {
	return am.muted
}

// ToggleMute toggles the muted state.
func (am *AudioManager) ToggleMute() {
	am.muted = !am.muted
}

// SetMasterVolume sets the master volume (0.0 to 1.0).
func (am *AudioManager) SetMasterVolume(volume float32) {
	if volume < 0 {
		volume = 0
	}
	if volume > 1 {
		volume = 1
	}
	am.masterVolume = volume
}

// GetMasterVolume returns the master volume.
func (am *AudioManager) GetMasterVolume() float32 {
	return am.masterVolume
}

// AdjustVolume adjusts volume by delta.
func (am *AudioManager) AdjustVolume(delta float32) {
	am.SetMasterVolume(am.masterVolume + delta)
}

// PlayAttract plays the attract interaction sound.
func (am *AudioManager) PlayAttract() {
	// Placeholder - will integrate with raylib audio
}

// PlayRepel plays the repel interaction sound.
func (am *AudioManager) PlayRepel() {
	// Placeholder - will integrate with raylib audio
}

// PlayTransition plays the preset transition sound.
func (am *AudioManager) PlayTransition() {
	// Placeholder - will integrate with raylib audio
}

// SetPreset changes the current audio preset.
func (am *AudioManager) SetPreset(name string) {
	am.currentPreset = name
	// Placeholder - will load and play new ambient track
}
