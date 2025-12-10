---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8]
inputDocuments: ['docs/prd.md']
workflowType: 'architecture'
lastStep: 8
project_name: 'showcase'
user_name: 'Deltatree'
date: '2025-12-10'
yolo_mode: true
---

# Architecture Decision Document - Particle Symphony

**Autor:** Deltatree  
**Datum:** 10. Dezember 2025  
**Status:** Finalisiert (YOLO MODE)  
**Version:** 1.0.0

---

## 1. Kontext & Systemübersicht

### 1.1 Architektur-Vision

**Particle Symphony** ist ein ECS-Showcase, der die Eleganz und Leistungsfähigkeit des `andygeiss/ecs` Frameworks demonstriert. Die Architektur folgt strikt dem **Entity-Component-System Pattern** und maximiert:

- **Klarheit:** Jede Komponente hat eine einzige Verantwortung
- **Performance:** Data-Oriented Design für Cache-Effizienz
- **Erweiterbarkeit:** Neue Features = neue Components + Systems
- **Lernwert:** Jeder Teil des Codes ist ein ECS-Lehrbuch

### 1.2 Systemkontext-Diagramm

```
┌─────────────────────────────────────────────────────────────────────────┐
│                         PARTICLE SYMPHONY                                │
│                                                                          │
│  ┌──────────┐    ┌──────────────────────────────────────────────────┐   │
│  │  User    │───▶│                  ECS Engine                       │   │
│  │ (Mouse/  │    │  ┌─────────────┐  ┌─────────────┐  ┌──────────┐  │   │
│  │ Keyboard)│    │  │Entity       │  │System       │  │Component │  │   │
│  └──────────┘    │  │Manager      │  │Manager      │  │Registry  │  │   │
│                  │  └─────────────┘  └─────────────┘  └──────────┘  │   │
│                  └──────────────────────────────────────────────────┘   │
│                                        │                                 │
│                                        ▼                                 │
│                  ┌──────────────────────────────────────────────────┐   │
│                  │              Raylib Renderer                      │   │
│                  │         (Window, Graphics, Input)                 │   │
│                  └──────────────────────────────────────────────────┘   │
│                                        │                                 │
│                                        ▼                                 │
│                  ┌──────────────────────────────────────────────────┐   │
│                  │                 Display                           │   │
│                  │         (60 FPS Visual Output)                    │   │
│                  └──────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────┘
```

### 1.3 Externe Abhängigkeiten

| Abhängigkeit | Version | Zweck | Risiko |
|--------------|---------|-------|--------|
| `github.com/andygeiss/ecs` | v0.3.12+ | ECS Framework | **Core** - Showcase-Fokus |
| `github.com/gen2brain/raylib-go/raylib` | v5.0+ | Rendering, Input, Window | Mittel - CGO erforderlich |

---

## 2. Architektur-Starter: Projekt-Setup

### 2.1 Go Module Initialisierung

```go
// go.mod
module github.com/deltatree/showcase

go 1.23

require (
    github.com/andygeiss/ecs v0.3.12
    github.com/gen2brain/raylib-go/raylib v0.0.0-20241117153000
)
```

### 2.2 Entry Point Design

```go
// main.go - Minimalistischer Entry Point
package main

import (
    "github.com/andygeiss/ecs"
    "github.com/deltatree/showcase/systems"
)

func main() {
    // ECS Setup
    em := ecs.NewEntityManager()
    sm := ecs.NewSystemManager()
    
    // Systems in korrekter Reihenfolge registrieren
    sm.Add(
        systems.NewInputSystem(),
        systems.NewEmitterSystem(),
        systems.NewGravitySystem(),
        systems.NewPhysicsSystem(),
        systems.NewLifetimeSystem(),
        systems.NewColorSystem(),
        systems.NewRenderSystem(),
    )
    
    // Engine starten
    engine := ecs.NewDefaultEngine(em, sm)
    engine.Setup()
    defer engine.Teardown()
    engine.Run()
}
```

---

## 3. Architektur-Entscheidungen (ADRs)

### ADR-001: Pure ECS ohne Game-State-Manager

**Kontext:** Viele Game Engines nutzen einen separaten Game-State-Manager für Menüs, Pause, etc.

