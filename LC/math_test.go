package lc

import (
	"log"
	"math"
	"testing"
)

func init() {
	log.Print("> Math")
}

// 3102h Minimize Manhattan Distances
func Test3102(t *testing.T) {
	minimumDistance := func(points [][]int) int {
		xdst := func(iSkip int) (d, i, j int) {
			x, n, xi, ni := math.MinInt, math.MaxInt, 0, 0
			xd, nd, xdi, ndi := math.MinInt, math.MaxInt, 0, 0

			for i, p := range points {
				if i == iSkip {
					continue
				}

				if x < p[0]+p[1] {
					x, xi = p[0]+p[1], i
				}
				if n > p[0]+p[1] {
					n, xi = p[0]+p[1], i
				}

				if xd < p[0]-p[1] {
					xd, xdi = p[0]-p[1], i
				}
				if nd > p[0]-p[1] {
					nd, xdi = p[0]-p[1], i
				}
			}

			if x-n > xd-nd {
				return x - n, xi, ni
			}
			return xd - nd, xdi, ndi
		}

		x, i, j := xdst(-1)
		log.Print(x, points[i], points[j])

		d1, _, _ := xdst(i)
		d2, _, _ := xdst(j)
		return min(d1, d2)
	}

	log.Print("0 ?= ", minimumDistance([][]int{{1, 1}, {1, 1}, {1, 1}}))
	log.Print("12 ?= ", minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
	log.Print("10 ?= ", minimumDistance([][]int{{3, 2}, {3, 9}, {7, 10}, {4, 4}, {8, 10}, {2, 7}}))
}
