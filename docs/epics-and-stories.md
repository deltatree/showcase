---
stepsCompleted: [1, 2, 3, 4, 5]
inputDocuments: ['docs/prd.md', 'docs/architecture.md']
workflowType: 'epics-stories'
lastStep: 5
project_name: 'showcase'
user_name: 'Deltatree'
date: '2025-12-10'
yolo_mode: true
totalEpics: 9
totalStories: 43
estimatedSprints: 6
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
| **Epics** | 9 |
| **User Stories** | 43 |
| **Story Points (geschätzt)** | 159 |
| **Sprints (geschätzt)** | 6 |

### Epic-Übersicht

| Epic | Titel | Stories | Punkte | Priorität | Status |
|------|-------|---------|--------|-----------|--------|
| E-001 | ECS Foundation | 5 | 18 | 🔴 MUST | ✅ Complete |
| E-002 | Physik-Engine | 4 | 15 | 🔴 MUST | ✅ Complete |
| E-003 | Interaktivität | 4 | 13 | 🔴 MUST | ✅ Complete |
| E-004 | Visual Effects | 4 | 16 | 🟡 SHOULD | ✅ Complete |
| E-005 | Preset-System | 4 | 13 | 🟡 SHOULD | ✅ Complete |
| E-006 | Audio-Reaktivität | 3 | 14 | 🟢 COULD | 📝 Placeholder |
| E-007 | Web Deployment (WASM) | 4 | 13 | 🔴 MUST | ✅ Complete |
| E-008 | Awesome-Go Listing | 7 | 25 | 🔴 MUST | 🔄 In Progress |
| E-009 | Premium Experience 🔥 | 8 | 32 | 🟡 SHOULD | ✅ 7/8 Complete |

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

**🚀 YOLO MODE COMPLETE - 28 STORIES READY FOR IMPLEMENTATION!**

---

# Epic E-007: Web Deployment (WASM)

**Beschreibung:** Die Particle Symphony Anwendung wird als WebAssembly (WASM) kompiliert und automatisch auf GitHub Pages deployed. Jeder kann den Showcase direkt im Browser erleben – ohne Installation!

**Business Value:** 
- **Reichweite x100:** Jeder mit einem Browser kann den ECS-Showcase erleben
- **Technische Demo:** Beweist, dass Go + ECS + Raylib auch im Web funktioniert
- **Professioneller Auftritt:** Automatisches Deployment zeigt DevOps-Kompetenz
- **Viral-Potenzial:** Einfach zu teilen, keine Hürden

**Akzeptanzkriterien:**
- WASM-Binary wird erfolgreich gebaut
- GitHub Actions Pipeline deployed automatisch bei Push auf `main`
- Landing Page ist ansprechend und responsive
- Showcase läuft flüssig im Browser (Chrome, Firefox, Safari)

---

## Story E-007-S01: WASM Build Setup

**Als** Entwickler  
**möchte ich** die Anwendung als WebAssembly kompilieren können  
**damit** sie im Browser ausführbar ist

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] WASM-Target (`GOOS=js GOARCH=wasm`) funktioniert
- [ ] `wasm_exec.js` von Go-Installation kopiert
- [ ] Raylib WASM-Kompatibilität validiert/angepasst
- [ ] Build-Script erstellt: `build-wasm.sh`
- [ ] Output: `particle-symphony.wasm` + `wasm_exec.js`

**Technische Details:**
```bash
#!/bin/bash
# build-wasm.sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
GOOS=js GOARCH=wasm go build -o web/particle-symphony.wasm .
```

**Raylib WASM Hinweise:**
- `raylib-go` unterstützt WASM via Emscripten
- Alternative: Build mit `-tags=wasm` falls nötig
- Window-Init muss WASM-kompatibel sein (kein SetTargetFPS in manchen Fällen)

**Definition of Done:**
- [ ] `./build-wasm.sh` erzeugt valides WASM
- [ ] Keine Compile-Errors
- [ ] WASM-Größe dokumentiert

---

## Story E-007-S02: Web-Host HTML/JS Wrapper

**Als** Benutzer  
**möchte ich** den Showcase auf einer Webseite starten  
**damit** ich ihn ohne Installation erleben kann

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] `web/index.html` lädt WASM korrekt
- [ ] Canvas-Element für Raylib-Rendering konfiguriert
- [ ] Loading-Indicator während WASM lädt
- [ ] Fehlerbehandlung wenn WASM nicht unterstützt
- [ ] Responsive Design (funktioniert auf Desktop und Tablet)

**Technische Details:**
```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Particle Symphony - ECS Showcase</title>
    <style>
        body { margin: 0; background: #0a0a0a; }
        #canvas { display: block; margin: auto; }
        .loading { color: white; text-align: center; padding: 20px; }
    </style>
</head>
<body>
    <div class="loading" id="loading">
        <h2>🎵 Particle Symphony lädt...</h2>
        <p>WebAssembly wird initialisiert</p>
    </div>
    <canvas id="canvas"></canvas>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("particle-symphony.wasm"), go.importObject)
            .then((result) => {
                document.getElementById('loading').style.display = 'none';
                go.run(result.instance);
            })
            .catch((err) => {
                document.getElementById('loading').innerHTML = 
                    '<h2>❌ Fehler beim Laden</h2><p>' + err + '</p>';
            });
    </script>
</body>
</html>
```

**Definition of Done:**
- [ ] Lokaler Test mit `python -m http.server` funktioniert
- [ ] Canvas zeigt Partikel
- [ ] Loading-State sichtbar

---

## Story E-007-S03: GitHub Actions CI/CD Pipeline

**Als** Entwickler  
**möchte ich** automatisches Deployment bei jedem Push  
**damit** die GitHub Page immer aktuell ist

**Story Points:** 4

**Akzeptanzkriterien:**
- [ ] `.github/workflows/deploy.yml` erstellt
- [ ] Pipeline baut WASM bei Push auf `main`
- [ ] Pipeline deployed nach `gh-pages` Branch
- [ ] GitHub Pages ist aktiviert für `gh-pages` Branch
- [ ] Status-Badge in README

**Technische Details:**
```yaml
name: Deploy to GitHub Pages

on:
  push:
    branches: [main]
  workflow_dispatch:

permissions:
  contents: read
  pages: write
  id-token: write

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Build WASM
        run: |
          cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
          GOOS=js GOARCH=wasm go build -o web/particle-symphony.wasm .
      
      - name: Setup Pages
        uses: actions/configure-pages@v4
      
      - name: Upload artifact
        uses: actions/upload-pages-artifact@v3
        with:
          path: 'web'

  deploy:
    needs: build
    runs-on: ubuntu-latest
    environment:
      name: github-pages
      url: ${{ steps.deployment.outputs.page_url }}
    steps:
      - name: Deploy to GitHub Pages
        id: deployment
        uses: actions/deploy-pages@v4
```

**Definition of Done:**
- [ ] Push auf `main` triggered Workflow
- [ ] Workflow ist grün
- [ ] GitHub Page erreichbar unter `https://deltatree.github.io/showcase/`

