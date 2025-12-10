//go:build js && wasm

// Particle Symphony - WebAssembly Version
//
// This is the WebAssembly entry point for Particle Symphony, enabling
// the particle simulation to run in web browsers. It uses Ebitengine
// for rendering instead of raylib for browser compatibility.
//
// Build: GOOS=js GOARCH=wasm go build -o main.wasm ./cmd/wasm
// Serve: Serve the web/ directory with index.html and wasm_exec.js
//
// Controls (same as native):
//   - Mouse: Move to guide particles
//   - Left Click: Attract particles
//   - Right Click: Repel particles
//   - 1-5: Switch between presets
//   - M: Toggle sound mute
package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var maxParticles = 10000 // Configurable via UI

// Particle represents a single particle entity
type Particle struct {
	X, Y                       float32
	VX, VY                     float32
	AX, AY                     float32
	R, G, B, A                 uint8
	StartR, StartG, StartB     uint8
	EndR, EndG, EndB           uint8
	Age, TTL                   float32
	Radius, StartSize, EndSize float32
	Active                     bool
}

// Preset represents a particle preset configuration
type Preset struct {
	Name                   string
	StartR, StartG, StartB uint8
	EndR, EndG, EndB       uint8
	MinSize, MaxSize       float32
	MinTTL, MaxTTL         float32
	MinVel, MaxVel         float32
	SpawnPattern           string
	SpawnRate              int
	// Premium color palette
	GlowR, GlowG, GlowB uint8
	GlowIntensity       float32
}

var presets = []Preset{
	{
		Name: "Galaxy", StartR: 100, StartG: 150, StartB: 255,
		EndR: 255, EndG: 100, EndB: 200,
		MinSize: 1.5, MaxSize: 4.0, MinTTL: 4.0, MaxTTL: 7.0,
		MinVel: -30, MaxVel: 30, SpawnPattern: "center", SpawnRate: 150,
		GlowR: 150, GlowG: 100, GlowB: 255, GlowIntensity: 0.8,
	},
	{
		Name: "Firework", StartR: 255, StartG: 200, StartB: 50,
		EndR: 255, EndG: 50, EndB: 0,
		MinSize: 2.0, MaxSize: 5.0, MinTTL: 1.5, MaxTTL: 3.0,
		MinVel: -150, MaxVel: 150, SpawnPattern: "center", SpawnRate: 200,
		GlowR: 255, GlowG: 180, GlowB: 50, GlowIntensity: 1.0,
	},
	{
		Name: "Swarm", StartR: 50, StartG: 255, StartB: 100,
		EndR: 0, EndG: 150, EndB: 50,
		MinSize: 1.0, MaxSize: 3.0, MinTTL: 5.0, MaxTTL: 8.0,
		MinVel: -20, MaxVel: 20, SpawnPattern: "random", SpawnRate: 100,
		GlowR: 100, GlowG: 255, GlowB: 150, GlowIntensity: 0.5,
	},
	{
		Name: "Fountain", StartR: 100, StartG: 200, StartB: 255,
		EndR: 50, EndG: 100, EndB: 200,
		MinSize: 2.0, MaxSize: 4.0, MinTTL: 2.0, MaxTTL: 4.0,
		MinVel: -80, MaxVel: 80, SpawnPattern: "bottom", SpawnRate: 180,
		GlowR: 100, GlowG: 180, GlowB: 255, GlowIntensity: 0.7,
	},
	{
		Name: "Chaos", StartR: 255, StartG: 50, StartB: 50,
		EndR: 50, EndG: 50, EndB: 255,
		MinSize: 1.0, MaxSize: 6.0, MinTTL: 2.0, MaxTTL: 5.0,
		MinVel: -100, MaxVel: 100, SpawnPattern: "edges", SpawnRate: 250,
		GlowR: 255, GlowG: 100, GlowB: 100, GlowIntensity: 0.9,
	},
}

