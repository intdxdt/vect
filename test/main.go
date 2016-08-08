package main

import (
	"fmt"
	"math"
	. "simplex/geom"
	. "simplex/vect"
	. "simplex/util/math"
	. "simplex/struct/item"
)


var A = NewPointXY(172.0, 224.0)
var B = NewPointXY(180.0, 158.0)
var C = NewPointXY(266.0, 46.0)
var D = NewPointXY(374.0, 38.0)
var E = NewPointXY(480.0, 100.0)
var F = NewPointXY(500.0, 200.0)
var G = NewPointXY(440.0, 300.0)
var H = NewPointXY(340.0, 280.0)
var I = NewPointXY(200.0, 240.0)

type Hull struct {
	H []*Point
}

func NewHull(coords []*Point) *Hull {
	return &Hull{H:coords}
}

func (self *Hull) Antipodal(i, j int) int {
	var fn = self.chainIndexer(i, len(self.H) - 1)
	var idxer = self.indexer(i, len(self.H) - 1)

	var hv = NewVect(&Options{A: self.H[i], B: self.H[j], })
	var start, end = fn(i), fn(i - 1)

	var mid = (start + end) / 2
	var pt, ptj = self.H[idxer(mid)], self.H[j]

	var uvect = func(m int) *Vector {
		return NewVector(ptj, self.H[m])
	}

	var angl = Deg2rad(90.0)
	var side = hv.SideOfPt(pt)

	if side.IsOn() {
		return end
	} else if side.IsLeft() {
		angl = Deg2rad(270.0)
	}

	vv := hv.Extvect(1e3, angl, true)
	fmt.Println(vv.V())
	orth := self.orthvector(hv.D(), angl)

	for {
		if start == end {
			mid = start
			break
		}
		mid = (start + end) / 2

		cur := self.othoffset(orth, uvect(idxer(mid)))
		next := self.othoffset(orth, uvect(idxer(mid + 1)))

		if cur.Compare(next) == 0 {
			mid += 1
			break
		} else {
			if cur.Compare(next) < 0 {
				start = mid + 1
			} else if cur.Compare(next) > 0 {
				end = mid
			} else {
				break
			}
		}
	}
	return idxer(mid)
}

func (self *Hull) othoffset(v, u *Vector) Float {
	return Float(v.Project(u))
}

func (self *Hull)  orthvector(direction, angle float64) *Vector {
	m := 1e3
	bβ := direction + Pi //back bearing
	fβ := bβ + angle //forward bearing
	if fβ > Tau {
		fβ -= Tau
	}
	return NewVectorXY(m * math.Cos(fβ), m * math.Sin(fβ))
}

func (self *Hull) indexer(origin, max int) func(k int) int {
	return func(k int) int {
		if k >= origin && k <= max {
			return k
		} else if k > max {
			return k - max - 1
		}
		panic("index out of bounds")
	}
}

func (self *Hull) chainIndexer(origin, max int) func(k int) int {
	return func(k int) int {
		if k >= origin && k <= max {
			return k
		} else if k < origin {
			return max + k + 1
		}
		panic("index out of bounds")
	}
}

func main() {
	coords := []*Point{A, B, C, D, E, F, G, H, I}
	fmt.Println(NewLineString(coords).WKT())
	hull := NewHull(coords)
	ext := hull.Antipodal(0, 1)
	fmt.Println(ext)

	v:= NewVector(A, A)
	fmt.Println(v.UnitVector())
}