**Entscheidung:** Wir verwenden **ausschließlich ECS** für alles - inklusive UI-State.

**Begründung:**
- Maximale Demonstration der ECS-Fähigkeiten
- Keine zusätzliche Komplexität
- State-Wechsel = Entity-Manipulation

**Konsequenzen:**
- ✅ Konsistente Architektur
- ✅ Besserer Lernwert
- ⚠️ Weniger Flexibilität für komplexe UI

---

### ADR-002: Bitmask-basierte Component-Registry

**Kontext:** Components brauchen eine Identifikation für effiziente Filterung.

**Entscheidung:** Wir nutzen das **native Bitmask-System** von andygeiss/ecs.

**Implementierung:**

```go
// components/masks.go
package components

const (
    MaskPosition     = uint64(1 << 0)
    MaskVelocity     = uint64(1 << 1)
    MaskAcceleration = uint64(1 << 2)
    MaskColor        = uint64(1 << 3)
    MaskLifetime     = uint64(1 << 4)
    MaskMass         = uint64(1 << 5)
    MaskSize         = uint64(1 << 6)
    MaskEmitter      = uint64(1 << 7)
    MaskAttractor    = uint64(1 << 8)
    MaskParticle     = uint64(1 << 9)  // Tag-Component
)

// Häufig verwendete Kombinationen
const (
    MaskMovable   = MaskPosition | MaskVelocity
    MaskPhysics   = MaskMovable | MaskAcceleration
    MaskRenderable = MaskPosition | MaskColor | MaskSize
    MaskFullParticle = MaskPhysics | MaskColor | MaskLifetime | MaskSize | MaskParticle
)
```

**Begründung:**
- O(1) Lookup für Entity-Filterung
- Memory-effizient (64 Components möglich)
- Native Unterstützung im Framework

---

### ADR-003: System-Reihenfolge als Architektur-Constraint

**Kontext:** ECS Systems müssen in einer definierten Reihenfolge ausgeführt werden.

**Entscheidung:** Die System-Pipeline ist **fest definiert** und dokumentiert.

**Pipeline-Definition:**

```
┌─────────────────────────────────────────────────────────────────┐
│                    SYSTEM EXECUTION ORDER                        │
├─────────────────────────────────────────────────────────────────┤
│  Phase 1: INPUT                                                  │
│  ├── InputSystem      → Liest Maus/Keyboard, aktualisiert State │
│                                                                  │
│  Phase 2: SPAWN                                                  │
│  ├── EmitterSystem    → Erstellt neue Partikel-Entities         │
│                                                                  │
│  Phase 3: PHYSICS                                                │
│  ├── GravitySystem    → Berechnet Acceleration aus Attraktoren  │
│  ├── PhysicsSystem    → Integriert Velocity und Position        │
│                                                                  │
│  Phase 4: LIFECYCLE                                              │
│  ├── LifetimeSystem   → Altert Entities, markiert zum Löschen   │
│  ├── CleanupSystem    → Entfernt "tote" Entities                │
│                                                                  │
│  Phase 5: VISUAL                                                 │
│  ├── ColorSystem      → Interpoliert Farben basierend auf Age   │
│  ├── RenderSystem     → Zeichnet alle Entities                  │
└─────────────────────────────────────────────────────────────────┘
```

**Begründung:**
- Deterministische Ausführung
- Keine Race Conditions
- Klare Verantwortlichkeiten

---

### ADR-004: Raylib-Integration über Wrapper-System

**Kontext:** Raylib-Calls sollten isoliert sein für Testbarkeit.

**Entscheidung:** Raylib-Funktionen werden **nur in RenderSystem und InputSystem** aufgerufen.

**Implementierung:**

```go
// systems/render.go
type renderSystem struct {
    width, height int32
    title         string
}

func (s *renderSystem) Setup() {
    rl.InitWindow(s.width, s.height, s.title)
    rl.SetTargetFPS(60)
}

func (s *renderSystem) Process(em ecs.EntityManager) int {
    if rl.WindowShouldClose() {
        return ecs.StateEngineStop
    }
    
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)
    
    // Partikel rendern
    for _, e := range em.FilterByMask(components.MaskRenderable) {
        pos := e.Get(components.MaskPosition).(*components.Position)
        col := e.Get(components.MaskColor).(*components.Color)
        size := e.Get(components.MaskSize).(*components.Size)
        
        rl.DrawCircle(int32(pos.X), int32(pos.Y), size.Radius, 
            rl.NewColor(col.R, col.G, col.B, col.A))
    }
    
    // FPS Overlay
    rl.DrawFPS(10, 10)
    rl.EndDrawing()
    
    return ecs.StateEngineContinue
}

func (s *renderSystem) Teardown() {
    rl.CloseWindow()
}
```

