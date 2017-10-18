package vect

import (
    "testing"
    "github.com/intdxdt/geom"
    "github.com/intdxdt/math"
    "github.com/franela/goblin"
)

func TestVector(t *testing.T) {
    g := goblin.Goblin(t)
    g.Describe("Point -  Unit Vector", func() {
        g.It("should test unit vector", func() {
            a := NewVectorXY(237, 289)
            b := NewVectorXY(462, 374)
            ab := a.Add(b)
            g.Assert(ab.X()).Equal(237 + 462.)
            g.Assert(ab.Y()).Equal(289 + 374.)

            v0 := &Vector{0, 0 }
            v1 := &Vector{-3, math.NaN()}
            v := &Vector{-3, 2}
            unit_v := v.UnitVector()
            for i, v := range *unit_v {
                (*unit_v)[i] = math.Round(v, 6)
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


        g.Describe("Vector - unit & Project", func() {
            var u = &Vector{0.88682, -1.06102}
            var v = &Vector{3.5, 1.0}
            g.It("should test projection", func() {
                g.Assert(math.Round(u.Project(v), 5)).Equal(0.56121)
            })
            g.It("should test Unit", func() {
                Z := &Vector{0., 0.}
                zv := Z.UnitVector()
                g.Assert(math.FloatEqual(zv.X(), 0)).IsTrue()
                g.Assert(math.FloatEqual(zv.Y(), 0)).IsTrue()
            })
        })
    })

    g.Describe("Point - Vector Dot Product", func() {
        g.It("should test dot product", func() {
            dot_prod := NewVectorXY(1.2, -4.2).DotProduct(NewVectorXY(1.2, -4.2))
            g.Assert(19.08).Equal(math.Round(dot_prod, 8))
        })
    })

}

func TestProject(t *testing.T) {
    g := goblin.Goblin(t)
    g.Describe("Vector - Project", func() {
        g.It("should test projection", func() {
            u := NewVectorXY(A.X(), A.Y())
            v := NewVectorXY(B.X(), B.Y())
            g.Assert(math.Round(u.Project(v), 5)).Equal(0.56121)
        })
    })
}