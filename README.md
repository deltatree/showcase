# Particle Symphony 🎵✨

[![Deploy to GitHub Pages](https://github.com/deltatree/showcase/actions/workflows/deploy-pages.yml/badge.svg)](https://github.com/deltatree/showcase/actions/workflows/deploy-pages.yml)

An Entity Component System (ECS) showcase demonstrating beautiful particle simulations in Go.

## 🌐 Live Demo

**[🎆 Try it in your Browser →](https://deltatree.github.io/showcase/)**

No installation required! The web version runs entirely in WebAssembly.

## 🚀 Quick Start

### Native Version (raylib)

For the best performance, run the native version:

```bash
# Clone the repository
git clone https://github.com/deltatree/showcase.git
cd showcase

# Build and run
go build -o particle-symphony . && ./particle-symphony
```

### WebAssembly Version (Ebitengine)

Build the WASM version locally:

```bash
./build-wasm.sh
cd web && python3 -m http.server 8080
# Open http://localhost:8080
```

## 🎮 Controls

| Key | Action |
|-----|--------|
| `1` | Galaxy Preset |
| `2` | Firework Preset |
| `3` | Swarm Preset |
| `4` | Fountain Preset |
| `5` | Chaos Preset |
| `LMB` | Attract Particles |
| `RMB` | Repel Particles |
| `2× Click` | Lock Attract/Repel |
| `F3` | Toggle Debug Overlay |
| `ESC` | Exit (Native only) |

## ✨ Features

- **Entity Component System** - Clean, modular architecture
- **Multiple Presets** - Fountain, Firework, Galaxy, Swarm, and Chaos effects
- **Real-time Physics** - Gravity, damping, and velocity simulation
- **Color Transitions** - Smooth gradient color animations
- **Lifetime Management** - Particle birth, aging, and death cycles

## 🏗️ Architecture

The project follows a pure ECS architecture:

- **Components** - Data containers (Position, Velocity, Color, Mass, etc.)
- **Systems** - Logic processors (Physics, Render, Emitter, etc.)
- **Entities** - Unique identifiers linking components together

## 📁 Project Structure

```
├── components/     # ECS component definitions
├── systems/        # ECS system implementations
├── presets/        # Particle effect presets
├── internal/       # Internal packages (config)
├── web/            # Web showcase page
└── docs/           # Project documentation
```

## 📝 License

MIT License - feel free to use this as a learning resource or starting point for your own ECS projects.