**Begründung:**
- Testbarkeit: Andere Systems können ohne Raylib getestet werden
- Separation of Concerns
- Einfacher Austausch des Renderers

---

### ADR-005: Konfiguration via JSON mit Sensiblen Defaults

**Kontext:** Parameter sollten ohne Recompile änderbar sein.

**Entscheidung:** JSON-Config mit **Struct-Defaults** als Fallback.

**Implementierung:**

```go
// internal/config/config.go
package config

type Config struct {
    Window    WindowConfig    `json:"window"`
    Particles ParticleConfig  `json:"particles"`
    Physics   PhysicsConfig   `json:"physics"`
    Presets   []PresetConfig  `json:"presets"`
}

type WindowConfig struct {
    Width  int32  `json:"width"`
    Height int32  `json:"height"`
    Title  string `json:"title"`
    FPS    int32  `json:"fps"`
}

type ParticleConfig struct {
    MaxCount      int     `json:"maxCount"`
    DefaultSize   float32 `json:"defaultSize"`
    SpawnRate     int     `json:"spawnRate"`
    DefaultTTL    float32 `json:"defaultTTL"`
}

type PhysicsConfig struct {
    Gravity       float32 `json:"gravity"`
    Damping       float32 `json:"damping"`
    MaxVelocity   float32 `json:"maxVelocity"`
}

// Default liefert Sensible Defaults
func Default() *Config {
    return &Config{
        Window: WindowConfig{
            Width:  1280,
            Height: 720,
            Title:  "Particle Symphony - ECS Showcase",
            FPS:    60,
        },
        Particles: ParticleConfig{
            MaxCount:    10000,
            DefaultSize: 3.0,
            SpawnRate:   100,
            DefaultTTL:  5.0,
        },
        Physics: PhysicsConfig{
            Gravity:     0.0,
            Damping:     0.99,
            MaxVelocity: 500.0,
        },
    }
}
```

---

### ADR-006: Preset-System als Entity-Templates

**Kontext:** Verschiedene "Modi" sollen schnell umschaltbar sein.

**Entscheidung:** Presets sind **Funktionen**, die den EntityManager konfigurieren.

**Implementierung:**

```go
// presets/presets.go
package presets

type Preset interface {
    Name() string
    Apply(em ecs.EntityManager, config *config.Config)
    Description() string
}

// presets/galaxy.go
type galaxyPreset struct{}

func (p *galaxyPreset) Name() string { return "Galaxy" }

func (p *galaxyPreset) Description() string {
    return "Spiral galaxy with central attractor"
}

func (p *galaxyPreset) Apply(em ecs.EntityManager, cfg *config.Config) {
    // Zentraler Attraktor
    em.Add(ecs.NewEntity("center", []ecs.Component{
        components.NewPosition().WithX(640).WithY(360),
        components.NewMass().WithValue(1000),
        components.NewAttractor(),
    }))
    
    // Spiralarme mit Partikeln
    for i := 0; i < 1000; i++ {
        angle := float32(i) * 0.1
        radius := float32(i) * 0.5
        x := 640 + radius*float32(math.Cos(float64(angle)))
        y := 360 + radius*float32(math.Sin(float64(angle)))
        
        em.Add(createParticle(x, y, angle))
    }
}
```

---

## 4. Component-Architektur

### 4.1 Component Interface

Jede Component implementiert das `ecs.Component` Interface:

```go
type Component interface {
    Mask() uint64
}
```

### 4.2 Component Definitionen

#### Position Component

```go
// components/position.go
package components

type Position struct {
    X, Y float32
}

func (p *Position) Mask() uint64 { return MaskPosition }

func NewPosition() *Position { return &Position{} }

func (p *Position) WithX(x float32) *Position { p.X = x; return p }
func (p *Position) WithY(y float32) *Position { p.Y = y; return p }
func (p *Position) With(x, y float32) *Position { p.X = x; p.Y = y; return p }
```

