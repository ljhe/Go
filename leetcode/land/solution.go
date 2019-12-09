package main

import (
	"fmt"
)

func find(grid [][]byte, i, j int)  {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
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
	row := len(grid)
	num := 0
	if row == 0 {
		return num
	}
	col := len(grid[0])
	if col == 0 {
		return num
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
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
	grid = [][]byte{{}}
	fmt.Println(numIslands(grid))
}


