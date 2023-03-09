package gen

import (
	"math/rand"

	"github.com/ereminiu/ginsudoku/solver"
)

type Generator struct {
	n    int
	grid [][]int
}

func randint(n int) int {
	return rand.Intn(n)
}

func NewGenerator() *Generator {
	n := 9
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			grid[i][j] = (3*i+i/3+j)%9 + 1
		}
	}

	return &Generator{9, grid}
}

func (g *Generator) transpose() {
	for i := 0; i < g.n; i++ {
		for j := 0; j < i; j++ {
			g.grid[i][j], g.grid[j][i] = g.grid[j][i], g.grid[i][j]
		}
	}
}

func (g *Generator) swapRows() {
	d := randint(3)
	x, y := 3*d+randint(3), 3*d+randint(3)

	for x == y {
		y = 3*d + randint(3)
	}

	g.grid[x], g.grid[y] = g.grid[y], g.grid[x]
}

func (g *Generator) swapCols() {
	g.transpose()
	g.swapRows()
	g.transpose()
}

func (g *Generator) swapBoxH() {
	x, y := 3*randint(3), 3*randint(3)

	for x == y {
		y = 3 * randint(3)
	}

	for i := 0; i < 3; i++ {
		g.grid[x+i], g.grid[y+i] = g.grid[y+i], g.grid[x+i]
	}
}

func (g *Generator) swapBoxV() {
	g.transpose()
	g.swapBoxH()
	g.transpose()
}

func (g *Generator) shuffle() {
	g.transpose()
	for it := 0; it < 20; it++ {
		for rep := 0; rep < 10; rep++ {
			idx := randint(4)
			switch idx {

			case 0:
				g.swapRows()
				continue

			case 1:
				g.swapCols()
				continue

			case 2:
				g.swapBoxH()
				continue

			case 3:
				g.swapBoxV()
				continue
			}
		}
	}
}

func copy(g *[][]int) [][]int {
	ret := make([][]int, len(*g))
	for i := 0; i < len(*g); i++ {
		ret[i] = make([]int, len(*g))
		for j := 0; j < len(*g); j++ {
			ret[i][j] = (*g)[i][j]
		}
	}
	return ret
}

func (g *Generator) removeCells(cells int) {
	for removed := 0; removed < cells; {
		i, j := randint(9), randint(9)
		val := g.grid[i][j]
		g.grid[i][j] = 0

		table := copy(&g.grid)

		if !solver.CheckSolvable(&table) {
			g.grid[i][j] = val
		} else {
			removed++
		}
	}
}

func (g *Generator) Get() [][]int {
	g.shuffle()
	g.removeCells(68)

	return g.grid
}