#### Velocity Component

```go
// components/velocity.go
package components

type Velocity struct {
    X, Y float32
}

func (v *Velocity) Mask() uint64 { return MaskVelocity }

func NewVelocity() *Velocity { return &Velocity{} }

func (v *Velocity) WithX(x float32) *Velocity { v.X = x; return v }
func (v *Velocity) WithY(y float32) *Velocity { v.Y = y; return v }

// Magnitude berechnet die Länge des Velocity-Vektors
func (v *Velocity) Magnitude() float32 {
    return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}
```

#### Acceleration Component

```go
// components/acceleration.go
package components

type Acceleration struct {
    X, Y float32
}

func (a *Acceleration) Mask() uint64 { return MaskAcceleration }

func NewAcceleration() *Acceleration { return &Acceleration{} }

func (a *Acceleration) Reset() { a.X = 0; a.Y = 0 }
func (a *Acceleration) Add(x, y float32) { a.X += x; a.Y += y }
```

#### Color Component

```go
// components/color.go
package components

type Color struct {
    R, G, B, A uint8
    // Für Interpolation
    StartR, StartG, StartB, StartA uint8
    EndR, EndG, EndB, EndA         uint8
}

func (c *Color) Mask() uint64 { return MaskColor }

func NewColor() *Color {
    return &Color{R: 255, G: 255, B: 255, A: 255}
}

func (c *Color) WithRGBA(r, g, b, a uint8) *Color {
    c.R, c.G, c.B, c.A = r, g, b, a
    c.StartR, c.StartG, c.StartB, c.StartA = r, g, b, a
    return c
}

func (c *Color) WithGradient(sr, sg, sb, sa, er, eg, eb, ea uint8) *Color {
    c.StartR, c.StartG, c.StartB, c.StartA = sr, sg, sb, sa
    c.EndR, c.EndG, c.EndB, c.EndA = er, eg, eb, ea
    c.R, c.G, c.B, c.A = sr, sg, sb, sa
    return c
}
```

#### Lifetime Component

```go
// components/lifetime.go
package components

type Lifetime struct {
    TTL     float32  // Time To Live (Sekunden)
    Age     float32  // Aktuelles Alter
    Expired bool     // Zum Löschen markiert
}

func (l *Lifetime) Mask() uint64 { return MaskLifetime }

func NewLifetime() *Lifetime { return &Lifetime{TTL: 5.0} }

func (l *Lifetime) WithTTL(ttl float32) *Lifetime { 
    l.TTL = ttl 
    return l 
}

func (l *Lifetime) Progress() float32 {
    if l.TTL <= 0 { return 1.0 }
    return l.Age / l.TTL
}
```

#### Size Component

```go
// components/size.go
package components

type Size struct {
    Radius    float32
    StartSize float32
    EndSize   float32
}

func (s *Size) Mask() uint64 { return MaskSize }

func NewSize() *Size { return &Size{Radius: 3.0, StartSize: 3.0, EndSize: 1.0} }

func (s *Size) WithRadius(r float32) *Size { 
    s.Radius = r
    s.StartSize = r
    return s
}
```

#### Mass Component

```go
// components/mass.go
package components

type Mass struct {
    Value float32
}

func (m *Mass) Mask() uint64 { return MaskMass }

func NewMass() *Mass { return &Mass{Value: 1.0} }

func (m *Mass) WithValue(v float32) *Mass { m.Value = v; return m }
```

---

## 5. System-Architektur

### 5.1 System Interface

Jedes System implementiert das `ecs.System` Interface:

```go
type System interface {
    Setup()
    Process(em EntityManager) int
    Teardown()
}
```

### 5.2 InputSystem