// QualityLevel for performance settings
type QualityLevel int

const (
	QualityLow QualityLevel = iota
	QualityMedium
	QualityHigh
)

func (q QualityLevel) String() string {
	switch q {
	case QualityLow:
		return "Low"
	case QualityMedium:
		return "Medium"
	case QualityHigh:
		return "High"
	}
	return "Unknown"
}

// QualitySettings defines performance parameters
type QualitySettings struct {
	MaxParticles int
	GlowEnabled  bool
	GlowPasses   int
	SpawnMult    float32
}

var qualityPresets = map[QualityLevel]QualitySettings{
	QualityLow:    {MaxParticles: 3000, GlowEnabled: false, GlowPasses: 0, SpawnMult: 0.5},
	QualityMedium: {MaxParticles: 7000, GlowEnabled: true, GlowPasses: 1, SpawnMult: 1.0},
	QualityHigh:   {MaxParticles: 12000, GlowEnabled: true, GlowPasses: 2, SpawnMult: 1.5},
}

// AudioEngine handles Web Audio API for interactive musical sound effects
type AudioEngine struct {
	ctx           js.Value
	muted         bool
	volume        float64
	lastNote      int
	scaleIndex    int
	chordIndex    int
	arpIndex      int
	interactTimer float32
}

// Musical scales for variety
var (
	// Different scales for different moods
	pentatonicMajor = []float64{261.63, 293.66, 329.63, 392.00, 440.00}         // C D E G A
	pentatonicMinor = []float64{261.63, 311.13, 349.23, 392.00, 466.16}         // C Eb F G Bb
	blues           = []float64{261.63, 311.13, 349.23, 369.99, 392.00, 466.16} // C Eb F F# G Bb
	japanese        = []float64{261.63, 277.18, 329.63, 392.00, 415.30}         // C Db E G Ab
	arabic          = []float64{261.63, 277.18, 329.63, 349.23, 392.00, 415.30} // C Db E F G Ab
	wholeTone       = []float64{261.63, 293.66, 329.63, 369.99, 415.30, 466.16} // C D E F# G# A#

	allScales = [][]float64{pentatonicMajor, pentatonicMinor, blues, japanese, arabic, wholeTone}

	// Chord types (intervals in semitones from root)
	chordMajor    = []float64{1.0, 1.26, 1.5}   // Root, Major 3rd, Perfect 5th
	chordMinor    = []float64{1.0, 1.19, 1.5}   // Root, Minor 3rd, Perfect 5th
	chordDim      = []float64{1.0, 1.19, 1.414} // Root, Minor 3rd, Tritone
	chordSus4     = []float64{1.0, 1.335, 1.5}  // Root, Perfect 4th, Perfect 5th
	chordPowerAdd = []float64{1.0, 1.5, 2.0}    // Root, 5th, Octave
)

func NewAudioEngine() *AudioEngine {
	ae := &AudioEngine{
		muted:      false,
		volume:     0.2,
		scaleIndex: 0,
	}
	// Create Web Audio context
	audioCtx := js.Global().Get("AudioContext")
	if audioCtx.IsUndefined() {
		audioCtx = js.Global().Get("webkitAudioContext")
	}
	if !audioCtx.IsUndefined() {
		ae.ctx = audioCtx.New()
	}
	return ae
}

func (ae *AudioEngine) IsReady() bool {
	return !ae.ctx.IsUndefined() && !ae.ctx.IsNull()
}

func (ae *AudioEngine) Resume() {
	if ae.IsReady() && ae.ctx.Get("state").String() == "suspended" {
		ae.ctx.Call("resume")
	}
}

