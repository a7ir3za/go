package lc

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.Print("> BFS/DFS")
}

// 1091m Shortest Path in Binary Matrix
func Test1091(t *testing.T) {
	shortestPathBinaryMatrix := func(grid [][]int) int {
		type P struct{ i, j int }

		Q := []P{}
		if grid[0][0] == 0 {
			Q = append(Q, P{0, 0})
			grid[0][0] = 1
		}

		m, n := len(grid), len(grid[0])

		var v P
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			if v.i == m-1 && v.j == n-1 {
				return grid[m-1][n-1]
			}

			for _, d := range [8]P{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}} {
				u := P{v.i + d.i, v.j + d.j}

				if u.i >= 0 && u.j >= 0 && m > u.i && n > u.j && grid[u.i][u.j] == 0 {
					grid[u.i][u.j] = 1 + grid[v.i][v.j]

					Q = append(Q, u)
				}
			}
		}

		return -1
	}

	gridSet := func() [][]int { return [][]int{{0, 0, 0, 0, 0}, {1, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 1, 1, 0}} }
	grid := gridSet()

	draw := func() {
		for i := range grid {
			for _, v := range grid[i] {
				fmt.Printf("| %d ", v)
			}
			fmt.Println("|")
		}
	}
	draw()
	log.Print("5 ?= ", shortestPathBinaryMatrix(grid))
	draw()
}