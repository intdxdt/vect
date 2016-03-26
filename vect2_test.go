package vect

import (
    . "github.com/intdxdt/simplex/util/math"
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/franela/goblin"
    "testing"
)

var A2 = &Point{0.88682, -1.06102}
var B2 = &Point{3.5, 1}
var C2 = &Point{-3, 1}
var D2 = &Point{-1.5, -3}

var m = 25.0
var dir = Deg2rad(165.0)
var va_opts = &Options{A: &Point{0, 0}, M: &m, D: &dir}
var va = NewVect(va_opts)
var va_b = Point{-24.148145657226706, 6.470476127563026}

//Test Init Vector
func TestDistToVect(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Dist2Vect", func() {
        g.It("should test distance vector", func() {
            a := &Point{16.82295, 10.44635}
            b := &Point{28.99656, 15.76452}
            on_ab := &Point{25.32, 14.16}

            tpoints := []*Point{
                {30., 0.},
                {15.78786, 25.26468},
                {-2.61504, -3.09018},
                {28.85125, 27.81773},
                a,
                b,
                on_ab,
            }

            t_dists := []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
            tvect := NewVect(&Options{A: a, B: b})
            dists := make([]float64, len(tpoints))

            for i, tp := range tpoints {
                dists[i] = tvect.Dist2Pt(tp)
            }

            for i, _ := range tpoints {
                g.Assert(Round(dists[i], 2)).Equal(Round(t_dists[i], 2))
            }
        })
    })
}

func TestSideOfVect(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector Sidedness", func() {
        g.It("should test side of point to vector", func() {
            k := &Point{-0.887, -1.6128}
            u := &Point{4.55309, 1.42996}

            testpoints := []*Point{
                {2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k,
            }

            v := NewVect(&Options{A: k, B: u})

            left, right := Sided.Left, Sided.Right

            sides := make([]Side, len(testpoints))
            for i, pnt := range testpoints {
                sides[i] = v.SideOfPt(pnt)
            }
            g.Assert(Sided.Left).Equal(v.SideOfPt(&Point{2, 2}))

            side_out := []Side{
                left, left, right, right, left,
                right, right, right,
            }

            for i, _ := range side_out {
                g.Assert(sides[i]).Equal(side_out[i])
            }
        })
    })

}

func TestSEDVect(t *testing.T) {
    g := Goblin(t)
    g.Describe("SEDVector", func() {
        g.It("should test side sed vector to point at time T", func() {

            a := &Vect3D{10, 150, 6.5}
            e := &Vect3D{280, 280, 12.8}
            i := &Vect3D{185, 155, 8.6}
            ai := &Point{i[x], i[y]}

            v := NewVect(&Options{
                A   : &Point{a[x], a[y]},
                B   : &Point{e[x], e[y]},
                At  : &a[2],
                Bt  : &e[2],
            })

            sed_v := v.SEDVector(ai, i[2])
            sed_v2 := v.SEDVector(ai, i[2])

            g.Assert(Round(sed_v.m, prec)).Equal(93.24400487)
            g.Assert(Round(sed_v2.m, prec)).Equal(93.24400487)
        })
    })

}

func TestExtVect(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Extend", func() {
        g.It("should test extending a vector", func() {

            va := NewVect(&Options{B: A2})
            vb := NewVect(&Options{B: B2})
            vc := NewVect(&Options{B: C2})
            vd := NewVect(&Options{B: D2})
            vdb := NewVect(&Options{A: D2, B: B2})

            g.Assert(Round(va.d, prec)).Equal(
                Round(Deg2rad(309.889497029295), prec),
            )
            g.Assert(Round(vb.d, prec)).Equal(
                Round(Deg2rad(15.945395900922854), prec),
            )
            g.Assert(Round(vc.d, prec)).Equal(
                Round(Deg2rad(161.565051177078), prec),
            )
            g.Assert(Round(vd.d, prec)).Equal(
                Round(Deg2rad(243.43494882292202), prec),
            )
            g.Assert( va.a[0]).Equal( 0.)
            g.Assert( vc.a[0]).Equal( vd.a[0])
            g.Assert(Round(vdb.m, 4)).Equal(
                Round(6.4031242374328485, 4),
            )
            g.Assert(Round(vdb.d, prec)).Equal(
                Round(Deg2rad(38.65980825409009), prec),
            )
            deflangle := 157.2855876468
            vo := vdb.Extvect(3.64005494464026, Deg2rad(180 + deflangle), true)
            vo_defl := vdb.DeflectVector(3.64005494464026, Deg2rad(-deflangle), true)
            // , "compare deflection and extending"
            g.Assert (vo.b).Eql(vo_defl.b)
            // "vo by extending vdb by angle to origin"
            g.Assert (Round(vo.b[0], prec)).Equal( 0.0)
            // "vo by extending vdb by angle to origin"
            g.Assert (Round(vo.b[1], 4)).Equal(Round(0.0, prec))
            deflangle_B := 141.34019174590992
            inclangle_D := 71.89623696549336
            // extend to c from end
            vextc := vdb.Extvect(6.5, Deg2rad(180 + deflangle_B), true)
            ////extend to c from begining
            vextC_fromD := vdb.Extvect(4.272001872658765, Deg2rad(inclangle_D), false)
            // deflect to c from begin
            vdeflC_fromD := vdb.DeflectVector(4.272001872658765, Deg2rad(180 - inclangle_D), false)
            // "comparing extend and deflect from begin point D"
            g.Assert (vextC_fromD.b).Eql(vdeflC_fromD.b)
            // "vextc from B and from D : extending vdb by angle to C"
            g.Assert (Round(vextC_fromD.b[0], prec)).Equal(Round(vextc.b[0], prec))
            // "vextc from B and from D : extending vdb by angle to C"
            g.Assert (Round(vextC_fromD.b[1], prec)).Equal(Round(vextc.b[1], prec))
            // "vextc by extending vdb by angle to C"
            g.Assert (Round(vextc.b[0], prec)).Equal( C[0])
            // "vextc by extending vdb by angle to C"
            g.Assert (Round(vextc.b[1], 4)).Equal(C[1])
            // "vextc with magnitudie extension from vdb C"
            g.Assert (Round(vextc.v[0], prec)).Equal(-vextc.m)
            // "vextc horizontal vector test:  extension from vdb C"
            g.Assert (Round(vextc.v[1], prec)).Equal( 0.)
        })
    })

}

func TestVectDirMag(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector - Direction - Magnitude", func() {
        g.It("should test vector direction and magnitude", func() {
            // "va endpoints equality: 0 "
            g.Assert(Round(va.b[0], prec)).Equal(
                Round(va_b[0], prec),
            )
            // "va endpoints equality: 1 "
            g.Assert(Round(va.b[1], prec)).Equal(
                Round(va_b[1], prec),
            )
            g.Assert( 25.).Equal(va.m)
            g.Assert( Deg2rad(165)).Equal( va.d)
            g.Assert( va.a[0]).Equal(0.0)
            g.Assert( va.a[1]).Equal(0.0)

            // "endpoint should be same as vector: 0 "
            g.Assert(Round(va.b[0], prec)).Equal(Round(va.v[0], prec), )
            // "endpoint should be same as vector: 1 "
            g.Assert(Round(va.b[1], prec)).Equal(Round(va.v[1], prec), )
        })
    })


}