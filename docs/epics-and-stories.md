---
stepsCompleted: [1, 2, 3, 4, 5]
inputDocuments: ['docs/prd.md', 'docs/architecture.md']
workflowType: 'epics-stories'
lastStep: 5
project_name: 'showcase'
user_name: 'Deltatree'
date: '2025-12-10'
yolo_mode: true
totalEpics: 6
totalStories: 24
estimatedSprints: 3
---

# Epics & User Stories - Particle Symphony

**Autor:** Deltatree  
**Datum:** 10. Dezember 2025  
**Status:** Finalisiert (YOLO MODE)  
**Quelle:** PRD + Architecture Document

---

## Übersicht

| Metrik | Wert |
|--------|------|
| **Epics** | 6 |
| **User Stories** | 24 |
| **Story Points (geschätzt)** | 89 |
| **Sprints (geschätzt)** | 3 |

### Epic-Übersicht

| Epic | Titel | Stories | Punkte | Priorität |
|------|-------|---------|--------|-----------|
| E-001 | ECS Foundation | 5 | 18 | 🔴 MUST |
| E-002 | Physik-Engine | 4 | 15 | 🔴 MUST |
| E-003 | Interaktivität | 4 | 13 | 🔴 MUST |
| E-004 | Visual Effects | 4 | 16 | 🟡 SHOULD |
| E-005 | Preset-System | 4 | 13 | 🟡 SHOULD |
| E-006 | Audio-Reaktivität | 3 | 14 | 🟢 COULD |

---

# Epic E-001: ECS Foundation

**Beschreibung:** Grundlegendes Setup des Entity-Component-System Frameworks mit allen Core-Components.

**Business Value:** Ohne dieses Fundament kann kein einziges Partikel existieren. Dies demonstriert die Basis-Architektur von `andygeiss/ecs`.

**Akzeptanzkriterien:**
- ECS Engine startet und läuft im Game Loop
- Alle Core-Components sind implementiert und getestet
- Bitmask-System funktioniert korrekt
- Mindestens ein Entity kann erstellt und gerendert werden

---

## Story E-001-S01: Projekt-Setup & Dependencies

**Als** Entwickler  
**möchte ich** ein korrekt konfiguriertes Go-Projekt  
**damit** ich sofort mit der ECS-Entwicklung beginnen kann

**Story Points:** 2

**Akzeptanzkriterien:**
- [ ] `go.mod` mit korrektem Module-Namen erstellt
- [ ] `github.com/andygeiss/ecs` als Dependency hinzugefügt
- [ ] `github.com/gen2brain/raylib-go/raylib` als Dependency hinzugefügt
- [ ] `go mod tidy` läuft ohne Fehler
- [ ] Projekt-Struktur gemäß Architecture-Doc angelegt

**Technische Details:**
```bash
go mod init github.com/deltatree/showcase
go get github.com/andygeiss/ecs@v0.3.12
go get github.com/gen2brain/raylib-go/raylib
```

**Definition of Done:**
- [ ] `go build` erfolgreich
- [ ] Ordner-Struktur erstellt

---

## Story E-001-S02: Bitmask Component Registry

**Als** Entwickler  
**möchte ich** ein zentrales Bitmask-Register  
**damit** alle Components eindeutig identifiziert werden können

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] `components/masks.go` mit allen Mask-Konstanten
- [ ] Keine Mask-Kollisionen (jede Mask ist unique)
- [ ] Composite-Masks für häufige Kombinationen definiert
- [ ] Dokumentation für jede Mask

**Technische Details:**
```go
const (
    MaskPosition     = uint64(1 << 0)
    MaskVelocity     = uint64(1 << 1)
    // ... etc
)
```

**Definition of Done:**
- [ ] Alle 10 Masks definiert
- [ ] Composite-Masks funktionieren

---

## Story E-001-S03: Core Components Implementation

