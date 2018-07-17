package vect

import (
	"github.com/intdxdt/side"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
)

const (
	x = iota
	y
	z
)

//Vector Type
type Vector geom.Point

//vector type
type Vect struct {
	A geom.Point
	B geom.Point
	V Vector
}

//New Vector given start and end point
func NewVector(a, b geom.Point) Vector {
	return Vector{b[x] - a[x], b[y] - a[y]}
}

//Creates A new vector with component x and y
func NewVectorXY(x, y float64) Vector {
	return Vector{x, y}
}

//New create A new Vector
func NewVect(a, b geom.Point) *Vect {
	return &Vect{A: a, B: b, V: NewVector(a, b)}
}

//M gets magnitude of Vector
func (v *Vect) Magnitude() float64 {
	return v.V.Magnitude()
}

//Computes the Direction of Vector
func (v *Vect) Direction() float64 {
	return v.V.Direction()
}

//Reversed direction of vector direction
func (v *Vect) ReverseDirection() float64 {
	return v.V.ReverseDirection()
}

//Computes the deflection angle from vector V to u
func (v *Vect) DeflectionAngle(u *Vect) float64 {
	return v.V.DeflectionAngle(u.V)
}

//At gets  time at begin point :number
func (v *Vect) At() float64 {
	return v.A[geom.Z]
}

//Bt gets Time at end point
func (v *Vect) Bt() float64 {
	return v.B[geom.Z]
}

//Dt computs the change in time
func (v *Vect) Dt() float64 {
	return math.Abs(v.Bt() - v.At())
}

//SideOfPt computes the relation of A point to A vector
func (v *Vect) SideOf(pnt *geom.Point) *side.Side {
	var ccw = pnt.Orientation2D(&v.A, &v.B)
	var s = side.NewSide().AsLeft()
	if math.FloatEqual(ccw, 0) {
		s.AsOn()
	} else if ccw > 0 {
		s.AsRight()
	}
	return s
}

//SEDvect computes the Synchronized Euclidean Distance - Vector
func (v *Vect) SEDVector(pnt geom.Point, t float64) *Vect {
	var m = (v.Magnitude() / v.Dt()) * (t - v.At())
	//var vb = v.ExtendVect(m, 0.0, false)
	var cx, cy = geom.Extend(
		v.V[x], v.V[y], m, 0, false,
	)
	cx, cy = v.A.Add(cx, cy)
	return NewVect(geom.Point{cx, cy}, pnt)
}

//Extvect extends vector from the from end or from begin of vector
func (v *Vect) ExtendVect(magnitude, angle float64, fromEnd bool) *Vect {
	var cx, cy = geom.Extend(v.V[x], v.V[y], magnitude, angle, fromEnd)
	var cv = NewVectorXY(cx, cy)
	var a = v.A
	if fromEnd {
		a = v.B
	}
	cx, cy = a.Add(cv[x], cv[y])
	return &Vect{A: a, B: geom.Point{cx, cy} , V: cv}
}

//Deflect_vector computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func (v *Vect) DeflectVector(magnitude, deflAngle float64, fromEnd bool) *Vect {
	var cx, cy = geom.Deflect(v.V[x], v.V[y], magnitude, deflAngle, fromEnd)
	var cv = NewVectorXY(cx, cy)
	var a = v.A
	if fromEnd {
		a = v.B
	}
	cx, cy = a.Add(cv[x], cv[y])
	return &Vect{
		A: a,
		B: geom.Point{cx, cy},
		V: cv,
	}
}

//Dist2Pt computes distance from A point to Vect
func (v *Vect) DistanceToPoint(pnt *geom.Point) float64 {
	return geom.DistanceToPoint(&v.A, &v.B, pnt)
}

//Project vector u on V
func (u *Vect) Project(onv *Vect) float64 {
	var a, b = geom.Point(u.V), geom.Point(onv.V)
	return geom.Project(&a, &b)
}


