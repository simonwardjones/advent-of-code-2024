package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/simonwardjones/advent-of-code-2024/pkg/grid"
)

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

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

type Region struct {
	area      int
	perimeter int
	points    map[[2]int]bool
}

func getRegions(grid grid.Grid[rune]) []Region {
	visited := make(map[[2]int]bool)
	toVisit := [][2]int{{0, 0}}
	regions := []Region{}
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		if visited[current] {
			continue
		}
		region := Region{area: 0, perimeter: 0, points: map[[2]int]bool{}}
		stack := [][2]int{current}
		for len(stack) > 0 {
			current = stack[0]
			stack = stack[1:]
			if visited[current] {
				continue
			}
			neighbours := grid.CrossNeighbours(current[0], current[1])
			perimeter := 4
			for _, n := range neighbours {
				if grid[n[0]][n[1]] == grid[current[0]][current[1]] {
					perimeter--
					stack = append(stack, n)
				} else {
					toVisit = append(toVisit, n)
				}
			}
			region.perimeter += perimeter
			region.area++
			region.points[current] = true
			visited[current] = true
		}
		regions = append(regions, region)
	}
	return regions
}

func part1(data []string) int {
	grid := passInput(data)
	grid.Print()
	regions := getRegions(grid)
	price := 0
	for _, region := range regions {
		price += region.area * region.perimeter
	}
	return price
}

func part2(data []string) int {
	regions := getRegions(passInput(data))
	price := 0
	for _, region := range regions {
		sides := 0
		for p := range region.points {
			// check for to if nothing above
			if !region.points[[2]int{p[0] - 1, p[1]}] {
				// and to the left is not top
				if !(region.points[[2]int{p[0], p[1] - 1}] &&
					!region.points[[2]int{p[0] - 1, p[1] - 1}]) {
					sides++
				}
			}
			// check for bottom
			if !region.points[[2]int{p[0] + 1, p[1]}] {
				// and to the right is not bottom
				if !(region.points[[2]int{p[0], p[1] + 1}] &&
					!region.points[[2]int{p[0] + 1, p[1] + 1}]) {
					sides++
				}
			}
			// check for left
			if !region.points[[2]int{p[0], p[1] - 1}] {
				// and to the bottom is not left
				if !(region.points[[2]int{p[0] + 1, p[1]}] &&
					!region.points[[2]int{p[0] + 1, p[1] - 1}]) {
					sides++
				}
			}
			// check for right
			if !region.points[[2]int{p[0], p[1] + 1}] {
				// and to the top is not right
				if !(region.points[[2]int{p[0] - 1, p[1]}] &&
					!region.points[[2]int{p[0] - 1, p[1] + 1}]) {
					sides++
				}
			}

		}
		price += region.area * sides
	}
	return price
}

func passInput(data []string) grid.Grid[rune] {
	g := make(grid.Grid[rune], len(data))
	for i, line := range data {
		g[i] = []rune(line)
	}
	return g
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
