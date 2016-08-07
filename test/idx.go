package main

import "fmt"

func main() {
	vals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fn := fwd(0, len(vals) - 1)
	idxr := idxer(0, len(vals) - 1)
	fmt.Println(fn(-1))
	fmt.Println(idxr(fn(-1)))
}

func fwd(origin, max int) func(k int) int {
	return func(k int) int {
		if k >= origin && k <= max {
			return k
		} else if k < origin {
			return max + k + 1
		}
		panic("index out of bounds")
	}
}

func idxer(origin, max int) func(k int) int {
	return func(k int) int {
		if k >= origin && k <= max {
			return k
		} else if k > max {
			return k - max - 1
		}
		panic("index out of bounds")
	}
}