```go
// systems/input.go
package systems

type inputSystem struct {
    mouseAttractorID string
    attracting       bool
    repelling        bool
}

func NewInputSystem() ecs.System {
    return &inputSystem{mouseAttractorID: "mouse-attractor"}
}

func (s *inputSystem) Setup() {}

func (s *inputSystem) Process(em ecs.EntityManager) int {
    mouseX := float32(rl.GetMouseX())
    mouseY := float32(rl.GetMouseY())
    
    // Maus-Attraktor finden oder erstellen
    attractor := em.Get(s.mouseAttractorID)
    if attractor == nil {
        attractor = ecs.NewEntity(s.mouseAttractorID, []ecs.Component{
            components.NewPosition().With(mouseX, mouseY),
            components.NewMass().WithValue(0), // Inaktiv bis Mausklick
            components.NewAttractor(),
        })
        em.Add(attractor)
    }
    
    // Position aktualisieren
    pos := attractor.Get(components.MaskPosition).(*components.Position)
    pos.X, pos.Y = mouseX, mouseY
    
    // Masse basierend auf Maustasten
    mass := attractor.Get(components.MaskMass).(*components.Mass)
    if rl.IsMouseButtonDown(rl.MouseLeftButton) {
        mass.Value = 500  // Anziehung
    } else if rl.IsMouseButtonDown(rl.MouseRightButton) {
        mass.Value = -500 // Abstoßung
    } else {
        mass.Value = 0    // Inaktiv
    }
    
    // Preset-Wechsel via Tasten 1-5
    if rl.IsKeyPressed(rl.KeyOne)   { switchPreset(em, 0) }
    if rl.IsKeyPressed(rl.KeyTwo)   { switchPreset(em, 1) }
    if rl.IsKeyPressed(rl.KeyThree) { switchPreset(em, 2) }
    
    return ecs.StateEngineContinue
}

func (s *inputSystem) Teardown() {}
```

### 5.3 EmitterSystem

```go
// systems/emitter.go
package systems

type emitterSystem struct {
    spawnRate   int
    spawnTimer  float32
    maxParticles int
    currentCount int
}

func NewEmitterSystem(spawnRate, maxParticles int) ecs.System {
    return &emitterSystem{
        spawnRate:    spawnRate,
        maxParticles: maxParticles,
    }
}

func (s *emitterSystem) Setup() {}

func (s *emitterSystem) Process(em ecs.EntityManager) int {
    dt := rl.GetFrameTime()
    s.spawnTimer += dt
    
    // Aktuelle Partikelzahl ermitteln
    particles := em.FilterByMask(components.MaskParticle)
    s.currentCount = len(particles)
    
    // Spawn-Intervall
    spawnInterval := 1.0 / float32(s.spawnRate)
    
    for s.spawnTimer >= spawnInterval && s.currentCount < s.maxParticles {
        s.spawnTimer -= spawnInterval
        s.spawnParticle(em)
        s.currentCount++
    }
    
    return ecs.StateEngineContinue
}

func (s *emitterSystem) spawnParticle(em ecs.EntityManager) {
    // Zufällige Position am Bildschirmrand
    x := rand.Float32() * 1280
    y := rand.Float32() * 720
    
    // Zufällige Geschwindigkeit
    vx := (rand.Float32() - 0.5) * 100
    vy := (rand.Float32() - 0.5) * 100
    
    id := fmt.Sprintf("particle-%d", time.Now().UnixNano())
    
    em.Add(ecs.NewEntity(id, []ecs.Component{
        components.NewPosition().With(x, y),
        components.NewVelocity().WithX(vx).WithY(vy),
        components.NewAcceleration(),
        components.NewColor().WithGradient(255, 100, 50, 255, 50, 50, 255, 0),
        components.NewLifetime().WithTTL(3.0 + rand.Float32()*2.0),
        components.NewSize().WithRadius(2.0 + rand.Float32()*3.0),
        &components.Particle{}, // Tag-Component
    }))
}

func (s *emitterSystem) Teardown() {}
```

### 5.4 GravitySystem

