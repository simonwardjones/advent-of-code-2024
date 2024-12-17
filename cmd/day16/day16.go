package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	inputFileName := getFileName()
	fmt.Println("Using input file:", inputFileName)

	data := loadInput(inputFileName)
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

type P struct {
	y, x int
}

type state struct {
	pos   P
	dir   P
	score int
	route []P
}

func getMinScore(grid map[P]rune, start, end P) (int, int) {
	toDo := []state{{start, P{0, 0}, 0, []P{start}}}
	visitedScore := make(map[[2]P]int)
	var routes []state = make([]state, 0)

	for len(toDo) > 0 {
		current := toDo[0]
		toDo = toDo[1:]

		if current.pos == end {
			routes = append(routes, current)
			continue
		}

		// if we have been here before in the same direction with a strictly better score, skip
		if bestScore, ok := visitedScore[[2]P{current.pos, current.dir}]; ok && bestScore < current.score {
			continue
		}
		visitedScore[[2]P{current.pos, current.dir}] = current.score

		for _, dir := range []P{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			neighbor := P{current.pos.y + dir.y, current.pos.x + dir.x}
			if val := grid[neighbor]; val != '#' {
				newScore := current.score + 1
				if dir != current.dir {
					newScore += 1000
				}
				newRoute := append([]P{}, current.route...)
				newRoute = append(newRoute, neighbor)
				toDo = append(toDo, state{neighbor, dir, newScore, newRoute})
			}
		}
	}
	minScore := routes[0].score
	for _, route := range routes {
		if route.score < minScore {
			minScore = route.score
		}
	}
	points := make(map[P]bool)
	for i := len(routes) - 1; i >= 0; i-- {
		if routes[i].score == minScore {
			for _, point := range routes[i].route {
				points[point] = true
			}
		}
	}
	return minScore, len(points)
}

func part1(data []string) int {
	grid, start, end := passInput(data)
	score, _ := getMinScore(grid, start, end)
	return score
}

func part2(data []string) int {
	grid, start, end := passInput(data)
	_, routeSquares := getMinScore(grid, start, end)
	return routeSquares
}

func passInput(data []string) (map[P]rune, P, P) {
	grid := make(map[P]rune)
	var start, end P
	for y, line := range data {
		for x, char := range line {
			pos := P{y, x}
			grid[pos] = char
			switch char {
			case 'S':
				start = pos
			case 'E':
				end = pos
			}
		}
	}
	return grid, start, end
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

func getFileName() string {
	inputFileName := "input.txt"
	for i, arg := range os.Args {
		if strings.Contains(arg, "test") || strings.Contains(arg, "-t") {
			for _, num := range "123456789" {
				if strings.Contains(os.Args[i], string(num)) ||
					(i+1 < len(os.Args) && strings.Contains(os.Args[i+1], string(num))) {
					return fmt.Sprintf("input_test_%s.txt", string(num))
				}
			}
			return "input_test_1.txt"
		}
	}
	return inputFileName
}
