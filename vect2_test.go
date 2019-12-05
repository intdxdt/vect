package vect

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
	"github.com/intdxdt/side"
	"testing"
)

var A2 = geom.Point{0.88682, -1.06102}
var B2 = geom.Point{3.5, 1}
var C2 = geom.Point{-3, 1}
var D2 = geom.Point{-1.5, -3}

var m = 25.0
var dir = math.Deg2rad(165.0)
var cx, cy = geom.Component(m, dir)
var va = NewVect(geom.Point{0, 0}, geom.Point{cx, cy})
var va_b = geom.Point{-24.148145657226706, 6.470476127563026}

//Test Init Vector
func TestDistToVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Dist2Vect", func() {
		g.It("should test distance vector", func() {
			var a = geom.Point{16.82295, 10.44635}
			var b = geom.Point{28.99656, 15.76452}
			var on_ab = geom.Point{25.32, 14.16}

			tpoints := []geom.Point{
				{30., 0.}, {15.78786, 25.26468}, {-2.61504, -3.09018}, {28.85125, 27.81773},
				a, b, on_ab,
			}

			var t_dists = []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
			var tvect = NewVect(a, b)
			var dists = make([]float64, len(tpoints))
			var dists2 = make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = tvect.DistanceToPoint(&tp)
				dists2[i] = tvect.DistanceToPoint(&tp)
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
			var k = geom.Point{-0.887, -1.6128}
			var u = geom.Point{4.55309, 1.42996}

			var testpoints = []geom.Point{
				{2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k,
			}

			var v = NewVect(k, u)
			var left, right, on = side.NewSide().AsLeft(),
				side.NewSide().AsRight(), side.NewSide().AsOn()

			sides := make([]*side.Side, len(testpoints))
			for i, pnt := range testpoints {
				sides[i] = v.SideOf(&pnt)
			}
			g.Assert(v.SideOf(&geom.Point{2, 2}).IsLeft()).IsTrue()

			var sideOut = []*side.Side{left, left, right, right, left, right, on, on}

			for i := range sideOut {
				g.Assert(sides[i]).Eql(sideOut[i])
			}
		})
	})

}

func TestSEDVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SEDVector", func() {
		g.It("should test side sed vector to point at time T", func() {

			var a = []float64{10, 150, 6.5}
			var e = []float64{280, 280, 12.8}
			var i = []float64{185, 155, 8.6}
			var ai = geom.Point{i[x], i[y]}
			var v = NewVect(
				geom.CreatePoint(a),
				geom.CreatePoint(e),
			)

			var sedV = v.SEDVector(ai, i[2])
			var sedV2 = v.SEDVector(ai, i[2])

			g.Assert(math.Round(sedV.Magnitude(), prec)).Equal(93.24400487)
			g.Assert(math.Round(sedV2.Magnitude(), prec)).Equal(93.24400487)
		})
	})

}

func TestExtVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Extend", func() {
		g.It("should test extending A vector", func() {

			var va = NewVect(geom.Point{}, A2)
			var vb = NewVect(geom.Point{}, B2)
			var vc = NewVect(geom.Point{}, C2)
			var vd = NewVect(geom.Point{}, D2)
			var vdb = NewVect(D2, B2)

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
			g.Assert(va.A[0]).Equal(0.)
			g.Assert(vc.A[0]).Equal(vd.A[0])
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
			g.Assert(vo.B).Eql(vo_defl.B)
			// "vo by extending vdb by angle to origin"
			g.Assert(math.Round(vo.B[0], prec)).Equal(0.0)
			// "vo by extending vdb by angle to origin"
			g.Assert(math.Round(vo.B[1], 4)).Equal(math.Round(0.0, prec))
			var deflangleB = 141.34019174590992
			var inclangleD = 71.89623696549336
			// extend to c from end
			vextc := vdb.ExtendVect(6.5, math.Deg2rad(180+deflangleB), true)
			////extend to c from begining
			var vextCFromD = vdb.ExtendVect(4.272001872658765, math.Deg2rad(inclangleD), false)
			// deflect to c from begin
			vdeflCFromD := vdb.DeflectVector(4.272001872658765, math.Deg2rad(180-inclangleD), false)
			// "comparing extend and deflect from begin point D"
			g.Assert(vextCFromD.B).Eql(vdeflCFromD.B)
			// "vextc from B and from D : extending vdb by angle to C"
			g.Assert(math.Round(vextCFromD.B[0], prec)).Equal(math.Round(vextc.B[0], prec))
			// "vextc from B and from D : extending vdb by angle to C"
			g.Assert(math.Round(vextCFromD.B[1], prec)).Equal(math.Round(vextc.B[1], prec))
			// "vextc by extending vdb by angle to C"
			g.Assert(math.Round(vextc.B[0], prec)).Equal(C[0])
			// "vextc by extending vdb by angle to C"
			g.Assert(math.Round(vextc.B[1], 4)).Equal(C[1])
			// "vextc with magnitudie extension from vdb C"
			g.Assert(math.Round(vextc.V[0], prec)).Equal(-vextc.Magnitude())
			// "vextc horizontal vector test:  extension from vdb C"
			g.Assert(math.Round(vextc.V[1], prec)).Equal(0.)
		})
	})

}

func TestVectDirMag(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Direction - Magnitude", func() {
		g.It("should test vector direction and magnitude", func() {
			// "va endpoints equality: 0 "
			g.Assert(math.Round(va.B[0], prec)).Equal(
				math.Round(va_b[0], prec),
			)
			// "va endpoints equality: 1 "
			g.Assert(math.Round(va.B[1], prec)).Equal(
				math.Round(va_b[1], prec),
			)
			g.Assert(math.FloatEqual(va.Magnitude(), 25.)).IsTrue()
			g.Assert(math.Deg2rad(165)).Equal(va.Direction())
			g.Assert(va.A[0]).Equal(0.0)
			g.Assert(va.A[1]).Equal(0.0)

			// "endpoint should be same as vector: 0 "
			g.Assert(math.Round(va.B[0], prec)).Equal(math.Round(va.V[0], prec), )
			// "endpoint should be same as vector: 1 "
			g.Assert(math.Round(va.B[1], prec)).Equal(math.Round(va.V[1], prec), )
		})
	})

}
