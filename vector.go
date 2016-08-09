package vect

import (
	"math"
	. "simplex/geom"
	. "simplex/util/math"
	"simplex/cart2d"
)
//Vector Type
type Vector [2]float64

//New Vector given start and end point
func NewVector(a, b *Point) *Vector {
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
	return &Vector{v[x] , v[y]}
}

//X gets the x compoent of vector
func (self *Vector) X() float64 {
	return self[x]
}

//Y gets the y component of vector
func (self *Vector) Y() float64 {
	return self[y]
}

//Add creates a new point by adding to other point
func (v *Vector) Add(o *Vector) *Vector {
	return &Vector{v[x] + o[x], v[y] + o[y]}
}

//Is a zero vector
func (self *Vector) IsZero() bool {
	return FloatEqual(self[x], 0.0) && FloatEqual(self[y], 0.0)
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
	return v.KProduct( -1.0)
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
//Returns a positive value, if AB-->BC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func (ab *Vector) SideOf(ac *Vector) float64 {
	return cart2d.CCWVector(ab, ac)
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
func (self *Vector) IsNull() bool {
	return math.IsNaN(self[x]) || math.IsNaN(self[y])
}



