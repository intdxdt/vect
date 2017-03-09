package vect

import (
	"math"
	"simplex/geom"
	umath "simplex/util/math"
	"simplex/side"
	"simplex/cart2d"
)
//Vector Type
type Vector [2]float64

//New Vector given start and end point
func NewVector(a, b *geom.Point) *Vector {
	return &Vector{b[x] - a[x], b[y] - a[y]}
}

//Creates a new vector with component x and y
func NewVectorXY(x, y float64) *Vector {
	return &Vector{x, y}
}

//Creates a new vector with component x and y
func NewVectorMagDir(m, d float64) *Vector {
	cx, cy := cart2d.Component(m, d)
	return &Vector{cx, cy}
}

//Clone Vector
func (v *Vector) Clone() *Vector {
	return &Vector{v[x], v[y]}
}

//X gets the x compoent of vector
func (v *Vector) X() float64 {
	return v[x]
}

//Y gets the y component of vector
func (v *Vector) Y() float64 {
	return v[y]
}

//Add creates a new point by adding to other point
func (v *Vector) Add(o *Vector) *Vector {
	cx, cy := cart2d.Add(v, o)
	return &Vector{cx, cy}
}

//Is a zero vector
func (v *Vector) IsZero() bool {
	return umath.FloatEqual(v[x], 0.0) && umath.FloatEqual(v[y], 0.0)
}

//Sub creates a new point by adding to other point
func (v *Vector) Sub(o *Vector) *Vector {
	cx, cy := cart2d.Sub(v, o)
	return NewVectorXY(cx, cy)
}

//KProduct create new point by multiplying point by a scaler  k
func (v *Vector) KProduct(k float64) *Vector {
	cx, cy := cart2d.KProduct(v, k)
	return NewVectorXY(cx, cy)
}

//Negate vector
func (v *Vector) Neg() *Vector {
	return v.KProduct(-1.0)
}

//Computes vector magnitude of pt as vector: x , y as components
func (v *Vector) Magnitude() float64 {
	return cart2d.Magnitude(v)
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func (v *Vector)  SquareMagnitude() float64 {
	return cart2d.SquareMagnitude(v)
}

//Dot Product of two points as vectors
func (v *Vector) DotProduct(o *Vector) float64 {
	return cart2d.DotProductXY(v[x], v[y], o[x], o[y])
}

//Unit vector of point
func (v *Vector) UnitVector() *Vector {
	cx, cy := cart2d.Unit(v)
	return NewVectorXY(cx, cy)
}

//Project vector u on v
func (u *Vector) Project(v *Vector) float64 {
	return cart2d.Project(u, v)
}

//2D cross product of AB and AC vectors,
//i.e. z-component of their 3D cross product.
//Computes A---B---C : location C is on Left , On , or Right of line AB
//Returns a positive value, if AB-->BC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func (ab *Vector) SideOf(ac *Vector) *side.Side {
	s:= side.NewSide()
	ccw := cart2d.CrossProduct(ab, ac)
	if umath.FloatEqual(ccw, 0){
		s.AsOn()
	} else if ccw > 0 {
		s.AsLeft()
	} else {
		s.AsRight()
	}
	return s
}

//Dir computes direction in radians - counter clockwise from x-axis.
func (v *Vector) Direction() float64 {
	return cart2d.DirectionXY(v[x], v[y])
}

//Reversed direction of vector direction
func (v *Vector)  ReverseDirection() float64 {
	return cart2d.ReverseDirection(v.Direction())
}

//Computes the deflection angle from vector v to u
func (v *Vector)  DeflectionAngle(u *Vector) float64 {
	return cart2d.DeflectionAngle(v.Direction(), u.Direction())
}

//Checks if vector has any component as NaN
func (v *Vector) IsNull() bool {
	return math.IsNaN(v[x]) || math.IsNaN(v[y])
}