```go
// systems/gravity.go
package systems

type gravitySystem struct{}

func NewGravitySystem() ecs.System {
    return &gravitySystem{}
}

func (s *gravitySystem) Setup() {}

func (s *gravitySystem) Process(em ecs.EntityManager) int {
    // Alle Attraktoren finden
    attractors := em.FilterByMask(components.MaskPosition | components.MaskMass | components.MaskAttractor)
    
    // Alle beweglichen Partikel
    particles := em.FilterByMask(components.MaskPosition | components.MaskAcceleration | components.MaskParticle)
    
    for _, particle := range particles {
        pPos := particle.Get(components.MaskPosition).(*components.Position)
        pAcc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
        
        // Reset Acceleration
        pAcc.Reset()
        
        for _, attractor := range attractors {
            aPos := attractor.Get(components.MaskPosition).(*components.Position)
            aMass := attractor.Get(components.MaskMass).(*components.Mass)
            
            if aMass.Value == 0 {
                continue // Inaktiver Attraktor
            }
            
            // Richtungsvektor
            dx := aPos.X - pPos.X
            dy := aPos.Y - pPos.Y
            
            // Distanz (mit Minimum um Division durch 0 zu vermeiden)
            dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
            if dist < 10 {
                dist = 10
            }
            
            // Gravitationskraft: F = G * m1 * m2 / r²
            // Vereinfacht: F = mass / dist²
            force := aMass.Value / (dist * dist) * 100
            
            // Normalisieren und auf Acceleration anwenden
            pAcc.Add(dx/dist*force, dy/dist*force)
        }
    }
    
    return ecs.StateEngineContinue
}

func (s *gravitySystem) Teardown() {}
```

### 5.5 PhysicsSystem

```go
// systems/physics.go
package systems

type physicsSystem struct {
    damping     float32
    maxVelocity float32
    width       float32
    height      float32
}

func NewPhysicsSystem(damping, maxVelocity, width, height float32) ecs.System {
    return &physicsSystem{
        damping:     damping,
        maxVelocity: maxVelocity,
        width:       width,
        height:      height,
    }
}

func (s *physicsSystem) Setup() {}

func (s *physicsSystem) Process(em ecs.EntityManager) int {
    dt := rl.GetFrameTime()
    
    entities := em.FilterByMask(components.MaskPosition | components.MaskVelocity)
    
    for _, e := range entities {
        pos := e.Get(components.MaskPosition).(*components.Position)
        vel := e.Get(components.MaskVelocity).(*components.Velocity)
        
        // Acceleration anwenden (falls vorhanden)
        if acc := e.Get(components.MaskAcceleration); acc != nil {
            a := acc.(*components.Acceleration)
            vel.X += a.X * dt
            vel.Y += a.Y * dt
        }
        
        // Damping
        vel.X *= s.damping
        vel.Y *= s.damping
        
        // Max Velocity begrenzen
        mag := vel.Magnitude()
        if mag > s.maxVelocity {
            vel.X = vel.X / mag * s.maxVelocity
            vel.Y = vel.Y / mag * s.maxVelocity
        }
        
        // Position aktualisieren
        pos.X += vel.X * dt
        pos.Y += vel.Y * dt
        
        // Screen Wrapping
        if pos.X < 0 { pos.X = s.width }
        if pos.X > s.width { pos.X = 0 }
        if pos.Y < 0 { pos.Y = s.height }
        if pos.Y > s.height { pos.Y = 0 }
    }
    
    return ecs.StateEngineContinue
}

func (s *physicsSystem) Teardown() {}
```

### 5.6 LifetimeSystem

```go
// systems/lifetime.go
package systems

type lifetimeSystem struct{}

func NewLifetimeSystem() ecs.System {
    return &lifetimeSystem{}
}

func (s *lifetimeSystem) Setup() {}

func (s *lifetimeSystem) Process(em ecs.EntityManager) int {
    dt := rl.GetFrameTime()
    
    entities := em.FilterByMask(components.MaskLifetime)
    
    var toRemove []string
    
    for _, e := range entities {
        life := e.Get(components.MaskLifetime).(*components.Lifetime)
        life.Age += dt
        
        if life.Age >= life.TTL {
            life.Expired = true
            toRemove = append(toRemove, e.ID())
        }
    }
    
    // Abgelaufene Entities entfernen
    for _, id := range toRemove {
        em.Remove(id)
    }
    
    return ecs.StateEngineContinue
}

func (s *lifetimeSystem) Teardown() {}
```

### 5.7 ColorSystem

