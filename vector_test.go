package vect

import (
    . "simplex/util/math"
    . "github.com/franela/goblin"
    "testing"
)

func TestUnit(t *testing.T) {
    g := Goblin(t)
    g.Describe("Point -  Unit Vector", func() {
        g.It("should test unit vector", func() {
            v := &Vector{-3, 2}
            unit_v := v.UnitVector()
            for i, v := range *unit_v {
                (*unit_v)[i] = Round(v, 6)
            }
            g.Assert(unit_v).Equal(&Vector{-0.83205, 0.5547})
        })
    })

}

