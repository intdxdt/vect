package vect

import (
	"github.com/intdxdt/geom"
	"math/rand"
	"testing"
	"time"
)

func RandPoint(size int) geom.Point {
	var seed = rand.NewSource(time.Now().UnixNano())
	var rnd = rand.New(seed)
	var x = rnd.Float64() * (100.0 - float64(size))
	var y = rnd.Float64() * (100.0 - float64(size))
	return geom.PointXY(x, y)
}

func GenData(N, size int) []geom.Point {
	var data = make([]geom.Point, N, N)
	for i := 0; i < N; i++ {
		data[i] = RandPoint(size)
	}
	return data
}

var N = int(1e5)
var a = geom.Point{16.82295, 10.44635}
var b = geom.Point{28.99656, 15.76452}

var tvect = NewVect(a, b)
var data = GenData(N, 1)
var distance = &[]float64{0}

func BenchmarkDist(b *testing.B) {
	var dist float64
	for i := 0; i < N; i++ {
		dist = tvect.DistanceToPoint(&data[i])
	}
	(*distance)[0] = dist
}

func BenchmarkDistanceToPoint(b *testing.B) {
	var dist float64
	for i := 0; i < N; i++ {
		dist = tvect.DistanceToPoint(&data[i])
	}
	(*distance)[0] = dist
}