```go
// systems/color.go
package systems

type colorSystem struct{}

func NewColorSystem() ecs.System {
    return &colorSystem{}
}

func (s *colorSystem) Setup() {}

func (s *colorSystem) Process(em ecs.EntityManager) int {
    entities := em.FilterByMask(components.MaskColor | components.MaskLifetime)
    
    for _, e := range entities {
        col := e.Get(components.MaskColor).(*components.Color)
        life := e.Get(components.MaskLifetime).(*components.Lifetime)
        
        t := life.Progress() // 0.0 bis 1.0
        
        // Lineare Interpolation
        col.R = lerp(col.StartR, col.EndR, t)
        col.G = lerp(col.StartG, col.EndG, t)
        col.B = lerp(col.StartB, col.EndB, t)
        col.A = lerp(col.StartA, col.EndA, t)
    }
    
    // Size auch interpolieren
    sizeEntities := em.FilterByMask(components.MaskSize | components.MaskLifetime)
    for _, e := range sizeEntities {
        size := e.Get(components.MaskSize).(*components.Size)
        life := e.Get(components.MaskLifetime).(*components.Lifetime)
        
        t := life.Progress()
        size.Radius = lerpF(size.StartSize, size.EndSize, t)
    }
    
    return ecs.StateEngineContinue
}

func lerp(a, b uint8, t float32) uint8 {
    return uint8(float32(a) + (float32(b)-float32(a))*t)
}

func lerpF(a, b, t float32) float32 {
    return a + (b-a)*t
}

func (s *colorSystem) Teardown() {}
```

### 5.8 RenderSystem

```go
// systems/render.go
package systems

type renderSystem struct {
    width, height int32
    title         string
    showDebug     bool
}

func NewRenderSystem(width, height int32, title string) ecs.System {
    return &renderSystem{
        width:  width,
        height: height,
        title:  title,
    }
}

func (s *renderSystem) Setup() {
    rl.InitWindow(s.width, s.height, s.title)
    rl.SetTargetFPS(60)
}

func (s *renderSystem) Process(em ecs.EntityManager) int {
    if rl.WindowShouldClose() {
        return ecs.StateEngineStop
    }
    
    // Toggle Debug Overlay mit F3
    if rl.IsKeyPressed(rl.KeyF3) {
        s.showDebug = !s.showDebug
    }
    
    rl.BeginDrawing()
    rl.ClearBackground(rl.NewColor(10, 10, 20, 255)) // Dunkles Blau
    
    // Partikel rendern
    particles := em.FilterByMask(components.MaskRenderable)
    for _, e := range particles {
        pos := e.Get(components.MaskPosition).(*components.Position)
        col := e.Get(components.MaskColor).(*components.Color)
        size := e.Get(components.MaskSize).(*components.Size)
        
        rl.DrawCircle(
            int32(pos.X), 
            int32(pos.Y), 
            size.Radius,
            rl.NewColor(col.R, col.G, col.B, col.A),
        )
    }
    
    // Debug Overlay
    if s.showDebug {
        rl.DrawFPS(10, 10)
        rl.DrawText(
            fmt.Sprintf("Entities: %d", len(particles)),
            10, 35, 20, rl.White,
        )
        rl.DrawText("F3: Toggle Debug | LMB: Attract | RMB: Repel", 
            10, s.height-30, 16, rl.Gray)
    }
    
    rl.EndDrawing()
    
    return ecs.StateEngineContinue
}

func (s *renderSystem) Teardown() {
    rl.CloseWindow()
}
```

---

## 6. Projektstruktur

```
showcase/
├── main.go                      # Entry Point
├── go.mod                       # Go Module Definition
├── go.sum                       # Dependency Lock
├── config.json                  # Runtime-Konfiguration
├── README.md                    # Dokumentation mit GIF
│
├── components/                  # ECS Components
│   ├── masks.go                 # Bitmask-Konstanten
│   ├── position.go
│   ├── velocity.go
│   ├── acceleration.go
│   ├── color.go
│   ├── lifetime.go
│   ├── mass.go
│   ├── size.go
│   ├── attractor.go             # Tag-Component
│   └── particle.go              # Tag-Component
│
├── systems/                     # ECS Systems
│   ├── input.go
│   ├── emitter.go
│   ├── gravity.go
│   ├── physics.go
│   ├── lifetime.go
│   ├── color.go
│   └── render.go
│
├── presets/                     # Vorkonfigurierte Modi
│   ├── presets.go               # Interface + Registry
│   ├── galaxy.go
│   ├── firework.go
│   └── swarm.go
│
├── internal/
│   └── config/
│       └── config.go            # Config-Loader
│
├── docs/
│   ├── prd.md
│   ├── architecture.md          # Dieses Dokument
│   └── screenshots/
│
└── .github/
    └── workflows/
        └── build.yaml           # CI/CD Pipeline
```

