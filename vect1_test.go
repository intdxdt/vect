package vect

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
	"testing"
)

const prec = 8

var A = geom.Point{0.88682, -1.06102}
var B = geom.Point{3.5, 1.0}
var C = geom.Point{-3, 1.0}

//Test Init Vector
func TestInitVector(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test Init Vector", func() {
		var v = NewVect(geom.Point{}, geom.Point{})
		g.It("should test zero vector", func() {
			g.Assert(v.A).Eql(geom.Point{0, 0})
			g.Assert(v.B).Eql(geom.Point{0, 0})
			g.Assert(v.V).Eql(Vector{0, 0})
			mdatbt := []float64{v.Magnitude(), v.Direction(), v.At(), v.Bt()}
			for _, o := range mdatbt {
				g.Assert(o).Equal(0.0)
			}
		})
		g.It("should test component", func() {
			v := NewVect(geom.Point{}, geom.PointXY(3, 4))
			g.Assert(v.A).Eql(geom.Point{0, 0})
			g.Assert(v.B).Eql(geom.Point{3, 4})
			g.Assert(v.V).Eql(Vector{3, 4})
			g.Assert(v.Magnitude()).Equal(5.0)
		})

	})

}

//Test Neg
func Test_Neg(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Negate Vector", func() {
		g.It("should test vector negation", func() {
			var a = []float64{10, 150, 6.5}
			var e = []float64{280, 280, 12.8}
			var v = NewVect(geom.CreatePoint(a), geom.CreatePoint(e))
			var pv = v.V
			var nv = v.V.Neg()
			var negA = geom.Point{}
			for i, v := range A {
				negA[i] = -v
			}
			g.Assert(nv).Eql(pv.KProduct(-1))
			g.Assert(A.Neg()).Eql(negA)

			//test immutability
			var va = v.A
			g.Assert(va).Eql(geom.Point{a[x], a[y], a[z]})
			va[x], va[y] = 31, 33
			//should not affect vector
			g.Assert(v.A).Eql(geom.Point{a[x], a[y], a[z]})

			var ve = v.B
			g.Assert(ve).Eql(geom.Point{e[x], e[y], e[z]})
			ve[x], ve[y] = 31, 33
			//should not affect vector
			g.Assert(v.B).Eql(geom.Point{e[x], e[y], e[z]})
		})
	})

}

//Test Vect
func TestVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Construct", func() {
		g.It("should test vector constructor", func() {
			var a = []float64{10, 150, 6.5}
			var e = []float64{280, 280, 12.8}
			var i = []float64{185, 155, 8.6}
			var v = NewVect(geom.CreatePoint(a), geom.CreatePoint(e))
			var vo = NewVect(geom.CreatePoint(a), geom.CreatePoint(e))
			var vi = NewVect(geom.CreatePoint(i[:]), geom.CreatePoint(e[:]))
			var m = 5.0
			var dir = math.Deg2rad(53.13010235415598)
			var cx, cy = geom.Component(m, dir)
			var vk = NewVect(geom.Point{}, geom.Point{cx, cy})

			g.Assert(vk.A).Eql(geom.Point{0, 0})
			g.Assert(math.Round(vk.B[x], 8)).Eql(3.0)
			g.Assert(math.Round(vk.B[y], 8)).Eql(4.0)

			g.Assert(v.A).Eql(vo.A)
			g.Assert(v.A).Eql(vo.A)
			g.Assert(v.B).Eql(vo.B)
			g.Assert(v.B).Eql(vo.B)

			g.Assert(v.Magnitude()).Equal(vo.Magnitude())
			g.Assert(v.Magnitude()).Equal(vo.Magnitude())
			g.Assert(v.Direction()).Equal(vo.Direction())
			g.Assert(v.Direction()).Equal(vo.Direction())

			g.Assert(v.A).Eql(geom.Point{a[0], a[1], a[2]})
			g.Assert(v.B).Eql(geom.Point{e[0], e[1], e[2]})
			g.Assert(vi.A).Eql(geom.Point{i[0], i[1], i[2]})
			g.Assert(vi.B).Eql(v.B)

			g.Assert(v.At()).Equal(a[2])
			g.Assert(v.Bt()).Equal(e[2])
			g.Assert(v.Dt()).Equal(e[2] - a[2])

			var aa = geom.Point{a[0], a[1]}
			var ee = geom.Point{e[0], e[1]}
			var d = ee.Magnitude(&aa)
			g.Assert(v.Magnitude()).Equal(d)
		})
	})

}

func TestDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Direction", func() {
		g.It("should test vector direction", func() {
			var v = NewVect(geom.Point{0, 0}, geom.Point{-1, 0})
			g.Assert(v.Direction()).Equal(math.Pi)
			g.Assert(NewVectorXY(1, 1).Direction()).Equal(0.7853981633974483)
			g.Assert(NewVectorXY(-1, 0).Direction()).Equal(math.Pi)
			g.Assert(NewVectorXY(1, math.Sqrt(3)).Direction()).Equal(math.Deg2rad(60))
			g.Assert(NewVectorXY(0, -1).Direction()).Equal(math.Deg2rad(270))
		})
	})

}

func TestReverseDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector RevDirection", func() {
		g.It("should test reverse vector direction", func() {
			var v = NewVect(geom.Point{0, 0}, geom.Point{-1, 0})
			g.Assert(v.Direction()).Equal(math.Pi)
			g.Assert(v.ReverseDirection()).Equal(0.0)
		})
	})

}

func TestDeflection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Deflection", func() {
		g.It("should test reverse vector direction", func() {
			var ln0 = []geom.Point{{0, 0}, {20, 30}}
			var ln1 = []geom.Point{{20, 30}, {40, 15}}
			var v0 = NewVect(ln0[0], ln0[1])
			var v1 = NewVect(ln1[0], ln1[1])

			g.Assert(math.Round(v0.DeflectionAngle(v1), 10)).Equal(math.Round(math.Deg2rad(93.17983011986422), 10))
			g.Assert(math.Round(v0.DeflectionAngle(v0), 10)).Equal(math.Deg2rad(0.0))

			ln1 = []geom.Point{{20, 30}, {20, 60}}
			v1 = NewVect(ln1[0], ln1[1])

			g.Assert(math.Round(v0.DeflectionAngle(v1), 10)).Equal(
				math.Round(math.Deg2rad(-33.690067525979806), 10),
			)
		})
	})

}

func TestProj(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("vect - Project", func() {
		g.It("should test projection", func() {
			u := NewVect(geom.PointXY(0, 0), A)
			v := NewVect(geom.PointXY(0, 0), B)
			g.Assert(math.Round(u.Project(v), 5)).Equal(0.56121)
		})
	})
}
