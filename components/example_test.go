package components_test

import (
	"fmt"

	"github.com/deltatree/showcase/components"
)

// ExampleNewPosition demonstrates creating a Position component.
func ExampleNewPosition() {
	pos := components.NewPosition().With(100, 200)
	fmt.Printf("Position: (%.0f, %.0f)\n", pos.X, pos.Y)
	// Output: Position: (100, 200)
}

// ExampleNewVelocity demonstrates creating a Velocity component with magnitude.
func ExampleNewVelocity() {
	vel := components.NewVelocity().With(3, 4)
	fmt.Printf("Velocity: (%.0f, %.0f), Magnitude: %.0f\n", vel.X, vel.Y, vel.Magnitude())
	// Output: Velocity: (3, 4), Magnitude: 5
}

// ExampleNewColor demonstrates creating a Color with gradient.
func ExampleNewColor() {
	col := components.NewColor().WithGradient(
		255, 0, 0, 255, // Start: Red
		0, 0, 255, 0, // End: Blue (fading)
	)
	fmt.Printf("Start: RGBA(%d,%d,%d,%d)\n", col.StartR, col.StartG, col.StartB, col.StartA)
	fmt.Printf("End: RGBA(%d,%d,%d,%d)\n", col.EndR, col.EndG, col.EndB, col.EndA)
	// Output:
	// Start: RGBA(255,0,0,255)
	// End: RGBA(0,0,255,0)
}

// ExampleNewLifetime demonstrates Lifetime progress calculation.
func ExampleNewLifetime() {
	life := components.NewLifetime().WithTTL(10.0)
	life.Age = 5.0 // Halfway through life
	fmt.Printf("Progress: %.1f%%\n", life.Progress()*100)
	// Output: Progress: 50.0%
}

// ExampleNewAcceleration demonstrates Acceleration usage.
func ExampleNewAcceleration() {
	acc := components.NewAcceleration()
	acc.Add(10, 20)
	acc.Add(5, 5)
	fmt.Printf("Acceleration: (%.0f, %.0f)\n", acc.X, acc.Y)
	acc.Reset()
	fmt.Printf("After Reset: (%.0f, %.0f)\n", acc.X, acc.Y)
	// Output:
	// Acceleration: (15, 25)
	// After Reset: (0, 0)
}

// ExampleNewSize demonstrates Size with end size for interpolation.
func ExampleNewSize() {
	size := components.NewSize().WithRadius(10.0).WithEndSize(2.0)
	fmt.Printf("Start: %.1f, End: %.1f\n", size.StartSize, size.EndSize)
	// Output: Start: 10.0, End: 2.0
}

// ExampleNewMass demonstrates Mass for gravitational calculations.
func ExampleNewMass() {
	mass := components.NewMass().WithValue(50000)
	fmt.Printf("Mass: %.0f\n", mass.Value)
	// Output: Mass: 50000
}

// Example demonstrates creating a full particle entity.
func Example() {
	// Create components for a complete particle
	pos := components.NewPosition().With(640, 360)
	vel := components.NewVelocity().With(100, -50)
	acc := components.NewAcceleration()
	col := components.NewColor().WithGradient(255, 200, 0, 255, 255, 50, 0, 0)
	life := components.NewLifetime().WithTTL(3.0)
	size := components.NewSize().WithRadius(5.0).WithEndSize(1.0)

	// Check bitmasks for filtering
	mask := pos.Mask() | vel.Mask() | acc.Mask() | col.Mask() | life.Mask() | size.Mask()
	fmt.Printf("Combined Mask: %d\n", mask)
	fmt.Printf("Includes Position: %v\n", mask&components.MaskPosition != 0)
	fmt.Printf("Includes Velocity: %v\n", mask&components.MaskVelocity != 0)
	// Output:
	// Combined Mask: 95
	// Includes Position: true
	// Includes Velocity: true
}
