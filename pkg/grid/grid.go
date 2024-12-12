package grid

import "fmt"

type Grid[T comparable] [][]T

func (grid *Grid[T]) n() int {
	return len(*grid)
}

func (grid *Grid[T]) m() int {
	return len((*grid)[0])
}

func (g *Grid[T]) ValidXY(x, y int) bool {
	n, m := g.n(), g.m()
	return x >= 0 && x < n && y >= 0 && y < m
}

func (grid *Grid[T]) Print() {
	for _, row := range *grid {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func (grid *Grid[T]) FindAll(x T) [][2]int {
	found := make([][2]int, 0)
	for i, row := range *grid {
		for j, cell := range row {
			if cell == x {
				found = append(found, [2]int{i, j})
			}
		}
	}
	return found
}

func (grid *Grid[T]) CrossNeighbours(x, y int) [][2]int {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	neighbours := [][2]int{}
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if grid.ValidXY(nx, ny) {
			neighbours = append(neighbours, [2]int{nx, ny})
		}
	}
	return neighbours
}
