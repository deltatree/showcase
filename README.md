# Particle Symphony 🎵✨

An Entity Component System (ECS) showcase demonstrating beautiful particle simulations in Go using the Ebitengine game engine.

## 🌐 Live Demo

**[View the Web Showcase →](https://deltatree.github.io/showcase/)**

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- A graphics environment (for running locally)

### Build & Run

```bash
# Clone the repository
git clone https://github.com/deltatree/showcase.git
cd showcase

# Build and run
go build -o particle-symphony . && ./particle-symphony
```

## 🎮 Controls

| Key | Action |
|-----|--------|
| `1` | Fountain Preset |
| `2` | Firework Preset |
| `3` | Galaxy Preset |
| `4` | Swarm Preset |
| `5` | Chaos Preset |
| `Space` | Toggle Physics |
| `ESC` | Exit |

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
