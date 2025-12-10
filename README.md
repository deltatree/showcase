# Particle Symphony 🎵✨

[![Deploy to GitHub Pages](https://github.com/deltatree/showcase/actions/workflows/deploy-pages.yml/badge.svg)](https://github.com/deltatree/showcase/actions/workflows/deploy-pages.yml)

A mesmerizing particle simulation showcasing the power and elegance of **[andygeiss/ecs](https://github.com/andygeiss/ecs)** – a lightweight, high-performance Entity Component System framework for Go.

> **🎯 This project exists to demonstrate how simple yet powerful ECS architecture can be.** Watch thousands of particles dance, interact, and create emergent beauty – all powered by clean, modular ECS patterns.

## 🌐 Live Demo

**[🎆 Try it in your Browser →](https://deltatree.github.io/showcase/)**

No installation required! The web version runs entirely in WebAssembly.

## ⚡ Powered By

| Technology | Purpose |
|------------|---------|
| **[andygeiss/ecs](https://github.com/andygeiss/ecs)** | 🏆 The star of the show! Lightweight ECS framework with bitmask-based entity filtering |
| **[raylib-go](https://github.com/gen2brain/raylib-go)** | Native desktop rendering (high performance) |
| **[Ebitengine](https://ebitengine.org)** | WebAssembly rendering (cross-platform) |
| **Go** | Because simplicity and performance matter |

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

The project follows a pure ECS architecture using **[andygeiss/ecs](https://github.com/andygeiss/ecs)**:

- **Components** - Pure data containers (Position, Velocity, Color, Mass, etc.)
- **Systems** - Logic processors (Physics, Render, Emitter, Gravity, etc.)
- **Entities** - Unique identifiers linking components together
- **Bitmask Filtering** - Blazingly fast entity queries using component masks

### Why andygeiss/ecs?

✅ **Zero Dependencies** - Pure Go, no external requirements  
✅ **Bitmask-based Filtering** - O(1) component lookups  
✅ **Clean API** - Simple, intuitive, Go-idiomatic  
✅ **Battle-tested** - Used in production game projects  
✅ **Minimal Footprint** - Small codebase, easy to understand

```go
// Example: How simple ECS can be!
particles := em.FilterByMask(MaskPosition | MaskVelocity | MaskColor)
for _, entity := range particles {
    pos := entity.Get(MaskPosition).(*Position)
    vel := entity.Get(MaskVelocity).(*Velocity)
    pos.X += vel.X * deltaTime
    pos.Y += vel.Y * deltaTime
}
```

## 📁 Project Structure

```
├── components/     # ECS component definitions
├── systems/        # ECS system implementations
├── presets/        # Particle effect presets
├── internal/       # Internal packages (config)
├── web/            # Web showcase page
├── cmd/wasm/       # WebAssembly version (Ebitengine)
└── docs/           # Project documentation
```

## 🙏 Credits & Acknowledgments

This project wouldn't exist without these amazing open-source projects:

- **[andygeiss/ecs](https://github.com/andygeiss/ecs)** by [@andygeiss](https://github.com/andygeiss) - The ECS framework that makes this all possible. Seriously, go check it out! ⭐
- **[raylib-go](https://github.com/gen2brain/raylib-go)** - Go bindings for the fantastic raylib game library
- **[Ebitengine](https://ebitengine.org)** by [@hajimehoshi](https://github.com/hajimehoshi) - Making Go games run everywhere, including the browser!

## 📝 License

MIT License - feel free to use this as a learning resource or starting point for your own ECS projects.

---

<div align="center">

**Built with ❤️ and mass quantities of ☕**

*If you found this useful, consider starring the repo and checking out [andygeiss/ecs](https://github.com/andygeiss/ecs)!*

</div>
