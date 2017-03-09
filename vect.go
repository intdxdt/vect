package vect

import (
	umath "simplex/util/math"
	"simplex/geom"
	. "simplex/side"
	"math"
	"simplex/cart2d"
)
const precision = 12

const (
	x = iota
	y
	z
)

type Options struct {
	A  cart2d.Cart2D
	B  cart2d.Cart2D
	M  *float64
	D  *float64
	At *float64
	Bt *float64
	V  cart2d.Cart2D
}

//vector type
type Vect struct {
	a  *geom.Point
	b  *geom.Point
	at float64
	bt float64
	v  *Vector
}

//New create a new Vector
func NewVect(opts *Options) *Vect {
	a := geom.NewPointXY(0.0, 0.0)
	b := geom.NewPointXY(math.NaN(), math.NaN())
	v := NewVectorXY(math.NaN(), math.NaN())

	var m, d, at, bt = math.NaN(), math.NaN(), math.NaN(), math.NaN()

	init_point2d(opts.A, a)
	init_point2d(opts.B, b)
	init_vect2d(opts.V, v)

	init_val(opts.M, &m)
	init_val(opts.D, &d)
	init_val(opts.At, &at)
	init_val(opts.Bt, &bt)

	//if not empty slice b , compute v
	if opts.B != nil {
		v = NewVector(a, b)
	}
	//compute v , given m & d
	if v.IsNull() && !math.IsNaN(m) && !math.IsNaN(d) {
		v = NewVectorMagDir(m, d)
	}

	//compute b given v and a
	if !v.IsNull() && b.IsNull() {
		b = a.Add(v)
	}

	//b is still empty
	if b.IsNull() {
		b[x], b[y] = a[x], a[y]
		at, bt = a[z], a[z]
		v = NewVectorXY(0.0, 0.0)
	}

	return &Vect{
		a: a, b: b,
		at: at, bt: bt,
		v: v,
	}
}

//A gets begin point [x, y]
func (v *Vect) A() *geom.Point {
	return v.a.Clone()
}

//B gets end point [x, y]
func (v *Vect) B() *geom.Point {
	return v.b.Clone()
}

//V gets component vector
func (v *Vect) Vector() *Vector {
	return v.v.Clone()
}

//M gets magnitude of Vector
func (v *Vect) Magnitude() float64 {
	return v.v.Magnitude()
}

//Computes the Direction of Vector
func (v *Vect) Direction() float64 {
	return v.v.Direction()
}

//Reversed direction of vector direction
func (v *Vect)  ReverseDirection() float64 {
	return v.v.ReverseDirection()
}

//Computes the deflection angle from vector v to u
func (v *Vect)  DeflectionAngle(u *Vect) float64 {
	return v.v.DeflectionAngle(u.v)
}

//At gets  time at begin point :number
func (v *Vect) At() float64 {
	return v.at
}

//Bt gets Time at end point
func (v *Vect) Bt() float64 {
	return v.bt
}

//Dt computs the change in time
func (v *Vect) Dt() float64 {
	return math.Abs(v.bt - v.at)
}

//SideOfPt computes the relation of a point to a vector
func (v *Vect) SideOf(pnt *geom.Point) *Side {
	s:= NewSide()
	ccw := cart2d.CCW(v.a, v.b, pnt)
	if umath.FloatEqual(ccw, 0){
		s.AsOn()
	} else if ccw > 0 {
		s.AsLeft()
	} else {
		s.AsRight()
	}
	return s
}

//SEDvect computes the Synchronized Euclidean Distance - Vector
func (v *Vect) SEDVector(pnt *geom.Point, t float64) *Vect {
	m := (v.Magnitude() / v.Dt()) * (t - v.at)
	vb := v.ExtendVect(m, 0.0, false)
	opts := &Options{A:vb.b, B:pnt}
	return NewVect(opts)
}

//Extvect extends vector from the from end or from begin of vector
func (v *Vect)  ExtendVect(magnitude, angle float64, from_end bool) *Vect {
	cx, cy := cart2d.Extend(v.Vector(), magnitude, angle, from_end)
	cv := NewVectorXY(cx, cy)
	a  := v.a
	if from_end {
		a = v.b
	}
	return &Vect{a:a.Clone(), b: a.Add(cv), v:cv}
}

//Deflect_vector computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func (v *Vect) DeflectVector(magnitude, defl_angle float64, from_end bool) *Vect {
	cx, cy:= cart2d.Deflect(v.Vector(), magnitude, defl_angle, from_end)
	cv := NewVectorXY(cx, cy)
	a  := v.a
	if from_end {
		a = v.b
	}
	return &Vect{a:a.Clone(), b: a.Add(cv), v:cv}
}

//Dist2Pt computes distance from a point to Vect
func (v *Vect) DistanceToPoint(pnt *geom.Point) float64 {
	return cart2d.DistanceToPoint(v.a, v.b, pnt)
}

//Project vector u on v
func (u *Vect) Project(onv *Vect) float64 {
	return cart2d.Project(u.v, onv.v)
}

//initval - initlialize values as numbers
func init_val(a  *float64, v *float64) {
	if a != nil {
		*v = *a
	}
}

//init_point2d
func init_point2d(a cart2d.Cart2D, v *geom.Point) {
	if a != nil && !a.IsNull() {
		v[x], v[y] = a.X(), a.Y()
	}
}
//init_vect2d
func init_vect2d(a cart2d.Cart2D, v *Vector) {
	if a != nil && !a.IsNull() {
		v[x], v[y] = a.X(), a.Y()
	}
}
