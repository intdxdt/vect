package vect

//import (
//    "testing"
//    "math/rand"
//    "time"
//    . "github.com/intdxdt/simplex/geom"
//)
//
//func RandPoint(size int) *Point {
//    _size := float64(size)
//    seed := rand.NewSource(time.Now().UnixNano())
//    rnd := rand.New(seed)
//    var x = rnd.Float64() * (100.0 - _size);
//    var y = rnd.Float64() * (100.0 - _size);
//    return NewPointXY(x, y)
//}
//
//func GenData(N, size int) []*Point {
//    var data = make([]*Point, N, N);
//    for i := 0; i < N; i++ {
//        data[i] = RandPoint(size);
//    }
//    return data;
//};
//
//var N = int(3e3)
//var a = &Point{16.82295, 10.44635}
//var b = &Point{28.99656, 15.76452}
//
//var tvect = NewVect(&Options{A: a, B: b})
//var data = GenData(N, 1)
//
//func Benchmark_Dist(b *testing.B) {
//    for i := 0; i < N; i++ {
//        tvect.Dist2Pt(data[i])
//    }
//}
//
//func Benchmark_DistanceToPoint(b *testing.B) {
//    for i := 0; i < N; i++ {
//        tvect.DistanceToPoint(data[i])
//    }
//}