**Als** Entwickler  
**möchte ich** alle Core-Components implementiert haben  
**damit** Entities mit Daten befüllt werden können

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] Position Component mit X, Y und Builder-Pattern
- [ ] Velocity Component mit X, Y und Magnitude()
- [ ] Acceleration Component mit Reset() und Add()
- [ ] Color Component mit RGBA und Gradient-Support
- [ ] Lifetime Component mit TTL, Age, Progress()
- [ ] Size Component mit Radius und Interpolation
- [ ] Mass Component für Gravitationsberechnung
- [ ] Tag-Components: Particle, Attractor, Emitter

**Technische Details:**
- Jede Component implementiert `ecs.Component` Interface
- Builder-Pattern für fluent API: `NewPosition().WithX(100).WithY(200)`
- Alle Components in separaten Dateien

**Definition of Done:**
- [ ] 8 Component-Dateien erstellt
- [ ] Alle Components haben Mask() Methode
- [ ] Unit Tests für kritische Methoden (Magnitude, Progress)

---

## Story E-001-S04: Engine Setup & Game Loop

**Als** Benutzer  
**möchte ich** dass die Anwendung startet und ein Fenster öffnet  
**damit** ich die Partikel-Simulation sehen kann

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] `main.go` initialisiert EntityManager und SystemManager
- [ ] Raylib-Fenster öffnet sich mit korrekter Größe (1280x720)
- [ ] Game Loop läuft mit 60 FPS
- [ ] Fenster schließt sauber bei ESC oder X-Button
- [ ] Schwarzer Hintergrund wird gerendert

**Technische Details:**
```go
engine := ecs.NewDefaultEngine(em, sm)
engine.Setup()
defer engine.Teardown()
engine.Run()
```

**Definition of Done:**
- [ ] Fenster öffnet sich
- [ ] ESC schließt Anwendung
- [ ] Kein Memory Leak bei Shutdown

---

## Story E-001-S05: Basis RenderSystem

**Als** Benutzer  
**möchte ich** Entities als farbige Kreise sehen  
**damit** die Simulation visuell wird

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] RenderSystem implementiert `ecs.System` Interface
- [ ] Setup() initialisiert Raylib-Fenster
- [ ] Process() rendert alle Entities mit MaskRenderable
- [ ] Teardown() schließt Fenster sauber
- [ ] Entities werden als farbige Kreise gezeichnet

**Technische Details:**
```go
for _, e := range em.FilterByMask(MaskRenderable) {
    pos := e.Get(MaskPosition).(*Position)
    col := e.Get(MaskColor).(*Color)
    rl.DrawCircle(int32(pos.X), int32(pos.Y), size.Radius, ...)
}
```

**Definition of Done:**
- [ ] Ein statisches Test-Entity wird gerendert
- [ ] Farbe und Größe korrekt

---

# Epic E-002: Physik-Engine

**Beschreibung:** Implementierung der Physik-Simulation für realistische Partikel-Bewegung.

**Business Value:** Die Physik-Engine ist das Herzstück der visuellen Demonstration. Sie zeigt, wie ECS-Systems Daten transformieren.

**Akzeptanzkriterien:**
- Partikel bewegen sich realistisch
- Velocity und Acceleration werden korrekt verarbeitet
- Frame-unabhängige Bewegung (DeltaTime)
- Screen-Wrapping funktioniert

---

## Story E-002-S01: PhysicsSystem - Bewegung

**Als** Benutzer  
**möchte ich** dass Partikel sich bewegen  
**damit** die Simulation lebendig wirkt

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] PhysicsSystem verarbeitet alle Entities mit Position + Velocity
- [ ] Acceleration wird auf Velocity angewendet
- [ ] Velocity wird auf Position angewendet
- [ ] DeltaTime wird für frame-unabhängige Bewegung genutzt
- [ ] Velocity wird mit Damping-Faktor reduziert

**Technische Details:**
```go
vel.X += acc.X * dt
vel.Y += acc.Y * dt
pos.X += vel.X * dt
pos.Y += vel.Y * dt
vel.X *= damping
vel.Y *= damping
```

**Definition of Done:**
- [ ] Partikel bewegen sich
- [ ] Bewegung ist smooth bei verschiedenen FPS

---

## Story E-002-S02: Screen Boundaries

