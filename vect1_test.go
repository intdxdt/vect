package vect

import (
    . "github.com/intdxdt/simplex/util/math"
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/franela/goblin"
    "testing"
    "math"
)

const prec = 8

var A = &Point{0.88682, -1.06102}
var B = &Point{3.5, 1.0}
var C = &Point{-3, 1.0}



//Test Init Vector
func TestInitVector(t *testing.T) {
    g := Goblin(t)
    g.Describe("Test Init Vector", func() {
        v := NewVect(&Options{})
        vals := []Point{v.A(), v.B(), v.V()}
        // assert equality
        g.It("should test zero vector", func() {
            for _, o := range vals {
                g.Assert(o).Eql(Point{0, 0})
            }
            mdatbt := []float64{v.m, v.d, v.at, v.bt}
            for _, o := range mdatbt {
                g.Assert(o).Equal(0.0)
            }
        })

    })

}

//Test Neg
func Test_Neg(t *testing.T) {
    g := Goblin(t)
    g.Describe("Negate Vector", func() {
        g.It("should test vector negation", func() {
            a := []float64{10, 150, 6.5}
            e := []float64{280, 280, 12.8}
            opts := &Options{
                A:NewPoint(a[:]),
                B:NewPoint(e[:]),
                At:  &a[2],
                Bt : &e[2],
            }
            v := NewVect(opts)
            pv := v.v
            nv := v.v.Neg()
            negA := &Point{0, 0}
            for i, v := range A {
                negA[i] = -v
            }
            g.Assert(nv).Eql(pv.KProduct(-1))
            g.Assert(A.Neg()).Eql(negA)

            //test immutability
            va := v.A()
            g.Assert(va).Eql(Point{a[x], a[y]})
            va[x], va[y] = 31, 33
            //should not affect vector
            g.Assert(v.A()).Eql(Point{a[x], a[y]})

            ve := v.B()
            g.Assert(ve).Eql(Point{e[x], e[y]})
            ve[x], ve[y] = 31, 33
            //should not affect vector
            g.Assert(v.B()).Eql(Point{e[x], e[y]})
        })
    })

}

//Test Vect
func TestVect(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Construct", func() {
        g.It("should test vector constructor", func() {
            a := []float64{10, 150, 6.5}
            e := []float64{280, 280, 12.8}
            i := []float64{185, 155, 8.6}

            v := NewVect(&Options{
                A:NewPoint(a[:]),
                B:NewPoint(e[:]),
                At:&a[z],
                Bt:&e[z],
            })
            vo := NewVect(&Options{
                A:NewPoint(a[:]),
                B:NewPoint(e[:]),
                At : &a[z],
                Bt : &e[z],
            })
            vi := NewVect(&Options{
                A:NewPoint(i[:]),
                B:NewPoint(e[:]),
                At : &a[z],
                Bt : &e[z],
            })
            m := 5.0
            dir := Deg2rad(53.13010235415598)
            vk := NewVect(&Options{D : &dir, M : &m, })

            g.Assert(vk.a).Eql(&Point{0, 0})
            g.Assert(Round(vk.b[x], 8)).Eql(3.0)
            g.Assert(Round(vk.b[y], 8)).Eql(4.0)

            g.Assert(v.a).Eql(vo.a)
            g.Assert(v.A()).Eql(*vo.a)
            g.Assert(v.b).Eql(vo.b)
            g.Assert(v.B()).Eql(*vo.b)

            g.Assert(v.m).Equal(vo.m)
            g.Assert(v.M()).Equal(vo.m)
            g.Assert(v.d).Equal(vo.d)
            g.Assert(v.D()).Equal(vo.d)

            g.Assert(v.a).Eql(&Point{a[0], a[1]})
            g.Assert(v.b).Eql(&Point{e[0], e[1]})
            g.Assert(vi.a).Eql(&Point{i[0], i[1]})
            g.Assert(vi.b).Eql(v.b)

            g.Assert(v.at).Equal(a[2])
            g.Assert(v.At()).Equal(a[2])
            g.Assert(v.bt).Equal(e[2])
            g.Assert(v.Bt()).Equal(e[2])
            g.Assert(v.Dt()).Equal(e[2] - a[2])

            _a := &Point{a[0], a[1]}
            _e := &Point{e[0], e[1]}
            d := _e.Sub(_a)
            g.Assert(v.m).Equal(d.Magnitude())
        })
    })

}

