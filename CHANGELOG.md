# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Future enhancements will be documented here

## [1.0.0] - 2025-12-10

### Added
- Initial release of Particle Symphony
- Core ECS architecture using [andygeiss/ecs](https://github.com/andygeiss/ecs)
- 5 built-in particle presets:
  - **Galaxy**: Spiral galaxy with orbital particle motion
  - **Firework**: Colorful explosion bursts with gravity
  - **Swarm**: Organic swarm behavior following attractors
  - **Fountain**: Water fountain shooting upward
  - **Chaos**: Random particles with varied colors and velocities
- Interactive mouse controls:
  - Left click/hold: Attract particles
  - Right click/hold: Repel particles
  - Double-click: Lock attract/repel mode
- Keyboard controls:
  - 1-5: Switch between presets
  - F3: Toggle debug overlay
- Component system with:
  - Position, Velocity, Acceleration (physics)
  - Color with gradient interpolation
  - Lifetime with TTL/Age tracking
  - Size with animated scaling
  - Mass for gravitational interactions
- System architecture:
  - PhysicsSystem: Movement and edge wrapping
  - GravitySystem: N-body gravitational attraction
  - ColorSystem: Smooth color and size transitions
  - LifetimeSystem: Particle aging and cleanup
  - EmitterSystem: Configurable particle spawning
  - InputSystem: Mouse and keyboard handling
  - RenderSystem: raylib-based rendering with debug overlay
- WebAssembly support via Ebitengine
- Configuration via JSON file
- Comprehensive test suite (100% coverage for components, config, presets)
- Documentation for pkg.go.dev

### Technical Details
- Native desktop build using raylib-go
- WebAssembly build using Ebitengine
- Supports up to 10,000+ particles at 60 FPS
- MIT License

[Unreleased]: https://github.com/deltatree/showcase/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/deltatree/showcase/releases/tag/v1.0.0
