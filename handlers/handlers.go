package handlers

import (
	"fmt"
	"net/http"

	"github.com/ereminiu/ginsudoku/gen"
	"github.com/ereminiu/ginsudoku/solver"
	"github.com/ereminiu/ginsudoku/tools"
	"github.com/gin-gonic/gin"
)

var grid [][]int
var n = 9

func ReadHandler(c *gin.Context) {
	grid = make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cell := "cell_" + tools.ToString(i) + tools.ToString(j)
			grid[i][j] = tools.ToInt(c.PostForm(cell))
		}
	}

	fmt.Println(grid)
	c.Redirect(http.StatusFound, "/getsol")
}

func SolveHandler(c *gin.Context) {
	grid = solver.GetSolution(grid)

	c.HTML(http.StatusOK, "viewer.tmpl", gin.H{
		"grid": grid,
		"flag": false,
	})
}

func GenGridHandler(c *gin.Context) {
	grid = gen.NewGenerator().Get()

	c.HTML(http.StatusOK, "viewer.tmpl", gin.H{
		"grid": grid,
		"flag": true,
	})
}

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func EnterGridHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "grid.tmpl", nil)
}
