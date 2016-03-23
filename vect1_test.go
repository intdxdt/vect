package vect

import (
    . "github.com/intdxdt/simplex/util/math"
    . "github.com/franela/goblin"
    "testing"
    "math"
)

const prec = 8
var A = &Vect2D{0.88682, -1.06102}
var B = &Vect2D{3.5, 1.0}
var C = &Vect2D{-3, 1.0}
var D = &Vect2D{-1.5, -3.0}



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
    g := Goblin(t)
    g.Describe("Test Init Vector", func() {
        v := NewVect(&Options{})
        vals := []Vect2D{v.A(), v.B(), v.V()}
        // assert equality
        g.It("should test zero vector", func() {
            for _, o := range vals {
                g.Assert(o).Eql(Vect2D{0, 0})
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
            g.Assert(nv).Eql(Mult(-1, pv))
            g.Assert(Neg(A)).Eql(negA)
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

            g.Assert(vk.a).Eql(&Vect2D{0, 0})
            g.Assert(Round(vk.b[x],8)).Eql(3.0)
            g.Assert(Round(vk.b[y],8)).Eql(4.0)

            g.Assert(v.a).Eql(vo.a)
            g.Assert(v.A()).Eql(*vo.a)
            g.Assert(v.b).Eql(vo.b)
            g.Assert(v.B()).Eql(*vo.b)

            g.Assert(v.m).Equal(vo.m)
            g.Assert(v.M()).Equal(vo.m)
            g.Assert(v.d).Equal(vo.d)
            g.Assert(v.D()).Equal(vo.d)

            g.Assert(v.a).Eql(&Vect2D{a[0], a[1]})
            g.Assert(v.b).Eql(&Vect2D{e[0], e[1]})
            g.Assert(vi.a).Eql(&Vect2D{i[0], i[1]})
            g.Assert(vi.b).Eql(v.b)

            g.Assert(v.at).Equal(a[2])
            g.Assert(v.At()).Equal(a[2])
            g.Assert(v.bt).Equal(e[2])
            g.Assert(v.Bt()).Equal(e[2])
            g.Assert(v.Dt()).Equal(e[2] - a[2])

            _a := &Vect2D{a[0], a[1]}
            _e := &Vect2D{e[0], e[1]}
            d := Sub(_e, _a)
            g.Assert(v.m).Equal(Mag(d[0], d[1]))
        })
    })

}

func TestMagDist(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Magnitude", func() {
        g.It("should test vector magnitude and distance", func() {
            a := &Vect2D{0, 0 }
            b := &Vect2D{3, 4 }

            g.Assert(Mag(1, 1)).Equal(math.Sqrt2)
            g.Assert(Round(Mag(-3, 2), 8)).Equal(
                Round(3.605551275463989, 8),
            )

            g.Assert(Mag(3, 4)).Equal(5.0)
            g.Assert(Dist(a, b)).Equal(5.0)

            g.Assert(Mag2(3, 4)).Equal(25.0)
            g.Assert(Dist2(a, b)).Equal(25.0)

            g.Assert(Mag(4.587, 0.)).Equal(4.587)
        })
    })

}

func TestDir(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Direction", func() {
        g.It("should test vector direction", func() {
            v := NewVect(&Options{
                A: &Vect2D{0, 0},
                B: &Vect2D{-1, 0},
            })
            g.Assert(Dir(1, 1)).Equal(0.7853981633974483)
            g.Assert(Dir(-1, 0)).Equal(math.Pi)
            g.Assert(Dir(v.v[0], v.v[1])).Equal(math.Pi)
            g.Assert(Dir(1, math.Sqrt(3))).Equal(Deg2rad(60))
            g.Assert(Dir(0, -1)).Equal(Deg2rad(270))
        })
    })

}

func TestRevdir(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector RevDirection", func() {
        g.It("should test reverse vector direction", func() {
            v := NewVect(&Options{
                A: &Vect2D{0, 0},
                B: &Vect2D{-1, 0},
            })
            g.Assert(Revdir(v.d)).Equal(0.0)
            g.Assert(Revdir(0.7853981633974483)).Equal(0.7853981633974483 + math.Pi)
            g.Assert(Revdir(0.7853981633974483 + math.Pi)).Equal(0.7853981633974483)
        })
    })

}

func TestDeflection(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Deflection", func() {
        g.It("should test reverse vector direction", func() {
            ln0 := []*Vect2D{{0, 0}, {20, 30}}
            ln1 := []*Vect2D{{20, 30}, {40, 15}}

            v0 := NewVect(&Options{A: ln0[0], B: ln0[1]})
            v1 := NewVect(&Options{A: ln1[0], B: ln1[1]})

            g.Assert(Round(Deflect(v0.d, v1.d), 10)).Equal(Round(Deg2rad(93.17983011986422), 10))
            g.Assert(Round(Deflect(v0.d, v0.d), 10)).Equal(Deg2rad(0.0))

            ln1 = []*Vect2D{{20, 30}, {20, 60}}
            v1 = NewVect(&Options{A: ln1[0], B: ln1[1]})

            g.Assert(Round(Deflect(v0.d, v1.d), 10)).Equal(
                Round(Deg2rad(-33.690067525979806), 10),
            )
        })
    })

}

func TestAngleAtPnt(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Angle at Point", func() {
        g.It("should test angle formed at point by vector", func() {

            a := &Vect2D{-1.28, 0.74}
            b := &Vect2D{1.9, 4.2}
            c := &Vect2D{3.16, -0.84}

            v := NewVect(&Options{A: b, B: c})
            g.Assert(Round(AngleAtPt(a, b, c), 8)).Equal(Round(1.1694239325184717, 8), )
            g.Assert(Round(AngleAtPt(a, v.a, v.b), 8)).Equal(Round(1.1694239325184717, 8), )
            g.Assert(Round(AngleAtPt(b, a, c), 8)).Equal(Round(0.9882331199311394, 8), )
        })
    })

}

func TestAdd(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Operators", func() {
        g.It("should test add", func() {
            negDplusB := Add(Neg(D), B)
            g.Assert(negDplusB).Eql(&Vect2D{5.0, 4.0})
        })
        g.It("should test sub", func() {
            CminusD := Sub(C, D)
            g.Assert(CminusD).Eql(&Vect2D{-1.5, 4})
            CminusD = Sub(C, D)
            g.Assert(CminusD).Eql(&Vect2D{-1.5, 4})
        })
    })

}

func TestUnit(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Unit", func() {
        g.It("should test unit vector", func() {
            v := &Vect2D{-3, 2}
            unit_v := Unit(v)
            for i, v := range *unit_v {
                (*unit_v)[i] = Round(v, 6)
            }
            g.Assert(unit_v).Equal(&Vect2D{-0.83205, 0.5547})
        })
    })

}

//dot perform dot in 2d even when 3d coords are passed
func TestDot(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Dot Product", func() {
        g.It("should test dot product", func() {
            dot_prod := Dot(&Vect2D{1.2, -4.2}, &Vect2D{1.2, -4.2})
            g.Assert(19.08).Equal(Round(dot_prod, 8))
        })
    })

}

func TestProj(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Project", func() {
        g.It("should test projection", func() {
            g.Assert( Round(Proj(A, B), 5)).Equal( 0.56121)
        })
    })
}


