package vect

import (
	. "github.com/intdxdt/simplex/util/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

var A2 = Vect2D{0.88682, -1.06102}
var B2 = Vect2D{3.5, 1}
var C2 = Vect2D{-3, 1}
var D2 = Vect2D{-1.5, -3}

var va_opts = map[string]interface{}{"a": Vect2D{0, 0}, "m": 25, "d": Torad(165)}
var va = New(va_opts)
var va_b = Vect2D{-24.148145657226706, 6.470476127563026}

//Test Init Vector
func TestDistToVect(t *testing.T) {

	a := Vect2D{16.82295, 10.44635}
	b := Vect2D{28.99656, 15.76452}
	on_ab := Vect2D{25.32, 14.16}

	tpoints := []Vect2D{
		{30., 0.},
		{15.78786, 25.26468},
		{-2.61504, -3.09018},
		{28.85125, 27.81773},
		a, b, on_ab,
	}

	t_dists := []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
	tvect := New(map[string]interface{}{"a": a, "b": b})
	dists := make([]float64, len(tpoints))

	for i, tp := range tpoints {
		dists[i] = tvect.DistToPt(tp)
	}

	for i, _ := range tpoints {
		assert.Equal(t, Round(dists[i], 2), Round(t_dists[i], 2), )
	}
}


func TestSideOfVect(t *testing.T) {
	k := Vect2D{-0.887, -1.6128}
	u := Vect2D{4.55309, 1.42996}
	testpoints := []Vect2D{
		{2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k,
	}
	v := New(map[string]interface{}{"a": k, "b": u})
	left, right := "left", "right"
	sides := make([]string, len(testpoints))
	for i, pnt := range testpoints {
		sides[i] = v.SideOfPt(pnt)
	}
	assert.Equal(t, "left", v.SideOfPt(Vect2D{2, 2}))
	side_out := []string{left, left, right, right, left, right, right, right}

	for i, _ := range side_out {
		assert.Equal(t, sides[i], side_out[i])
	}
}

func TestSEDVect(t *testing.T) {
	a  := Vect3D{10, 150, 6.5}
	e  := Vect3D{280, 280, 12.8}
	i  := Vect3D{185, 155, 8.6}
	ai := Vect2D{i[0], i[1]}
	v  := New(map[string]interface{}{
		"a": a[0:2],
		"b": e[0:2],
		"at": a[2],
		"bt": e[2],
	})

	sed_v  := v.SEDvect(ai, i[2])
	sed_v2 := v.SEDvect(ai, i[2])


	assert.Equal(t, Round(sed_v.m, prec), 93.24400487)
	assert.Equal(t, Round(sed_v2.m, prec), 93.24400487)

}

func TestExtVect(t *testing.T) {
	va := New(map[string]interface{}{"b": A2})
	vb := New(map[string]interface{}{"b": B2})
	vc := New(map[string]interface{}{"b": C2})
	vd := New(map[string]interface{}{"b": D2})
	vdb := New(map[string]interface{}{"a": D2, "b": B2})
	assert.Equal(t,
		Round(va.d, prec),
		Round(Torad(309.889497029295), prec),
	)
	assert.Equal(t,
		Round(vb.d, prec),
		Round(Torad(15.945395900922854), prec),
	)
	assert.Equal(t,
		Round(vc.d, prec),
		Round(Torad(161.565051177078), prec),
	)
	assert.Equal(t,
		Round(vd.d, prec),
		Round(Torad(243.43494882292202), prec),
	)
	assert.Equal(t, va.a[0], 0.)
	assert.Equal(t, vc.a[0], vd.a[0])
	assert.Equal(t,
		Round(vdb.m, 4),
		Round(6.4031242374328485, 4),
	)
	assert.Equal(t,
		Round(vdb.d, prec),
		Round(Torad(38.65980825409009), prec),
	)
	deflangle := 157.2855876468
	vo := vdb.Extvect(3.64005494464026, Torad(180+deflangle), true)
	vo_defl := vdb.Deflectvect(3.64005494464026, Torad(-deflangle), true)
	// , "compare deflection and extending"
	assert.Equal(t, vo.b, vo_defl.b)
	// "vo by extending vdb by angle to origin"
	assert.Equal(t, Round(vo.b[0], prec), 0.0)
	// "vo by extending vdb by angle to origin"
	assert.Equal(t, Round(vo.b[1], 4), Round(0.0, prec))
	deflangle_B := 141.34019174590992
	inclangle_D := 71.89623696549336
	// extend to c from end
	vextc := vdb.Extvect(6.5, Torad(180+deflangle_B), true)
	////extend to c from begining
	vextC_fromD := vdb.Extvect(4.272001872658765, Torad(inclangle_D), false)
	// deflect to c from begin
	vdeflC_fromD := vdb.Deflectvect(4.272001872658765, Torad(180-inclangle_D), false)
	// "comparing extend and deflect from begin point D"
	assert.True(t, cmp_array(vextC_fromD.b, vdeflC_fromD.b))
	// "vextc from B and from D : extending vdb by angle to C"
	assert.Equal(t, Round(vextC_fromD.b[0], prec), Round(vextc.b[0], prec))
	// "vextc from B and from D : extending vdb by angle to C"
	assert.Equal(t, Round(vextC_fromD.b[1], prec), Round(vextc.b[1], prec))
	// "vextc by extending vdb by angle to C"
	assert.Equal(t, Round(vextc.b[0], prec), C[0])
	// "vextc by extending vdb by angle to C"
	assert.Equal(t, Round(vextc.b[1], 4), C[1])
	// "vextc with magnitudie extension from vdb C"
	assert.Equal(t, Round(vextc.v[0], prec), -vextc.m)
	// "vextc horizontal vector test:  extension from vdb C"
	assert.Equal(t, Round(vextc.v[1], prec), 0.)
}


func TestVectDirMag(t *testing.T){

     // "va endpoints equality: 0 "
     assert.Equal(t,
         Round(va.b[0], prec),
         Round(va_b[0], prec),
     )
     // "va endpoints equality: 1 "
     assert.Equal(t,
         Round(va.b[1], prec),
         Round(va_b[1], prec),
     )
     assert.Equal(t,25., va.m)
     assert.Equal(t,Torad(165), va.d)
     assert.Equal(t,0., va.a[0])
     assert.Equal(t,0., va.a[1])
     // "endpoint should be same as vector: 0 "
     assert.Equal(t,
         Round(va.b[0], prec),
         Round(va.v[0], prec),
     )
     // "endpoint should be same as vector: 1 "
     assert.Equal(t,
         Round(va.b[1], prec),
         Round(va.v[1], prec),
     )
}