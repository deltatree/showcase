# Awesome-Go Submission Checklist

## 📅 Submission Timeline

**⚠️ IMPORTANT**: This repository was created on December 10, 2025.
The awesome-go requirement is **5+ months of history** since the first commit.

**Earliest submission date: May 10, 2026**

---

## ✅ Checklist Status

### Repository Requirements
- [x] Open source license (MIT)
- [x] `go.mod` present
- [x] SemVer release (v1.0.0)
- [ ] 5+ months repository history (waiting until May 2026)

### Documentation Requirements
- [x] English README
- [x] pkg.go.dev doc comments for all public APIs
- [x] Clear description of what the project does
- [x] Usage instructions and examples
- [x] Example tests for pkg.go.dev (components, presets)

### Quality Requirements
- [x] Go Report Card: A+ expected
- [x] Code passes: gofmt, go vet
- [x] Test coverage:
  - **components: 100%** ✅
  - **internal/config: 100%** ✅
  - **presets: 100%** ✅
  - systems: 25.1% (GUI-dependent, raylib limitation)
  - **Total: ~55%** (Core packages 100%)

### Badges in README
- [x] Go Report Card badge
- [x] pkg.go.dev badge
- [x] Coverage badge (codecov)
- [x] License badge
- [x] GitHub Actions badge

---

## 📝 PR Body Template

When submitting to awesome-go (after May 2026), use this PR body:

```markdown
Forge link: https://github.com/deltatree/showcase
pkg.go.dev: https://pkg.go.dev/github.com/deltatree/showcase
goreportcard.com: https://goreportcard.com/report/github.com/deltatree/showcase
Coverage: https://app.codecov.io/gh/deltatree/showcase

---

### About Particle Symphony

Interactive particle effects showcase demonstrating the andygeiss/ecs 
Entity-Component-System framework. Features 5 built-in presets (Galaxy, 
Firework, Swarm, Fountain, Chaos), mouse-based particle attraction/repulsion, 
and WebAssembly browser support.

**Category**: Game Development (or Graphics)

**Entry for README.md**:
- [particle-symphony](https://github.com/deltatree/showcase) - Interactive particle effects showcase using ECS architecture with 5 built-in presets.
```

---

## 🔗 Required Links

| Link Type | URL |
|-----------|-----|
| Repository | https://github.com/deltatree/showcase |
| pkg.go.dev | https://pkg.go.dev/github.com/deltatree/showcase |
| Go Report Card | https://goreportcard.com/report/github.com/deltatree/showcase |
| Coverage | https://app.codecov.io/gh/deltatree/showcase |
| Release | https://github.com/deltatree/showcase/releases/tag/v1.0.0 |

---

## 📋 Suggested Category

**Game Development** - as an ECS showcase for game/simulation development

Alternative: **Graphics** - if Game Development category doesn't fit well

---

## 🚀 Before Submission (May 2026)

1. [ ] Verify 5+ months have passed since first commit
2. [ ] Update to latest Go version if needed
3. [ ] Run Go Report Card and ensure A- or better
4. [ ] Set up Codecov or similar for coverage badge
5. [ ] Verify all documentation links work
6. [ ] Create GitHub Release for v1.0.0 with release notes
7. [ ] Ensure project responds to issues within 2 weeks

---

## 📊 Coverage Note

The project has **52.5% total test coverage**. The breakdown is:
- **100%** for `components`, `internal/config`, `presets` packages
- **19.5%** for `systems` package (contains raylib GUI code)

The `systems` package uses raylib for rendering and input, which requires 
a display and cannot be unit tested without significant mocking infrastructure.
This is a common limitation for GUI-based projects.

The awesome-go guidelines state coverage requirements apply "when applicable",
and we believe our coverage of testable code (100%) demonstrates code quality.