**Als** Benutzer  
**möchte ich** dass Partikel am Bildschirmrand wrappen  
**damit** keine Partikel verloren gehen

**Story Points:** 2

**Akzeptanzkriterien:**
- [ ] Partikel die links rausgehen, erscheinen rechts
- [ ] Partikel die oben rausgehen, erscheinen unten
- [ ] Wrapping funktioniert in beide Richtungen
- [ ] Keine visuellen Artefakte beim Wrapping

**Technische Details:**
```go
if pos.X < 0 { pos.X = width }
if pos.X > width { pos.X = 0 }
// ... analog für Y
```

**Definition of Done:**
- [ ] Partikel wrappen korrekt
- [ ] Kein "Teleport-Flicker"

---

## Story E-002-S03: EmitterSystem - Partikel Spawning

**Als** Benutzer  
**möchte ich** dass kontinuierlich neue Partikel entstehen  
**damit** die Simulation nicht leer wird

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] EmitterSystem spawnt Partikel basierend auf SpawnRate
- [ ] Neue Partikel haben zufällige Positionen
- [ ] Neue Partikel haben zufällige Velocities
- [ ] MaxParticles-Limit wird respektiert
- [ ] Spawn-Timer ist frame-unabhängig

**Technische Details:**
- SpawnRate: 100 Partikel/Sekunde (konfigurierbar)
- MaxParticles: 10.000 (konfigurierbar)
- Zufällige Start-Farbe aus Preset

**Definition of Done:**
- [ ] Partikel spawnen kontinuierlich
- [ ] Limit wird eingehalten
- [ ] Performance bleibt stabil

---

## Story E-002-S04: LifetimeSystem - Partikel Cleanup

**Als** Entwickler  
**möchte ich** dass alte Partikel automatisch entfernt werden  
**damit** der Speicher nicht überläuft

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] LifetimeSystem altert alle Entities mit Lifetime-Component
- [ ] Age wird um DeltaTime erhöht
- [ ] Wenn Age >= TTL, wird Entity markiert (Expired)
- [ ] Abgelaufene Entities werden aus EntityManager entfernt
- [ ] Keine Memory Leaks durch nicht-entfernte Entities

**Technische Details:**
```go
life.Age += dt
if life.Age >= life.TTL {
    life.Expired = true
    toRemove = append(toRemove, e.ID())
}
// Nach Loop:
for _, id := range toRemove {
    em.Remove(id)
}
```

**Definition of Done:**
- [ ] Partikel verschwinden nach TTL
- [ ] Entity-Count stabilisiert sich

---

# Epic E-003: Interaktivität

**Beschreibung:** Maus- und Tastatur-Interaktion für dynamische Benutzererfahrung.

**Business Value:** Interaktivität macht aus einer Demo ein Erlebnis. Der Benutzer fühlt sich als Teil der Simulation.

**Akzeptanzkriterien:**
- Maus erzeugt Gravitationsfelder
- Tastatur wechselt Presets
- Debug-Overlay ist togglebar

---

## Story E-003-S01: InputSystem - Maus Tracking

**Als** Benutzer  
**möchte ich** dass die Mausposition einen Attraktor erzeugt  
**damit** ich mit den Partikeln interagieren kann

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] InputSystem erstellt/aktualisiert Mouse-Attractor Entity
- [ ] Attractor-Position folgt der Maus
- [ ] Linke Maustaste: Positive Masse (Anziehung)
- [ ] Rechte Maustaste: Negative Masse (Abstoßung)
- [ ] Ohne Maustaste: Masse = 0 (inaktiv)

**Technische Details:**
- Attractor-Entity: Position + Mass + Attractor-Tag
- Mass-Wert: ±500 bei Klick, 0 sonst

**Definition of Done:**
- [ ] Maus wird getrackt
- [ ] Attraktor reagiert auf Klicks

---

## Story E-003-S02: GravitySystem - Attraktor-Physik

