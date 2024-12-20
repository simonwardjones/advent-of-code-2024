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

func manhattanDistance(a, b [2]int) int {
	return absDiff(a[0], b[0]) + absDiff(a[1], b[1])
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getPositions(grid map[[2]int]bool, start,  end [2]int) map[[2]int]int {
	positions := map[[2]int]int{start: 0}
	current := start
	for current != end {
		for _, next := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newPoint := [2]int{current[0] + next[0], current[1] + next[1]}
			if _, visited := positions[newPoint]; !visited && grid[newPoint] {
				positions[newPoint] = positions[current] + 1
				current = newPoint
				break
			}
		}
	}
	return positions
}

func part1(data []string) int {
	grid, start, end := passInput(data)
	positions := getPositions(grid, start, end)
	cheats := 0
	for pos1, num1 := range positions {
		for pos2, num2 := range positions {
			if manhattanDistance(pos1, pos2) == 2 && absDiff(num1, num2) >= 20 {
				cheats += 1
			}
		}
	}
	return cheats / 2
}

func part2(data []string) int {
	grid, start, end := passInput(data)
	positions := getPositions(grid, start, end)
	savings:= map[int]int{}
	totalOver := 0
	for pos1, num1 := range positions {
		for pos2, num2 := range positions {
			cheatDistance := manhattanDistance(pos1, pos2)
			saving := num2 - num1 - cheatDistance
			if cheatDistance <= 20 && saving >= 100 {
				savings[saving] += 1
				totalOver += 1
			}
		}
	}
	return totalOver
}

func passInput(data []string) (grid map[[2]int]bool, start, end [2]int) {
	grid = make(map[[2]int]bool)
	for row, line := range data {
		for col, char := range line {
			if char == 'S' {
				start = [2]int{row, col}
			}
			if char == 'E' {
				end = [2]int{row, col}
			}
			if char == '.' || char == 'S' || char == 'E' {
				grid[[2]int{row, col}] = true
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
