package vect

import (
	. "github.com/intdxdt/simplex/util/math"
	"math"
)

const π = math.Pi
const τ = 2 * π
const dim2D = 2
const dim3D = 3
const (
	x = iota
	y
	z
)

//vector type
type vector struct {
	a  Vect2D
	b  Vect2D
	m  float64
	d  float64
	at float64
	bt float64
	v  Vect2D
}

//Get begin point [x, y]
func (v *vector) A() Vect2D { return v.a }

//Get end point [x, y]
func (v *vector) B() Vect2D { return v.b }

//Get component vector
func (v *vector) V() Vect2D { return v.v }

//Get Magnitude of Vector
func (v *vector) M() float64 { return v.m }

//Get Direction of Vector
func (v *vector) D() float64 { return v.d }

//Get Time at begin point :number
func (v *vector) At() float64 { return v.at }

//Get Time at end point
func (v *vector) Bt() float64 { return v.bt }

//Change in time
func (v *vector) Dt() float64 {
	return math.Abs(v.bt - v.at)
}

//Side of pnt to vector
func (v *vector) SideOfPt(pnt Vect2D) string {
	//	p := pnt[:dim2D]
	ax, ay := v.a[x], v.a[y]
	bx, by := v.b[x], v.b[y]
	cx, cy := pnt[x], pnt[y]

	mat := Mat2D{
		{ax - cx, ay - cy},
		{bx - cx, by - cy},
	}
	if Sign(Det2(mat)) > 0 {
		return "left"
	}
	return "right"
}

//Synchronized Euclidean Distance - Vector
func (v *vector) SEDvect(pnt Vect2D, t float64) vector {
	//	a := pnt[:dim2D]
	m := (v.m / v.Dt()) * (t - v.at)
	vb := v.Extvect(m, 0.0, false)
	return New(map[string]interface{}{
		"a" : pnt, "b" : vb.b,
	})
}

//Extend vector from the end = true or from the begin point
func (v *vector)  Extvect(mag, angl float64, end bool) vector {
	//from a of v back direction innitiates as fwd v direction anticlockwise
	backdir := v.d
	a 		:= v.a
	if end {
		if v.d >= π {
			backdir = v.d - π
		}else {
			backdir = v.d + π
		}
		a = v.b
	}
	fwddir := backdir + angl
	if fwddir > τ {
		fwddir = fwddir - τ
	}
	return New(map[string]interface{}{
		"a" : a, "m": mag, "d" : fwddir,
	})
}

//Deflect vector
func (v *vector) Deflectvect(mag, deflangl float64, end bool) vector {
	angl := π - deflangl
	return v.Extvect(mag, angl, end)
}

//Distance to Vect
// Minimum distance to vector from a point
// compute the minimum distance between point and vector
// if points outside the range of the vector the minimum distance
// is not perperndicular to the vector
// Ref: http://www.mappinghacks.com/code/PolyLineReduction/
func (v *vector) DistToPt(pnt Vect2D) float64 {
	precision := 12
	u := New(map[string]interface{}{"a":v.a, "b":pnt})
	dist_uv := Proj(u.v, v.v)

	rstate := false
	result := 0.0

	if dist_uv < 0 {  // if negative
		result = u.m
		rstate = true
	}else {
		negv := Neg(v.v)
		negv_pnt := Add(negv, u.v)
		if Proj(negv_pnt, negv) < 0.0 {
			result = Mag(negv_pnt[x], negv_pnt[y])
			rstate = true
		}
	}

	if rstate == false {
		// avoid floating point imprecision
		h := Round(math.Abs(u.m), precision)
		a := Round(math.Abs(dist_uv), precision)

		if FloatEq(h, 0.0) && FloatEq(a, 0.0) {
			result = 0.0
		}else {
			r := Round(a / h, precision)
			// to avoid numeric overflow
			result = h * math.Sqrt(1 - r * r)
		}
	}
	// opposite distance to hypotenus
	return result
}


//Vector constructor
func New(opts map[string]interface{}) vector {
	var a, b, v Vect2D
	var slice_a []float64 = init2d(opts, "a")
	var slice_b []float64 = init2d(opts, "b")
	var m float64 = initval(opts, "m")
	var d float64 = initval(opts, "d")
	var at float64 = initval(opts, "at")
	var bt float64 = initval(opts, "bt")

	if !is_empty(slice_a) {
		a[x], a[y] = slice_a[x], slice_a[y]
		if len(slice_a) > 2 && at == 0 {
			at = slice_a[z]
		}
	}



	if !is_empty(slice_b) {
		b[x], b[y] = slice_b[x], slice_b[y]
		if len(slice_b) > 2 && bt == 0 {
			bt = slice_b[z]
		}
		//if not empty slice b , compute v
		v = Sub(b, a)
	}
	if is_zero(v) && (m != 0) && (d != 0) {
		v = Comp(m, d)
	}
	//d direction
	if !is_zero(v) && d == 0 {
		d = Dir(v[x], v[y])
	}

	//m magnitude
	if !is_zero(v) && m == 0 {
		m = Mag(v[x], v[y])
	}

	//compute b
	if !is_zero(v) && is_zero(b) {
		b = Add(a, v)
	}

	//b is still empty
	if is_zero(b) {
		b[x], b[y] = a[x], a[y]
		m, d = 0.0, 0.0
	}
	return vector{a: a, b: b, m: m, d: d, at: at, bt: bt, v: v}
}