**Als** Benutzer  
**möchte ich** dass Partikel von der Maus angezogen/abgestoßen werden  
**damit** ich die Simulation steuern kann

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] GravitySystem findet alle Attraktoren
- [ ] Für jeden Partikel wird Gravitationskraft berechnet
- [ ] Kraft = mass / distance² (inverse square law)
- [ ] Kraftvektor wird auf Acceleration addiert
- [ ] Minimum-Distanz verhindert Singularitäten

**Technische Details:**
```go
dx := attractor.X - particle.X
dy := attractor.Y - particle.Y
dist := sqrt(dx*dx + dy*dy)
if dist < 10 { dist = 10 }  // Prevent singularity
force := mass / (dist * dist) * 100
acc.X += dx/dist * force
acc.Y += dy/dist * force
```

**Definition of Done:**
- [ ] Partikel werden angezogen bei LMB
- [ ] Partikel werden abgestoßen bei RMB
- [ ] Keine Division-by-Zero Crashes

---

## Story E-003-S03: Keyboard Input - Preset Switching

**Als** Benutzer  
**möchte ich** mit Tasten 1-5 zwischen Presets wechseln  
**damit** ich verschiedene Effekte sehen kann

**Story Points:** 2

**Akzeptanzkriterien:**
- [ ] Taste 1: Galaxy Preset
- [ ] Taste 2: Firework Preset
- [ ] Taste 3: Swarm Preset
- [ ] Taste 4: Fountain Preset
- [ ] Taste 5: Chaos Preset
- [ ] Preset-Wechsel entfernt alte Partikel und startet neu

**Technische Details:**
```go
if rl.IsKeyPressed(rl.KeyOne) { switchPreset(em, "galaxy") }
```

**Definition of Done:**
- [ ] Alle 5 Presets erreichbar
- [ ] Sauberer Übergang

---

## Story E-003-S04: Debug Overlay Toggle

**Als** Entwickler/Präsentierer  
**möchte ich** Debug-Informationen ein/ausblenden  
**damit** ich Performance analysieren kann

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] F3 toggled Debug-Overlay
- [ ] Overlay zeigt: FPS, Entity-Count
- [ ] Overlay zeigt: Aktives Preset
- [ ] Overlay zeigt: Maus-Position
- [ ] Overlay ist semi-transparent

**Technische Details:**
- Render am Ende von RenderSystem
- Nur wenn `showDebug = true`

**Definition of Done:**
- [ ] F3 toggled korrekt
- [ ] Alle Infos werden angezeigt

---

# Epic E-004: Visual Effects

**Beschreibung:** Visuelle Verbesserungen für "WOW"-Effekt.

**Business Value:** Der Showcase muss visuell beeindrucken, um als Demo zu überzeugen.

**Akzeptanzkriterien:**
- Farben interpolieren über Lifetime
- Partikel schrumpfen beim Verschwinden
- Optionale Trails für Bewegungsspuren

---

## Story E-004-S01: ColorSystem - Farbinterpolation

**Als** Benutzer  
**möchte ich** dass Partikel ihre Farbe über die Zeit ändern  
**damit** die Simulation visuell ansprechend ist

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] ColorSystem interpoliert von StartColor zu EndColor
- [ ] Interpolation basiert auf Lifetime.Progress (0-1)
- [ ] Alpha-Wert faded out am Ende (sanftes Verschwinden)
- [ ] Lineare Interpolation für alle RGBA-Kanäle

**Technische Details:**
```go
t := life.Progress()  // 0.0 bis 1.0
col.R = lerp(col.StartR, col.EndR, t)
// ... analog für G, B, A
```

**Definition of Done:**
- [ ] Partikel ändern Farbe
- [ ] Alpha fadet korrekt

---

## Story E-004-S02: Size Interpolation

**Als** Benutzer  
**möchte ich** dass Partikel beim Verschwinden schrumpfen  
**damit** das Verschwinden natürlich wirkt

**Story Points:** 2

**Akzeptanzkriterien:**
- [ ] Size.Radius interpoliert von StartSize zu EndSize
- [ ] Interpolation folgt Lifetime.Progress
- [ ] Partikel sind am Anfang groß, am Ende klein

**Technische Details:**
```go
size.Radius = lerp(size.StartSize, size.EndSize, t)
```