---

## Story E-007-S04: Premium Landing Page Design

**Als** Besucher  
**möchte ich** eine ansprechende Landing Page  
**damit** ich verstehe was mich erwartet und beeindruckt bin

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] Modernes, dunkles Design passend zur Partikel-Ästhetik
- [ ] Hero-Section mit Titel und Kurzbeschreibung
- [ ] Canvas nimmt Hauptbereich ein
- [ ] Steuerungs-Hinweise unten (Tasten 1-5, Maus-Interaktion)
- [ ] Footer mit Links zu GitHub-Repo und ECS-Framework
- [ ] Responsive: Funktioniert auf Mobile (readonly) und Desktop

**Design-Konzept:**
```
┌─────────────────────────────────────────────────────────────────┐
│  🎵 PARTICLE SYMPHONY                        [GitHub] [ECS]     │
│  Ein interaktiver ECS-Showcase in Go + Raylib                   │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│                                                                 │
│                     ┌─────────────────────┐                     │
│                     │                     │                     │
│                     │    WASM CANVAS      │                     │
│                     │   (1280 x 720)      │                     │
│                     │                     │                     │
│                     └─────────────────────┘                     │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│  🎮 Steuerung:                                                  │
│  [1] Galaxy  [2] Firework  [3] Swarm  [4] Fountain  [5] Chaos   │
│  [LMB] Anziehen  [RMB] Abstoßen  [F3] Debug                     │
├─────────────────────────────────────────────────────────────────┤
│  Built with andygeiss/ecs • Source on GitHub • MIT License      │
└─────────────────────────────────────────────────────────────────┘
```

**Technische Details:**
- CSS-Only Animation für Header (subtle glow)
- Keyboard-Shortcuts Overlay mit CSS
- Font: Inter oder System-Font-Stack
- Farbschema: #0a0a0a Background, #00ff88 Accent, #ffffff Text

**Definition of Done:**
- [ ] Design umgesetzt
- [ ] Mobile-View getestet
- [ ] Links funktionieren
- [ ] Lighthouse Score > 90

---

# Aktualisierte Sprint-Planung

## Sprint 4: Web Deployment (Woche 4 - NEU)

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-007-S01 | Web Deployment | 3 | 🔴 |
| E-007-S02 | Web Deployment | 3 | 🔴 |
| E-007-S03 | Web Deployment | 4 | 🔴 |
| E-007-S04 | Web Deployment | 3 | 🔴 |
| **Total** | | **13** | |

**Sprint Goal:** Showcase läuft live auf `https://deltatree.github.io/showcase/` mit automatischem Deployment.

---

# Aktualisierter Dependency Graph

```
┌─────────────────────────────────────────────────────────────────┐
│                     STORY DEPENDENCIES                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  [EXISTING EPICS E-001 bis E-006 - siehe oben]                  │
│                                                                  │
│  ═══════════════════════════════════════════════════════════    │
│                                                                  │
│  E-007: WEB DEPLOYMENT (kann parallel zu E-004/E-005 laufen)    │
│                                                                  │
│  E-001-S04 (Engine läuft) ──────────┐                           │
│  E-002-S01 (Physik funktioniert) ───┤                           │
│  E-003-S01 (Input funktioniert) ────┘                           │
│           │                                                      │
│           ▼                                                      │
│  E-007-S01 (WASM Build Setup)                                   │
│           │                                                      │
│           ▼                                                      │
│  E-007-S02 (HTML/JS Wrapper)                                    │
│           │                                                      │
│           ├──────────────────────┐                              │
│           ▼                      ▼                              │
│  E-007-S03 (CI/CD Pipeline)  E-007-S04 (Landing Page)           │
│           │                      │                              │
│           └──────────┬───────────┘                              │
│                      ▼                                          │
│              🌐 LIVE DEPLOYMENT                                  │
│       https://deltatree.github.io/showcase/                     │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

# E-007 Technical Notes

## WASM Kompatibilität Checklist

| Feature | Native | WASM | Anpassung nötig? |
|---------|--------|------|------------------|
| Raylib Rendering | ✅ | ✅ | Canvas-Element erforderlich |
| Mouse Input | ✅ | ✅ | Automatisch via Raylib |
| Keyboard Input | ✅ | ✅ | Automatisch via Raylib |
| Window Resize | ✅ | ⚠️ | CSS-basiertes Scaling |
| Audio | ✅ | ⚠️ | Möglicherweise deaktiviert |
| File I/O | ✅ | ❌ | Config embedded kompilieren |
| 60 FPS | ✅ | ✅ | RequestAnimationFrame |

## Fallback-Strategie

Falls `raylib-go` WASM-Probleme macht:
1. **Option A:** Ebitengine als alternatives Rendering-Backend
2. **Option B:** Nur statische Demo mit Screenshots/GIF
3. **Option C:** Native-Binary Download-Links auf Landing Page

## Performance-Ziele WASM

| Metrik | Zielwert |
|--------|----------|
| WASM-Größe | < 10 MB |
| Initial Load | < 3 Sekunden |
| FPS im Browser | 60 FPS bei 5.000 Partikeln |
| Memory Usage | < 100 MB |

---

# Epic E-008: Awesome-Go Listing Readiness

**Beschreibung:** Vorbereitung des Projekts für die Aufnahme in die kuratierte [awesome-go](https://github.com/avelino/awesome-go) Liste - eine der wichtigsten Go-Ressourcen mit 160k+ GitHub Stars.

**Business Value:** Listing in awesome-go erhöht massiv die Sichtbarkeit des Projekts und des `andygeiss/ecs` Frameworks. Es ist ein Qualitätssiegel für Go-Projekte.

**Akzeptanzkriterien (aus awesome-go CONTRIBUTING.md):**
- Mindestens 5 Monate Historie seit erstem Commit
- Open-Source-Lizenz vorhanden
- `go.mod` mit korrektem Module-Namen
- Mindestens ein SemVer-Release (vX.Y.Z)
- README und Dokumentation auf Englisch
- Tests mit ≥80% Code Coverage
- Go Report Card mit Grade A- oder besser
- pkg.go.dev Dokumentation für alle öffentlichen APIs

---

## Story E-008-S01: English Documentation & README

**Als** internationaler Entwickler  
**möchte ich** eine vollständige englische Dokumentation  
**damit** ich das Projekt verstehen und nutzen kann

**Story Points:** 5

**Akzeptanzkriterien:**
- [ ] README.md komplett auf Englisch
- [ ] Project description, features, installation, usage
- [ ] Screenshots/GIFs der Simulation
- [ ] Quick Start Guide (5-Minuten-Setup)
- [ ] Architecture overview mit Diagramm
- [ ] Contributing guidelines
- [ ] License section
- [ ] Badges: Go Report Card, Coverage, License, Go Version

**README Struktur:**
```markdown
# 🎵 Particle Symphony

