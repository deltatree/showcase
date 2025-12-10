---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]
inputDocuments: []
documentCounts:
  briefs: 0
  research: 1
  brainstorming: 0
  projectDocs: 0
workflowType: 'prd'
lastStep: 11
project_name: 'showcase'
user_name: 'Deltatree'
date: '2025-12-10'
yolo_mode: true
---

# Product Requirements Document - Particle Symphony

**Autor:** Deltatree  
**Datum:** 10. Dezember 2025  
**Status:** Finalisiert (YOLO MODE)

---

## 1. Executive Summary

### 1.1 Produktvision

**Particle Symphony** ist ein visuell beeindruckender, interaktiver Partikel-Simulator, der die Stärken des `andygeiss/ecs` Frameworks in einer unvergesslichen Demonstration präsentiert. Das Projekt kombiniert Echtzeit-Physik, emergentes Verhalten und optionale Audio-Reaktivität zu einem hypnotischen visuellen Erlebnis.

### 1.2 Warum dieser Showcase?

| ECS-Feature | Demonstration im Showcase |
|-------------|---------------------------|
| **Bitmask-Filterung** | Tausende Partikel mit unterschiedlichen Komponenten-Kombinationen werden in Echtzeit gefiltert |
| **Data-Oriented Design** | Saubere Trennung: Components = pure Daten, Systems = pure Logik |
| **Skalierbarkeit** | Von 100 bis 100.000 Partikel - die Architektur skaliert |
| **Null Dependencies** | Das ECS selbst braucht nichts - nur Raylib für Rendering |
| **System-Pipeline** | Mehrere Systems arbeiten sequentiell: Physics → Behavior → Render |

### 1.3 Das "WOW" Moment

Ein Benutzer startet die Anwendung und sieht sofort:
- Tausende Partikel, die sich organisch bewegen
- Emergente Schwarm-Muster durch einfache Regeln
- Interaktion mit der Maus erzeugt Gravitationsfelder
- Optional: Musik-Visualisierung wo Beats die Partikel beeinflussen

**Das bleibt hängen:** Die visuelle Komplexität entsteht aus simplen ECS-Patterns.

---

## 2. Erfolgskriterien

### 2.1 Primäre Erfolgsmetriken

| Metrik | Zielwert | Messmethode |
|--------|----------|-------------|
| **Performance** | 60 FPS bei 10.000 Partikeln | In-App FPS Counter |
| **Code-Klarheit** | Jedes System < 100 LOC | Code Review |
| **ECS-Pattern-Reinheit** | 100% Component/System Trennung | Architektur-Review |
| **Visueller Impact** | "Wow"-Reaktion bei Demo | User Feedback |
| **Build-Simplizität** | `go run .` funktioniert | CI/CD Pipeline |

### 2.2 Sekundäre Erfolgsmetriken

- README mit GIF/Video zeigt den Showcase in Aktion
- Benchmark-Vergleich: Mit vs. ohne ECS-Optimierungen
- Dokumentierter Code als Lernressource

---

## 3. User Journeys

### 3.1 Journey: Der neugierige Go-Entwickler

**Persona:** Alex, 28, Backend-Entwickler mit Go-Erfahrung, interessiert an Game Development

**Szenario:** Alex findet `andygeiss/ecs` auf GitHub und klickt auf den Showcase-Link

1. **Entdeckung** → Sieht README mit animiertem GIF → "Das sieht cool aus!"
2. **Installation** → `git clone && go run .` → Läuft sofort
3. **Interaktion** → Bewegt Maus, Partikel reagieren → "Wie funktioniert das?"
4. **Exploration** → Öffnet `components/` und `systems/` → "Ah, so simpel!"
5. **Modifikation** → Ändert Parameter, sieht sofort Ergebnis
6. **Adoption** → Nutzt ECS für eigenes Projekt

### 3.2 Journey: Der Showcase-Präsentierer

**Persona:** Deltatree, präsentiert ECS auf einem Meetup

1. **Vorbereitung** → Startet Showcase auf Präsentations-Laptop
2. **Demo** → Zeigt Partikel-Simulation, erklärt ECS-Konzepte
3. **Live-Coding** → Fügt neues System hinzu während Präsentation
4. **Impact** → Publikum sieht sofort den Effekt der Änderung

---

## 4. Domain Model

### 4.1 Core Entities

```
┌─────────────────────────────────────────────────────────────────┐
│                        ENTITY TYPES                              │
├─────────────────────────────────────────────────────────────────┤
│  Particle        │ Position + Velocity + Color + Lifetime       │
│  Attractor       │ Position + Mass + (kein Velocity)            │
│  Emitter         │ Position + EmissionRate + ParticleTemplate   │
│  AudioReactor    │ FrequencyBand + Intensity + TargetMask       │
└─────────────────────────────────────────────────────────────────┘
```

### 4.2 Components (Bitmask-Design)