**Definition of Done:**
- [ ] Partikel schrumpfen
- [ ] Effekt ist subtil aber sichtbar

---

## Story E-004-S03: Trail-Rendering (Optional)

**Als** Benutzer  
**möchte ich** optionale Bewegungsspuren für Partikel  
**damit** die Bewegung noch dynamischer wirkt

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] Trail-Component speichert letzte N Positionen
- [ ] TrailSystem aktualisiert Positionen jedes Frame
- [ ] RenderSystem zeichnet Trail als verblassende Linie
- [ ] Trail-Länge ist konfigurierbar
- [ ] Trail kann per Taste F4 getoggled werden

**Technische Details:**
- Ring-Buffer für Positionen (letzte 10 Frames)
- Alpha verblasst entlang des Trails

**Definition of Done:**
- [ ] Trails werden gerendert
- [ ] Toggle funktioniert
- [ ] Performance bleibt akzeptabel

---

## Story E-004-S04: Glow/Bloom Effect (Optional)

**Als** Benutzer  
**möchte ich** einen subtilen Glow-Effekt  
**damit** die Partikel "leuchtend" wirken

**Story Points:** 6

**Akzeptanzkriterien:**
- [ ] Partikel haben einen weichen Glow um sich
- [ ] Glow-Intensität basiert auf Partikel-Alpha
- [ ] Glow-Farbe entspricht Partikel-Farbe
- [ ] Effekt ist togglebar mit F5

**Technische Details:**
- Additives Blending für Glow
- Mehrere Kreise mit abnehmender Opacity übereinander

**Definition of Done:**
- [ ] Glow-Effekt sichtbar
- [ ] Toggle funktioniert

---

# Epic E-005: Preset-System

**Beschreibung:** Vordefinierte Konfigurationen für verschiedene visuelle Stile.

**Business Value:** Presets zeigen die Vielseitigkeit des ECS-Patterns und machen die Demo abwechslungsreich.

**Akzeptanzkriterien:**
- Mindestens 5 Presets implementiert
- Jedes Preset hat eigene Charakteristik
- Presets sind über Tasten erreichbar

---

## Story E-005-S01: Preset Interface & Registry

**Als** Entwickler  
**möchte ich** ein einheitliches Preset-Interface  
**damit** neue Presets einfach hinzugefügt werden können

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] Preset-Interface mit Name(), Description(), Apply()
- [ ] Registry für alle verfügbaren Presets
- [ ] GetPreset(name) Funktion
- [ ] SwitchPreset() bereinigt alte Entities

**Technische Details:**
```go
type Preset interface {
    Name() string
    Description() string
    Apply(em ecs.EntityManager, cfg *config.Config)
}
```

**Definition of Done:**
- [ ] Interface definiert
- [ ] Registry funktioniert

---

## Story E-005-S02: Galaxy Preset

**Als** Benutzer  
**möchte ich** eine Spiral-Galaxie Simulation  
**damit** ich einen kosmischen Effekt erlebe

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] Zentraler massiver Attraktor
- [ ] Partikel starten in Spiralarmen
- [ ] Partikel haben tangentiale Anfangsgeschwindigkeit
- [ ] Blaue/weiße Farbpalette

**Definition of Done:**
- [ ] Spiralmuster sichtbar
- [ ] Partikel umkreisen Zentrum

---

## Story E-005-S03: Firework Preset

**Als** Benutzer  
**möchte ich** Feuerwerk-artige Explosionen  
**damit** ich einen festlichen Effekt erlebe

**Story Points:** 4

**Akzeptanzkriterien:**
- [ ] Partikel starten von unten, explodieren oben
- [ ] Explosion = viele Partikel in alle Richtungen
- [ ] Schwerkraft nach unten
- [ ] Bunte Farbpalette (Rot, Gold, Grün)
- [ ] Partikel haben kurze Lifetime

**Definition of Done:**
- [ ] Explosionen sichtbar
- [ ] Realistische Schwerkraft

---

## Story E-005-S04: Swarm Preset

