package vect

import (
	. "simplex/util/math"
	. "github.com/franela/goblin"
	"testing"
	"math"
	"simplex/geom"
)

func TestVector(t *testing.T) {
	g := Goblin(t)
	g.Describe("Point -  Unit Vector", func() {
		g.It("should test unit vector", func() {
			v0 := &Vector{0, 0 }
			v1 := &Vector{-3, math.NaN()}
			v := &Vector{-3, 2}
			unit_v := v.UnitVector()
			for i, v := range *unit_v {
				(*unit_v)[i] = Round(v, 6)
			}
			g.Assert(unit_v).Equal(&Vector{-0.83205, 0.5547})
			g.Assert(v.IsZero()).IsFalse()
			g.Assert(v0.IsZero()).IsTrue()
			g.Assert(v0.IsNull()).IsFalse()
			g.Assert(v1.IsNull()).IsTrue()

			v3 := NewVector(geom.NewPointXY(0, 0), geom.NewPointXY(3, 4))
			g.Assert(v3.Magnitude()).Equal(5.0)
			v4 := v3.Sub(NewVectorXY(2, 3))
			g.Assert(v4.SquareMagnitude()).Equal(2.0)
			g.Assert(v4.Magnitude()).Equal(math.Sqrt2)

		})

		g.Describe("ccw turn", func() {
			g.It("turn ccw", func() {
				a := geom.NewPointXY(237, 289)
				b := geom.NewPointXY(354.47839239412275, 333.38072601555746)
				c := geom.NewPointXY(462, 374)

				d := geom.NewPointXY(297.13043478260863, 339.30434782608694)
				e := geom.NewPointXY(445.8260869565217, 350.17391304347825)

				ab := NewVector(a, b)
				ac := NewVector(a, c)
				ad := NewVector(a, d)
				ae := NewVector(a, e)

				g.Assert(FloatEqual(ab.SideOf(ac), 0)).IsTrue()
				g.Assert(ab.SideOf(ad) > 0).IsTrue()
				g.Assert(ab.SideOf(ae) < 0).IsTrue()
			})
		})

		g.Describe("Vector - unit & Project", func() {
			var u = &Vector{0.88682, -1.06102}
			var v = &Vector{3.5, 1.0}
			g.It("should test projection", func() {
				g.Assert(Round(u.Project(v), 5)).Equal(0.56121)
			})
			g.It("should test Unit", func() {
				Z := &Vector{0., 0.}
				zv := Z.UnitVector()
				g.Assert(FloatEqual(zv.X(), 0)).IsTrue()
				g.Assert(FloatEqual(zv.Y(), 0)).IsTrue()
			})
		})
	})

	g.Describe("Point - Vector Dot Product", func() {
		g.It("should test dot product", func() {
			dot_prod := NewVectorXY(1.2, -4.2).DotProduct(NewVectorXY(1.2, -4.2))
			g.Assert(19.08).Equal(Round(dot_prod, 8))
		})
	})

}


