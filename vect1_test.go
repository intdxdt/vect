package vect

import (
    . "github.com/intdxdt/simplex/util/math"
    "github.com/stretchr/testify/assert"
    "testing"
    "math"
)

const prec = 8
const eps = 1.0e-12

var A = &Vect2D{0.88682, -1.06102}
var B = &Vect2D{3.5, 1}
var C = &Vect2D{-3, 1}
var D = &Vect2D{-1.5, -3}

func min_val(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func float_cmp(a, b []float64) bool {
    bl := true
    N := min_val(len(a), len(b))
    for i := 0; i < N; i++ {
        bl = bl && (a[i] == b[i])
    }
    return bl
}

func make_slice(a interface{}) []float64 {
    var val []float64
    switch v := a.(type) {
    case [2]float64:
        val = v[:]
    case [3]float64:
        val = v[:]
    case []float64:
        val = v[:]
    }
    return val
}

func cmp_array(a, b interface{}) bool {
    _a := make_slice(a)
    _b := make_slice(b)
    return float_cmp(_a, _b)
}

//Test Init Vector
func TestInitVector(t *testing.T) {
    v := NewVect(&Options{})
    vals := []Vect2D{v.A(), v.B(), v.V()}
    // assert equality
    for _, v := range vals {
        assert.Equal(t, v, Vect2D{0, 0}, "array should be equal")
    }
    mdatbt := []float64{v.m, v.d, v.at, v.bt}
    for _, v := range mdatbt {
        assert.Equal(t, v, 0.0)
    }
}

//Test Neg
func Test_Neg(t *testing.T) {
    a := []float64{10, 150, 6.5}
    e := []float64{280, 280, 12.8}
    opts := &Options{
        A:&Vect2D{a[x], a[y]},
        B:&Vect2D{e[x], e[y]},
        At:  &a[2],
        Bt : &e[2],
    }
    v := NewVect(opts)
    pv := v.v
    nv := Neg(v.v)
    negA := &Vect2D{}
    for i, v := range A {
        negA[i] = -v
    }
    assert.Equal(t, nv, Mult(-1, pv), "negation should be equal")
    assert.Equal(t, Neg(A), negA, "array should be equal")
}

//Test Vect
func TestVect(t *testing.T) {
    a := []float64{10, 150, 6.5}
    e := []float64{280, 280, 12.8}
    i := []float64{185, 155, 8.6}

    v := NewVect(&Options{
        A:&Vect2D{a[x], a[y]},
        B:&Vect2D{e[x], e[y]},
        At:&a[z],
        Bt:&e[z],
    })
    vo := NewVect(&Options{
        A  : &Vect2D{a[x], a[y]},
        B  : &Vect2D{e[x], e[y]},
        At : &a[z],
        Bt : &e[z],
    })
    vi := NewVect(&Options{
            A  : &Vect2D{i[x], i[y]},
            B  : &Vect2D{e[x], e[y]},
            At : &a[z],
            Bt : &e[z],
        })
    m := 5.0
    dir := Deg2rad(53.13010235415598)
    vk := NewVect(&Options{
            D  : &dir,
            M  : &m,
        })

    assert.True(t, cmp_array(vk.a, Vect2D{0, 0}))
    assert.True(t, cmp_array(vk.b, Vect2D{3.0, 4.0}))

    assert.True(t, cmp_array(v.a, vo.a))
    assert.True(t, cmp_array(v.A(), vo.a))
    assert.True(t, cmp_array(v.b, vo.b))
    assert.True(t, cmp_array(v.B(), vo.b))

    assert.Equal(t, v.m, vo.m)
    assert.Equal(t, v.M(), vo.m)
    assert.Equal(t, v.d, vo.d)
    assert.Equal(t, v.D(), vo.d)

    assert.True(t, cmp_array(v.a, a[0:2]))
    assert.True(t, cmp_array(v.b, e[0:2]))
    assert.True(t, cmp_array(vi.a, i[0:2]))
    assert.True(t, cmp_array(vi.b, vi.a))

    assert.Equal(t, v.at, a[2])
    assert.Equal(t, v.At(), a[2])
    assert.Equal(t, v.bt, e[2])
    assert.Equal(t, v.Bt(), e[2])
    assert.Equal(t, v.Dt(), e[2] - a[2])

    _a := &Vect2D{a[0], a[1]}
    _e := &Vect2D{e[0], e[1]}
    d := Sub(_e, _a)
    assert.Equal(t, v.m, Mag(d[0], d[1]))
}

func TestMagDist(t *testing.T) {
    a := &Vect2D{0, 0 }
    b := &Vect2D{3, 4 }

    assert.Equal(t, Mag(1, 1), math.Sqrt(2))
    assert.Equal(t,
        Round(Mag(-3, 2), 8),
        Round(3.605551275463989, 8),
    )

    assert.Equal(t, Mag(3, 4), 5.0)
    assert.Equal(t, Dist(a, b), 5.0)

    assert.Equal(t, Mag2(3, 4), 25.0)
    assert.Equal(t, Dist2(a, b), 25.0)

    assert.Equal(t, Mag(4.587, 0.), 4.587)
}

func TestDir(t *testing.T) {
    v := NewVect(&Options{
        A: &Vect2D{0, 0},
        B: &Vect2D{-1, 0},
    })
    assert.Equal(t, Dir(1, 1), 0.7853981633974483)
    assert.Equal(t, Dir(-1, 0), math.Pi)
    assert.Equal(t, Dir(v.v[0], v.v[1]), math.Pi)
    assert.Equal(t, Dir(1, math.Sqrt(3)), Deg2rad(60))
    assert.Equal(t, Dir(0, -1), Deg2rad(270))
}

func TestRevdir(t *testing.T) {
    v := NewVect(&Options{
        A: &Vect2D{0, 0},
        B: &Vect2D{-1, 0},
    })
    assert.Equal(t, Revdir(v.d), 0.0)
    assert.Equal(t, Revdir(0.7853981633974483), 0.7853981633974483 + math.Pi)
    assert.Equal(t, Revdir(0.7853981633974483 + math.Pi), 0.7853981633974483)
}

func TestDeflection(t *testing.T) {
    ln0 := []*Vect2D{{0, 0}, {20, 30}}
    ln1 := []*Vect2D{{20, 30}, {40, 15}}

    v0 := NewVect(&Options{A: ln0[0], B: ln0[1]})
    v1 := NewVect(&Options{A: ln1[0], B: ln1[1]})

    assert.Equal(t, Round(Deflect(v0.d, v1.d), 10), Round(Deg2rad(93.17983011986422), 10))
    assert.Equal(t, Round(Deflect(v0.d, v0.d), 10), Deg2rad(0.0))

    ln1 = []*Vect2D{{20, 30}, {20, 60}}
    v1 = NewVect(&Options{A: ln1[0], B: ln1[1]})

    assert.Equal(t,
        Round(Deflect(v0.d, v1.d), 10),
        Round(Deg2rad(-33.690067525979806), 10),
    )
}

func TestAngleAtPnt(t *testing.T) {
    a := &Vect2D{-1.28, 0.74}
    b := &Vect2D{1.9, 4.2}
    c := &Vect2D{3.16, -0.84}
    v := NewVect(&Options{A: b, B: c})
    assert.Equal(t,
        Round(AngleAtPt(a, b, c), 8),
        Round(1.1694239325184717, 8),
    )
    assert.Equal(t,
        Round(AngleAtPt(a, v.a, v.b), 8),
        Round(1.1694239325184717, 8),
    )
    assert.Equal(t,
        Round(AngleAtPt(b, a, c), 8),
        Round(0.9882331199311394, 8),
    )
}

func TestAdd(t *testing.T) {
    negDplusB := Add(Neg(D), B)
    assert.Equal(t, negDplusB, &Vect2D{5.0, 4.0})
}

func TestSub(t *testing.T) {
    CminusD := Sub(C, D)
    assert.Equal(t, CminusD, &Vect2D{-1.5, 4})
    CminusD = Sub(C, D)
    assert.Equal(t, CminusD, &Vect2D{-1.5, 4})
}

func TestUnit(t *testing.T) {

    v := &Vect2D{-3, 2}
    unit_v := Unit(v)
    for i, v := range *unit_v {
        (*unit_v)[i] = Round(v, 6)
    }
    assert.Equal(t,
        &Vect2D{-0.83205, 0.5547}, unit_v,
    )
}

//dot perform dot in 2d even when 3d coords are passed
func TestDot(t *testing.T) {
    dot_prod := Dot(&Vect2D{1.2, -4.2}, &Vect2D{1.2, -4.2});
    assert.Equal(t, 19.08, Round(dot_prod, 8));
}

func TestProj(t *testing.T) {
    assert.Equal(t, Round(Proj(A, B), 5), 0.56121);
}


