package vect

import (
    . "github.com/intdxdt/simplex/util/math"
    "math"
)

const π = Pi
const τ = Tau

const (
    x = iota
    y
    z
)
const (
    i = iota
    j
    k
)

type Side int

var Sided = struct {
    Left  Side
    Right Side
}{-1, 1}

type Options struct {
    A  *Vect2D
    B  *Vect2D
    M  *float64
    D  *float64
    At *float64
    Bt *float64
    V  *Vect2D
}
//vector type
type Vect struct {
    a  *Vect2D
    b  *Vect2D
    m  float64
    d  float64
    at float64
    bt float64
    v  *Vect2D
}

//New create a new Vector
func NewVect(opts *Options) *Vect {
    a := &Vect2D{0.0, 0.0}
    b := &Vect2D{0.0, 0.0}
    v := &Vect2D{0.0, 0.0}
    var m, d, at, bt float64

    init_vect2d(opts.A, a)
    init_vect2d(opts.B, b)
    init_vect2d(opts.V, v)

    init_val(opts.M, &m)
    init_val(opts.D, &d)
    init_val(opts.At, &at)
    init_val(opts.Bt, &bt)

    //if not empty slice b , compute v
    if opts.B != nil {
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
    return &Vect{
        a: a, b: b,
        m: m, d: d,
        at: at, bt: bt,
        v: v,
    }
}


//A gets begin point [x, y]
func (v *Vect) A() Vect2D {
    return *v.a
}

//B gets end point [x, y]
func (v *Vect) B() Vect2D {
    return *v.b
}

//V gets component vector
func (v *Vect) V() Vect2D {
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
func (v *Vect) SideOfPt(pnt *Vect2D) Side {
    ax, ay := v.a[x], v.a[y]
    bx, by := v.b[x], v.b[y]
    cx, cy := pnt[x], pnt[y]

    mat := &Mat2D{
        {ax - cx, ay - cy},
        {bx - cx, by - cy},
    }
    if Sign(Det2(mat)) > 0 {
        return Sided.Left
    }
    return Sided.Right
}

//SEDvect computes the Synchronized Euclidean Distance - Vector
func (v *Vect) SEDVector(pnt *Vect2D, t float64) *Vect {
    m := (v.m / v.Dt()) * (t - v.at)
    vb := v.Extvect(m, 0.0, false)
    opts := &Options{A:pnt, B:vb.b}
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
        bβ = v.d + π
        a = v.b
    }
    fβ := bβ + angle
    if fβ > τ {
        fβ -= τ
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
    angl := π - defl_angle
    return v.Extvect(mag, angl, from_end)
}

//Dist2Pt computes distance from a point to Vect
// Minimum distance to vector from a point
// compute the minimum distance between point and vector
// if points outside the range of the vector the minimum distance
// is not perperndicular to the vector
// Ref: http://www.mappinghacks.com/code/PolyLineReduction/
func (v *Vect) Dist2Pt(pnt *Vect2D) float64 {
    precision := 12
    opts := &Options{
        A: v.a,
        B : pnt,
    }
    u := NewVect(opts)
    dist_uv := Proj(u.v, v.v)

    rstate := false
    result := 0.0

    if dist_uv < 0 {
        // if negative
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

        if FloatEqual(h, 0.0) && FloatEqual(a, 0.0) {
            result = 0.0
        } else {
            r := Round(a / h, precision)
            // to avoid numeric overflow
            result = h * math.Sqrt(1 - r * r)
        }
    }
    // opposite distance to hypotenus
    return result
}



//Is point empty.
func is_zero(vect *Vect2D) bool {
    b := true
    v := *vect
    for _, v := range v {
        b = (b && (v == 0.0))
    }
    return b
}


//initval - initlialize values as numbers
func init_val(a  *float64, v *float64) {
    if a != nil {
        *v = *a
    }
}

//init_vect2d
func init_vect2d(a, v *Vect2D) {
    if a != nil {
        *v = *a
    } else {
        *v = Vect2D{0.0, 0.0}
    }
}

//Dir computes direction in radians - counter clockwise from x-axis.
func Dir(x, y float64) float64 {
    d := math.Atan2(y, x)
    if d < 0 {
        d += Tau
    }
    return d
}

//Revdir computes the reversed direction from a foward direction
func Revdir(d float64) float64 {
    if d < π {
        return d + π
    }
    return d - π
}

//Unit vector
func Unit(p *Vect2D) *Vect2D {
    m := Mag(p[x], p[y])
    res := &Vect2D{0, 0}
    for i, v := range p {
        res[i] = v / m
    }
    return res
}
//Project vector u on v
func Proj(u, onv *Vect2D) float64 {
    return Dot(u, Unit(onv))
}

func AngleAtPt(atp1, p2, p3 *Vect2D) float64 {
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
func Dist(va, vb *Vect2D) float64 {
    a, b := *va, *vb
    return Mag(a[x] - b[x], a[y] - b[y])
}

//Distance squared , direct squred distance between two points , may overflow or underflow
func Dist2(va, vb *Vect2D) float64 {
    a, b := *va, *vb
    return Mag2(a[x] - b[x], a[y] - b[y])
}

//Component vector
func Comp(m, d float64) *Vect2D {
    return &Vect2D{m * math.Cos(d), m * math.Sin(d)}
}

//vector dot product
func Dot(va, vb *Vect2D) float64 {
    sum := 0.0
    a, b := *va, *vb
    sum += a[i] * b[i]
    sum += a[j] * b[j]
    return sum
}

//negate point
func Neg(v *Vect2D) *Vect2D {
    var nv Vect2D
    for i, a := range *v {
        nv[i] = -a
    }
    return &nv
}

//Multiply k by point
func Mult(k float64, vect *Vect2D) *Vect2D {
    var va Vect2D
    for i, v := range vect {
        va[i] = k * v
    }
    return &va
}

//Subtract two points{Vect2D}.
func Sub(va, vb *Vect2D) *Vect2D {
    return &Vect2D{va[x] - vb[x], va[y] - vb[y]}
}

//Add two points{Vect2D}.
func Add(a, b *Vect2D) *Vect2D {
    return &Vect2D{a[x] + b[x], a[y] + b[y]}
}