**Als** Benutzer  
**möchte ich** Schwarm-Verhalten wie bei Vögeln oder Fischen  
**damit** ich emergentes Verhalten erlebe

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] Partikel folgen simplen Regeln (Separation, Alignment, Cohesion)
- [ ] Schwarm bewegt sich organisch
- [ ] Maus-Attraktor beeinflusst Schwarm
- [ ] Einheitliche Farbpalette

**Definition of Done:**
- [ ] Schwarmverhalten erkennbar
- [ ] Maus-Interaktion funktioniert

---

# Epic E-006: Audio-Reaktivität (Optional)

**Beschreibung:** Partikel reagieren auf Musik/Audio-Input.

**Business Value:** Audio-Reaktivität macht den Showcase zu einem einzigartigen audiovisuellen Erlebnis.

**Akzeptanzkriterien:**
- System-Audio oder Mikrofon-Input
- FFT-Analyse für Frequenzbänder
- Partikel reagieren auf Beat/Bass

---

## Story E-006-S01: Audio Input System

**Als** Entwickler  
**möchte ich** Audio-Daten vom System erfassen  
**damit** ich sie für Visualisierung nutzen kann

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] Audio-Input über System-Audio oder Mikrofon
- [ ] Sample-Rate: 44100 Hz
- [ ] Buffer-Size: 1024 Samples
- [ ] Graceful Degradation wenn kein Audio verfügbar
- [ ] Toggle mit F6

**Technische Details:**
- Library: `beep` oder `portaudio`
- Alternative: Vorgefertigte Audio-Datei

**Definition of Done:**
- [ ] Audio-Daten werden empfangen
- [ ] Kein Crash ohne Audio-Device

---

## Story E-006-S02: FFT-Analyse

**Als** Entwickler  
**möchte ich** Frequenzanalyse des Audio-Signals  
**damit** ich Bass, Mitten und Höhen unterscheiden kann

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] FFT auf Audio-Buffer anwenden
- [ ] Frequenzbänder extrahieren:
  - Bass: 20-200 Hz
  - Mitten: 200-2000 Hz
  - Höhen: 2000-20000 Hz
- [ ] Amplitude pro Band normalisiert (0-1)

**Technische Details:**
- Standard FFT-Algorithmus
- Hamming-Window für bessere Ergebnisse

**Definition of Done:**
- [ ] Frequenzbänder werden berechnet
- [ ] Werte sind normalisiert

---

## Story E-006-S03: AudioSystem - Visual Mapping

**Als** Benutzer  
**möchte ich** dass Partikel auf Musik reagieren  
**damit** ein audiovisuelles Erlebnis entsteht

**Story Points:** 4

**Akzeptanzkriterien:**
- [ ] Bass → Partikel-Größe (größer bei mehr Bass)
- [ ] Bass → Spawn-Rate (mehr Partikel bei Beat)
- [ ] Höhen → Partikel-Geschwindigkeit
- [ ] Mitten → Farbintensität
- [ ] Reaktion ist subtil aber sichtbar

**Technische Details:**
```go
if bass > threshold {
    size.Radius *= 1 + bass * 0.5
    emitter.SpawnRate *= 1 + bass
}
```

**Definition of Done:**
- [ ] Partikel "tanzen" zur Musik
- [ ] Effekt ist angenehm, nicht übertrieben

---

# Sprint-Planung

## Sprint 1: Foundation (Woche 1)

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-001-S01 | Foundation | 2 | 🔴 |
| E-001-S02 | Foundation | 3 | 🔴 |
| E-001-S03 | Foundation | 5 | 🔴 |
| E-001-S04 | Foundation | 5 | 🔴 |
| E-001-S05 | Foundation | 3 | 🔴 |
| **Total** | | **18** | |

**Sprint Goal:** ECS Engine läuft, Fenster öffnet sich, ein Test-Partikel wird gerendert.

---

## Sprint 2: Core Gameplay (Woche 2)

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-002-S01 | Physik | 3 | 🔴 |
| E-002-S02 | Physik | 2 | 🔴 |
| E-002-S03 | Physik | 5 | 🔴 |
| E-002-S04 | Physik | 5 | 🔴 |
| E-003-S01 | Interaktivität | 3 | 🔴 |
| E-003-S02 | Interaktivität | 5 | 🔴 |
| **Total** | | **23** | |