// PlayTone plays a single tone with specified waveform
func (ae *AudioEngine) PlayTone(freq float64, duration float64, volume float64, waveType string) {
	if ae.muted || !ae.IsReady() {
		return
	}
	osc := ae.ctx.Call("createOscillator")
	gain := ae.ctx.Call("createGain")

	osc.Get("frequency").Set("value", freq)
	osc.Set("type", waveType)

	now := ae.ctx.Get("currentTime").Float()
	gain.Get("gain").Call("setValueAtTime", volume*ae.volume, now)
	gain.Get("gain").Call("exponentialRampToValueAtTime", 0.001, now+duration)

	osc.Call("connect", gain)
	gain.Call("connect", ae.ctx.Get("destination"))
	osc.Call("start", now)
	osc.Call("stop", now+duration)
}

// PlayChord plays multiple notes simultaneously
func (ae *AudioEngine) PlayChord(baseFreq float64, intervals []float64, duration float64, volume float64, waveType string) {
	for i, interval := range intervals {
		// Slight detuning for richness
		detune := 1.0 + float64(i)*0.002
		ae.PlayTone(baseFreq*interval*detune, duration, volume/float64(len(intervals)), waveType)
	}
}

// PlayArpeggio plays notes in sequence
func (ae *AudioEngine) PlayArpeggio(baseFreq float64, intervals []float64, noteLen float64, volume float64, waveType string) {
	if ae.muted || !ae.IsReady() {
		return
	}
	now := ae.ctx.Get("currentTime").Float()
	for i, interval := range intervals {
		osc := ae.ctx.Call("createOscillator")
		gain := ae.ctx.Call("createGain")

		osc.Get("frequency").Set("value", baseFreq*interval)
		osc.Set("type", waveType)

		startTime := now + float64(i)*noteLen*0.8
		gain.Get("gain").Call("setValueAtTime", 0.001, now)
		gain.Get("gain").Call("setValueAtTime", volume*ae.volume, startTime)
		gain.Get("gain").Call("exponentialRampToValueAtTime", 0.001, startTime+noteLen)

		osc.Call("connect", gain)
		gain.Call("connect", ae.ctx.Get("destination"))
		osc.Call("start", now)
		osc.Call("stop", startTime+noteLen+0.1)
	}
}

// PlayInteraction - called during particle interaction, creates varied musical response
func (ae *AudioEngine) PlayInteraction(isAttract bool, intensity float64, particleCount int) {
	if ae.muted || !ae.IsReady() {
		return
	}

	// Choose scale based on attract/repel and randomness
	scale := allScales[ae.scaleIndex]

	// Base note from scale
	noteIdx := (ae.lastNote + 1 + rand.Intn(3)) % len(scale)
	ae.lastNote = noteIdx
	baseFreq := scale[noteIdx]

	// Octave based on intensity
	octave := 1.0
	if intensity > 0.5 {
		octave = 2.0
	}
	if particleCount > 5000 {
		octave *= 0.5 // Lower octave for many particles
	}
	baseFreq *= octave

	// Volume based on intensity
	vol := 0.1 + intensity*0.3
	if vol > 0.4 {
		vol = 0.4
	}

	// Waveform variety
	waveTypes := []string{"sine", "triangle", "square", "sawtooth"}
	waveType := waveTypes[rand.Intn(len(waveTypes))]

	// Different musical responses
	responseType := rand.Intn(5)

	switch responseType {
	case 0:
		// Single note with harmonics
		ae.PlayTone(baseFreq, 0.2, vol, waveType)
		ae.PlayTone(baseFreq*2, 0.15, vol*0.3, "sine")
	case 1:
		// Major or minor chord based on attract/repel
		if isAttract {
			ae.PlayChord(baseFreq, chordMajor, 0.25, vol, waveType)
		} else {
			ae.PlayChord(baseFreq, chordMinor, 0.25, vol, waveType)
		}
	case 2:
		// Power chord
		ae.PlayChord(baseFreq, chordPowerAdd, 0.3, vol, "sawtooth")
	case 3:
		// Quick arpeggio up or down
		intervals := []float64{1.0, 1.26, 1.5, 2.0}
		if !isAttract {
			// Reverse for repel
			intervals = []float64{2.0, 1.5, 1.26, 1.0}
		}
		ae.PlayArpeggio(baseFreq, intervals, 0.08, vol, waveType)
	case 4:
		// Suspended chord (ethereal)
		ae.PlayChord(baseFreq, chordSus4, 0.3, vol, "triangle")
	}

	// Occasionally change scale for variety
	if rand.Intn(20) == 0 {
		ae.scaleIndex = (ae.scaleIndex + 1) % len(allScales)
	}
}