```go
const (
    MaskPosition     = uint64(1 << 0)   // X, Y Koordinaten
    MaskVelocity     = uint64(1 << 1)   // VX, VY Geschwindigkeit
    MaskAcceleration = uint64(1 << 2)   // AX, AY Beschleunigung
    MaskColor        = uint64(1 << 3)   // R, G, B, A Farbe
    MaskLifetime     = uint64(1 << 4)   // TTL, Age
    MaskMass         = uint64(1 << 5)   // Masse für Gravitation
    MaskEmitter      = uint64(1 << 6)   // Partikel-Emitter
    MaskTrail        = uint64(1 << 7)   // Trail-Rendering
    MaskAudioReact   = uint64(1 << 8)   // Audio-Reaktivität
    MaskSize         = uint64(1 << 9)   // Partikel-Größe
)
```

### 4.3 Systems Pipeline

```
┌────────────────────────────────────────────────────────────────┐
│                      SYSTEM PIPELINE                            │
│                    (Ausführungsreihenfolge)                     │
├────────────────────────────────────────────────────────────────┤
│  1. InputSystem        │ Liest Maus/Keyboard → Setzt Attraktoren │
│  2. EmitterSystem      │ Spawnt neue Partikel                    │
│  3. GravitySystem      │ Berechnet Anziehung zu Attraktoren      │
│  4. PhysicsSystem      │ Velocity += Acceleration, Pos += Vel    │
│  5. LifetimeSystem     │ Altert Partikel, entfernt "tote"        │
│  6. ColorSystem        │ Interpoliert Farben basierend auf Age   │
│  7. TrailSystem        │ Aktualisiert Trail-Positionen           │
│  8. AudioSystem        │ (Optional) Modifiziert basierend auf FFT│
│  9. RenderSystem       │ Zeichnet alles mit Raylib               │
└────────────────────────────────────────────────────────────────┘
```

---

## 5. Innovation & Differenzierung

### 5.1 Was macht diesen Showcase besonders?

| Aspekt | Standard-Demo | Particle Symphony |
|--------|---------------|-------------------|
| **Visuell** | Bewegende Rechtecke | Tausende organische Partikel |
| **Interaktivität** | Keyboard-Input | Maus-Gravitation in Echtzeit |
| **Emergenz** | Vordefiniertes Verhalten | Schwarm-Intelligenz aus simplen Regeln |
| **Skalierung** | 10-100 Entities | 10.000+ Entities |
| **Lernwert** | "So funktioniert ECS" | "Deshalb ist ECS mächtig" |

### 5.2 Technische Innovationen

1. **Hot-Reload Config:** JSON/YAML Konfiguration für Parameter → Änderungen ohne Neustart
2. **Preset-System:** Verschiedene "Modi" (Firework, Galaxy, Swarm, etc.)
3. **Performance-Overlay:** Zeigt Entity-Count, FPS, System-Timings
4. **Screenshot/GIF-Export:** Für README und Social Media

---

## 6. Projekt-Typ & Scope

### 6.1 Technologie-Stack

| Komponente | Technologie | Begründung |
|------------|-------------|------------|
| **ECS Framework** | `andygeiss/ecs` | Das ist der Showcase-Fokus! |
| **Rendering** | `raylib-go` | Leichtgewichtig, Cross-Platform, Go-Bindings |
| **Audio (Optional)** | `portaudio` oder `beep` | FFT für Audio-Reaktivität |
| **Config** | JSON | Einfach, keine Dependencies |

### 6.2 Scope-Definition

#### In Scope (MVP)

- [x] Basis-Partikel-System mit Position, Velocity, Color
- [x] Maus-Interaktion als Gravitationsquelle
- [x] 3-5 vordefinierte Presets
- [x] Performance-Counter (FPS, Entity-Count)
- [x] Keyboard-Shortcuts für Presets
- [x] README mit GIF/Screenshot

#### In Scope (Enhanced)

- [ ] Trail-Rendering für Partikel
- [ ] Audio-Reaktivität (Beat-Detection)
- [ ] Mehr Presets (Galaxy, Firework, DNA-Helix)
- [ ] Config-Hot-Reload
- [ ] Screenshot-Export

#### Out of Scope

- Multiplayer/Networking
- 3D-Rendering
- Mobile Plattformen
- Persistenz/Speichern

---

## 7. Funktionale Anforderungen

### 7.1 Core Features

#### F-001: Partikel-Rendering
**Priorität:** MUST  
**Beschreibung:** Das System rendert bis zu 10.000 Partikel mit 60 FPS  
**Akzeptanzkriterien:**
- Partikel werden als farbige Kreise/Punkte gezeichnet
- Farbe interpoliert von Start- zu Endfarbe über Lifetime
- Alpha-Blending für weiche Übergänge

#### F-002: Physik-Simulation
**Priorität:** MUST  
**Beschreibung:** Realistische Bewegung durch Velocity und Acceleration  
**Akzeptanzkriterien:**
- Frame-unabhängige Bewegung (DeltaTime)
- Dämpfung für natürliches Abbremsen
- Boundary-Handling (Wrap oder Bounce)

#### F-003: Maus-Interaktion
**Priorität:** MUST  
**Beschreibung:** Maus erzeugt Gravitationsfeld  
**Akzeptanzkriterien:**
- Linke Maustaste: Anziehung
- Rechte Maustaste: Abstoßung
- Stärke proportional zur Entfernung

