package vect

import (
	. "github.com/intdxdt/simplex/util/math"
	"github.com/stretchr/testify/assert"
	"testing"
//	"fmt"
	"math"
)

const prec = 8
const eps = 1.0e-12

var A = []float64{0.88682, -1.06102}
var B = []float64{3.5, 1}
var C = []float64{-3, 1}
var D = []float64{-1.5, -3}

//Test Init Vector
func TestInitVector(t *testing.T) {
	v := New(map[string]interface{}{})
	vals := [][]float64{v.A(), v.B(), v.V()}
	// assert equality
	for _, v := range vals {
		assert.InDeltaSlice(t, v, []float64{0, 0}, eps)
	}
	mdatbt := []float64{v.m, v.d, v.at, v.bt}
	for _, v := range mdatbt {
		assert.Equal(t, v, 0.0, "should be equal")
	}
}

//Test Neg
func Test_Neg(t *testing.T) {
	a := []float64{10, 150, 6.5}
	e := []float64{280, 280, 12.8}
	v := New(map[string]interface{}{
		"a": a, "b": e,
		"at": a[2], "bt": e[2],
	})
	pv := v.v
	nv := Neg(v.v)
	negA := make([]float64, len(A))
	for i, v := range A {
		negA[i] = -v
	}
	assert.Equal(t, nv, Mult(-1, pv), "negation should be equal")
	assert.InDeltaSlice(t, Neg(A), negA, eps)
}

//Test Vect
func TestVect(t *testing.T) {
	a := []float64{10, 150, 6.5}
	e := []float64{280, 280, 12.8}
	i := []float64{185, 155, 8.6}

	v := New(map[string]interface{}{"a": a, "b": e})
	vo := New(map[string]interface{}{"a": a, "b": e})
	vi := New(map[string]interface{}{"a": i, "i": e})
	vk := New(map[string]interface{}{"d": Torad(53.13010235415598), "m": 5.0})

	assert.InDeltaSlice(t, vk.a, []float64{0,0}, eps)
	assert.InDeltaSlice(t, vk.b, []float64{3.0,4.0}, eps)

	assert.InDeltaSlice(t, v.a, vo.a, eps)
	assert.InDeltaSlice(t, v.A(), vo.a, eps)
	assert.InDeltaSlice(t, v.b, vo.b, eps)
	assert.InDeltaSlice(t, v.B(), vo.b, eps)

	assert.InDelta(t, v.m, vo.m, eps)
	assert.InDelta(t, v.M(), vo.m, eps)
	assert.InDelta(t, v.d, vo.d, eps)
	assert.InDelta(t, v.D(), vo.d, eps)

	assert.InDeltaSlice(t, v.a, a[0:2], eps)
	assert.InDeltaSlice(t, v.b, e[0:2], eps)
	assert.InDeltaSlice(t, vi.a, i[0:2], eps)
	assert.InDeltaSlice(t, vi.b, vi.a, eps)

	assert.Equal(t, v.at, a[2])
	assert.Equal(t, v.At(), a[2])
	assert.Equal(t, v.bt, e[2])
	assert.Equal(t, v.Bt(), e[2])
	assert.Equal(t, v.Dt(), e[2]-a[2])


	d := Sub(e, a)
	assert.Equal(t, v.m, Mag(d[0], d[1]))
}

func TestMagDist(t *testing.T) {
	a := []float64{0, 0 }
	b := []float64{3, 4 }
	c := []float64{3}

	assert.Equal(t, Mag(1, 1), math.Sqrt(2))
	assert.Equal(t,
		Round(Mag(-3, 2), 8),
		Round(3.605551275463989, 8),
	)
	assert.Equal(t, Mag(3, 4), 5.0)
	assert.Equal(t, Dist(a, b), 5.0)
	assert.Equal(t, math.IsNaN(Dist(a, c)), true)
	assert.Equal(t, math.IsNaN(Dist(c, a)), true)

	assert.Equal(t, Mag2(3, 4), 25.0)
	assert.Equal(t, Dist2(a, b), 25.0)
	assert.Equal(t, math.IsNaN(Dist2(c, a)), true)
	assert.Equal(t, math.IsNaN(Dist2(a, c)), true)

	assert.Equal(t, Mag(4.587, 0.), 4.587)
}

func TestDir(t *testing.T) {
	v := New(map[string]interface{}{
		"a": []float64{0, 0},
		"b": []float64{-1, 0},
	})
	assert.Equal(t, Dir(1, 1), 0.7853981633974483)
	assert.Equal(t, Dir(-1, 0), math.Pi)
	assert.Equal(t, Dir(v.v[0], v.v[1]), math.Pi)
	assert.Equal(t, Dir(1, math.Sqrt(3)), Torad(60))
	assert.Equal(t, Dir(0, -1), Torad(270))
}

func TestRevdir(t *testing.T) {
	v := New(map[string]interface{}{
		"a": []float64{0, 0},
		"b": []float64{-1, 0},
	})
	assert.Equal(t, Revdir(v.d), 0.0)
	assert.Equal(t, Revdir(0.7853981633974483), 0.7853981633974483+math.Pi)
	assert.Equal(t, Revdir(0.7853981633974483+math.Pi), 0.7853981633974483)
}

func TestDeflection(t *testing.T) {
	ln0 := [][]float64{{0, 0}, {20, 30}}
	ln1 := [][]float64{{20, 30}, {40, 15}}

	v0 := New(map[string]interface{}{"a": ln0[0], "b": ln0[1]})
	v1 := New(map[string]interface{}{"a": ln1[0], "b": ln1[1]})

	assert.Equal(t, Round(Deflect(v0.d, v1.d), 10), Round(Torad(93.17983011986422), 10))
	assert.Equal(t, Round(Deflect(v0.d, v0.d), 10), Torad(0.0))

	ln1 = [][]float64{{20, 30}, {20, 60}}
	v1 = New(map[string]interface{}{"a": ln1[0], "b": ln1[1]})

	assert.Equal(t,
		Round(Deflect(v0.d, v1.d), 10),
		Round(Torad(-33.690067525979806), 10),
	)
}

func TestAngleAtPnt(t *testing.T) {
	a := []float64{-1.28, 0.74}
	b := []float64{1.9, 4.2}
	c := []float64{3.16, -0.84}
	v := New(map[string]interface{}{"a": b, "b": c})
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
	assert.InDeltaSlice(t, negDplusB, []float64{5.0, 4.0}, eps)
}

func TestSub(t *testing.T) {
	CminusD := Sub(C, D)
	assert.InDeltaSlice(t, CminusD, []float64{-1.5, 4}, eps)
	CminusD = Sub(C, D)
	assert.InDeltaSlice(t, CminusD, []float64{-1.5, 4}, eps)
}

func TestUnit(t *testing.T) {

	v := []float64{-3, 2, 4}
	unit_v := Unit(v)
	for i, v := range unit_v {
		unit_v[i] = Round(v, 6)
	}
	assert.InDeltaSlice(t,
		[]float64{-0.83205, 0.5547}, unit_v, eps,
	)
}

//dot perform dot in 2d even when 3d coords are passed
func TestDot(t *testing.T) {
	dot_prod := Dot([]float64{1.2, -4.2, 3.5}, []float64{1.2, -4.2, 3.5});
	assert.Equal(t, 19.08, Round(dot_prod, 8));
}

func TestProj(t *testing.T) {
	assert.Equal(t, Round(Proj(A, B), 5), 0.56121);
}