---

## 7. Validierung & Qualitätssicherung

### 7.1 Test-Strategie

| Ebene | Was wird getestet | Tool |
|-------|-------------------|------|
| **Unit** | Components, Physics-Berechnungen | `go test` |
| **Integration** | System-Interaktion | `go test` + Mock-EntityManager |
| **Visual** | Rendering-Output | Manuell + Screenshots |
| **Performance** | FPS bei verschiedenen Entity-Counts | In-App Benchmark |

### 7.2 Beispiel Unit Test

```go
// components/velocity_test.go
func TestVelocity_Magnitude(t *testing.T) {
    tests := []struct {
        name string
        x, y float32
        want float32
    }{
        {"zero", 0, 0, 0},
        {"unit x", 1, 0, 1},
        {"unit y", 0, 1, 1},
        {"3-4-5 triangle", 3, 4, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            v := NewVelocity().WithX(tt.x).WithY(tt.y)
            got := v.Magnitude()
            if math.Abs(float64(got-tt.want)) > 0.001 {
                t.Errorf("Magnitude() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 7.3 Checkliste vor Release

- [ ] Alle Tests grün
- [ ] `golangci-lint run` ohne Errors
- [ ] 60 FPS bei 10k Partikeln bestätigt
- [ ] README mit aktuellen Screenshots/GIFs
- [ ] Cross-Platform Build funktioniert (macOS, Windows, Linux)

---

## 8. Implementierungsreihenfolge

### Phase 1: Fundament (Tag 1)
1. Projekt-Setup (`go mod init`, Dependencies)
2. Alle Components implementieren
3. Basis-RenderSystem mit leerem Fenster

### Phase 2: Core-Loop (Tag 2)
4. PhysicsSystem implementieren
5. EmitterSystem für Partikel-Spawning
6. LifetimeSystem für Cleanup

### Phase 3: Interaktivität (Tag 3)
7. InputSystem für Maus-Interaktion
8. GravitySystem für Attraktor-Physik
9. ColorSystem für visuelle Effekte

### Phase 4: Polish (Tag 4)
10. Preset-System implementieren
11. Debug-Overlay verfeinern
12. Config-System finalisieren

### Phase 5: Release (Tag 5)
13. README schreiben
14. Screenshots/GIFs erstellen
15. CI/CD Pipeline aufsetzen

---

## 9. Appendix: Vollständiger main.go

```go
package main

import (
    "log"
    
    "github.com/andygeiss/ecs"
    "github.com/deltatree/showcase/components"
    "github.com/deltatree/showcase/internal/config"
    "github.com/deltatree/showcase/presets"
    "github.com/deltatree/showcase/systems"
)

func main() {
    // Config laden
    cfg, err := config.Load("config.json")
    if err != nil {
        log.Printf("Config nicht gefunden, nutze Defaults: %v", err)
        cfg = config.Default()
    }
    
    // ECS Manager initialisieren
    em := ecs.NewEntityManager()
    sm := ecs.NewSystemManager()
    
    // Systems in korrekter Reihenfolge registrieren
    sm.Add(
        systems.NewInputSystem(),
        systems.NewEmitterSystem(cfg.Particles.SpawnRate, cfg.Particles.MaxCount),
        systems.NewGravitySystem(),
        systems.NewPhysicsSystem(
            cfg.Physics.Damping,
            cfg.Physics.MaxVelocity,
            float32(cfg.Window.Width),
            float32(cfg.Window.Height),
        ),
        systems.NewLifetimeSystem(),
        systems.NewColorSystem(),
        systems.NewRenderSystem(cfg.Window.Width, cfg.Window.Height, cfg.Window.Title),
    )
    
    // Default-Preset laden
    presets.Galaxy().Apply(em, cfg)
    
    // Engine starten
    engine := ecs.NewDefaultEngine(em, sm)
    engine.Setup()
    defer engine.Teardown()
    engine.Run()
}
```

---

**🏗️ ARCHITECTURE COMPLETE - READY TO BUILD!**