// PlayPresetChange plays a distinctive sound when switching presets
func (ae *AudioEngine) PlayPresetChange(presetIndex int) {
	if ae.muted || !ae.IsReady() {
		return
	}

	// Each preset gets a unique chord/sound
	presetSounds := []struct {
		freq  float64
		chord []float64
		wave  string
	}{
		{329.63, chordMajor, "triangle"},  // Fountain - bright, major
		{261.63, chordMinor, "sawtooth"},  // Firework - explosive
		{440.00, chordSus4, "sine"},       // Galaxy - ethereal
		{349.23, chordPowerAdd, "square"}, // Chaos - aggressive
		{392.00, chordDim, "triangle"},    // Swarm - mysterious
	}

	if presetIndex < len(presetSounds) {
		ps := presetSounds[presetIndex]
		ae.PlayChord(ps.freq, ps.chord, 0.3, 0.35, ps.wave)
		// Add shimmer
		ae.PlayTone(ps.freq*4, 0.2, 0.1, "sine")
	}

	// Switch to a scale that matches the preset mood
	ae.scaleIndex = presetIndex % len(allScales)
}

// UpdateInteraction should be called each frame during interaction
func (ae *AudioEngine) UpdateInteraction(isInteracting bool, isAttract bool, intensity float64, particleCount int, dt float32) {
	if !isInteracting {
		ae.interactTimer = 0
		return
	}

	ae.interactTimer -= dt
	if ae.interactTimer <= 0 {
		ae.PlayInteraction(isAttract, intensity, particleCount)
		// Variable timing based on intensity
		ae.interactTimer = 0.1 + (1.0-float32(intensity))*0.2 + rand.Float32()*0.1
	}
}

func (ae *AudioEngine) ToggleMute() {
	ae.muted = !ae.muted
}

func (ae *AudioEngine) IsMuted() bool {
	return ae.muted
}

// Game implements ebiten.Game interface
type Game struct {
	particles      []Particle
	rng            *rand.Rand
	spawnTimer     float32
	currentPreset  int
	preset         Preset
	mouseX, mouseY int
	attractorMass  float32
	lockedMode     int
	showDebug      bool
	lastClickTime  time.Time
	activeCount    int
	audio          *AudioEngine
	lastAttract    bool
	lastRepel      bool
	// Premium features
	quality         QualityLevel
	qualitySettings QualitySettings
	// Touch/Mobile support
	touchIDs      []ebiten.TouchID
	lastTouchTime time.Time
	isMobile      bool
}

func NewGame() *Game {
	// Detect mobile via user agent
	isMobile := false
	navigator := js.Global().Get("navigator")
	if !navigator.IsUndefined() {
		ua := navigator.Get("userAgent").String()
		isMobile = contains(ua, "Mobile") || contains(ua, "Android") || contains(ua, "iPhone") || contains(ua, "iPad")
	}

	// Default quality is HIGH for best experience
	defaultQuality := QualityHigh
	if isMobile {
		defaultQuality = QualityMedium // Mobile gets Medium for performance
	}

	g := &Game{
		particles:       make([]Particle, maxParticles),
		rng:             rand.New(rand.NewSource(time.Now().UnixNano())),
		currentPreset:   0,
		showDebug:       !isMobile, // Hide debug on mobile by default
		lockedMode:      0,
		audio:           NewAudioEngine(),
		quality:         defaultQuality,
		qualitySettings: qualityPresets[defaultQuality],
		isMobile:        isMobile,
	}
	g.preset = presets[0]
	return g
}

