package vect

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
	"testing"
)

func TestVector(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Point -  Unit Vector", func() {
		g.It("should test unit vector", func() {
			var a = NewVectorXY(237, 289)
			var b = NewVectorXY(462, 374)
			var ab = a.Add(b)
			g.Assert(ab[x]).Equal(237 + 462.)
			g.Assert(ab[y]).Equal(289 + 374.)

			var v0 = &Vector{0, 0}
			var v1 = &Vector{-3, math.NaN()}
			var v = &Vector{-3, 2}
			unitV := v.UnitVector()
			for i, v := range unitV {
				unitV[i] = math.Round(v, 6)
			}
			g.Assert(unitV).Equal(Vector{-0.83205, 0.5547})
			g.Assert(v.IsZero()).IsFalse()
			g.Assert(v0.IsZero()).IsTrue()
			g.Assert(v0.IsNull()).IsFalse()
			g.Assert(v1.IsNull()).IsTrue()

			v3 := NewVector(geom.PointXY(0, 0), geom.PointXY(3, 4))
			g.Assert(v3.Magnitude()).Equal(5.0)
			v4 := v3.Sub(NewVectorXY(2, 3))
			g.Assert(v4.SquareMagnitude()).Equal(2.0)
			g.Assert(v4.Magnitude()).Equal(math.Sqrt2)

		})

		g.Describe("Vector - unit & Project", func() {
			var u = Vector{0.88682, -1.06102}
			var v = Vector{3.5, 1.0}
			g.It("should test projection", func() {
				g.Assert(math.Round(u.Project(v), 5)).Equal(0.56121)
			})
			g.It("should test Unit", func() {
				var Z = Vector{0, 0}
				var zv = Z.UnitVector()
				g.Assert(math.FloatEqual(zv[x], 0)).IsTrue()
				g.Assert(math.FloatEqual(zv[y], 0)).IsTrue()
			})
		})
	})

	g.Describe("Point - Vector Dot Product", func() {
		g.It("should test dot product", func() {
			var dotProd = NewVectorXY(1.2, -4.2).DotProduct(
				NewVectorXY(1.2, -4.2),
			)
			g.Assert(19.08).Equal(math.Round(dotProd, 8))
		})
	})

}

func TestProject(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("Vector - Project", func() {
		g.It("should test projection", func() {
			var u = NewVectorXY(A[x], A[y])
			var v = NewVectorXY(B[x], B[y])
			g.Assert(math.Round(u.Project(v), 5)).Equal(0.56121)
		})
	})
}
