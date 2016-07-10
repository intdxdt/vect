package vect

import (
    . "simplex/util/math"
    . "simplex/geom"
    . "simplex/side"
    "math"
)

const (
    x = iota
    y
    z
)

type Options struct {
    A  *Point
    B  *Point
    M  *float64
    D  *float64
    At *float64
    Bt *float64
    V  *Point
}

//vector type
type Vect struct {
    a  *Point
    b  *Point
    m  float64
    d  float64
    at float64
    bt float64
    v  *Point
}

//New create a new Vector
func NewVect(opts *Options) *Vect {
    a := NewPointXY(0.0, 0.0)
    b := NewPointXY(math.NaN(), math.NaN())
    v := NewPointXY(math.NaN(), math.NaN())

    var m, d, at, bt = math.NaN(), math.NaN(), math.NaN(), math.NaN()

    init_vect2d(opts.A, a)
    init_vect2d(opts.B, b)
    init_vect2d(opts.V, v)

    init_val(opts.M, &m)
    init_val(opts.D, &d)
    init_val(opts.At, &at)
    init_val(opts.Bt, &bt)

    //if not empty slice b , compute v
    if opts.B != nil {
        v = b.Sub(a)
    }
    //compute v , given m & d
    if v.IsNull() && !math.IsNaN(m) && !math.IsNaN(d) {
        v = Component(m, d)
    }

    //compute d given v
    if !v.IsNull() && math.IsNaN(d) {
        d = Direction(v[x], v[y])
    }

    //compute m given v
    if !v.IsNull() && math.IsNaN(m) {
        m = v.Magnitude()
    }

    //compute b given v and a
    if !v.IsNull() && b.IsNull() {
        b = a.Add(v)
    }

    //b is still empty
    if b.IsNull() {
        b[x], b[y] = a[x], a[y]
        m, d = 0.0, 0.0
        at, bt = 0.0, 0.0
        v = NewPointXY(0.0, 0.0)
    }

    return &Vect{
        a: a, b: b,
        m: m, d: d,
        at: at, bt: bt,
        v: v,
    }
}


//A gets begin point [x, y]
func (v *Vect) A() Point {
    return *v.a
}

//B gets end point [x, y]
func (v *Vect) B() Point {
    return *v.b
}

//V gets component vector
func (v *Vect) V() Point {
    return *v.v
}

//M gets magnitude of Vector
func (v *Vect) M() float64 {
    return v.m
}

//D gets Direction of Vector
func (v *Vect) D() float64 {
    return v.d
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
func (v *Vect) SideOfPt(pnt *Point) *Side {
    ax, ay := v.a[x], v.a[y]
    bx, by := v.b[x], v.b[y]
    cx, cy := pnt[x], pnt[y]

    mat := &Mat2D{
        {ax - cx, ay - cy},
        {bx - cx, by - cy},
    }
    if Sign(Det2(mat)) > 0 {
        return NewSide().AsLeft()
    }
    return NewSide().AsRight()
}

//SEDvect computes the Synchronized Euclidean Distance - Vector
func (v *Vect) SEDVector(pnt *Point, t float64) *Vect {
    m := (v.m / v.Dt()) * (t - v.at)
    vb := v.Extvect(m, 0.0, false)
    opts := &Options{A:vb.b, B:pnt}
    return NewVect(opts)
}

//Extvect extends vector from the from end (from_end is true) else from begin of vector
func (v *Vect)  Extvect(magnitude, angle float64, from_end bool) *Vect {
    //from a of v back direction innitiates as fwd v direction anticlockwise
    //bβ - back bearing
    //fβ - forward bearing
    bβ := v.d
    a := v.a
    if from_end {
        bβ = v.d + Pi
        a = v.b
    }
    fβ := bβ + angle
    if fβ > Tau {
        fβ -= Tau
    }

    opts := &Options{
        A : a,
        M : &magnitude,
        D : &fβ,
    }
    return NewVect(opts)
}

//Deflect_vector computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func (v *Vect) DeflectVector(mag, defl_angle float64, from_end bool) *Vect {
    angl := Pi - defl_angle
    return v.Extvect(mag, angl, from_end)
}

//Dist2Pt computes distance from a point to Vect
// Minimum distance to vector from a point
// compute the minimum distance between point and vector
// if points outside the range of the vector the minimum distance
// is not perperndicular to the vector
// modified @Ref: http://www.mappinghacks.com/code/PolyLineReduction/
func (v *Vect) DistanceToPoint(pnt *Point) float64 {
    precision := 12
    opts := &Options{A: v.a, B : pnt, }
    u := NewVect(opts)
    dist_uv := Project(u.v, v.v)

    rstate := false
    result := 0.0

    if dist_uv < 0 {
        // if negative
        result = u.m
        rstate = true
    } else {
        negv := v.v.Neg()
        negv_pnt := negv.Add(u.v)
        if Project(negv_pnt, negv) < 0.0 {
            result = negv_pnt.Magnitude()
            rstate = true
        }
    }

    if rstate == false {
        // avoid floating point imprecision
        h := Round(math.Abs(u.m), precision)
        a := Round(math.Abs(dist_uv), precision)

        if FloatEqual(h, 0.0) && FloatEqual(a, 0.0) {
            result = 0.0
        } else {
            r := Round(a / h, precision)
            // to avoid numeric overflow
            result = h * math.Sqrt(1 - r * r)
        }
    }
    //opposite distance to hypotenus
    return result
}


//initval - initlialize values as numbers
func init_val(a  *float64, v *float64) {
    if a != nil {
        *v = *a
    }
}

//init_vect2d
func init_vect2d(a, v *Point) {
    if a != nil && !a.IsNull() {
        v[x], v[y] = a[x], a[y]
    }
}

//Dir computes direction in radians - counter clockwise from x-axis.
func Direction(x, y float64) float64 {
    d := math.Atan2(y, x)
    if d < 0 {
        d += Tau
    }
    return d
}

//Revdir computes the reversed direction from a foward direction
func ReverseDirection(d float64) float64 {
    if d < Pi {
        return d + Pi
    }
    return d - Pi
}

//Project vector u on v
func Project(u, onv *Point) float64 {
    return u.DotProduct(onv.UnitVector())
}

func DeflectionAngle(bearing1, bearing2 float64) float64 {
    a := bearing2 - ReverseDirection(bearing1)
    if a < 0.0 {
        a = a + Tau
    }
    return Pi - a
}

//Component vector
func Component(m, d float64) *Point {
    return NewPointXY(m * math.Cos(d), m * math.Sin(d))
}