// Simple string contains helper
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func (g *Game) spawnParticle() {
	var p *Particle
	for i := range g.particles {
		if !g.particles[i].Active {
			p = &g.particles[i]
			break
		}
	}
	if p == nil {
		return
	}

	preset := g.preset
	var x, y float32

	switch preset.SpawnPattern {
	case "center":
		x = screenWidth/2 + (g.rng.Float32()-0.5)*100
		y = screenHeight/2 + (g.rng.Float32()-0.5)*100
	case "bottom":
		x = screenWidth/2 + (g.rng.Float32()-0.5)*200
		y = screenHeight - 50
	case "edges":
		side := g.rng.Intn(4)
		switch side {
		case 0:
			x, y = g.rng.Float32()*screenWidth, 0
		case 1:
			x, y = g.rng.Float32()*screenWidth, screenHeight
		case 2:
			x, y = 0, g.rng.Float32()*screenHeight
		case 3:
			x, y = screenWidth, g.rng.Float32()*screenHeight
		}
	default:
		x = g.rng.Float32() * screenWidth
		y = g.rng.Float32() * screenHeight
	}

	vx := preset.MinVel + g.rng.Float32()*(preset.MaxVel-preset.MinVel)
	vy := preset.MinVel + g.rng.Float32()*(preset.MaxVel-preset.MinVel)
	size := preset.MinSize + g.rng.Float32()*(preset.MaxSize-preset.MinSize)
	ttl := preset.MinTTL + g.rng.Float32()*(preset.MaxTTL-preset.MinTTL)

	*p = Particle{
		X: x, Y: y, VX: vx, VY: vy, AX: 0, AY: 0,
		R: preset.StartR, G: preset.StartG, B: preset.StartB, A: 255,
		StartR: preset.StartR, StartG: preset.StartG, StartB: preset.StartB,
		EndR: preset.EndR, EndG: preset.EndG, EndB: preset.EndB,
		Age: 0, TTL: ttl,
		Radius: size, StartSize: size, EndSize: size * 0.3,
		Active: true,
	}
}