func TestDirection(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Direction", func() {
        g.It("should test vector direction", func() {
            v := NewVect(&Options{
                A: &Point{0, 0},
                B: &Point{-1, 0},
            })
            g.Assert(Direction(1, 1)).Equal(0.7853981633974483)
            g.Assert(Direction(-1, 0)).Equal(math.Pi)
            g.Assert(Direction(v.v[0], v.v[1])).Equal(math.Pi)
            g.Assert(Direction(1, math.Sqrt(3))).Equal(Deg2rad(60))
            g.Assert(Direction(0, -1)).Equal(Deg2rad(270))
        })
    })

}

func TestReverseDirection(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector RevDirection", func() {
        g.It("should test reverse vector direction", func() {
            v := NewVect(&Options{
                A: &Point{0, 0},
                B: &Point{-1, 0},
            })
            g.Assert(ReverseDirection(v.d)).Equal(0.0)
            g.Assert(ReverseDirection(0.7853981633974483)).Equal(0.7853981633974483 + math.Pi)
            g.Assert(ReverseDirection(0.7853981633974483 + math.Pi)).Equal(0.7853981633974483)
        })
    })

}

func TestDeflection(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Deflection", func() {
        g.It("should test reverse vector direction", func() {
            ln0 := []*Point{{0, 0}, {20, 30}}
            ln1 := []*Point{{20, 30}, {40, 15}}

            v0 := NewVect(&Options{A: ln0[0], B: ln0[1]})
            v1 := NewVect(&Options{A: ln1[0], B: ln1[1]})

            g.Assert(Round(DeflectionAngle(v0.d, v1.d), 10)).Equal(Round(Deg2rad(93.17983011986422), 10))
            g.Assert(Round(DeflectionAngle(v0.d, v0.d), 10)).Equal(Deg2rad(0.0))

            ln1 = []*Point{{20, 30}, {20, 60}}
            v1 = NewVect(&Options{A: ln1[0], B: ln1[1]})

            g.Assert(Round(DeflectionAngle(v0.d, v1.d), 10)).Equal(
                Round(Deg2rad(-33.690067525979806), 10),
            )
        })
    })

}

func TestAngleAtPnt(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Angle at Point", func() {
        g.It("should test angle formed at point by vector", func() {

            a := &Point{-1.28, 0.74}
            b := &Point{1.9, 4.2}
            c := &Point{3.16, -0.84}

            v := NewVect(&Options{A: b, B: c})
            g.Assert(Round(AngleAtPoint(a, b, c), 8)).Equal(Round(1.1694239325184717, 8), )
            g.Assert(Round(AngleAtPoint(a, v.a, v.b), 8)).Equal(Round(1.1694239325184717, 8), )
            g.Assert(Round(AngleAtPoint(b, a, c), 8)).Equal(Round(0.9882331199311394, 8), )
        })
    })

}

func TestUnit(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Unit", func() {
        g.It("should test unit vector", func() {
            v := &Point{-3, 2}
            unit_v := Unit(v)
            for i, v := range *unit_v {
                (*unit_v)[i] = Round(v, 6)
            }
            g.Assert(unit_v).Equal(&Point{-0.83205, 0.5547})
        })
    })

}

//dot perform dot in 2d even when 3d coords are passed
func TestDot(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Dot Product", func() {
        g.It("should test dot product", func() {
            dot_prod := Dot(&Point{1.2, -4.2}, &Point{1.2, -4.2})
            g.Assert(19.08).Equal(Round(dot_prod, 8))
        })
    })

}

func TestProj(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Project", func() {
        g.It("should test projection", func() {
            g.Assert(Round(Project(A, B), 5)).Equal(0.56121)
        })
    })
}


