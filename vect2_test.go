package vect

import (
	"testing"
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/side"
	"github.com/franela/goblin"
)

var A2 = &geom.Point{0.88682, -1.06102}
var B2 = &geom.Point{3.5, 1}
var C2 = &geom.Point{-3, 1}
var D2 = &geom.Point{-1.5, -3}

var m = 25.0
var dir = math.Deg2rad(165.0)
var va_opts = &Options{A: &geom.Point{0, 0}, M: &m, D: &dir}
var va = NewVect(va_opts)
var va_b = geom.Point{-24.148145657226706, 6.470476127563026}

//Test Init Vector
func TestDistToVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Dist2Vect", func() {
		g.It("should test distance vector", func() {
			a := &geom.Point{16.82295, 10.44635}
			b := &geom.Point{28.99656, 15.76452}
			on_ab := &geom.Point{25.32, 14.16}

			tpoints := []*geom.Point{
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
			dists2 := make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = tvect.DistanceToPoint(tp)
				dists2[i] = tvect.DistanceToPoint(tp)
			}

			for i := range tpoints {
				g.Assert(math.Round(dists[i], 2)).Equal(math.Round(t_dists[i], 2))
			}
		})
	})
}

func TestSideOfVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Sidedness", func() {
		g.It("should test side of point to vector", func() {
			k := &geom.Point{-0.887, -1.6128}
			u := &geom.Point{4.55309, 1.42996}

			testpoints := []*geom.Point{
				{2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k,
			}

			v := NewVect(&Options{A: k, B: u})

			left, right, on := side.NewSide().AsLeft(), side.NewSide().AsRight(), side.NewSide().AsOn()

			sides := make([]*side.Side, len(testpoints))
			for i, pnt := range testpoints {
				sides[i] = v.SideOf(pnt)
			}
			g.Assert(v.SideOf(&geom.Point{2, 2}).IsLeft()).IsTrue()

			side_out := []*side.Side{
				left, left, right, right, left,
				right, on, on,
			}

			for i := range side_out {
				g.Assert(sides[i]).Eql(side_out[i])
			}
		})
	})

}

func TestSEDVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SEDVector", func() {
		g.It("should test side sed vector to point at time T", func() {

			a := []float64{10, 150, 6.5}
			e := []float64{280, 280, 12.8}
			i := []float64{185, 155, 8.6}
			ai := &geom.Point{i[x], i[y]}

			v := NewVect(&Options{
				A:  &geom.Point{a[x], a[y]},
				B:  &geom.Point{e[x], e[y]},
				At: &a[2],
				Bt: &e[2],
			})

			sed_v := v.SEDVector(ai, i[2])
			sed_v2 := v.SEDVector(ai, i[2])

			g.Assert(math.Round(sed_v.Magnitude(), prec)).Equal(93.24400487)
			g.Assert(math.Round(sed_v2.Magnitude(), prec)).Equal(93.24400487)
		})
	})

}

func TestExtVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Extend", func() {
		g.It("should test extending a vector", func() {

			va := NewVect(&Options{B: A2})
			vb := NewVect(&Options{B: B2})
			vc := NewVect(&Options{B: C2})
			vd := NewVect(&Options{B: D2})
			vdb := NewVect(&Options{A: D2, B: B2})

			g.Assert(math.Round(va.Direction(), prec)).Equal(
				math.Round(math.Deg2rad(309.889497029295), prec),
			)
			g.Assert(math.Round(vb.Direction(), prec)).Equal(
				math.Round(math.Deg2rad(15.945395900922854), prec),
			)
			g.Assert(math.Round(vc.Direction(), prec)).Equal(
				math.Round(math.Deg2rad(161.565051177078), prec),
			)
			g.Assert(math.Round(vd.Direction(), prec)).Equal(
				math.Round(math.Deg2rad(243.43494882292202), prec),
			)
			g.Assert(va.a[0]).Equal(0.)
			g.Assert(vc.a[0]).Equal(vd.a[0])
			g.Assert(math.Round(vdb.Magnitude(), 4)).Equal(
				math.Round(6.4031242374328485, 4),
			)
			g.Assert(math.Round(vdb.Direction(), prec)).Equal(
				math.Round(math.Deg2rad(38.65980825409009), prec),
			)
			deflangle := 157.2855876468
			vo := vdb.ExtendVect(3.64005494464026, math.Deg2rad(180+deflangle), true)
			vo_defl := vdb.DeflectVector(3.64005494464026, math.Deg2rad(-deflangle), true)
			// , "compare deflection and extending"
			g.Assert(vo.b).Eql(vo_defl.b)
			// "vo by extending vdb by angle to origin"
			g.Assert(math.Round(vo.b[0], prec)).Equal(0.0)
			// "vo by extending vdb by angle to origin"
			g.Assert(math.Round(vo.b[1], 4)).Equal(math.Round(0.0, prec))
			deflangle_B := 141.34019174590992
			inclangle_D := 71.89623696549336
			// extend to c from end
			vextc := vdb.ExtendVect(6.5, math.Deg2rad(180+deflangle_B), true)
			////extend to c from begining
			vextC_fromD := vdb.ExtendVect(4.272001872658765, math.Deg2rad(inclangle_D), false)
			// deflect to c from begin
			vdeflC_fromD := vdb.DeflectVector(4.272001872658765, math.Deg2rad(180-inclangle_D), false)
			// "comparing extend and deflect from begin point D"
			g.Assert(vextC_fromD.b).Eql(vdeflC_fromD.b)
			// "vextc from B and from D : extending vdb by angle to C"
			g.Assert(math.Round(vextC_fromD.b[0], prec)).Equal(math.Round(vextc.b[0], prec))
			// "vextc from B and from D : extending vdb by angle to C"
			g.Assert(math.Round(vextC_fromD.b[1], prec)).Equal(math.Round(vextc.b[1], prec))
			// "vextc by extending vdb by angle to C"
			g.Assert(math.Round(vextc.b[0], prec)).Equal(C[0])
			// "vextc by extending vdb by angle to C"
			g.Assert(math.Round(vextc.b[1], 4)).Equal(C[1])
			// "vextc with magnitudie extension from vdb C"
			g.Assert(math.Round(vextc.v[0], prec)).Equal(-vextc.Magnitude())
			// "vextc horizontal vector test:  extension from vdb C"
			g.Assert(math.Round(vextc.v[1], prec)).Equal(0.)
		})
	})

}

func TestVectDirMag(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Direction - Magnitude", func() {
		g.It("should test vector direction and magnitude", func() {
			// "va endpoints equality: 0 "
			g.Assert(math.Round(va.b[0], prec)).Equal(
				math.Round(va_b[0], prec),
			)
			// "va endpoints equality: 1 "
			g.Assert(math.Round(va.b[1], prec)).Equal(
				math.Round(va_b[1], prec),
			)
			g.Assert(math.FloatEqual(va.Magnitude(), 25.)).IsTrue()
			g.Assert(math.Deg2rad(165)).Equal(va.Direction())
			g.Assert(va.a[0]).Equal(0.0)
			g.Assert(va.a[1]).Equal(0.0)

			// "endpoint should be same as vector: 0 "
			g.Assert(math.Round(va.b[0], prec)).Equal(math.Round(va.v[0], prec), )
			// "endpoint should be same as vector: 1 "
			g.Assert(math.Round(va.b[1], prec)).Equal(math.Round(va.v[1], prec), )
		})
	})

}
