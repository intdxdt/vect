package vect

import (
	"testing"
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
)

const prec = 8

var A = &geom.Point{0.88682, -1.06102}
var B = &geom.Point{3.5, 1.0}
var C = &geom.Point{-3, 1.0}



//Test Init Vector
func TestInitVector(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test Init Vector", func() {
		v := NewVect(&Options{})
		g.It("should test zero vector", func() {
			g.Assert(v.A()).Eql(&geom.Point{0, 0})
			g.Assert(v.B()).Eql(&geom.Point{0, 0})
			g.Assert(v.Vector()).Eql(&Vector{0, 0})
			mdatbt := []float64{v.Magnitude(), v.Direction(), v.at, v.bt}
			for _, o := range mdatbt {
				g.Assert(o).Equal(0.0)
			}
		})
		g.It("should test compoent", func() {
			v := NewVect(&Options{V:geom.NewPointXY(3, 4)})
			g.Assert(v.A()).Eql(&geom.Point{0, 0})
			g.Assert(v.B()).Eql(&geom.Point{3, 4})
			g.Assert(v.Vector()).Eql(&Vector{3, 4})
			g.Assert(v.Magnitude()).Equal(5.0)
		})

	})

}

//Test Neg
func Test_Neg(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Negate Vector", func() {
		g.It("should test vector negation", func() {
			a := []float64{10, 150, 6.5}
			e := []float64{280, 280, 12.8}
			opts := &Options{
				A:geom.NewPoint(a[:]),
				B:geom.NewPoint(e[:]),
				At:  &a[2],
				Bt : &e[2],
			}
			v := NewVect(opts)
			pv := v.v
			nv := v.v.Neg()
			negA := &geom.Point{0, 0}
			for i, v := range A {
				negA[i] = -v
			}
			g.Assert(nv).Eql(pv.KProduct(-1))
			g.Assert(A.Neg()).Eql(negA)

			//test immutability
			va := v.A()
			g.Assert(va).Eql(&geom.Point{a[x], a[y]})
			va[x], va[y] = 31, 33
			//should not affect vector
			g.Assert(v.A()).Eql(&geom.Point{a[x], a[y]})

			ve := v.B()
			g.Assert(ve).Eql(&geom.Point{e[x], e[y]})
			ve[x], ve[y] = 31, 33
			//should not affect vector
			g.Assert(v.B()).Eql(&geom.Point{e[x], e[y]})
		})
	})

}

//Test Vect
func TestVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Construct", func() {
		g.It("should test vector constructor", func() {
			a := []float64{10, 150, 6.5}
			e := []float64{280, 280, 12.8}
			i := []float64{185, 155, 8.6}

			v := NewVect(&Options{
				A:geom.NewPoint(a[:]),
				B:geom.NewPoint(e[:]),
				At:&a[z],
				Bt:&e[z],
			})
			vo := NewVect(&Options{
				A:geom.NewPoint(a[:]),
				B:geom.NewPoint(e[:]),
				At : &a[z],
				Bt : &e[z],
			})
			vi := NewVect(&Options{
				A:geom.NewPoint(i[:]),
				B:geom.NewPoint(e[:]),
				At : &a[z],
				Bt : &e[z],
			})
			m := 5.0
			dir := math.Deg2rad(53.13010235415598)
			vk := NewVect(&Options{D : &dir, M : &m, })

			g.Assert(vk.a).Eql(&geom.Point{0, 0})
			g.Assert(math.Round(vk.b[x], 8)).Eql(3.0)
			g.Assert(math.Round(vk.b[y], 8)).Eql(4.0)

			g.Assert(v.a).Eql(vo.a)
			g.Assert(v.A()).Eql(vo.a)
			g.Assert(v.b).Eql(vo.b)
			g.Assert(v.B()).Eql(vo.b)

			g.Assert(v.Magnitude()).Equal(vo.Magnitude())
			g.Assert(v.Magnitude()).Equal(vo.Magnitude())
			g.Assert(v.Direction()).Equal(vo.Direction())
			g.Assert(v.Direction()).Equal(vo.Direction())

			g.Assert(v.a) .Eql(&geom.Point{a[0], a[1]})
			g.Assert(v.b) .Eql(&geom.Point{e[0], e[1]})
			g.Assert(vi.a).Eql(&geom.Point{i[0], i[1]})
			g.Assert(vi.b).Eql(v.b)

			g.Assert(v.at).Equal(a[2])
			g.Assert(v.At()).Equal(a[2])
			g.Assert(v.bt).Equal(e[2])
			g.Assert(v.Bt()).Equal(e[2])
			g.Assert(v.Dt()).Equal(e[2] - a[2])

			_a := &geom.Point{a[0], a[1]}
			_e := &geom.Point{e[0], e[1]}
			d := _e.Magnitude(_a)
			g.Assert(v.Magnitude()).Equal(d)
		})
	})

}

func TestDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Direction", func() {
		g.It("should test vector direction", func() {
			v := NewVect(&Options{
				A: &geom.Point{0, 0},
				B: &geom.Point{-1, 0},
			})
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
			v := NewVect(&Options{
				A: &geom.Point{0, 0},
				B: &geom.Point{-1, 0},
			})
			g.Assert(v.Direction()).Equal(math.Pi)
			g.Assert(v.ReverseDirection()).Equal(0.0)
		})
	})

}

func TestDeflection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Deflection", func() {
		g.It("should test reverse vector direction", func() {
			ln0 := []*geom.Point{{0, 0}, {20, 30}}
			ln1 := []*geom.Point{{20, 30}, {40, 15}}

			v0 := NewVect(&Options{A: ln0[0], B: ln0[1]})
			v1 := NewVect(&Options{A: ln1[0], B: ln1[1]})

			g.Assert(math.Round(v0.DeflectionAngle(v1), 10)).Equal(math.Round(math.Deg2rad(93.17983011986422), 10))
			g.Assert(math.Round(v0.DeflectionAngle(v0), 10)).Equal(math.Deg2rad(0.0))

			ln1 = []*geom.Point{{20, 30}, {20, 60}}
			v1 = NewVect(&Options{A: ln1[0], B: ln1[1]})

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
			u := NewVect(&Options{A:geom.NewPointXY(0, 0),  B:A})
			v := NewVect(&Options{A:geom.NewPointXY(0, 0),  B:B})
			g.Assert(math.Round(u.Project(v), 5)).Equal(0.56121)
		})
	})
}