func (g *Game) Update() error {
	dt := float32(1.0 / 60.0)

	// Handle touch input for mobile
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	touchActive := len(ebiten.AppendTouchIDs(nil)) > 0

	// Get input position (touch or mouse)
	if touchActive {
		touches := ebiten.AppendTouchIDs(nil)
		if len(touches) > 0 {
			g.mouseX, g.mouseY = ebiten.TouchPosition(touches[0])
		}
	} else {
		g.mouseX, g.mouseY = ebiten.CursorPosition()
	}

	// Resume audio context on first interaction (browser requirement)
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) ||
		len(g.touchIDs) > 0 {
		g.audio.Resume()
	}

	now := time.Now()

	// Touch: single tap = attract, double tap = lock, two-finger = repel
	if len(g.touchIDs) > 0 {
		touches := ebiten.AppendTouchIDs(nil)
		if len(touches) >= 2 {
			// Two fingers = repel
			g.attractorMass = -5000
		} else if now.Sub(g.lastTouchTime) < 300*time.Millisecond {
			// Double tap = toggle lock
			if g.lockedMode == 1 {
				g.lockedMode = 0
			} else {
				g.lockedMode = 1
			}
		}
		g.lastTouchTime = now
	}

	// Mouse input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if now.Sub(g.lastClickTime) < 300*time.Millisecond {
			if g.lockedMode == 1 {
				g.lockedMode = 0
			} else {
				g.lockedMode = 1
			}
		}
		g.lastClickTime = now
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		if now.Sub(g.lastClickTime) < 300*time.Millisecond {
			if g.lockedMode == -1 {
				g.lockedMode = 0
			} else {
				g.lockedMode = -1
			}
		}
		g.lastClickTime = now
	}

	// Determine attractor mass
	if g.lockedMode == 1 {
		g.attractorMass = 8000
	} else if g.lockedMode == -1 {
		g.attractorMass = -8000
	} else if touchActive {
		touches := ebiten.AppendTouchIDs(nil)
		if len(touches) >= 2 {
			g.attractorMass = -5000 // Two fingers = repel
		} else if len(touches) == 1 {
			g.attractorMass = 5000 // One finger = attract
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.attractorMass = 5000
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.attractorMass = -5000
	} else {
		g.attractorMass = 0
	}

	// Interactive music - plays varied sounds during interaction
	isInteracting := g.attractorMass != 0
	isAttracting := g.attractorMass > 0
	intensity := math.Abs(float64(g.attractorMass)) / 8000.0
	if intensity > 1.0 {
		intensity = 1.0
	}
	g.audio.UpdateInteraction(isInteracting, isAttracting, intensity, g.activeCount, dt)

	g.lastAttract = isAttracting
	g.lastRepel = g.attractorMass < 0

	keys := []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5}
	for i, key := range keys {
		if inpututil.IsKeyJustPressed(key) {
			g.switchPreset(i)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		g.showDebug = !g.showDebug
	}

	// Toggle mute with M key
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.audio.ToggleMute()
	}

	// Toggle quality with Q key
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.cycleQuality()
	}

	// Spawn with quality-adjusted rate
	effectiveRate := float32(g.preset.SpawnRate) * g.qualitySettings.SpawnMult
	g.spawnTimer += dt
	spawnInterval := 1.0 / effectiveRate
	for g.spawnTimer >= spawnInterval {
		g.spawnTimer -= spawnInterval
		// Only spawn if under quality limit
		if g.activeCount < g.qualitySettings.MaxParticles {
			g.spawnParticle()
		}
	}

	g.activeCount = 0
	for i := range g.particles {
		p := &g.particles[i]
		if !p.Active {
			continue
		}
		g.activeCount++

		if g.attractorMass != 0 {
			dx := float32(g.mouseX) - p.X
			dy := float32(g.mouseY) - p.Y
			dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
			if dist < 10 {
				dist = 10
			}
			force := g.attractorMass / (dist * dist) * 500
			p.AX = dx / dist * force
			p.AY = dy / dist * force
		} else {
			p.AX, p.AY = 0, 0
		}

		p.VX += p.AX * dt
		p.VY += p.AY * dt
		p.VX *= 0.99
		p.VY *= 0.99

		mag := float32(math.Sqrt(float64(p.VX*p.VX + p.VY*p.VY)))
		if mag > 500 {
			p.VX = p.VX / mag * 500
			p.VY = p.VY / mag * 500
		}

		p.X += p.VX * dt
		p.Y += p.VY * dt

		if p.X < 0 {
			p.X = screenWidth
		}
		if p.X > screenWidth {
			p.X = 0
		}
		if p.Y < 0 {
			p.Y = screenHeight
		}
		if p.Y > screenHeight {
			p.Y = 0
		}

		p.Age += dt
		if p.Age >= p.TTL {
			p.Active = false
			continue
		}

		t := p.Age / p.TTL
		p.R = lerp(p.StartR, p.EndR, t)
		p.G = lerp(p.StartG, p.EndG, t)
		p.B = lerp(p.StartB, p.EndB, t)
		p.A = uint8(255 * (1 - t*t))
		p.Radius = p.StartSize + (p.EndSize-p.StartSize)*t
	}

	return nil
}

func lerp(a, b uint8, t float32) uint8 {
	return uint8(float32(a) + (float32(b)-float32(a))*t)
}

func (g *Game) switchPreset(index int) {
	if index < 0 || index >= len(presets) {
		return
	}
	g.audio.PlayPresetChange(index)
	g.currentPreset = index
	g.preset = presets[index]
	for i := range g.particles {
		g.particles[i].Active = false
	}
}