**Sprint Goal:** Partikel spawnen, bewegen sich, reagieren auf Maus.

---

## Sprint 3: Polish & Extras (Woche 3)

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-003-S03 | Interaktivität | 2 | 🟡 |
| E-003-S04 | Interaktivität | 3 | 🟡 |
| E-004-S01 | Visual | 3 | 🟡 |
| E-004-S02 | Visual | 2 | 🟡 |
| E-005-S01 | Presets | 3 | 🟡 |
| E-005-S02 | Presets | 3 | 🟡 |
| E-005-S03 | Presets | 4 | 🟡 |
| E-005-S04 | Presets | 3 | 🟡 |
| **Total** | | **23** | |

**Sprint Goal:** Alle Presets, visuelle Effekte, Debug-Overlay.

---

## Backlog: Nice-to-Have

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-004-S03 | Visual (Trail) | 5 | 🟢 |
| E-004-S04 | Visual (Glow) | 6 | 🟢 |
| E-006-S01 | Audio | 5 | 🟢 |
| E-006-S02 | Audio | 5 | 🟢 |
| E-006-S03 | Audio | 4 | 🟢 |
| **Total** | | **25** | |

**Backlog Goal:** Optionale Features für erweiterten Showcase.

---

# Dependency Graph

```
┌─────────────────────────────────────────────────────────────────┐
│                     STORY DEPENDENCIES                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  E-001-S01 (Setup)                                               │
│      │                                                           │
│      ▼                                                           │
│  E-001-S02 (Masks) ──────────────────┐                          │
│      │                                │                          │
│      ▼                                ▼                          │
│  E-001-S03 (Components)          E-005-S01 (Preset Interface)   │
│      │                                │                          │
│      ▼                                ▼                          │
│  E-001-S04 (Engine) ◀────────── E-005-S02/S03/S04 (Presets)     │
│      │                                                           │
│      ▼                                                           │
│  E-001-S05 (RenderSystem)                                       │
│      │                                                           │
│      ├──────────────────────┬──────────────────┐                │
│      ▼                      ▼                  ▼                │
│  E-002-S03 (Emitter)   E-002-S01 (Physics)  E-003-S01 (Input)   │
│      │                      │                  │                │
│      ▼                      ▼                  ▼                │
│  E-002-S04 (Lifetime)  E-002-S02 (Bounds)  E-003-S02 (Gravity)  │
│      │                                         │                │
│      ▼                                         ▼                │
│  E-004-S01 (Color)                        E-003-S03 (Keys)      │
│      │                                         │                │
│      ▼                                         ▼                │
│  E-004-S02 (Size)                         E-003-S04 (Debug)     │
│      │                                                           │
│      ├──────────────────────┐                                   │
│      ▼                      ▼                                   │
│  E-004-S03 (Trail)     E-004-S04 (Glow)                         │
│                                                                  │
│  [OPTIONAL PATH]                                                 │
│  E-006-S01 (Audio In) ──▶ E-006-S02 (FFT) ──▶ E-006-S03 (Viz)   │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

# Definition of Ready (DoR)

Eine Story ist "Ready" für Entwicklung wenn:

- [ ] User Story ist formuliert (Als...möchte ich...damit)
- [ ] Akzeptanzkriterien sind vollständig
- [ ] Story Points sind geschätzt
- [ ] Dependencies sind erfüllt
- [ ] Technische Details sind klar
- [ ] Keine offenen Fragen

---

# Definition of Done (DoD)

Eine Story ist "Done" wenn:

- [ ] Alle Akzeptanzkriterien erfüllt
- [ ] Code compiles ohne Warnings
- [ ] Unit Tests geschrieben und grün
- [ ] Code Review bestanden
- [ ] In `main` Branch gemerged
- [ ] Demo-fähig

---

**🚀 YOLO MODE COMPLETE - 24 STORIES READY FOR IMPLEMENTATION!**
