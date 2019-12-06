package main

import (
	"fmt"
)

var num int = 0
var row int = 0
var col int = 0

func find(grid [][]byte, i, j int)  {
	if i < 0 || i >= row || j < 0 || j >= col {
		return
	}

	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	find(grid, i + 1, j)
	find(grid, i, j - 1)
	find(grid, i - 1, j)
	find(grid, i, j + 1)
}

func numIslands(grid [][]byte) int {
	row = len(grid)
	col = len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Println(grid[i][j])
			if grid[i][j] == '1' {
				find(grid, i, j)
				num++
			}
		}
	}
	return num
}

func main()  {
	var grid = [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
	fmt.Println(numIslands(grid))
}