//is slice empty
func is_empty(s []float64) bool {
	return len(s) == 0
}

//Is point empty.
func is_zero(v Vect2D) bool {
	b := true
	for _, v := range v {
		b = b && (v == 0.0)
	}
	return b
}


//Initialize 2d point from user input options.
func init2d(opts map[string]interface{}, attr string) []float64 {
	var pt = make([]float64, dim3D)
	if _, ok := opts[attr]; ok {
		if sa, _ok := opts[attr].(Vect2D); _ok {
			pt[x], pt[y] = sa[x], sa[y]
		}else if sa, _ok := opts[attr].(Vect3D); _ok {
			pt[x], pt[y], pt[z] = sa[x], sa[y], sa[z]
		} else if sa, _ok := opts[attr].([]float64); _ok {
			pt[x], pt[y] = sa[x], sa[y]
			if len(sa) > 2 {
				pt[z] = sa[z]
			}
		}
	} else {
			pt = make([]float64, 0)
	}

	return pt
}

//Initlialize values as numbers
func initval(opts map[string]interface{}, attr string) float64 {
	var v float64

	if _, ok := opts[attr]; ok {
		if _a, _ok := opts[attr].(float64); _ok {
			v = _a
		}else if _a, _ok := opts[attr].(int); _ok {
			v = float64(_a)
		}
	}
	return v
}

//Direction in radians - counter clockwise from x-axis.
func Dir(x, y float64) float64 {
	d := math.Atan2(y, x)
	if d < 0 {
		d += 2 * math.Pi
	}
	return d
}

//Reverse direction
func Revdir(d float64) float64 {
	if d < π {
		return d + π
	}
	return d - π
}

//Unit vector
func Unit(p Vect2D) Vect2D {
	m := Mag(p[x], p[y])
	res := Vect2D{0, 0}
	for i, v := range p {
		res[i] = v / m
	}
	return res
}
//Project vector u on v
func Proj(u, onv Vect2D) float64 {
	return Dot(u, Unit(onv))
}

func AngleAtPt(atp1, p2, p3 Vect2D) float64 {
	da := Dist(atp1, p2)
	db := Dist(atp1, p3)
	dc := Dist(p2, p3)
	// keep product units small to avoid overflow
	return math.Acos(((da / db) * 0.5) + ((db / da) * 0.5) - ((dc / db) * (dc / da) * 0.5))
}


func Deflect(b0, b1 float64) float64 {
	a := b1 - Revdir(b0)
	if a < 0.0 {
		a = a + τ
	}
	return π - a
}

//Vector magnitude
func Mag(dx, dy float64) float64 {
	return math.Hypot(dx, dy)
}

//direct squred distance given dx, dy , may overflow or underflow
func Mag2(dx, dy float64) float64 {
	return dx * dx + dy * dy
}

//Distance between two points , uses internal math.Hypot for overflow and underflow
func Dist(a, b Vect2D) float64 {
	return Mag(a[x] - b[x], a[y] - b[y])
}

//Distance squared , direct squred distance between two points , may overflow or underflow
func Dist2(a, b Vect2D) float64 {
	if len(a) < dim2D || len(b) < dim2D {
		return math.NaN()
	}
	return Mag2(a[x] - b[x], a[y] - b[y])
}

//Component vector
func Comp(m, d float64) Vect2D {
	return Vect2D{m * math.Cos(d), m * math.Sin(d)}
}

//vector dot product
func Dot(va, vb Vect2D) float64 {
	sum := 0.0
	for i, _ := range va {
		sum += va[i] * vb[i]
	}
	return sum
}

//negate point
func Neg(v Vect2D) Vect2D {
	var nv Vect2D
	for i, _ := range v {
		nv[i] = -v[i]
	}
	return nv
}

//Multiply k by point
func Mult(k float64, v Vect2D) Vect2D {
	var va Vect2D
	for i, v := range v {
		va[i] = k * v
	}
	return va
}

//Subtract two points{Vect2D}.
func Sub(va, vb Vect2D) Vect2D {
	return Vect2D{va[x] - vb[x], va[y] - vb[y]}
}

//Add two points{Vect2D}.
func Add(a, b Vect2D) Vect2D {
	return Vect2D{a[x] + b[x], a[y] + b[y]}
}
