package main

import (
    . "github.com/intdxdt/simplex/geom"
    . "github.com/intdxdt/simplex/vect"
    "fmt"
)

const (
    x = iota
    y
)


// isLeft(): test if a point is Left|On|Right of an infinite line.
//    Input:  three points a, b, and pt
//    Return: >0 for pt left of the line through a and b
//            =0 for pt on the line
//            <0 for pt right of the line
//    See: Algorithm 1 on Area of Triangles
//side of point in relation to line segment formed by a and b
func SideOf(a, b, pt *Point) int{
    v := (b[x] - a[x]) * (pt[y] - a[y]) - (pt[x] - a[x]) * (b[y] - a[y]);
    var side int;
    if v > 0 {
        side = -1
    } else if v < 0 {
        side = 1
    }
    return side
}

func main() {
    k := &Point{-0.887, -1.6128}
    u := &Point{4.55309, 1.42996}
    testpoints := []*Point{
        {2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k,
    }
    v := NewVect(&Options{A: k, B: u})

    sides := make([]Side, len(testpoints))
    sides2 := make([]int, len(testpoints))

    for i, pnt := range testpoints {
        sides[i] = v.SideOfPt(pnt)
        sides2[i] = SideOf(k, u, pnt)
    }
    fmt.Println(sides)
    fmt.Println(sides2)
    fmt.Println(k)
    fmt.Println(u)

    fmt.Println(NewLineString([]*Point{k, u}))

    //left, right, On := 'L', 'R', 'O'
    //fmt.Println(left)
    //fmt.Println(right)
    //fmt.Println(On)

}