func (g *Game) cycleQuality() {
	switch g.quality {
	case QualityLow:
		g.quality = QualityMedium
	case QualityMedium:
		g.quality = QualityHigh
	case QualityHigh:
		g.quality = QualityLow
	}
	g.qualitySettings = qualityPresets[g.quality]
}

func (g *Game) nextPreset() {
	g.switchPreset((g.currentPreset + 1) % len(presets))
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 20, 255})

	preset := g.preset

	for i := range g.particles {
		p := &g.particles[i]
		if !p.Active {
			continue
		}

		// Glow effect (if enabled and quality allows)
		if g.qualitySettings.GlowEnabled && preset.GlowIntensity > 0 {
			glowAlpha := uint8(float32(p.A) * preset.GlowIntensity * 0.3)
			glowCol := color.RGBA{preset.GlowR, preset.GlowG, preset.GlowB, glowAlpha}

			// Draw glow passes
			for pass := 0; pass < g.qualitySettings.GlowPasses; pass++ {
				glowSize := p.Radius * (2.0 + float32(pass)*1.2)
				drawCircleFast(screen, p.X, p.Y, glowSize, glowCol)
			}
		}

		// Main particle
		col := color.RGBA{p.R, p.G, p.B, p.A}
		drawCircleFast(screen, p.X, p.Y, p.Radius, col)
	}

	// Debug overlay
	if g.showDebug {
		soundStatus := "🔊"
		if g.audio.IsMuted() {
			soundStatus = "🔇"
		}
		info := fmt.Sprintf("FPS: %.0f | Particles: %d/%d | %s | Quality: %s %s",
			ebiten.ActualFPS(), g.activeCount, g.qualitySettings.MaxParticles,
			g.preset.Name, g.quality.String(), soundStatus)
		if g.lockedMode == 1 {
			info += " [LOCKED]"
		} else if g.lockedMode == -1 {
			info += " [REPEL]"
		}
		ebitenutil.DebugPrint(screen, info)

		// Controls help
		helpText := "Q: Quality | M: Mute | 1-5: Presets"
		if g.isMobile {
			helpText = "Tap: Attract | 2-Finger: Repel | Double-Tap: Lock"
		}
		ebitenutil.DebugPrintAt(screen, helpText, 10, screenHeight-20)
	}

	// Mobile UI: Touch buttons overlay
	if g.isMobile {
		g.drawMobileUI(screen)
	}
}

func (g *Game) drawMobileUI(screen *ebiten.Image) {
	// Preset buttons at bottom
	btnWidth := 50
	btnHeight := 40
	startX := screenWidth/2 - (len(presets)*btnWidth)/2
	y := screenHeight - btnHeight - 5

	for i, p := range presets {
		x := startX + i*btnWidth
		btnCol := color.RGBA{60, 60, 80, 200}
		if i == g.currentPreset {
			btnCol = color.RGBA{100, 100, 200, 255}
		}

		// Draw button background
		for dy := 0; dy < btnHeight; dy++ {
			for dx := 0; dx < btnWidth-2; dx++ {
				screen.Set(x+dx, y+dy, btnCol)
			}
		}

		// Check if button is tapped
		for _, tid := range g.touchIDs {
			tx, ty := ebiten.TouchPosition(tid)
			if tx >= x && tx < x+btnWidth && ty >= y && ty < y+btnHeight {
				g.switchPreset(i)
			}
		}

		// Button label (first letter)
		label := string(p.Name[0])
		ebitenutil.DebugPrintAt(screen, label, x+btnWidth/2-4, y+btnHeight/2-8)
	}

	// Quality button top-right
	qx, qy := screenWidth-60, 10
	qCol := color.RGBA{60, 60, 80, 200}
	for dy := 0; dy < 30; dy++ {
		for dx := 0; dx < 50; dx++ {
			screen.Set(qx+dx, qy+dy, qCol)
		}
	}
	ebitenutil.DebugPrintAt(screen, g.quality.String()[:1], qx+20, qy+8)

	// Check quality tap
	for _, tid := range g.touchIDs {
		tx, ty := ebiten.TouchPosition(tid)
		if tx >= qx && tx < qx+50 && ty >= qy && ty < qy+30 {
			g.cycleQuality()
		}
	}

	// Mute button
	mx, my := screenWidth-120, 10
	mCol := color.RGBA{60, 60, 80, 200}
	for dy := 0; dy < 30; dy++ {
		for dx := 0; dx < 50; dx++ {
			screen.Set(mx+dx, my+dy, mCol)
		}
	}
	mLabel := "🔊"
	if g.audio.IsMuted() {
		mLabel = "🔇"
	}
	ebitenutil.DebugPrintAt(screen, mLabel, mx+18, my+8)

	// Check mute tap
	for _, tid := range g.touchIDs {
		tx, ty := ebiten.TouchPosition(tid)
		if tx >= mx && tx < mx+50 && ty >= my && ty < my+30 {
			g.audio.ToggleMute()
		}
	}
}

