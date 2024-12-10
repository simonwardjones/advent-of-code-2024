package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/simonwardjones/advent-of-code-2024/pkg/stack"
)

var (
	Up    = [2]int{-1, 0}
	Down  = [2]int{1, 0}
	Left  = [2]int{0, -1}
	Right = [2]int{0, 1}
)
var Directions = [4][2]int{Up, Down, Left, Right}

func main() {
	inputFileName := "input.txt"
	for _, arg := range os.Args {
		if arg == "--test" || arg == "test" {
			inputFileName = "input_test_1.txt"
			break
		} else if arg == "--test2" || arg == "test2" {
			inputFileName = "input_test_2.txt"
		}
	}

	data := loadInput(inputFileName)
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

func part1(data []string) int {
	grid := passInput(data)
	grid.Print()
	zeroLocations := grid.FindAll(0)
	fmt.Println(zeroLocations)
	total := 0
	for _, loc := range zeroLocations {
		total += grid.NinesVisitable(loc[0], loc[1])
	}
	return total
}

func part2(data []string) int {
	grid := passInput(data)
	grid.Print()
	zeroLocations := grid.FindAll(0)
	fmt.Println(zeroLocations)
	total := 0
	for _, loc := range zeroLocations {
		total += grid.RoutesToNine(loc[0], loc[1])
	}
	return total
}

type Grid [][]int

func (grid *Grid) n() int {
	return len(*grid)
}

func (grid *Grid) m() int {
	return len((*grid)[0])
}

func (g *Grid) ValidXY(x, y int) bool {
	n, m := g.n(), g.m()
	return x >= 0 && x < n && y >= 0 && y < m
}

func (grid *Grid) Print() {
	for _, row := range *grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func (grid *Grid) FindAll(x int) [][2]int {
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

func (grid *Grid) NinesVisitable(x, y int) int {
	total := 0
	stack := stack.New[[2]int]()
	stack.Push([2]int{x, y})
	visited := make(map[[2]int]struct{})
	for !stack.IsEmpty() {
		loc := stack.Pop()
		visited[loc] = struct{}{}
		x, y := loc[0], loc[1]
		currentVal := (*grid)[x][y]
		if currentVal == 9 {
			total++
			continue
		}
		for _, dir := range Directions {
			newX, newY := x+dir[0], y+dir[1]
			if grid.ValidXY(newX, newY) && (*grid)[newX][newY] == currentVal+1 {
				if _, ok := visited[[2]int{newX, newY}]; !ok {
					stack.Push([2]int{newX, newY})
				}
			}
		}
	}
	return total
}

func (grid *Grid) RoutesToNine(x, y int) int {
	total := 0
	stack := stack.New[[2]int]()
	stack.Push([2]int{x, y})
	for !stack.IsEmpty() {
		loc := stack.Pop()
		x, y := loc[0], loc[1]
		currentVal := (*grid)[x][y]
		if currentVal == 9 {
			total++
			continue
		}
		for _, dir := range Directions {
			newX, newY := x+dir[0], y+dir[1]
			if grid.ValidXY(newX, newY) && (*grid)[newX][newY] == currentVal+1 {
				stack.Push([2]int{newX, newY})
			}
		}
	}
	return total
}

func passInput(data []string) Grid {
	grid := make(Grid, len(data))
	for i, line := range data {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			value, _ := strconv.Atoi(string(char))
			grid[i][j] = value
		}
	}
	return grid
}

func loadInput(fileName string) []string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)
	inputFilePath := filepath.Join(currentDir, fileName)

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
