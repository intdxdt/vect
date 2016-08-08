package vect

import (
	"math"
	. "simplex/geom"
	. "simplex/util/math"
)

const ε = 1e-12

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
    return &Vector{v[x] - o[x], v[y] - o[y]}
}

//KProduct create new point by multiplying point by a scaler  k
func (v *Vector) KProduct(k float64) *Vector {
    return &Vector{k * v[x], k * v[y]}
}

//Computes vector magnitude of pt as vector: x , y as components
func (v *Vector) Magnitude() float64 {
	return math.Hypot(v[x], v[y])
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func (self *Vector)  SquareMagnitude() float64 {
    return (self[x] * self[x]) + (self[y] * self[y])
}

//Dot Product of two points as vectors
func (v *Vector) DotProduct(o *Vector) float64 {
	return (v[x] * o[x]) + (v[y] * o[y])
}

//Unit vector of point
func (v *Vector) UnitVector() *Vector {
	m := v.Magnitude()
	if FloatEqual(m, 0.0) {
		m = ε
	}
	return NewVectorXY(v[x] / m, v[y] / m)
}

//Project vector u on v
func (v *Vector) Project(u *Vector) float64 {
	return u.DotProduct(v.UnitVector())
}


//2D cross product of OA and OB vectors,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if OAB makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func (v *Vector) CCW(a, b *Vector) float64 {
	return (b[x] - a[x]) * (v[y] - a[y]) - (b[y] - a[y]) * (v[x] - a[x])
}


//is null
func (self *Vector) IsNull() bool {
    return math.IsNaN(self[x]) || math.IsNaN(self[y])
}



