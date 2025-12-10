//go:build js && wasm

package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	maxParticles = 10000
)

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
}

var presets = []Preset{
	{
		Name: "Galaxy", StartR: 100, StartG: 150, StartB: 255,
		EndR: 255, EndG: 100, EndB: 200,
		MinSize: 1.5, MaxSize: 4.0, MinTTL: 4.0, MaxTTL: 7.0,
		MinVel: -30, MaxVel: 30, SpawnPattern: "center", SpawnRate: 150,
	},
	{
		Name: "Firework", StartR: 255, StartG: 200, StartB: 50,
		EndR: 255, EndG: 50, EndB: 0,
		MinSize: 2.0, MaxSize: 5.0, MinTTL: 1.5, MaxTTL: 3.0,
		MinVel: -150, MaxVel: 150, SpawnPattern: "center", SpawnRate: 200,
	},
	{
		Name: "Swarm", StartR: 50, StartG: 255, StartB: 100,
		EndR: 0, EndG: 150, EndB: 50,
		MinSize: 1.0, MaxSize: 3.0, MinTTL: 5.0, MaxTTL: 8.0,
		MinVel: -20, MaxVel: 20, SpawnPattern: "random", SpawnRate: 100,
	},
	{
		Name: "Fountain", StartR: 100, StartG: 200, StartB: 255,
		EndR: 50, EndG: 100, EndB: 200,
		MinSize: 2.0, MaxSize: 4.0, MinTTL: 2.0, MaxTTL: 4.0,
		MinVel: -80, MaxVel: 80, SpawnPattern: "bottom", SpawnRate: 180,
	},
	{
		Name: "Chaos", StartR: 255, StartG: 50, StartB: 50,
		EndR: 50, EndG: 50, EndB: 255,
		MinSize: 1.0, MaxSize: 6.0, MinTTL: 2.0, MaxTTL: 5.0,
		MinVel: -100, MaxVel: 100, SpawnPattern: "edges", SpawnRate: 250,
	},
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
}

func NewGame() *Game {
	g := &Game{
		particles:     make([]Particle, maxParticles),
		rng:           rand.New(rand.NewSource(time.Now().UnixNano())),
		currentPreset: 0,
		showDebug:     true,
		lockedMode:    0,
	}
	g.preset = presets[0]
	return g
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

	g.mouseX, g.mouseY = ebiten.CursorPosition()

	now := time.Now()
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

	if g.lockedMode == 1 {
		g.attractorMass = 8000
	} else if g.lockedMode == -1 {
		g.attractorMass = -8000
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.attractorMass = 5000
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.attractorMass = -5000
	} else {
		g.attractorMass = 0
	}

	keys := []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5}
	for i, key := range keys {
		if inpututil.IsKeyJustPressed(key) {
			g.switchPreset(i)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		g.showDebug = !g.showDebug
	}

	g.spawnTimer += dt
	spawnInterval := 1.0 / float32(g.preset.SpawnRate)
	for g.spawnTimer >= spawnInterval {
		g.spawnTimer -= spawnInterval
		g.spawnParticle()
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
	g.currentPreset = index
	g.preset = presets[index]
	for i := range g.particles {
		g.particles[i].Active = false
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 20, 255})

	for i := range g.particles {
		p := &g.particles[i]
		if !p.Active {
			continue
		}
		col := color.RGBA{p.R, p.G, p.B, p.A}
		drawCircle(screen, p.X, p.Y, p.Radius, col)
	}

	if g.showDebug {
		info := fmt.Sprintf("FPS: %.0f\nParticles: %d\nPreset: %s\nMouse: (%d, %d)",
			ebiten.ActualFPS(), g.activeCount, g.preset.Name, g.mouseX, g.mouseY)
		if g.lockedMode == 1 {
			info += "\n[ATTRACT LOCKED]"
		} else if g.lockedMode == -1 {
			info += "\n[REPEL LOCKED]"
		}
		ebitenutil.DebugPrint(screen, info)
		ebitenutil.DebugPrintAt(screen, "F3: Toggle Debug | LMB: Attract | RMB: Repel | 1-5: Presets | 2x Click: Lock", 10, screenHeight-20)
	}
}

func drawCircle(screen *ebiten.Image, cx, cy, radius float32, col color.RGBA) {
	r := int(radius)
	if r < 1 {
		r = 1
	}
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				px, py := int(cx)+x, int(cy)+y
				if px >= 0 && px < screenWidth && py >= 0 && py < screenHeight {
					screen.Set(px, py, col)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particle Symphony - ECS Showcase")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
