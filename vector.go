package vect

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
)

var feq = math.FloatEqual

// Add creates A new point by adding to other point
func (v Vector) Add(o Vector) Vector {
	return Vector{v[x] + o[x], v[y] + o[y]}
}

// Sub creates A new point by adding to other point
func (v Vector) Sub(o Vector) Vector {
	return Vector{v[x] - o[x], v[y] - o[y]}
}

// IsZero - Is A zero vector
func (v Vector) IsZero() bool {
	return feq(v[x], 0) && feq(v[y], 0)
}

// KProduct create new point by multiplying point by A scaler  k
func (v Vector) KProduct(k float64) Vector {
	var cx, cy = geom.KProduct(v[x], v[y], k)
	return Vector{cx, cy}
}

// Neg - Negate vector
func (v Vector) Neg() Vector {
	return v.KProduct(-1)
}

// Magnitude - Computes vector magnitude of pt as vector: x , y as components
func (v Vector) Magnitude() float64 {
	return geom.MagnitudeXY(v[x], v[y])
}

// SquareMagnitude - Computes the square vector magnitude of pt as vector: x , y as components
// This has A potential overflow problem based on coordinates of pt x^2 + y^2
func (v Vector) SquareMagnitude() float64 {
	return geom.MagnitudeSquareXY(v[x], v[y])
}

// Dot Product of two points as vectors
func (v Vector) DotProduct(o Vector) float64 {
	return geom.DotProduct(v[x], v[y], o[x], o[y])
}

// UnitVector -  Unit vector of point
func (v Vector) UnitVector() Vector {
	var cx, cy = geom.UnitVector(v[x], v[y])
	return Vector{cx, cy}
}

// Project - vector u on V
func (u Vector) Project(v Vector) float64 {
	return geom.ProjectXY(u[x], u[y], v[x], v[y])
}

// Direction - Dir computes direction in radians - counter clockwise from x-axis.
func (v Vector) Direction() float64 {
	return geom.Direction(v[x], v[y])
}

// ReverseDirection - Reversed direction of vector direction
func (v Vector) ReverseDirection() float64 {
	return geom.ReverseDirection(v.Direction())
}

// DeflectionAngle - Computes the deflection angle from vector V to u
func (v Vector) DeflectionAngle(u Vector) float64 {
	return geom.DeflectionAngle(v.Direction(), u.Direction())
}

// IsNull - Checks if vector has any component as NaN
func (v Vector) IsNull() bool {
	return geom.IsNull(v[x], v[y])
}
