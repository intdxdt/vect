package main

import (
	"fmt"
	"github.com/intdxdt/simplex/vect"
)

func main() {
	a := []float64{10, 150, 6.5}
	e := []float64{280, 280, 12.8}
	v := vect.New(map[string]interface{}{
		"a": a, "b": e,
		"at": a[2], "bt": e[2],
	})
	pv := v.V()
	nv := vect.Neg(v.V())
	fmt.Println(nv)
	fmt.Println(pv)
	fmt.Println(a)
	fmt.Println(e)
}