#### F-004: Preset-System
**Priorität:** SHOULD  
**Beschreibung:** Vordefinierte Konfigurationen  
**Akzeptanzkriterien:**
- Mindestens 3 Presets (Galaxy, Firework, Swarm)
- Keyboard-Shortcuts (1, 2, 3...)
- Smooth Transition zwischen Presets

#### F-005: Performance-Overlay
**Priorität:** SHOULD  
**Beschreibung:** Debug-Informationen on-screen  
**Akzeptanzkriterien:**
- FPS-Anzeige
- Entity-Count
- Toggle mit F3 oder ähnlich

#### F-006: Audio-Reaktivität
**Priorität:** COULD  
**Beschreibung:** Partikel reagieren auf Musik  
**Akzeptanzkriterien:**
- FFT-Analyse des Audio-Inputs
- Bass → Größe/Intensität
- Höhen → Geschwindigkeit/Farbe

---

## 8. Nicht-Funktionale Anforderungen

### 8.1 Performance

| Anforderung | Zielwert | Messung |
|-------------|----------|---------|
| **FPS bei 1k Entities** | ≥ 60 FPS | In-App |
| **FPS bei 10k Entities** | ≥ 60 FPS | In-App |
| **FPS bei 50k Entities** | ≥ 30 FPS | In-App |
| **Startup Time** | < 2 Sekunden | Stoppuhr |
| **Memory Usage** | < 100 MB bei 10k | Task Manager |

### 8.2 Code-Qualität

| Anforderung | Zielwert |
|-------------|----------|
| **Test Coverage** | > 80% für Core-Logic |
| **Cyclomatic Complexity** | < 10 pro Funktion |
| **Documentation** | Jedes exportierte Symbol dokumentiert |
| **Linting** | golangci-lint ohne Errors |

### 8.3 Portabilität

- **Primär:** macOS, Windows, Linux
- **Build:** `go build` ohne spezielle Toolchains (außer Raylib-Dependencies)
- **Dependencies:** Minimal, alles via `go mod`

---

## 9. Projekt-Struktur

```
showcase/
├── main.go                 # Entry Point
├── go.mod
├── go.sum
├── config.json             # Runtime-Konfiguration
├── README.md               # Mit GIF und Quickstart
├── components/
│   ├── components.go       # Mask-Konstanten
│   ├── position.go
│   ├── velocity.go
│   ├── acceleration.go
│   ├── color.go
│   ├── lifetime.go
│   ├── mass.go
│   └── size.go
├── systems/
│   ├── input.go
│   ├── emitter.go
│   ├── gravity.go
│   ├── physics.go
│   ├── lifetime.go
│   ├── color.go
│   ├── render.go
│   └── audio.go            # Optional
├── presets/
│   ├── galaxy.go
│   ├── firework.go
│   └── swarm.go
├── internal/
│   └── config/
│       └── loader.go
└── docs/
    ├── prd.md              # Dieses Dokument
    └── screenshots/
```

---

## 10. Risiken & Mitigationen

| Risiko | Wahrscheinlichkeit | Impact | Mitigation |
|--------|-------------------|--------|------------|
| Raylib-Installation komplex | Mittel | Hoch | Detaillierte README, Docker-Option |
| Performance-Probleme bei vielen Entities | Niedrig | Mittel | Spatial Partitioning als Fallback |
| Audio-Integration zu komplex | Mittel | Niedrig | Als optionales Feature markiert |
| Cross-Platform Issues | Niedrig | Mittel | CI/CD für alle Plattformen |

---

## 11. Definition of Done

Ein Feature gilt als "Done" wenn:

- [ ] Code implementiert und compiles
- [ ] Unit Tests geschrieben und grün
- [ ] Code Review bestanden
- [ ] Dokumentation aktualisiert
- [ ] Performance-Ziele erreicht (falls relevant)
- [ ] In README/Demo sichtbar

---

## 12. Appendix

### A. Referenzen

- [andygeiss/ecs GitHub](https://github.com/andygeiss/ecs)
- [engine-example Repository](https://github.com/andygeiss/engine-example)
- [Raylib](https://www.raylib.com/)
- [raylib-go Bindings](https://github.com/gen2brain/raylib-go)

### B. Inspiration

- [Particle Life](https://particle-life.com/) - Emergentes Verhalten
- [Boids Algorithm](https://www.red3d.com/cwr/boids/) - Schwarm-Simulation
- [Audio Visualizers](https://www.youtube.com/results?search_query=music+visualizer) - Visual Impact

### C. Glossar

| Begriff | Definition |
|---------|------------|
| **ECS** | Entity-Component-System, Architektur-Pattern für Game Development |
| **Entity** | Container für Components, hat eine ID |
| **Component** | Pure Data, keine Logik |
| **System** | Pure Logic, operiert auf Entities mit bestimmten Components |
| **Bitmask** | Effiziente Filterung von Entities nach Component-Kombination |

---

**🚀 YOLO MODE COMPLETE - LET'S BUILD THIS!**