// Fast circle drawing using Ebitengine's optimized methods
func drawCircleFast(screen *ebiten.Image, cx, cy, radius float32, col color.RGBA) {
	r := int(radius)
	if r < 1 {
		r = 1
	}
	// Use squared distance check for filled circle
	r2 := r * r
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r2 {
				px, py := int(cx)+x, int(cy)+y
				if px >= 0 && px < screenWidth && py >= 0 && py < screenHeight {
					screen.Set(px, py, col)
				}
			}
		}
	}
}

func drawCircle(screen *ebiten.Image, cx, cy, radius float32, col color.RGBA) {
	drawCircleFast(screen, cx, cy, radius, col)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Fixed logical size - Ebitengine handles coordinate transformation
	return screenWidth, screenHeight
}

// Global game reference for JS callbacks
var gameInstance *Game

// setParticleCount is called from JavaScript to update max particles
func setParticleCount(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		count := args[0].Int()
		if count < 500 {
			count = 500
		}
		if count > 20000 {
			count = 20000
		}
		maxParticles = count

		// Resize particle slice if game exists
		if gameInstance != nil && len(gameInstance.particles) < maxParticles {
			newParticles := make([]Particle, maxParticles)
			copy(newParticles, gameInstance.particles)
			gameInstance.particles = newParticles
		}
	}
	return nil
}

// getParticleCount returns current max particles for JS
func getParticleCount(this js.Value, args []js.Value) interface{} {
	return maxParticles
}

// getActiveParticleCount returns current active particle count
func getActiveParticleCount(this js.Value, args []js.Value) interface{} {
	if gameInstance != nil {
		return gameInstance.activeCount
	}
	return 0
}

// setQuality sets quality level from JS (0=Low, 1=Medium, 2=High)
func setQualityLevel(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 && gameInstance != nil {
		level := args[0].Int()
		switch level {
		case 0:
			gameInstance.quality = QualityLow
		case 1:
			gameInstance.quality = QualityMedium
		case 2:
			gameInstance.quality = QualityHigh
		}
		gameInstance.qualitySettings = qualityPresets[gameInstance.quality]
	}
	return nil
}

func main() {
	// Register JavaScript API
	js.Global().Set("setParticleCount", js.FuncOf(setParticleCount))
	js.Global().Set("getParticleCount", js.FuncOf(getParticleCount))
	js.Global().Set("getActiveParticleCount", js.FuncOf(getActiveParticleCount))
	js.Global().Set("setQualityLevel", js.FuncOf(setQualityLevel))

	// Set fixed window size for WASM - CSS controls actual display
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particle Symphony - ECS Showcase")
	// Enable resizing mode for proper coordinate handling
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewGame()
	gameInstance = game // Store for JS callbacks
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