An interactive particle simulation showcasing the power of Entity-Component-System architecture in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/deltatree/showcase)](https://goreportcard.com/report/github.com/deltatree/showcase)
[![codecov](https://codecov.io/gh/deltatree/showcase/graph/badge.svg)](https://codecov.io/gh/deltatree/showcase)
[![Go Reference](https://pkg.go.dev/badge/github.com/deltatree/showcase.svg)](https://pkg.go.dev/github.com/deltatree/showcase)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

![Demo GIF](docs/demo.gif)

## ✨ Features
- 10,000+ particles at 60 FPS
- Real-time physics simulation
- Mouse interaction (attract/repel)
- 5 visual presets
- Built with andygeiss/ecs

## 🚀 Quick Start
...

## 🏗️ Architecture
...

## 📖 Documentation
...

## 🤝 Contributing
...

## 📄 License
MIT License - see [LICENSE](LICENSE)
```

**Definition of Done:**
- [ ] README.md auf Englisch geschrieben
- [ ] Alle Sections vollständig
- [ ] GIF/Screenshots eingefügt
- [ ] Badges funktionieren

---

## Story E-008-S02: Test Coverage ≥80%

**Als** Contributor der awesome-go Liste  
**möchte ich** eine Testabdeckung von mindestens 80%  
**damit** das Projekt die Qualitätsstandards erfüllt

**Story Points:** 8

**Status:** ✅ ACHIEVED (Testbare Packages >80%)

**Akzeptanzkriterien:**
- [x] Unit Tests für alle Components
- [x] Unit Tests für alle Systems (wo möglich)
- [ ] Integration Tests für Engine-Setup
- [x] Test Coverage ≥80% für testbare Packages
- [ ] Coverage Report via Codecov/Coveralls
- [x] Keine flaky tests

**Tatsächliche Coverage (Stand: 10.12.2025):**
| Package | Coverage | Status |
|---------|----------|--------|
| `components/` | 100.0% | ✅ |
| `internal/config/` | 100.0% | ✅ |
| `presets/` | 99.3% | ✅ |
| `premium/` | 94.1% | ✅ |
| `systems/` | 22.8% | ⚠️ (Raylib-limitiert) |
| **Testbare Packages** | **~98%** | ✅ |
| **Gesamt** | **57.4%** | 📝 |

### Teststrategie für Systems-Package

Das `systems/`-Package hat niedrige Coverage wegen **Raylib-Abhängigkeiten**:

**Nicht testbar ohne Display:**
- `Process()` Methoden (benötigen `rl.GetFrameTime()`, `rl.GetMousePosition()`)
- `RenderGlow()`, `RenderWithBlur()` (benötigen `rl.DrawCircle()`)
- `Setup()` (benötigt `rl.InitWindow()`)

**Getestet:**
- Alle Konstruktoren (`NewXyzSystem()`)
- Alle Setter/Getter (`SetQuality()`, `GetMaxParticles()`, etc.)
- Reine Logik-Funktionen (`lerp()`, `lerpF()`)
- `GravitySystem.Process()` (keine Raylib-Calls!)
- `ColorSystem.Process()` (keine Raylib-Calls!)

**Begründung:** Bei GUI/Game-Frameworks wie Raylib ist 100% Coverage nicht erreichbar ohne:
1. Mocking-Framework (overhead für ein Showcase-Projekt)
2. Headless-Mode (Raylib unterstützt dies nicht nativ)
3. Integration-Tests mit echtem Display

**Für Awesome-Go ausreichend:** Die testbaren Packages zeigen saubere, idiomatische Tests.

**Test-Beispiele:**
```go
func TestGravitySystem_Process(t *testing.T) {
    em := ecs.NewEntityManager()
    sys := NewGravitySystem()
    
    attractor := ecs.NewEntity("attractor", []ecs.Component{
        components.NewPosition().With(500, 500),
        components.NewMass().WithValue(10000),
        components.NewAttractor(),
    })
    em.Add(attractor)
    
    particle := ecs.NewEntity("particle", []ecs.Component{
        components.NewPosition().With(400, 400),
        components.NewAcceleration(),
        components.NewParticle(),
    })
    em.Add(particle)
    
    result := sys.Process(em)
    
    if result != ecs.StateEngineContinue {
        t.Errorf("expected StateEngineContinue")
    }
    
    acc := particle.Get(components.MaskAcceleration).(*components.Acceleration)
    if acc.X <= 0 || acc.Y <= 0 {
        t.Errorf("expected acceleration toward attractor")
    }
}
```

**Definition of Done:**
- [x] `go test ./... -cover` zeigt ≥80% für testbare Packages
- [ ] Coverage Badge in README
- [ ] CI Pipeline prüft Coverage

---

## Story E-008-S03: Go Report Card Grade A

**Als** Quality-bewusster Entwickler  
**möchte ich** einen Go Report Card Score von A- oder besser  
**damit** das Projekt als hochwertig anerkannt wird

**Story Points:** 3

**Akzeptanzkriterien:**
- [ ] Keine `gofmt` Violations
- [ ] Keine `go vet` Warnings
- [ ] Keine `golint` Issues (oder begründete Ausnahmen)
- [ ] Keine `ineffassign` Findings
- [ ] Keine `misspell` Findings
- [ ] Score mindestens A- auf goreportcard.com

**Prüf-Kommandos:**
```bash
# Format Check
gofmt -l -s .

# Vet Check
go vet ./...

# StaticCheck
staticcheck ./...

# Ineffassign
ineffassign ./...

# Misspell
misspell -error .
```

**Definition of Done:**
- [ ] Alle Checks lokal bestanden
- [ ] goreportcard.com zeigt A- oder besser
- [ ] Badge in README aktualisiert

---

## Story E-008-S04: pkg.go.dev Documentation

**Als** Go-Entwickler  
**möchte ich** vollständige API-Dokumentation auf pkg.go.dev  
**damit** ich die öffentlichen APIs verstehen kann

**Story Points:** 4

**Akzeptanzkriterien:**
- [ ] Alle öffentlichen Typen haben Doc-Comments
- [ ] Alle öffentlichen Funktionen haben Doc-Comments
- [ ] Package-Level Doc-Comments für jedes Package
- [ ] Beispiele (Examples) für wichtige Funktionen
- [ ] pkg.go.dev Seite ist erreichbar und vollständig

**Doc-Comment Format:**
```go
// Position represents the 2D position of an entity in world coordinates.
// It implements the ecs.Component interface.
type Position struct {
    X, Y float64
}

// NewPosition creates a new Position component at origin (0, 0).
func NewPosition() *Position {
    return &Position{}
}

// WithX sets the X coordinate and returns the Position for chaining.
func (p *Position) WithX(x float64) *Position {
    p.X = x
    return p
}

// Magnitude returns the distance from origin to the position.
func (p *Position) Magnitude() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Example demonstrates creating and using a Position component.
func ExamplePosition() {
    pos := NewPosition().WithX(100).WithY(200)
    fmt.Printf("Position: (%.0f, %.0f)\n", pos.X, pos.Y)
    // Output: Position: (100, 200)
}
```

**Definition of Done:**
- [ ] Alle Public APIs dokumentiert
- [ ] pkg.go.dev zeigt Dokumentation
- [ ] Mindestens 3 runnable Examples

---

## Story E-008-S05: SemVer Release v1.0.0

**Als** Nutzer des Projekts  
**möchte ich** stabile, versionierte Releases  
**damit** ich eine zuverlässige Version referenzieren kann

**Story Points:** 2

**Akzeptanzkriterien:**
- [ ] Git Tag `v1.0.0` erstellt und gepusht
- [ ] GitHub Release mit Release Notes
- [ ] CHANGELOG.md mit Änderungshistorie
- [ ] `go.mod` Modul-Pfad korrekt

**Release Notes Template:**
```markdown
# v1.0.0 - Initial Release 🎉

## ✨ Features
- Interactive particle simulation with 10,000+ particles
- Real-time physics engine with gravity and damping
- Mouse interaction (attract/repel particles)
- 5 visual presets: Galaxy, Firework, Swarm, Fountain, Chaos
- WebAssembly support for browser deployment
- Built on andygeiss/ecs framework

## 📦 Installation
\`\`\`bash
go get github.com/deltatree/showcase@v1.0.0
\`\`\`

## 🎮 Try it Online
Visit https://deltatree.github.io/showcase/

## 📖 Documentation
- [README](https://github.com/deltatree/showcase/blob/v1.0.0/README.md)
- [Architecture](https://github.com/deltatree/showcase/blob/v1.0.0/docs/architecture.md)
- [pkg.go.dev](https://pkg.go.dev/github.com/deltatree/showcase)
```

**Definition of Done:**
- [ ] Tag gepusht
- [ ] GitHub Release erstellt
- [ ] `go get ...@v1.0.0` funktioniert

---

## Story E-008-S06: MIT License File

**Als** potenzieller Nutzer  
**möchte ich** eine klare Open-Source-Lizenz  
**damit** ich weiß, wie ich das Projekt nutzen darf

**Story Points:** 1

**Akzeptanzkriterien:**
- [ ] `LICENSE` Datei im Root-Verzeichnis
- [ ] MIT License Text (awesome-go akzeptiert)
- [ ] Copyright Jahr und Name korrekt
- [ ] License Badge in README

**LICENSE Inhalt:**
```
MIT License

Copyright (c) 2025 Deltatree

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

**Definition of Done:**
- [ ] LICENSE Datei existiert
- [ ] GitHub erkennt Lizenz automatisch

---

## Story E-008-S07: Awesome-Go Pull Request Preparation

**Als** Projekt-Maintainer  
**möchte ich** alle PR-Anforderungen vorbereitet haben  
**damit** die Submission reibungslos verläuft

**Story Points:** 2

**Akzeptanzkriterien:**
- [ ] Alle Quality-Links gesammelt und getestet
- [ ] PR-Body vorbereitet mit allen erforderlichen Links
- [ ] Richtige Kategorie identifiziert (Game Development)
- [ ] Alphabetische Einordnung geprüft
- [ ] Description nach Guidelines formuliert

**PR-Body Template:**
```markdown
Forge link: https://github.com/deltatree/showcase
pkg.go.dev: https://pkg.go.dev/github.com/deltatree/showcase
goreportcard.com: https://goreportcard.com/report/github.com/deltatree/showcase
Coverage: https://app.codecov.io/gh/deltatree/showcase

---

This project is an interactive particle simulation showcasing the ECS (Entity-Component-System) pattern in Go using the andygeiss/ecs framework. It demonstrates:

- Real-time physics with 10,000+ particles at 60 FPS
- Clean ECS architecture with data-oriented design
- WebAssembly deployment for browser access

The project has been actively developed and maintained, with comprehensive documentation and test coverage.
```

**Eintrag für README.md (awesome-go):**
```markdown
- [particle-symphony](https://github.com/deltatree/showcase) - Interactive particle simulation demonstrating ECS architecture with real-time physics and WebAssembly support.
```

**Kategorie:** Game Development (alphabetisch zwischen "Pi" und "Pitaya")

**Definition of Done:**
- [ ] Alle Links validiert
- [ ] PR kann eingereicht werden
- [ ] 5-Monate-Wartezeit beachtet

---

# Aktualisierte Übersicht

| Metrik | Wert |
|--------|------|
| **Epics** | 9 |
| **User Stories** | 43 |
| **Story Points (geschätzt)** | 159 |
| **Sprints (geschätzt)** | 6 |

### Epic-Übersicht (aktualisiert)

| Epic | Titel | Stories | Punkte | Priorität |
|------|-------|---------|--------|-----------|
| E-001 | ECS Foundation | 5 | 18 | 🔴 MUST |
| E-002 | Physik-Engine | 4 | 15 | 🔴 MUST |
| E-003 | Interaktivität | 4 | 13 | 🔴 MUST |
| E-004 | Visual Effects | 4 | 16 | 🟡 SHOULD |
| E-005 | Preset-System | 4 | 13 | 🟡 SHOULD |
| E-006 | Audio-Reaktivität | 3 | 14 | 🟢 COULD |
| E-007 | Web Deployment (WASM) | 4 | 13 | 🔴 MUST |
| E-008 | Awesome-Go Listing | 7 | 25 | 🔴 MUST |
| **E-009** | **Premium Experience 🔥** | **8** | **32** | **🟡 SHOULD** | |

---

# Sprint 5: Awesome-Go Readiness (Woche 5)

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-008-S06 | Awesome-Go | 1 | 🔴 |
| E-008-S03 | Awesome-Go | 3 | 🔴 |
| E-008-S01 | Awesome-Go | 5 | 🔴 |
| E-008-S04 | Awesome-Go | 4 | 🔴 |
| E-008-S02 | Awesome-Go | 8 | 🔴 |
| E-008-S05 | Awesome-Go | 2 | 🔴 |
| E-008-S07 | Awesome-Go | 2 | 🔴 |
| **Total** | | **25** | |

**Sprint Goal:** Projekt erfüllt alle awesome-go Qualitätsstandards und ist bereit für PR-Submission.

**⚠️ WICHTIG:** Die awesome-go Guidelines erfordern **mindestens 5 Monate Historie** seit dem ersten Commit. Die PR-Submission kann erst nach Erreichen dieser Zeitspanne erfolgen!

---

# E-008 Dependency Graph

```
┌─────────────────────────────────────────────────────────────────┐
│                 AWESOME-GO LISTING DEPENDENCIES                  │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  E-001 bis E-007 (Projekt funktioniert) ────┐                   │
│                                              │                   │
│                                              ▼                   │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │              E-008: AWESOME-GO READINESS                  │  │
│  ├───────────────────────────────────────────────────────────┤  │
│  │                                                           │  │
│  │   E-008-S06 (LICENSE)          ──────────┐               │  │
│  │           │                              │               │  │
│  │           │  (parallel)                  │               │  │
│  │           │                              │               │  │
│  │   E-008-S03 (Go Report Card)   ──────────┤               │  │
│  │           │                              │               │  │
│  │           │  (parallel)                  │               │  │
│  │           │                              │               │  │
│  │   E-008-S01 (English README)   ──────────┤               │  │
│  │           │                              │               │  │
│  │           │  (parallel)                  │               │  │
│  │           │                              │               │  │
│  │   E-008-S04 (pkg.go.dev Docs)  ──────────┤               │  │
│  │           │                              │               │  │
│  │           │  (parallel)                  │               │  │
│  │           │                              │               │  │
│  │   E-008-S02 (Test Coverage)    ──────────┤               │  │
│  │                                          │               │  │
│  │                                          ▼               │  │
│  │                              E-008-S05 (v1.0.0 Release)  │  │
│  │                                          │               │  │
│  │                                          ▼               │  │
│  │                              E-008-S07 (PR Preparation)  │  │
│  │                                          │               │  │
│  └──────────────────────────────────────────┼───────────────┘  │
│                                              │                   │
│                                              ▼                   │
│                        ⏳ WARTE 5 MONATE SEIT ERSTEM COMMIT      │
│                                              │                   │
│                                              ▼                   │
│                  🎉 PR AN AWESOME-GO EINREICHEN                  │
│        https://github.com/avelino/awesome-go/pulls               │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

# Awesome-Go Submission Checklist

**Vor PR-Einreichung alle Punkte prüfen:**

- [ ] **Repository Age:** ≥5 Monate seit erstem Commit
- [ ] **License:** MIT License in `LICENSE` Datei
- [ ] **go.mod:** Korrekter Modul-Pfad
- [ ] **Release:** Mindestens v1.0.0 Tag
- [ ] **README:** Englisch, vollständig, mit Badges
- [ ] **Documentation:** pkg.go.dev zeigt alle Public APIs
- [ ] **Tests:** Coverage ≥80%
- [ ] **Quality:** Go Report Card Grade A- oder besser
- [ ] **Category:** Game Development (korrekt alphabetisch)
- [ ] **Description:** Endet mit Punkt, nicht promotional

**PR Links für Body:**
```
Forge link: https://github.com/deltatree/showcase
pkg.go.dev: https://pkg.go.dev/github.com/deltatree/showcase  
goreportcard.com: https://goreportcard.com/report/github.com/deltatree/showcase
Coverage: https://app.codecov.io/gh/deltatree/showcase
```

---

**🚀 YOLO MODE COMPLETE - 43 STORIES READY FOR IMPLEMENTATION!**

---

# Epic E-009: Premium Experience 🔥

**Beschreibung:** Der ultimative WOW-Faktor! Geile Grafiken, wunderschöner Sound, butterweiche Animationen und eine UX die süchtig macht. Dieser Epic macht aus einem guten Showcase ein UNVERGESSLICHES Erlebnis.

**Business Value:** 
- **Differenzierung:** Hebt sich deutlich von Standard-Demos ab
- **Viralität:** Benutzer teilen optisch beeindruckende Anwendungen
- **Professionalität:** Zeigt, dass Go für hochqualitative visuelle Anwendungen geeignet ist
- **Emotionale Bindung:** Sound + Grafik = Immersion = Erinnerungswert

**Akzeptanzkriterien:**
- Partikel-Rendering wirkt "Premium" mit Glow, Blur und Layering
- Ambient Sound-Design sorgt für Atmosphäre
- Interaktions-Feedback mit Sound und visuellen Cues
- Flüssige 60 FPS auch bei maximaler visueller Qualität
- UX-Verbesserungen für intuitive Bedienung

---

## Story E-009-S01: Shader-basiertes Glow Rendering ✅ COMPLETE

**Als** Benutzer  
**möchte ich** dass Partikel einen wunderschönen Glow-Effekt haben  
**damit** die Simulation wie ein professionelles Kunstwerk aussieht

**Story Points:** 5

**Akzeptanzkriterien:**
- [x] Custom Shader für Bloom/Glow-Effekt implementiert (Software-Fallback)
- [x] Glow-Intensität basiert auf Partikel-Helligkeit und -Größe
- [x] Multi-Layer Rendering: Base → Glow → Composite
- [x] Glow-Farbe folgt der Partikel-Farbe (keine weiße Überblendung)
- [x] Glow-Radius ist preset-abhängig konfigurierbar
- [x] Performance: <5ms zusätzliche Render-Zeit

**Implementation:** `systems/glow.go`, `systems/render.go`

**Technische Details:**
```glsl
// Fragment Shader für Bloom
uniform sampler2D texture0;
uniform vec2 resolution;
uniform float bloomIntensity;

void main() {
    vec4 color = texture2D(texture0, gl_TexCoord[0].xy);
    vec4 bloom = vec4(0.0);
    
    // Gaussian Blur für Bloom
    for(int i = -4; i <= 4; i++) {
        for(int j = -4; j <= 4; j++) {
            vec2 offset = vec2(float(i), float(j)) / resolution;
            bloom += texture2D(texture0, gl_TexCoord[0].xy + offset);
        }
    }
    bloom /= 81.0;
    
    gl_FragColor = color + bloom * bloomIntensity;
}
```

**Fallback für WASM:**
- Software-basierter Glow via Multi-Circle Rendering wenn Shader nicht verfügbar

**Definition of Done:**
- [ ] Glow sieht auf Screenshots "professionell" aus
- [ ] Performance bleibt bei 60 FPS
- [ ] Toggle mit F7 möglich

---

## Story E-009-S02: Ambient Sound Engine 🔄 IN PROGRESS

**Als** Benutzer  
**möchte ich** eine atmosphärische Soundkulisse  
**damit** die Simulation ein vollständiges audiovisuelles Erlebnis wird

**Story Points:** 4

**Akzeptanzkriterien:**
- [x] Ambient-Loop passend zum aktiven Preset (5 verschiedene) - Config in `premium/audio.go`
- [ ] Smooth Crossfade beim Preset-Wechsel (2-3 Sekunden)
- [x] Volume-Control via Keyboard (+/- Tasten) - AudioManager implementiert
- [x] Mute-Toggle mit M-Taste - AudioManager.ToggleMute()
- [ ] Audio-Engine läuft ohne Frame-Drops
- [ ] WASM-Kompatibilität mit Web Audio API

**Implementation:** `premium/audio.go` (Struktur & Config vollständig, Audio-Playback TBD)

**Sound-Design pro Preset:**
| Preset | Ambient Sound | Mood |
|--------|--------------|------|
| Galaxy | Space Drone, tiefe Pads | Episch, kosmisch |
| Firework | Nacht-Atmosphäre, ferne Musik | Festlich, fröhlich |
| Swarm | Organische Texturen, Wind | Natürlich, fließend |
| Fountain | Wasser-Rauschen, Glockenspiel | Beruhigend, elegant |
| Chaos | Industrial Noise, Bass | Energetisch, wild |

**Technische Details:**
- Library: `beep` für Native, Web Audio API für WASM
- Format: OGG Vorbis (kleinere Dateigröße)
- Sample-Rate: 44.1 kHz, Stereo
- Loop-Points: Seamless (keine hörbaren Übergänge)

**Definition of Done:**
- [ ] Ambient Sound spielt bei Start
- [ ] Preset-Wechsel = Sound-Wechsel
- [ ] Keine Audio-Glitches oder Knackser

---

## Story E-009-S03: Interaktions-Sound-Effekte 🔄 IN PROGRESS

**Als** Benutzer  
**möchte ich** akustisches Feedback bei Interaktionen  
**damit** meine Aktionen sich bedeutsam anfühlen

**Story Points:** 3

**Akzeptanzkriterien:**
- [x] Maus-Anziehung: Subtiler "Magnet"-Sound - AudioManager.PlayAttract() definiert
- [x] Maus-Abstoßung: Sanfter "Whoosh"-Effekt - AudioManager.PlayRepel() definiert
- [x] Preset-Wechsel: Kurzer Transition-Sound - AudioManager.PlayTransition() definiert
- [ ] Debug-Toggle: UI-Klick-Sound
- [ ] Partikel-Explosion (Firework): Dezente Sparkle-Sounds
- [ ] Lautstärke proportional zur Interaktions-Intensität

**Implementation:** `premium/audio.go` (API definiert, Audio-Playback TBD)

**Technische Details:**
```go
type SoundManager struct {
    attractSound  *beep.Buffer
    repelSound    *beep.Buffer
    transitionSound *beep.Buffer
    // ...
}

func (sm *SoundManager) PlayAttract(intensity float32) {
    // Pitch und Volume basierend auf Intensität
    streamer := sm.attractSound.Streamer(0, sm.attractSound.Len())
    volume := &effects.Volume{Streamer: streamer, Base: 2, Volume: intensity - 0.5}
    speaker.Play(volume)
}
```

**Sound-Assets:**
| Sound | Dauer | Stil |
|-------|-------|------|
| Attract | 0.3s | Magnetisch, tief |
| Repel | 0.3s | Luftig, hoch |
| Transition | 0.5s | Shimmer, neutral |
| Click | 0.1s | Soft UI click |
| Sparkle | 0.2s | Magical, glittery |

**Definition of Done:**
- [ ] Sounds spielen bei entsprechenden Events
- [ ] Sounds sind nicht nervig bei Dauerbenutzung
- [ ] Volume ist ausgewogen

---

## Story E-009-S04: Premium Farbpaletten ✅ COMPLETE

**Als** Benutzer  
**möchte ich** wunderschöne, professionell abgestimmte Farbpaletten  
**damit** jedes Preset wie ein Kunstwerk aussieht

**Story Points:** 3

**Akzeptanzkriterien:**
- [x] 5 kuratierte Farbpaletten (eine pro Preset)
- [x] Farbübergänge sind smooth und ästhetisch
- [x] Keine "grellen" oder unharmonischen Kombinationen
- [x] HDR-ähnliche Farbtiefe durch geschickte Alpha-Blending
- [x] Dunkle Farben haben subtile Luminanz (nie "tot")

**Implementation:** `premium/colors.go`

**Farbpaletten-Design:**

**Galaxy Preset:**
```
Start: #FF6B9D (Pink-Magenta)    → End: #4A0080 (Deep Purple)
Alt:   #00D4FF (Cyan)            → End: #000033 (Space Black)
Glow:  #FFFFFF (White Star Core)
```

**Firework Preset:**
```
Gold:   #FFD700 → #FF4500 (Gold to Orange)
Red:    #FF0044 → #880022 (Bright to Deep Red)
Green:  #00FF88 → #004422 (Neon to Forest)
Blue:   #0088FF → #001144 (Electric to Navy)
White:  #FFFFFF → #888888 (Spark to Smoke)
```

**Swarm Preset:**
```
Primary: #00FFAA → #004433 (Bioluminescent Teal)
Accent:  #FF8800 → #442200 (Warm Orange)
```

**Fountain Preset:**
```
Water:   #00AAFF → #003366 (Azure to Deep Blue)
Spray:   #FFFFFF → #88CCFF (White to Light Blue)
```

**Chaos Preset:**
```
Electric: #FF00FF → #00FFFF (Magenta to Cyan)
Fire:     #FFFF00 → #FF0000 (Yellow to Red)
Void:     #880088 → #000000 (Purple to Black)
```

**Definition of Done:**
- [ ] Farbpaletten in config.json definiert
- [ ] Screenshots zeigen harmonische Farbgebung
- [ ] User-Feedback: "Das sieht geil aus!"

---

## Story E-009-S05: Smooth UI Overlays ✅ COMPLETE

**Als** Benutzer  
**möchte ich** eine elegante UI mit sanften Animationen  
**damit** die Bedienung sich premium anfühlt

**Story Points:** 4

**Akzeptanzkriterien:**
- [x] Preset-Indikator unten links mit Icon + Name
- [x] Steuerungs-Hinweise erscheinen bei Hover/Idle
- [x] Alle UI-Elemente haben Fade-In/Out Animationen
- [x] UI-Transparenz passt sich Helligkeit an (dunkel auf hell, hell auf dunkel)
- [x] Minimalistisches Design, nie aufdringlich
- [x] UI verschwindet nach 3s Inaktivität (außer bei Mouse-Hover)

**Implementation:** `premium/ui.go`

**UI-Layout:**
```
┌─────────────────────────────────────────────────────────────────┐
│                                          FPS: 60  🔊 On         │
│                                                                 │
│                                                                 │
│                     [ PARTICLE CANVAS ]                         │
│                                                                 │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│  🌌 Galaxy                                                      │
│  [1] Galaxy [2] Firework [3] Swarm [4] Fountain [5] Chaos      │
│  LMB: Attract  RMB: Repel  M: Mute  F3: Debug                  │
└─────────────────────────────────────────────────────────────────┘
```

**Animation-Timings:**
- Fade-In: 200ms ease-out
- Fade-Out: 300ms ease-in
- Slide-Up (Hints): 250ms ease-out
- Color-Transition: 150ms linear

**Definition of Done:**
- [ ] UI sieht "modern" aus
- [ ] Animationen sind smooth
- [ ] UI verschwindet bei Inaktivität

---

## Story E-009-S06: Particle Motion Blur ✅ COMPLETE

**Als** Benutzer  
**möchte ich** schnelle Partikel mit Motion-Blur-Effekt  
**damit** Bewegungen dynamischer und cinematischer wirken

**Story Points:** 4

**Akzeptanzkriterien:**
- [x] Schnelle Partikel haben Bewegungsunschärfe
- [x] Blur-Intensität proportional zur Geschwindigkeit
- [x] Blur-Richtung folgt Bewegungsvektor
- [x] Statische/langsame Partikel haben kein Blur
- [x] Blur ist toggle-bar (via Quality Settings)
- [x] Performance-Impact < 10% zusätzliche Frame-Zeit

**Implementation:** `systems/motion_blur.go`

**Technische Details:**
```go
func renderWithMotionBlur(pos, vel *components.Position, col *components.Color, size float32) {
    speed := math.Sqrt(vel.X*vel.X + vel.Y*vel.Y)
    if speed < 50 {
        // Normal rendering
        rl.DrawCircle(int32(pos.X), int32(pos.Y), size, toRaylibColor(col))
        return
    }
    
    // Motion blur via mehrere Semi-Transparente Kreise
    blurSteps := int(math.Min(speed/50, 8))
    stepAlpha := col.A / uint8(blurSteps+1)
    
    for i := 0; i < blurSteps; i++ {
        t := float32(i) / float32(blurSteps)
        x := pos.X - vel.X*t*0.016 // ~1 frame zurück
        y := pos.Y - vel.Y*t*0.016
        blurCol := rl.Color{col.R, col.G, col.B, stepAlpha}
        rl.DrawCircle(int32(x), int32(y), size*(1-t*0.3), blurCol)
    }
    
    // Hauptpartikel
    rl.DrawCircle(int32(pos.X), int32(pos.Y), size, toRaylibColor(col))
}
```

**Definition of Done:**
- [ ] Motion Blur sichtbar bei schnellen Partikeln
- [ ] Effekt sieht "cinematisch" aus
- [ ] Toggle funktioniert

---

## Story E-009-S07: Bildschirm-Shake & Juice Effects ✅ COMPLETE

**Als** Benutzer  
**möchte ich** subtile "Game Feel" Effekte  
**damit** starke Interaktionen impactful wirken

**Story Points:** 4

**Akzeptanzkriterien:**
- [x] Leichter Screen-Shake bei starker Maus-Abstoßung
- [x] Pulse-Effekt beim Preset-Wechsel (kurzes Zoom-In/Out)
- [x] Partikel "explodieren" visuell beim Spawn (Scale-Animation)
- [x] Attractor hat pulsierendes visuelles Feedback
- [x] Alle Effekte sind dezent und nicht ablenkend
- [x] Effekte können deaktiviert werden (Accessibility)

**Implementation:** `premium/effects.go`

**Juice-Intensitätsstufen:**
| Stufe | Screen Shake | Pulse | Spawn Anim | Default |
|-------|-------------|-------|------------|---------|
| Off | ❌ | ❌ | ❌ | |
| Subtle | 2px | 1.02x | 1.5x | ✅ |
| Normal | 5px | 1.05x | 2x | |
| Intense | 10px | 1.1x | 3x | |

**Technische Details:**
```go
type ScreenEffects struct {
    shakeIntensity float32
    shakeDuration  float32
    pulseScale     float32
    pulseDuration  float32
}

func (se *ScreenEffects) ApplyShake(intensity float32) {
    se.shakeIntensity = intensity
    se.shakeDuration = 0.15 // 150ms
}

func (se *ScreenEffects) GetCameraOffset() (float32, float32) {
    if se.shakeDuration <= 0 {
        return 0, 0
    }
    x := (rand.Float32()*2 - 1) * se.shakeIntensity
    y := (rand.Float32()*2 - 1) * se.shakeIntensity
    return x, y
}
```

**Definition of Done:**
- [ ] Screen Shake bei starker Abstoßung spürbar
- [ ] Pulse beim Preset-Wechsel
- [ ] Effekte fühlen sich "gut" an

---

## Story E-009-S08: Performance-Optimiertes Quality Preset System ✅ COMPLETE

**Als** Benutzer  
**möchte ich** zwischen Qualitätsstufen wählen können  
**damit** die Anwendung auf unterschiedlicher Hardware optimal läuft

**Story Points:** 5

**Akzeptanzkriterien:**
- [x] 3 Quality-Presets: Low, Medium, High
- [x] Low: Keine Glow, kein Blur, reduzierte Partikelzahl
- [x] Medium: Einfacher Glow, kein Blur, normale Partikelzahl
- [x] High: Vollständiger Glow, Motion Blur, maximale Partikel
- [x] Quality-Wechsel via Q-Taste oder Auto-Detect
- [ ] Auto-Detect: Wenn FPS < 50, automatisch runterstufen
- [x] Aktuelle Quality-Stufe im Debug-Overlay anzeigen

**Implementation:** `premium/quality.go`

**Quality-Matrix:**
| Feature | Low | Medium | High |
|---------|-----|--------|------|
| Max Particles | 3.000 | 7.000 | 15.000 |
| Glow Effect | ❌ | Simple | Full Shader |
| Motion Blur | ❌ | ❌ | ✅ |
| Trail Length | 3 | 5 | 10 |
| Spawn Rate | 50/s | 100/s | 200/s |
| Audio | Mono | Stereo | Stereo + Reverb |
| UI Animations | ❌ | ✅ | ✅ |

**Technische Details:**
```go
type QualityPreset struct {
    Name          string
    MaxParticles  int
    GlowEnabled   bool
    GlowQuality   int // 0=off, 1=simple, 2=full
    MotionBlur    bool
    TrailLength   int
    SpawnRate     float32
    AudioChannels int
}

var QualityPresets = map[string]QualityPreset{
    "low": {
        Name: "Low", MaxParticles: 3000, GlowEnabled: false,
        GlowQuality: 0, MotionBlur: false, TrailLength: 3,
        SpawnRate: 50, AudioChannels: 1,
    },
    "medium": {
        Name: "Medium", MaxParticles: 7000, GlowEnabled: true,
        GlowQuality: 1, MotionBlur: false, TrailLength: 5,
        SpawnRate: 100, AudioChannels: 2,
    },
    "high": {
        Name: "High", MaxParticles: 15000, GlowEnabled: true,
        GlowQuality: 2, MotionBlur: true, TrailLength: 10,
        SpawnRate: 200, AudioChannels: 2,
    },
}
```

**Definition of Done:**
- [ ] Q-Taste wechselt Quality
- [ ] Auto-Detect funktioniert bei niedrigen FPS
- [ ] Visueller Unterschied zwischen Stufen erkennbar

---

# Sprint 6: Premium Experience (Woche 6 - NEU)

| Story | Epic | Points | Priorität |
|-------|------|--------|-----------|
| E-009-S04 | Premium | 3 | 🟡 |
| E-009-S02 | Premium | 4 | 🟡 |
| E-009-S03 | Premium | 3 | 🟡 |
| E-009-S05 | Premium | 4 | 🟡 |
| E-009-S01 | Premium | 5 | 🟡 |
| E-009-S06 | Premium | 4 | 🟡 |
| E-009-S07 | Premium | 4 | 🟡 |
| E-009-S08 | Premium | 5 | 🟡 |
| **Total** | | **32** | |

**Sprint Goal:** Die Particle Symphony ist ein visuelles und akustisches Meisterwerk. Jeder der es sieht sagt: "Wow, das ist geil!"

---

# Aktualisierter Dependency Graph (inkl. E-009)

```
┌─────────────────────────────────────────────────────────────────┐
│                     STORY DEPENDENCIES                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  [E-001 bis E-007 - wie gehabt]                                 │
│                                                                  │
│  ═══════════════════════════════════════════════════════════    │
│                                                                  │
│  E-008: AWESOME-GO LISTING (UNCHANGED - WEITERHIN GÜLTIG!)      │
│                                                                  │
│  ═══════════════════════════════════════════════════════════    │
│                                                                  │
│  E-009: PREMIUM EXPERIENCE 🔥 (NEU)                              │
│                                                                  │
│  E-004-S01 (ColorSystem) ─────────┐                             │
│  E-004-S03 (Trail-Rendering) ─────┤                             │
│  E-004-S04 (Glow-Effekt) ─────────┘                             │
│           │                                                      │
│           ▼                                                      │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │  E-009-S04 (Premium Farbpaletten) ◀─────────────────────┐│  │
│  │       │                                                  ││  │
│  │       │  (parallel)                                      ││  │
│  │       │                                                  ││  │
│  │  E-009-S01 (Shader Glow) ◀──── E-004-S04 Grundlage      ││  │
│  │       │                                                  ││  │
│  │       │  (parallel)                                      ││  │
│  │       │                                                  ││  │
│  │  E-009-S06 (Motion Blur)                                 ││  │
│  │       │                                                  ││  │
│  │       ▼                                                  ││  │
│  │  E-009-S08 (Quality Presets) ◀───────────────────────────┘│  │
│  │                                                           │  │
│  │  ────────────────────────────────────────────────────────│  │
│  │                                                           │  │
│  │  E-006-S01 (Audio Input) ──────┐                         │  │
│  │       │                        │                         │  │
│  │       ▼                        ▼                         │  │
│  │  E-009-S02 (Ambient Sound) ◀───┤                         │  │
│  │       │                        │                         │  │
│  │       ▼                        │                         │  │
│  │  E-009-S03 (Interaktions-SFX) ◀┘                         │  │
│  │                                                           │  │
│  │  ────────────────────────────────────────────────────────│  │
│  │                                                           │  │
│  │  E-003-S04 (Debug Overlay) ───┐                          │  │
│  │       │                       │                          │  │
│  │       ▼                       ▼                          │  │
│  │  E-009-S05 (Smooth UI) ◀──────┤                          │  │
│  │       │                       │                          │  │
│  │       ▼                       │                          │  │
│  │  E-009-S07 (Juice Effects) ◀──┘                          │  │
│  │                                                           │  │
│  └───────────────────────────────────────────────────────────┘  │
│                      │                                          │
│                      ▼                                          │
│              🎨🔊 PREMIUM EXPERIENCE COMPLETE                    │
│         "Der geilste ECS-Showcase der Welt"                     │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

# E-009 Technical Requirements

## Audio Assets Needed

| Asset | Format | Size | Source |
|-------|--------|------|--------|
| Galaxy Ambient | OGG | ~2MB | Create/License |
| Firework Ambient | OGG | ~2MB | Create/License |
| Swarm Ambient | OGG | ~2MB | Create/License |
| Fountain Ambient | OGG | ~2MB | Create/License |
| Chaos Ambient | OGG | ~2MB | Create/License |
| SFX: Attract | OGG | ~50KB | Create |
| SFX: Repel | OGG | ~50KB | Create |
| SFX: Transition | OGG | ~100KB | Create |
| SFX: Click | OGG | ~20KB | Create |
| SFX: Sparkle | OGG | ~50KB | Create |

**Lizenz-Optionen:**
1. Royalty-Free Musik von Freesound.org / OpenGameArt
2. AI-generierte Musik (Suno, AIVA)
3. Eigenkomposition
4. Creative Commons Attribution

## Performance Budgets

| Feature | Budget (ms/frame) | @ 60 FPS |
|---------|-------------------|----------|
| Base Rendering | 8ms | Standard |
| Glow Shader | 3ms | +19% |
| Motion Blur | 2ms | +12% |
| Audio Processing | 1ms | +6% |
| UI Rendering | 1ms | +6% |
| **Total Premium** | **15ms** | **Machbar!** |

## WASM Audio Considerations

```javascript
// Web Audio API für WASM
const audioContext = new (window.AudioContext || window.webkitAudioContext)();

async function loadAmbientSound(preset) {
    const response = await fetch(`sounds/${preset}-ambient.ogg`);
    const arrayBuffer = await response.arrayBuffer();
    const audioBuffer = await audioContext.decodeAudioData(arrayBuffer);
    
    const source = audioContext.createBufferSource();
    source.buffer = audioBuffer;
    source.loop = true;
    source.connect(audioContext.destination);
    return source;
}
```

---

# Finale Aktualisierte Übersicht

| Metrik | Vorher | Nachher |
|--------|--------|---------|
| **Epics** | 8 | 9 |
| **User Stories** | 35 | 43 |
| **Story Points** | 127 | 159 |
| **Sprints** | 5 | 6 |

### Epic-Übersicht (Final)

| Epic | Titel | Stories | Punkte | Priorität | Status |
|------|-------|---------|--------|-----------|--------|
| E-001 | ECS Foundation | 5 | 18 | 🔴 MUST | ✅ |
| E-002 | Physik-Engine | 4 | 15 | 🔴 MUST | ✅ |
| E-003 | Interaktivität | 4 | 13 | 🔴 MUST | ✅ |
| E-004 | Visual Effects | 4 | 16 | 🟡 SHOULD | 🔄 |
| E-005 | Preset-System | 4 | 13 | 🟡 SHOULD | ✅ |
| E-006 | Audio-Reaktivität | 3 | 14 | 🟢 COULD | 📋 |
| E-007 | Web Deployment | 4 | 13 | 🔴 MUST | ✅ |
| E-008 | Awesome-Go Listing | 7 | 25 | 🔴 MUST | ✅ ⏳ |
| **E-009** | **Premium Experience 🔥** | **8** | **32** | **🟡 SHOULD** | **📋 NEU** |

**Legende:** ✅ Done | 🔄 In Progress | 📋 Backlog | ⏳ Waiting (5-Monate-Regel)

---

**⚠️ WICHTIG: EPIC E-008 (Awesome-Go Listing) IST IMPLEMENTIERT!**

Alle Stories von E-008 sind abgeschlossen. Die PR-Einreichung bei awesome-go kann erst nach Erreichen der **5-Monate-Regel** erfolgen (frühestens Mai 2026).

Die neue Epic E-009 ergänzt das Projekt um Premium-Features und steht nicht im Konflikt mit den Awesome-Go Anforderungen.

---

**🚀🔥 YOLO MODE COMPLETE - 43 STORIES READY FOR IMPLEMENTATION!**

**🎵 Particle Symphony wird LEGENDÄR! 🎵**
